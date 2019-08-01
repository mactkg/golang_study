package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		printError()
		return
	}

	v := os.Args[len(os.Args)-1]
	hashtype := "sha256"

	if t := os.Args[1]; t == "sha256" || t == "sha384" || t == "sha512" {
		hashtype = t
	}

	switch hashtype {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(v)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(v)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(v)))
	}
}

func printError() {
	fmt.Print("This command require at least one argument.\n" +
		"Usage: go run main.go [<hashtype>] <input>\n\n" +
		"hashtype: sha256, sha284, sha512")
}
