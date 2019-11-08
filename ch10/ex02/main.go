package main

import (
	"fmt"
	"log"
	"m/myarchive"
	_ "m/myarchive/tar"
	_ "m/myarchive/zip"
	"os"
)

func main() {
	fmt.Println("Registered types:", myarchive.RegisteredTypes())
	files, err := myarchive.Unarchive(os.Stdin)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, f := range files {
		log.Println(f.Name)
	}
}
