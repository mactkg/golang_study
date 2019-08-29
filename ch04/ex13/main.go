package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const APIURL = "http://www.omdbapi.com"

type Comic struct {
	Num         int
	transscript string
	img         string
}

func getJSONPath(id int) string {
	return "./data/" + strconv.FormatInt(int64(id), 10) + ".json"
}

func getComic(id int) (comic *Comic, err error) {
	// checkCache
	_, err = os.Stat(getJSONPath(id))
	if err == nil {
		return load(id)
	}

	// if can't find, fetch it
	url := fmt.Sprintf("%s/id/info.0.json", APIURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error: %d", resp.StatusCode)
		return nil, err
	}

	var result Comic
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	save(&result)

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

	buf := bytes.NewBufferString("")
	err = json.NewEncoder(buf).Encode(comic)
	if err != nil {
		return false
	}

	writer := bufio.NewWriter(f)
	_, err = writer.ReadFrom(buf)
	if err != nil {
		return false
	}

	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: go run main.go <movie-title> <api-key>\n")
		os.Exit(0)
	}

	title := os.Args[1]
	key := os.Args[2]
	movie, err := getMovie(title, key)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(2)
	}
	fmt.Printf("A movie is found...: %s\n", movie.Title)

	path, err := fetchPoster(movie, key)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(2)
	}

	fmt.Printf("Check it out: ./%s\n", path)
}
