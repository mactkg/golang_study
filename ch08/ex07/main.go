package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type page struct {
	url	string
	depth int
}

var workerToken = make(chan struct{}, 20)
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}
var wg = sync.WaitGroup{}

func crawl(p page, maxDepth int) {
	defer wg.Done()

	log.Printf("[%d] %s\n", p.depth, p.url)

	workerToken <- struct{}{}
	urls, err := VisitAndSave(p)
	<-workerToken
	if err != nil {
		log.Print(err)
	}

	// skip if too deep to dig
	if p.depth + 1 > maxDepth {
		return
	}

	base, err := url.Parse(p.url)
	for _, u := range urls {
		parsed, err := url.Parse(u)
		if err != nil {
			log.Printf("href parse error: %v\n", err)
			continue
		}

		if parsed.Hostname() != base.Hostname() {
			//log.Printf("Skip %v\n", parsed)
			continue
		}

		// seen / unseen check
		seenLock.Lock()
		if !seen[u] {
			seen[u] = true
			nextPage := page{url: u, depth: p.depth+1}
			wg.Add(1)
			go crawl(nextPage, maxDepth)
		}
		seenLock.Unlock()
	}
}

func Save(url url.URL, body io.Reader) error {
	// where to save
	dir := url.Host
	var filename string
	if filepath.Ext(url.Path) == "" {
		// get index page
		dir = filepath.Join("./data", dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		// fetch index page
		filename = filepath.Join("./data", dir, url.Path)
		dir = filepath.Dir(filename)
	}

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("Failed make directory: %v", dir)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed create file: %v", err)
	}


	_, err = io.Copy(file, body)
	if err != nil {
		return fmt.Errorf("Failed writing data")
	}
	file.Close()


	return nil
}

func VisitAndSave(p page) ([]string, error) {
	// fetch and parse
	resp, err := http.Get(p.url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", p.url, resp.Status)
	}

	// find links and edit
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link" || n.Data == "script" || n.Data == "img") {
			for i, a := range n.Attr {
				if a.Key != "href" && a.Key != "src" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())

				host := link.Host
				link.Scheme = ""
				link.Host = ""
				link.User = nil
				a.Val = "/" + host + link.String()
				//log.Printf("ReWrite: %v\n", a.Val)
				n.Attr[i] = a
			}
		}
	}

	buf := &bytes.Buffer{}
	if strings.HasPrefix(resp.Header.Get("content-type"), "text/html") {
		// Parse HTML
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", p.url, err)
		}

		// Edit links
		forEachNode(doc, visitNode, nil)
		err = html.Render(buf, doc)
		if err != nil {
			return nil, fmt.Errorf("rendering html Node: %v", err)
		}
	} else {
		io.Copy(buf, resp.Body)
		resp.Body.Close()
	}

	url, _ := url.Parse(p.url)
	err = Save(*url, buf)
	if err != nil {
		return nil, fmt.Errorf("saving %s to file: %v", p.url, err)
	}
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//!+
func main() {
	var depth int
	flag.IntVar(&depth, "depth", 3, "Depth")
	flag.Parse()

	for _, url := range flag.Args() {
		p := page{url, 0}
		wg.Add(1)
		go crawl(p, depth)
	}

	wg.Wait()
}

//!-
