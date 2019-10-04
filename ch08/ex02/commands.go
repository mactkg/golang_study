package main

import (
	"fmt"
	"net"
)

type connection struct {
	user net.Conn
	data net.Conn
}

func (c connection) cd(dest string) (pwd string, err error) {
	return "", nil
}

func (_ connection) ls() (list []string, err error) {
	return []string{}, nil
}

func (c connection) get(files []string) (err error) {
	// c.data.Write() // write soemthing
	return nil
}

func (c connection) close() (err error) {
	err_user := c.user.Close()
	err_data := c.data.Close()

	if err_user != nil || err_data != nil {
		err = fmt.Errorf("Connection closing error. user: %v, data: %v", err_user, err_data)
	}
	return
}
