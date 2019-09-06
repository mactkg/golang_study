package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	buf := bytes.NewBufferString("")
	var visit func(*tree)
	var depth = 0
	visit = func(t *tree) {
		fmt.Fprintf(buf, "%*s%v\n", depth*2, "", t.value)
		if t.left != nil {
			depth++
			fmt.Fprintln(buf, "left")
			visit(t.left)
			depth--
		}
		if t.right != nil {
			depth++
			fmt.Fprintln(buf, "right")
			visit(t.right)
			depth--
		}
	}
	visit(t)

	return buf.String()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
}
