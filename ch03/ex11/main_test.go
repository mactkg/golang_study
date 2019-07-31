package main

import "testing"

func TestComma(t *testing.T) {
	if str := Comma("1234567"); str != "1,234,567" {
		t.Fatalf("Error. %s", str)
	}
}

func TestCommaWithMiniNumber(t *testing.T) {
	if str := Comma("123"); str != "123" {
		t.Fatalf("Comma should return same number when an argument less than 3 digits: %s", str)
	}
}

func TestCommaWithFloatNumber(t *testing.T) {
	if str := Comma("1234567.89"); str != "1,234,567.89" {
		t.Fatalf("Comma should return same number when an argument less than 3 digits: %s", str)
	}
}
