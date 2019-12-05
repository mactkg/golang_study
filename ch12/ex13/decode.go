// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 344.

// Package sexpr provides a means for converting Go objects to and
// from S-expressions.
package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

var Interfaces map[string]reflect.Type

func init() {
	Interfaces = make(map[string]reflect.Type)
}

func RegisterInterface(name string, i reflect.Type) {
	Interfaces[name] = i
}

func checkInterfaceSupport(name string) (ok bool) {
	_, ok = Interfaces[name]
	return
}

// low level API
type Token interface{}
type Symbol string
type String string
type Int int
type StartList struct{}
type EndList struct{}

func (s Symbol) String() string {
	return string(s)
}
func (s String) String() string {
	return string(s)
}
func (s Int) String() string {
	return strconv.FormatInt(int64(s), 10)
}

// Token() returns next Token
func (dec *Decoder) Token() (Token, error) {
	dec.lex.next()
	switch dec.lex.token {
	case scanner.EOF:
		return nil, io.EOF
	case scanner.Ident:
		return Symbol(dec.lex.text()), nil
	case scanner.String:
		s, _ := strconv.Unquote(dec.lex.text())
		return String(s), nil
	case scanner.Int:
		i, _ := strconv.ParseInt(dec.lex.text(), 10, 64)
		return Int(i), nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	}

	return nil, fmt.Errorf("unexpected token %q", dec.lex.text())
}

// Stream API to decode
type Decoder struct {
	reader io.Reader
	lex    lexer
}

func NewDecoder(r io.Reader) *Decoder {
	lex := lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	return &Decoder{reader: r, lex: lex}
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	dec.lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
		}
	}()
	read(&dec.lex, reflect.ValueOf(v).Elem())
	return nil
}

//!+Unmarshal
// Unmarshal parses S-expression data and populates the variable
// whose address is in the non-nil pointer out.
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

//!-Unmarshal

//!+lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

//!-lexer

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
// - that v is always a variable of the appropriate type for the
//   S-expression value.  For example, v must not be a boolean,
//   interface, channel, or function, and if v is an array, the input
//   must have the correct number of elements.
// - that v in the top-level call to read has the zero value of its
//   type and doesn't need clearing.
// - that if v is a numeric variable, it is a signed integer.

//!+read
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		} else if lex.text() == "t" {
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(float64(f))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	case '-': // TODO: support negative value
		lex.next()
		read(lex, v)
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

//!-read

//!+readlist
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, fieldByName(v, name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	case reflect.Interface:
		t, ok := Interfaces[strings.Trim(lex.text(), `"`)]
		if !ok {
			panic(fmt.Sprintf("Not supported type: %v", lex.text()))
		}
		d := reflect.Indirect(reflect.New(t))
		lex.next()
		read(lex, d)
		v.Set(d)

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

func fieldByName(v reflect.Value, name string) reflect.Value {
	if v.Kind() != reflect.Struct {
		panic("fieldByName: not supported type")
	}

	res := v.FieldByName(name)
	if res.Kind() != reflect.Invalid {
		return res
	}

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		customName := fieldInfo.Tag.Get("sexpr")
		if customName == name {
			return v.Field(i)
		}
	}

	return reflect.Value{}
}

//!-readlist
