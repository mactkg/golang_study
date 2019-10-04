package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
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

func sendWelcome(c net.Conn) {
	fmt.Fprintln(c, "220 Service ready for new user.")
}

func replyInvalidParamsError(c net.Conn) {
	fmt.Fprintln(c, "501 Syntax error in parameters or arguments.")
}

func replyInvalidActionError(c net.Conn) {
	fmt.Fprintln(c, "502 Command not implemented.")
}

func replyParseParamsError(c net.Conn) {
	fmt.Fprintln(c, "504 Command not implemented for that parameter.")
}

func parsePort(in string) (addr string, err error) {
	// split addr and port
	// TODO: we have to check range of ip or port
	//       especially > 0 check
	r := regexp.MustCompile(`(\d+),(\d+),(\d+),(\d+),(\d+),(\d+)`)
	res := r.FindStringSubmatch(in)
	if len(res) != 7 {
		err = fmt.Errorf("Invalid format")
		return
	}

	ip, p := res[1:5], res[6:7]
	port_high, err := strconv.ParseInt(p[0], 0, 64)
	if err != nil {
		return
	}
	port_low, err := strconv.ParseInt(p[1], 0, 64)
	if err != nil {
		return
	}
	port := port_high*256 + port_low

	addr = fmt.Sprintf("%s.%s.%s.%s:%d", ip[0], ip[1], ip[2], ip[3], port)
	return
}

func handleConn(c net.Conn) {
	defer c.Close()

	sendWelcome(c)
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, " ")
		log.Println(tokens, len(tokens))

		switch strings.ToUpper(tokens[0]) {
		case "USER":
			if len(tokens) != 2 {
				replyInvalidParamsError(c)
				continue
			}
			user := tokens[1]
			log.Printf("Login user '%s'\n", user)
		case "QUIT":
			if len(tokens) != 1 {
				replyInvalidParamsError(c)
				continue
			}
			log.Printf("Bye~\n")
			return
		case "PORT":
			if len(tokens) != 2 {
				replyInvalidParamsError(c)
				continue
			}
			addr, err := parsePort(tokens[1])
			if err != nil {
				replyParseParamsError(c)
				log.Println(tokens[1])
				continue
			}
			log.Printf("Create data connection to %s\n", addr)
		case "TYPE":
			if len(tokens) != 2 {
				replyInvalidParamsError(c)
				continue
			}
		case "MODE":
			if len(tokens) != 2 {
				replyInvalidParamsError(c)
				continue
			}
		case "STRU":
			if len(tokens) != 2 {
				replyInvalidParamsError(c)
				continue
			}
		case "RETR":
			if len(tokens) != 2 {
				replyInvalidParamsError(c)
				continue
			}
		case "STOR":
			if len(tokens) != 2 {
				replyInvalidParamsError(c)
				continue
			}
		case "NOOP":
			if len(tokens) != 1 {
				replyInvalidParamsError(c)
				continue
			}
		default:
			replyInvalidActionError(c)
		}
	}
}
