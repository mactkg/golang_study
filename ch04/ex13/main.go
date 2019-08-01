package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const APIURL = "http://www.omdbapi.com"

type Movie struct {
	Title  string
	Poster string
}

func getMovie(title, key string) (movie *Movie, err error) {
	url := fmt.Sprintf("%s?t=%s&apikey=%s", APIURL, url.QueryEscape(title), key)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error: %d", resp.StatusCode)
		return nil, err
	}

	var result Movie
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func fetchPoster(movie *Movie, key string) (path string, err error) {
	url := movie.Poster
	resp, err := http.Get(url + "?apikey=" + key)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error: %d", resp.StatusCode)
		return "", err
	}

	file, err := os.Create(movie.Title + filepath.Ext(movie.Poster))
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	err = writer.Flush()
	if err != nil {
		return "", err
	}
	return "./" + file.Name(), nil
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
