package main

import "testing"

func TestComma(t *testing.T) {
	if str := Comma("1234567"); str != "1,234,567" {
		t.Fatalf("Error. %s", str)
	}
	if str := Comma("12345678"); str != "12,345,678" {
		t.Fatalf("Error. %s", str)
	}
	if str := Comma("123456789"); str != "123,456,789" {
		t.Fatalf("Error. %s", str)
	}
}

func TestCommaWithMiniNumber(t *testing.T) {
	if str := Comma("123"); str != "123" {
		t.Fatalf("Comma should return same number when an argument less than 3 digits: %s", str)
	}
}
