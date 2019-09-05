package main

import "testing"

func TestFetch(t *testing.T) {
	resO, nO, _ := fetchOld("http://gopl.io")
	resN, nN, _ := fetchNew("http://gopl.io")

	if resO != resN {
		t.Fatalf("Responces are different: %v, %v", resO, resN)
	}

	if nO != nN {
		t.Fatalf("Fetched bytes are different: %v, %v", nO, nN)
	}
}
