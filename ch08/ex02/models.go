package main

import (
	"fmt"
	"net"
)

type FTPStructure int

const (
	FILE FTPStructure = iota
	RECORD
)

type FTPMode int

const (
	STREAM FTPMode = iota
)

type FTPConnection struct {
	loggedIn  bool
	user      string
	cwd       string
	structure FTPStructure
	mode      FTPMode
	dataAddr  string
	dataConn  net.Conn
	net.Conn
}

func (c FTPConnection) loginRequired() error {
	if !c.loggedIn {
		c.loginRequired()
		return fmt.Errorf("Not logged in.")
	}
	return nil
}
