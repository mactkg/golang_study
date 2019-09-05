package main

import (
	"bytes"
	"fmt"
)

func join(sep string, strings ...string) string {
	buf := bytes.Buffer{}
	for i, v := range strings {
		buf.WriteString(v)
		if i < len(strings)-1 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(join("-", "080", "1234", "4321"))
}
