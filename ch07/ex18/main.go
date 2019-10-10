// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // *Element or CharData

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	buf := &bytes.Buffer{}
	visit(buf, e, 0)
	return buf.String()
}

func visit(w io.Writer, n Node, depth int) {
	switch n := n.(type) {
	case CharData:
		fmt.Fprintf(w, "%*s%s", depth*2, "", n)
	case *Element:
		fmt.Fprintf(w, "%*s%s\t%s", depth*2, "", n.Type.Local, n.Attr)
		for _, c := range n.Children {
			visit(w, c, depth+1)
		}
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []*Element // stack of element names
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) == 0 {
				root = elem
			} else {
				latest := stack[len(stack)-1]
				latest.Children = append(latest.Children, elem)
			}
			stack = append(stack, elem)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			latest := stack[len(stack)-1]
			latest.Children = append(latest.Children, CharData(tok))
		}
	}
	fmt.Println(root)
}
