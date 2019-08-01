package main

import "testing"

func TestCheckAnagram(t *testing.T) {
	if CheckAnagram("BAC", "CAB") != true {
		t.Fatal("Error: BAC / CAB")
	}

	if CheckAnagram("AAAAAA", "AAAAAA") != true {
		t.Fatal("Error: AAAAAA / AAAAAA")
	}

	if CheckAnagram("ABC", "AAA") != false {
		t.Fatal("Error: ABC / AAA")
	}
}
