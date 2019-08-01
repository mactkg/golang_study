package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/mactkg/golang_study/ch04/ex11/github"
)

func main() {
	if len(os.Args) < 3 {
		help()
		return
	}

	method := os.Args[1]
	switch method {
	case "get":
		get(os.Args[2:])
	case "list":
		list(os.Args[2:])
	case "close":
		close(os.Args[2:])
	case "create":
		create(os.Args[2:])
	case "edit":
		edit(os.Args[2:])
	case "search":
		search(os.Args[2:])
	default:
		help()
	}
}

func help() {
	fmt.Printf("Usage: <command> [<owner>] [<repo>] [<issue-number>] [<query>]\n" +
		"commands:\n" +
		"\tsearch(query)\n" +
		"\tlist(owner, repo)\n" +
		"\tcreate(owner, repo)\n" +
		"\tget(owner, repo, issue-number)\n" +
		"\tedit(owner, repo, issue-number)\n" +
		"\tclose(owner, repo, issue-number)\n")
}

func printIssue(issue *github.Issue) {
	fmt.Printf("[#%d] %s by %s (%s)\n", issue.Number, issue.Title, issue.User, issue.State)
	fmt.Printf("%s\n", issue.Body)
	fmt.Printf("Link: %s\n", issue.HTMLURL)
}

func get(args []string) {
	if len(args) < 3 {
		help()
		return
	}

	owner, repo := args[0], args[1]
	var number int64
	if n := args[2]; n != "" {
		number, _ = strconv.ParseInt(n, 10, 64)
	}
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	printIssue(issue)
}

func list(args []string) {
	if len(args) < 2 {
		help()
		return
	}

	owner, repo := args[0], args[1]
	issues, err := github.ListIssue(owner, repo)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	printIssues("Issues in "+owner+"/"+repo, issues)
}

func close(args []string) {
	if len(args) < 3 {
		help()
		return
	}

	owner, repo := args[0], args[1]
	var number int64
	if n := args[2]; n != "" {
		number, _ = strconv.ParseInt(n, 10, 64)
	}
	issue, err := github.CloseIssue(owner, repo, number)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	printIssue(issue)
}

func create(args []string) {
	if len(args) < 2 {
		help()
		return
	}
	owner, repo := args[0], args[1]

	// launch editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	path, err := exec.LookPath(editor)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	tmp, err := ioutil.TempFile("", "github_issue")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	cmd := &exec.Cmd{
		Path:   path,
		Args:   []string{editor, tmp.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	// load changes
	tmp.Seek(0, 0)
	data, err := ioutil.ReadAll(tmp)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	lines := strings.Split(string(data), "\n")
	var title, body string
	if t := lines[0]; t != "" {
		title = t
	}
	if b := lines[1]; b != "" {
		body = b
	}

	// send it
	fields := map[string]string{
		"title": title,
		"body":  body,
	}
	fmt.Println(fields)
	issue, err := github.CreateIssue(owner, repo, fields)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	printIssue(issue)
}

func edit(args []string) {
	// fetch issue to edit
	owner, repo := args[0], args[1]
	var number int64
	if n := args[2]; n != "" {
		number, _ = strconv.ParseInt(n, 10, 64)
	}
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	// launch editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	path, err := exec.LookPath(editor)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	tmp, err := ioutil.TempFile("", "github_issue")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	tmp.WriteString(issue.Title + "\n" + issue.Body)
	cmd := &exec.Cmd{
		Path:   path,
		Args:   []string{editor, tmp.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	// load changes
	tmp.Seek(0, 0)
	data, err := ioutil.ReadAll(tmp)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	lines := strings.Split(string(data), "\n")
	var title, body string
	if t := lines[0]; t != "" {
		title = t
	}
	if b := lines[1]; b != "" {
		body = b
	}

	// send it
	fields := map[string]string{
		"title": title,
		"body":  body,
	}
	edited, err := github.EditIssue(owner, repo, number, fields)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	printIssue(edited)
}

func search(args []string) {
	result, err := github.SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()

	fmt.Printf("%d issues are found:\n", result.TotalCount)

	thisMonth := make([]*github.Issue, 0)
	thisYear := make([]*github.Issue, 0)
	pastYear := make([]*github.Issue, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(now.AddDate(0, -1, 0)):
			thisMonth = append(thisMonth, item)
		case item.CreatedAt.After(now.AddDate(-1, 0, 0)):
			thisYear = append(thisYear, item)
		default:
			pastYear = append(pastYear, item)
		}
	}

	printIssues("This Month", thisMonth)
	printIssues("This Year", thisYear)
	printIssues("Past Year", pastYear)
}

func printIssues(title string, issues []*github.Issue) {
	fmt.Printf("\n%s\n", title)
	fmt.Printf("%d issues:\n", len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
