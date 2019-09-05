package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	"linear algebra": {"calculus"}, // added
}

func main() {
	res, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, course := range res {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string, string) error

	visitAll = func(items []string, from string) error {
		for _, item := range items {
			if item == from {
				return fmt.Errorf("Error: Circular reference happend between %v and %v", item, from)
			}

			if !seen[item] {
				seen[item] = true

				if from == "" {
					from = items[0]
				}

				err := visitAll(m[item], from)
				if err != nil {
					return err
				}
				order = append(order, item)
			}
		}

		return nil
	}

	for k, _ := range m {
		err := visitAll([]string{k}, "")
		if err != nil {
			return nil, err
		}
	}
	return order, nil
}
