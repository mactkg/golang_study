package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	data := bytes.NewBufferString("123ABCハロー🤪")
	_, types, _, _ := CharCount(data)
	if c := types["digit"]; c != 3 {
		t.Fatal("wrong digit count", c)
	}
	if c := types["latter"]; c != 6 {
		t.Fatal("wrong latter count", c)
	}
	if c := types["graphic"]; c != 1 {
		t.Fatal("wrong graphic count", c)
	}
}
