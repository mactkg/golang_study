package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const data = `
<html>
html
<head>
	<link rel="stylesheet" href="style.css">
</head>
<body>
	body
	<h1><a href="/">Test Page</a></h1>
	<div>
		<h2><a href="/blogs/1">Post #1</a></h2>
		<div>
			<p>Hello, world!</p>
			<p><a href="http://golang.org/">Golang Official Web Page</a></p>
		</div>
	</div>
	<footer><a href="https://contact.me/golang">contact</a></footer>

	<script type="text/javascript">
		console.log("hello")
	</script>
</body>
</html>
`

func TestGetTextContent(t *testing.T) {
	r := strings.NewReader(data)
	doc, err := html.Parse(r)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"html",
		"body",
		"Test Page",
		"Post #1",
		"Hello, world!",
		"Golang Official Web Page",
		"contact",
	}
	res := GetTextContent(nil, doc)

	for i, v := range res {
		fmt.Printf("%v\t%v\n", i, v)
	}

	for i, v := range res {
		if v != expected[i] {
			t.Fatalf("expected %q, but got %q", expected[i], v)
		}
	}

}
