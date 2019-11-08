package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"m/myarchive"
)

func init()  {
	myarchive.RegisterType("zip", Unarchive, "a")
}

func Unarchive(r io.Reader) ([]myarchive.File, error) {
	buff := bytes.NewBuffer([]byte{})
	size, err := io.Copy(buff, r)
	if err != nil {
		return nil, err
	}
	br := bytes.NewReader(buff.Bytes())
	reader, err := zip.NewReader(br, size)
	if err != nil {
		return nil, err
	}

	for _, r := range reader.File {
		log.Println(r.Name)
	}

	return nil, nil
}
