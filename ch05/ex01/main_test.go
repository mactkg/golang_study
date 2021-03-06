package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const data = `
<html>
<body>
	<h1><a href="/">Test Page</a></h1>
	<div>
		<h2><a href="/blogs/1">Post #1</a></h2>
		<div>
			<p>Hello, world!</p>
			<p><a href="http://golang.org/">Golang Offical Web Page</a></p>
		</div>
	</div>
	<footer><a href="https://contact.me/golang">contact</a></footer>
</body>
</html>
`

func TestVisit(t *testing.T) {
	r := strings.NewReader(data)
	doc, err := html.Parse(r)
	if err != nil {
		t.Fatal(err)
	}

	results := Visit(nil, doc)
	check := 0
	for _, res := range results {
		switch res {
		case "/":
			check |= 1
		case "/blogs/1":
			check |= 2
		case "http://golang.org/":
			check |= 4
		case "https://contact.me/golang":
			check |= 8
		default:
			t.Fatal("Found unexpected link", res)
		}
	}
	if check != 15 {
		t.Fatal("Some links aren't found", results)
	}
}
