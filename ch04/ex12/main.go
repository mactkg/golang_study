package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const APIURL = "https://xkcd.com"
const MAX_ID = 3000

type Comic struct {
	Num        int
	Transcript string
	Img        string
	Title      string `json:safe_title`
}

func getJSONPath(id int) string {
	return "./data/" + strconv.FormatInt(int64(id), 10) + ".json"
}

func getIndexJSONPath() string {
	return "./data/index.json"
}

func fetchComic(id int) (comic *Comic, err error) {
	if _, err := os.Stat(getJSONPath(id)); err == nil {
		return load(id)
	}

	url := fmt.Sprintf("%s/%d/info.0.json", APIURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("Error: %d", resp.StatusCode)
		return nil, err
	}

	var result Comic
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	save(&result)

	time.Sleep(200 * time.Millisecond)
	return &result, nil
}

func load(id int) (*Comic, error) {
	f, err := os.Open(getJSONPath(id))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var result Comic
	err = json.NewDecoder(f).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func save(comic *Comic) bool {
	f, err := os.Create(getJSONPath(comic.Num))
	if err != nil {
		return false
	}
	defer f.Close()

	b, err := json.Marshal(comic)
	if err != nil {
		return false
	}

	_, err = f.Write(b)
	if err != nil {
		return false
	}

	return true
}

func printComic(comic *Comic) {
	fmt.Printf("Title: %s\nTransScript: %s\nURL: %s/\n\n%d",
		comic.Title, comic.Transcript, APIURL, comic.Num)
}

// commands
func index() {
	var items []Comic

	// download
	skip := 0
	for i := 1; i < MAX_ID; i++ {
		comic, err := fetchComic(i)
		if err != nil {
			skip++
			if skip > 2 {
				break
			}
			continue
		}
		skip = 0
		fmt.Printf("#%d %s\n", comic.Num, comic.Title)
		items = append(items, *comic)
	}

	// write index
	f, err := os.Create(getIndexJSONPath())
	if err != nil {
		return
	}
	defer f.Close()

	fmt.Printf("fetched %d comics", len(items))
	b, err := json.Marshal(items)
	if err != nil {
		return
	}

	_, err = f.Write(b)
	if err != nil {
		return
	}
}

func search(query string) {
	f, err := os.Open(getIndexJSONPath())
	if err != nil {
		return
	}
	defer f.Close()

	var items []Comic
	err = json.NewDecoder(f).Decode(&items)
	if err != nil {
		return
	}

	var found []Comic
	for _, c := range items {
		if strings.Index(c.Title, query) >= 0 {
			found = append(found, c)
			continue
		}

		if strings.Index(c.Transcript, query) >= 0 {
			found = append(found, c)
			continue
		}
	}

	for _, c := range found {
		printComic(&c)
	}
}

func help() {
	fmt.Printf("Usage: go run main.go <command> [<query>]\n" +
		"commands:\n" +
		"\tindex: Download all comics data from xkcd.com and index it\n" +
		"\tsearch(query): Search comics using query with local index\n")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}

	method := os.Args[1]
	switch method {
	case "index":
		index()
	case "search":
		search(os.Args[2])
	default:
		help()
	}

	/*
		var query int64 = 571
		query, err := strconv.ParseInt(os.Args[1], 0, 10)
		if err != nil {
			fmt.Printf("It looks like your argument is wrong\n")
			os.Exit(2)
		}

		comic, err := fetchComic(int(query))
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(2)
		}
		printComic(comic)
	*/
}
