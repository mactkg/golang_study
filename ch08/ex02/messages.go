package main

import "fmt"

// 150
func (c FTPConnection) replyOpeningDataConn() {
	fmt.Fprint(c, "150 File status okay; about to open data connection.\r\n")
}

// 200
func (c FTPConnection) replyOkay() {
	fmt.Fprint(c, "200 Command okay.\r\n")
}

// 220
func (c FTPConnection) sendWelcome() {
	fmt.Fprint(c, "220 Service ready for new user.\r\n")
}

// 221
func (c FTPConnection) replyClosingConn() {
	fmt.Fprint(c, "221 Service closing control connection.\r\n")
}

// 226
func (c FTPConnection) replyClosingDataConn() {
	fmt.Fprint(c, "226 Closing data connection.\r\n")
}

// 230
func (c FTPConnection) replyLoggedIn() {
	fmt.Fprint(c, "230 User logged in, proceed.\r\n")
}

// 425
func (c FTPConnection) replyCantOpenDataConn() {
	fmt.Fprint(c, "425 Can't open data connection.\r\n")
}

// 426
func (c FTPConnection) replyTransferAborted() {
	fmt.Fprint(c, "426 Connection closed; transfer aborted.\r\n")
}

// 451
func (c FTPConnection) replyLocalError() {
	fmt.Fprint(c, "451 Requested action aborted. Local error in processing.\r\n")
}

// 501
func (c FTPConnection) replyInvalidParamsError() {
	fmt.Fprint(c, "501 Syntax error in parameters or arguments.\r\n")
}

// 502
func (c FTPConnection) replyInvalidActionError() {
	fmt.Fprint(c, "502 Command not implemented.\r\n")
}

// 504
func (c FTPConnection) replyParseParamsError() {
	fmt.Fprint(c, "504 Command not implemented for that parameter.\r\n")
}

// 530
func (c FTPConnection) replyNotLoggedIn() {
	fmt.Fprint(c, "530 Not logged in.\r\n")
}

// 550
func (c FTPConnection) replyRequestedActionNotTaken() {
	fmt.Fprintf(c, "550 Requested action not taken.\r\n")
}
