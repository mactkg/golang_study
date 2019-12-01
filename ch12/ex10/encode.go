// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{writer: w}
}

func (enc *Encoder) Encode(v interface{}) (err error) {
	if err := encode(enc.writer, reflect.ValueOf(v), 0); err != nil {
		return err
	}
	return nil
}

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}

	if err := encode(buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf io.Writer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.Write([]byte("nil"))

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "t")
		} else {
			fmt.Fprint(buf, "nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex128, reflect.Complex64:
		comp := v.Complex()
		fmt.Fprintf(buf, "#C(%.1f, %.1f)", real(comp), imag(comp))

	case reflect.Array, reflect.Slice: // (value ...)
		buf.Write([]byte{'('})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, " ")
			} else {
				indent += 1
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}
		buf.Write([]byte{')'})

	case reflect.Struct: // ((name value) ...)
		buf.Write([]byte{'('})
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).IsZero() {
				continue
			}

			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, " ")
			}
			indent += len(v.Type().Field(i).Name) + 3 // '(' + ' '

			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent); err != nil {
				return err
			}
			buf.Write([]byte{')'})

			indent -= len(v.Type().Field(i).Name) + 3
		}
		buf.Write([]byte{')'})

	case reflect.Map: // ((key value) ...)
		buf.Write([]byte{'('})
		indent += 1
		for i, key := range v.MapKeys() {
			if v.MapIndex(key).IsZero() {
				continue
			}

			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, " ")
			}
			buf.Write([]byte{'('})
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.Write([]byte{' '})
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.Write([]byte{')'})
		}
		buf.Write([]byte{')'})

	case reflect.Interface:
		if v.IsNil() {
			buf.Write([]byte("nil"))
			return nil
		}

		fmt.Fprintf(buf, "(\"%s\" ", v.Elem().Type())
		encode(buf, v.Elem(), indent)
		buf.Write([]byte{')'})

	default: // chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
