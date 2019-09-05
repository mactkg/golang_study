package main

import (
	"fmt"
	"io"
	"log"
	"m/links"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

func save(raw string) error {
	url, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("URL Parse Error")
	}
	dir := url.Host
	var filename string
	if filepath.Ext(url.Path) == "" {
		// get index page
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		// fetch index page
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = filepath.Join(url.Host, url.Path)
	}
	fmt.Printf("%v %v\n", dir, filename)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("Failed make directory: %v", dir)
	}

	res, err := http.Get(raw)
	if err != nil {
		return fmt.Errorf("Fetch failed: %v", raw)
	}
	defer res.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed create file: %v", filename)
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("Failed writing data")
	}
	file.Close()

	return nil
}

//!+main
func main() {
	base, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hostname := base.Scheme + "://" + base.Hostname()

	crawl := func(url string) []string {
		if strings.HasPrefix(url, hostname) {
			fmt.Println(url)
			err := save(url)
			if err != nil {
				fmt.Println(err)
			}
		}

		list, err := links.Extract(url)
		if err != nil {
			log.Print(err)
		}
		return list
	}

	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
