package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

func sendHTTP(method, url, authToken string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendHTTPWithJSON(method, url string, body io.Reader, authToken string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+authToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := sendHTTP("GET", SearchURL+"?q="+q, os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func ListIssue(owner, repo string) ([]*Issue, error) {
	resp, err := sendHTTP("GET", BaseURL+path.Join("repos", owner, repo, "issues"), os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("list query failed: %s", resp.Status)
	}

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func GetIssue(owner, repo string, number int64) (*Issue, error) {
	resp, err := sendHTTP("GET", BaseURL+path.Join("repos", owner, repo, "issues", strconv.FormatInt(number, 10)), os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CreateIssue(owner, repo string, fields map[string]string) (*Issue, error) {
	if fields["title"] == "" {
		err := fmt.Errorf("Title should be filled")
		return nil, err
	}

	buf := bytes.NewBufferString("")
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	resp, err := sendHTTPWithJSON("POST", BaseURL+path.Join("repos", owner, repo, "issues"), buf, os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		resp.Body.Close()
		return nil, fmt.Errorf("create issue failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func EditIssue(owner, repo string, number int64, fields map[string]string) (*Issue, error) {
	if fields["title"] == "" && fields["state"] == "" {
		err := fmt.Errorf("Title should be filled")
		return nil, err
	}

	buf := bytes.NewBufferString("")
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	resp, err := sendHTTPWithJSON("PATCH", BaseURL+path.Join("repos", owner, repo, "issues", strconv.FormatInt(number, 10)), buf, os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("edit issue failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CloseIssue(owner, repo string, number int64) (*Issue, error) {
	return EditIssue(owner, repo, number, map[string]string{"state": "closed"})
}
