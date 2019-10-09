// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"sync"
)

type page struct {
	url	string
	depth int
}

func crawl(p page) []string {
	fmt.Printf("[%d] %s\n", p.depth, p.url)
	list, err := links.Extract(p.url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	var depth int
	flag.IntVar(&depth, "depth", 3, "Depth")
	flag.Parse()

	wg := sync.WaitGroup{}

	worklist := make(chan []page)  // lists of URLs, may have duplicates
	unseenPages := make(chan page) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		pages := []page{}
		for _, url := range flag.Args() {
			pages = append(pages, page{url, 0})
		}
		worklist <- pages
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			wg.Add(1)
			for p := range unseenPages {

				foundLinks := crawl(p)

				if p.depth >= depth {
					break
				}

				go func(depth int) {
					pages := []page{}
					for _, link := range foundLinks {
						pages = append(pages, page{link, depth})
					}
					worklist <- pages
				}(p.depth+1)
			}
			wg.Done()
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	go func() {
		for list := range worklist {
			for _, p := range list {
				if !seen[p.url] {
					seen[p.url] = true
					unseenPages <- p
				}
			}
		}
	}()
	wg.Wait()
}

//!-
