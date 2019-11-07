package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type GoDeps struct {
	Deps []string
}

func fetchDeps(target string) (*GoDeps, error) {
	result, err := exec.Command("go", "list", "-json", target).Output()
	if err != nil {
		return nil, err
	}

	var deps GoDeps
	err = json.Unmarshal(result, &deps)
	if err != nil {
		return nil, err
	}

	return &deps, nil
}

func main() {
	var target string
	flag.StringVar(&target, "target", "", "target package")
	flag.Parse()

	// 1st fetch
	result := map[string]struct{}{}
	deps, err := fetchDeps(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error!: %v\n", err)
		os.Exit(1)
	}

	// 2nd fetch loop
	resCh := make(chan GoDeps)
	for _, dep := range deps.Deps {
		go func(target string) {
			res, err := fetchDeps(target)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error happened while fetching"+target+": %v\n", err)
				resCh <- GoDeps{}
			}
			resCh <- *res
		}(dep)
	}

	for i := 0; i < len(deps.Deps); i++ {
		res := <-resCh
		for _, dep := range res.Deps {
			result[dep] = struct{}{}
		}
	}

	// print results
	keys := make([]string, 0, len(result))
	for key := range result {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(strings.Join(keys, "\n"))
}
