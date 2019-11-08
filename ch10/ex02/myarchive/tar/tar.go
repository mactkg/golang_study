package myarchive

import (
	"archive/tar"
	"bytes"
	"io"
	"m/myarchive"
)

func init()  {
	myarchive.RegisterType("tar", Unarchive, "a")
}

func Unarchive(r io.Reader) ([]myarchive.File, error) {
	var files []myarchive.File
	reader := tar.NewReader(r)
	for {
		hdr, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// load file
		buf := bytes.Buffer{}
		io.Copy(&buf, reader)
		f := myarchive.File{
			Body: buf,
			Name: hdr.Name,
			Size: hdr.Size,
			ModTime: hdr.ModTime,
		}
		files = append(files, f)
	}

	return files, nil
}
