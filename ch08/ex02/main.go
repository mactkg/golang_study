package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
5.1. 最小の実装
不要なエラーメッセージなしに FTP を動作させるために、全てのサーバーに最低限以下の実装が必須である：
		型 - ASCII Non-print
		モード - ストリーム
		構造 - ファイル、レコード
		コマンド - USER, QUIT, PORT,
			TYPE, MODE, STRU,(デフォルト値のためのもの)
			RETR, STOR,
			NOOP.
転送パラメータのデフォルト値：
		TYPE - ASCII Non-print
		MODE - ストリーム
		STRU - ファイル
全てのホストは標準的なデフォルトとして上記の値を受け入れなければならない。
http://srgia.com/docs/rfc959j.html
https://scrapbox.io/mactkg-pub/Go%E3%81%A7FTP%E3%82%B5%E3%83%BC%E3%83%90%E6%9B%B8%E3%81%8F
*/

// 8081, 8080
func main() {
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

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
	structure FTPStructure
	mode      FTPMode
	dataAddr  string
	dataConn  net.Conn
	net.Conn
}

// 150
func (c FTPConnection) replyOpeningDataConn() {
	fmt.Fprintln(c, "150 File status okay; about to open data connection.")
}

// 200
func (c FTPConnection) replyOkay() {
	fmt.Fprintln(c, "200 Command okay.")
}

// 220
func (c FTPConnection) sendWelcome() {
	fmt.Fprintln(c, "220 Service ready for new user.")
}

// 221
func (c FTPConnection) replyClosingConn() {
	fmt.Fprintln(c, "221 Service closing control connection.")
}

// 226
func (c FTPConnection) replyClosingDataConn() {
	fmt.Fprintln(c, "226 Closing data connection.")
}

// 230
func (c FTPConnection) replyLoggedIn() {
	fmt.Fprintln(c, "230 User logged in, proceed.")
}

// 425
func (c FTPConnection) replyCantOpenDataConn() {
	fmt.Fprintln(c, "425 Can't open data connection.")
}

// 426
func (c FTPConnection) replyTransferAborted() {
	fmt.Fprintln(c, "426 Connection closed; transfer aborted.")
}

// 451
func (c FTPConnection) replyLocalError() {
	fmt.Fprintln(c, "451 Requested action aborted. Local error in processing.")
}

// 501
func (c FTPConnection) replyInvalidParamsError() {
	fmt.Fprintln(c, "501 Syntax error in parameters or arguments.")
}

// 502
func (c FTPConnection) replyInvalidActionError() {
	fmt.Fprintln(c, "502 Command not implemented.")
}

// 504
func (c FTPConnection) replyParseParamsError() {
	fmt.Fprintln(c, "504 Command not implemented for that parameter.")
}

// 530
func (c FTPConnection) replyNotLoggedIn() {
	fmt.Fprintf(c, "530 Not logged in.")
}

func (c FTPConnection) loginRequired() error {
	if !c.loggedIn {
		c.loginRequired()
		return fmt.Errorf("Not logged in.")
	}
	return nil
}

func (c FTPConnection) ls(path string) error {
	if path == "" {
		path = "."
	}

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
	filenames, err := filepath.Glob(path + "/*")
	if err != nil {
		c.replyLocalError()
		return fmt.Errorf("Local Error: (%v)", err)
	}

	_, err = dataConn.Write([]byte(strings.Join(filenames, "\t") + "\n"))
	if err != nil {
		c.replyTransferAborted()
	}

	return nil
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

// FIXME:
// it's ambigous whether FTPConnection send error message or not.
// error should be returned because control flow of switch-case but
// FTPConnection have a net.Conn, so it's able to reply error message.
func handleConn(conn net.Conn) {
	defer conn.Close()

	c := FTPConnection{
		Conn: conn,
	}

	c.sendWelcome()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, " ")
		log.Println(tokens, len(tokens))

		switch strings.ToUpper(tokens[0]) {
		case "USER":
			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}
			c.user = tokens[1]
			c.loggedIn = true
			log.Printf("Login user '%s'\n", c.user)
			c.replyLoggedIn()
		case "QUIT":
			if len(tokens) != 1 {
				c.replyInvalidParamsError()
				continue
			}
			log.Printf("Bye~\n")
			c.replyClosingConn()
		case "PORT":
			err := c.loginRequired()
			if err != nil {
				continue
			}

			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}
			err = c.parsePort(tokens[1])
			if err != nil {
				c.replyParseParamsError()
				log.Println(tokens[1])
				continue
			}
			log.Printf("Create data connection to %s\n", c.dataAddr)
			c.replyOkay()
		case "TYPE":
			err := c.loginRequired()
			if err != nil {
				continue
			}

			// only supported ascii non print, so
			// length of tokens should be 3 ("TYPE A N")
			if len(tokens) != 3 {
				c.replyInvalidParamsError()
				continue
			}
			err = c.parseType(tokens[1], tokens[2])
			if err != nil {
				c.replyParseParamsError()
				log.Println(err)
				continue
			}
			c.replyOkay()
		case "MODE":
			err := c.loginRequired()
			if err != nil {
				continue
			}

			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}
			err = c.parseMode(tokens[1])
			if err != nil {
				c.replyParseParamsError()
				log.Println(err)
				continue
			}
			c.replyOkay()
		case "STRU":
			err := c.loginRequired()
			if err != nil {
				continue
			}

			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}
			err = c.parseStructure(tokens[1])
			if err != nil {
				c.replyParseParamsError()
				log.Println(err)
				continue
			}
			c.replyOkay()
		case "RETR":
			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}
		case "STOR":
			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}
		case "NOOP":
			if len(tokens) != 1 {
				c.replyInvalidParamsError()
				continue
			}
			c.replyOkay()
		case "LIST":
			err := c.loginRequired()
			if err != nil {
				continue
			}
			switch len(tokens) {
			case 1:
				c.ls("")
			case 2:
				c.ls(tokens[1])
			default:
				c.replyInvalidParamsError()
				continue
			}
		case "PWD":
			err := c.loginRequired()
			if err != nil {
				continue
			}
			if len(tokens) != 1 {
				c.replyInvalidParamsError()
				continue
			}
			c.ls("")
		default:
			c.replyInvalidActionError()
		}

	}
}
