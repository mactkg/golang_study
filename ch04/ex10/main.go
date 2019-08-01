package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mactkg/golang_study/ch04/ex10/github"
)

func main() {

	result, err := github.SearchIssues(os.Args[1:])
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
