package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var jpeg_out, gif_out, png_out bool

func main() {
	flag.BoolVar(&jpeg_out, "jpeg", true, "convert to .jpeg")
	flag.BoolVar(&png_out, "png", false, "convert to .png")
	flag.BoolVar(&gif_out, "gif", false, "convert to .gif")
	flag.Parse()

	buf := &bytes.Buffer{}
	io.Copy(buf, os.Stdin)
	reader := bytes.NewReader(buf.Bytes())

	if jpeg_out {
		f, err := os.OpenFile("./out.jpeg", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not open file: %v", err)
			os.Exit(1)
		}
		defer f.Close()

		if err := toJPEG(buf, f); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	}

	if png_out {
		reader.Seek(0, 0)
		f, err := os.OpenFile("./out.png", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not open file: %v", err)
			os.Exit(1)
		}
		defer f.Close()

		if err := toPNG(reader, f); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	}

	if gif_out {
		reader.Seek(0, 0)
		f, err := os.OpenFile("./out.gif", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not open file: %v", err)
			os.Exit(1)
		}
		defer f.Close()

		if err := toGIF(reader, f); err != nil {
			fmt.Fprintf(os.Stderr, "gif: %v\n", err)
			os.Exit(1)
		}
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toGIF(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	return gif.Encode(out, img, &gif.Options{NumColors: 256})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	return png.Encode(out, img)
}
