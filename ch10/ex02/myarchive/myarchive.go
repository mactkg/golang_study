package myarchive

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"
)

type Type struct {
	name      string
	unarchive func(rd io.Reader) ([]File, error)
	header    string
}
func (t Type) String() string {
	return t.name + " header: " + t.header
}

type File struct {
	Body bytes.Buffer
	Name string
	Size int64
	ModTime time.Time
}

var archiveTypes []Type
func init() {
	archiveTypes = []Type{}
}

func RegisterType(name string, unarchive func(rd io.Reader) ([]File, error), header string) {
	archiveTypes = append(archiveTypes, Type{name, unarchive, header})
}

func Unarchive(r io.Reader) ([]File, error) {
	buf := &bytes.Buffer{}
	io.Copy(buf, r)
	reader := bytes.NewReader(buf.Bytes())

	for _, v := range archiveTypes {
		reader.Seek(0, 0)
		log.Println(v.name)
		files, err := v.unarchive(reader)
		if err != nil{
			log.Println(err)
			continue
		}
		return files, nil
	}

	return nil, fmt.Errorf("suitable archive type not found")
}

func RegisteredTypes() string {
	return fmt.Sprintf("%v", archiveTypes)
}
