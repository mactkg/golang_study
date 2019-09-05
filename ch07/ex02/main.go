package main

import "io"

func main() {

}

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	_, err := c.w.Write(p)
	if err != nil {
		return 0, err
	}

	c.count += int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &ByteCounter{w, 0}
	return c, &c.count
}
