package main

import "regexp"

func main() {
	return
}

func Expand(s string, f func(string) string) string {
	var r = regexp.MustCompile(`\$\w+`)
	return r.ReplaceAllStringFunc(s, func(in string) string {
		return f(in[1:])
	})
}
