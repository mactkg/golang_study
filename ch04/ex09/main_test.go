package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	data := bytes.NewBufferString("now it supports only English sorry sorry!!!")
	result := Wordfreq(data)
	if c := result["sorry"]; c != 2 {
		t.Fatal("wrong 'sorry' count", c)
	}

	if c := result["now"]; c != 1 {
		t.Fatal("wrong 'now' count", c)
	}
}
