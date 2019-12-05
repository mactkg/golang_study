package bzip

import (
	"fmt"
	"io"
	"os/exec"
)

type writer struct {
	cmd       *exec.Cmd
	cmdWriter io.WriteCloser
}

func NewWriter(out io.Writer) (io.WriteCloser, error) {
	cmd := exec.Command("/usr/bin/bzip2")
	cmd.Stdout = out
	cmdWriter, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	fmt.Println(cmd)

	w := &writer{cmd: cmd, cmdWriter: cmdWriter}
	return w, nil
}

func (w *writer) Write(data []byte) (int, error) {
	return w.cmdWriter.Write(data)
}

func (w *writer) Close() error {
	err := w.cmdWriter.Close()
	w.cmd.Process.Wait() // wait for flushing stdout buffer
	if err != nil {
		return err
	}
	return nil
}
