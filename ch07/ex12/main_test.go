package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sendRequest(db *database, function func(database, http.ResponseWriter, *http.Request), method string, path string) (rec *httptest.ResponseRecorder, req *http.Request) {
	rec = httptest.NewRecorder()
	req = httptest.NewRequest(method, path, nil)
	function(*db, rec, req)
	return
}

func TestDBList(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	rec, _ := sendRequest(&db, database.list, "GET", "/list")

	if rec.Code != http.StatusOK {
		t.Fatalf("HTTP Status: expected: 200, got: %v(%v)", rec.Code, rec.Body)
	}

	// TODO: write DOM test if I have enough time to write
	if strings.Index(rec.Body.String(), "<table>") == -1 {
		t.Fatalf("Returned body doesn't have a table node: %v", rec.Body.String())
	}
}
