package main

import (
	"bufio"
	"log"
	"net"
	"strings"
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
			err := c.loginRequired()
			if err != nil {
				continue
			}
			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}

			err = c.get(tokens[1])
			if err != nil {
				c.replyParseParamsError()
				log.Println(err)
				continue
			}

		case "STOR":
			err := c.loginRequired()
			if err != nil {
				continue
			}
			if len(tokens) != 2 {
				c.replyInvalidParamsError()
				continue
			}

			err = c.put(tokens[1])
			if err != nil {
				c.replyParseParamsError()
				log.Println(err)
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
