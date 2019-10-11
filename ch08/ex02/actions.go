package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (c FTPConnection) ls(path string) error {
	if !filepath.IsAbs(path) {
		p, err := filepath.Abs(filepath.Join(c.cwd, path))
		path = p
		if err != nil {
			c.replyRequestedActionNotTaken()
			return err
		}
	}
	fmt.Println("ls: " + path)

	// opening data connection
	if c.dataAddr == "" {
		c.replyCantOpenDataConn()
		return fmt.Errorf("User should be set data address")
	}
	dataConn, err := net.DialTimeout("tcp", c.dataAddr, 3*time.Second)
	if err != nil {
		c.replyCantOpenDataConn()
		return fmt.Errorf("Can't open connection: %v", c.dataAddr)
	}
	defer func() {
		c.replyClosingDataConn()
		dataConn.Close()
	}()
	c.replyOpeningDataConn()

	// ls
	filenames, err := filepath.Glob(filepath.Join(path, "/*"))
	if err != nil {
		c.replyLocalError()
		return fmt.Errorf("Local Error: (%v)", err)
	}

	_, err = dataConn.Write([]byte(strings.Join(filenames, "\t") + "\r\n"))
	if err != nil {
		c.replyTransferAborted()
	}

	return nil
}

func (c FTPConnection) get(path string) error {
	if !filepath.IsAbs(path) {
		p, err := filepath.Abs(filepath.Join(c.cwd, path))
		path = p
		if err != nil {
			c.replyRequestedActionNotTaken()
			return err
		}
	}

	// opening file
	f, err := os.Open(path)
	if err != nil {
		c.replyRequestedActionNotTaken()
		return fmt.Errorf("Can't load %v (%v)", path, err)
	}
	defer f.Close()

	// opening data connection
	if c.dataAddr == "" {
		c.replyCantOpenDataConn()
		return fmt.Errorf("User should be set data address")
	}
	dataConn, err := net.DialTimeout("tcp", c.dataAddr, 3*time.Second)
	if err != nil {
		c.replyCantOpenDataConn()
		return fmt.Errorf("Can't open connection: %v", c.dataAddr)
	}
	defer func() {
		c.replyClosingDataConn()
		dataConn.Close()
	}()
	c.replyOpeningDataConn()

	// send
	_, err = io.Copy(dataConn, f)
	if err != nil {
		c.replyTransferAborted()
	}

	_, err = dataConn.Write([]byte("\r\n")) // TODO: find more great way to tell EOF
	if err != nil {
		c.replyTransferAborted()
	}
	return nil
}

func (c FTPConnection) put(path string) error {
	if !filepath.IsAbs(path) {
		p, err := filepath.Abs(filepath.Join(c.cwd, path))
		path = p
		if err != nil {
			c.replyRequestedActionNotTaken()
			return err
		}
	}


	// opening file
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.replyRequestedActionNotTaken()
		return fmt.Errorf("Can't load %v (%v)", path, err)
	}
	defer f.Close()

	// opening data connection
	if c.dataAddr == "" {
		c.replyCantOpenDataConn()
		return fmt.Errorf("User should be set data address")
	}
	dataConn, err := net.DialTimeout("tcp", c.dataAddr, 3*time.Second)
	if err != nil {
		c.replyCantOpenDataConn()
		return fmt.Errorf("Can't open connection: %v", c.dataAddr)
	}
	defer func() {
		c.replyClosingDataConn()
		dataConn.Close()
	}()
	c.replyOpeningDataConn()

	// send
	_, err = io.Copy(f, dataConn)
	if err != nil {
		c.replyTransferAborted()
	}
	return nil
}

func (c *FTPConnection) cd(path string) error {
	nwd, err := filepath.Abs(filepath.Join(c.cwd, path))
	if err != nil {
		c.replyRequestedActionNotTaken()
		return err
	}

	check, err := exists(nwd)
	if err != nil {
		c.replyRequestedActionNotTaken()
		return err
	} else if !check {
		c.replyRequestedActionNotTaken()
		return fmt.Errorf("%v not found", nwd)
	}

	c.cwd = nwd
	c.replyCompleted()
	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func (c *FTPConnection) parsePort(in string) (err error) {
	// split addr and port
	// TODO: we have to check range of ip or port
	//       especially > 0 check
	r := regexp.MustCompile(`(\d+),(\d+),(\d+),(\d+),(\d+),(\d+)`)
	res := r.FindStringSubmatch(in)
	if len(res) != 7 {
		err = fmt.Errorf("Invalid format")
		return
	}

	ip, p := res[1:5], res[5:7]
	port_high, err := strconv.ParseInt(p[0], 0, 64)
	if err != nil {
		return
	}
	port_low, err := strconv.ParseInt(p[1], 0, 64)
	if err != nil {
		return
	}
	port := port_high*256 + port_low

	c.dataAddr = fmt.Sprintf("%s.%s.%s.%s:%d", ip[0], ip[1], ip[2], ip[3], port)
	return
}

func (c FTPConnection) parseType(t, f string) (err error) {
	// only supported ascii non print
	switch strings.ToUpper(t) {
	case "A":
		switch strings.ToUpper(f) {
		case "N":
			return nil
		}
	}

	return fmt.Errorf("Wrong type: %v, %v", t, f)
}

func (c *FTPConnection) parseMode(mode string) (err error) {
	// only supported stream
	switch strings.ToUpper(mode) {
	case "S":
		c.mode = STREAM
		return nil
	}

	return fmt.Errorf("Wrong mode: %v", mode)
}

func (c *FTPConnection) parseStructure(structure string) (err error) {
	switch strings.ToUpper(structure) {
	case "F":
		c.structure = FILE
	case "R":
		c.structure = RECORD
	default:
		return fmt.Errorf("Wrong structure: %v", structure)
	}

	return nil
}
