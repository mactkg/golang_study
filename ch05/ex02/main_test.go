package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const data = `
<html>
<head></head>
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

func wrongCount(t *testing.T, k string, count, shouldBe int) {
	t.Fatalf("Wrong count of <%s>(got %d, but it should be %d)", k, count, shouldBe)
}

func TestOutline(t *testing.T) {
	r := strings.NewReader(data)
	doc, err := html.Parse(r)
	if err != nil {
		t.Fatal(err)
	}

	res := Outline(make(map[string]int), doc)
	for k, count := range res {
		switch k {
		case "html", "body", "head", "h1", "h2", "footer":
			if count != 1 {
				wrongCount(t, k, count, 1)
			}
		case "div", "p":
			if count != 2 {
				wrongCount(t, k, count, 2)
			}
		case "a":
			if count != 4 {
				wrongCount(t, k, count, 4)
			}
		default:
			t.Fatal("Found unexpected node.", k)
		}
	}
}
