package main

import "io"

func main() {
	// io.LimitReader
}

type limitReader struct {
	r     io.Reader
	limit int64
}

func (r limitReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p[:r.limit])
	if err != nil {
		return 0, err
	}

	return n, nil
}

func MyLimitReader(r io.Reader, n int64) io.Reader {
	return limitReader{r: r, limit: n}
}
