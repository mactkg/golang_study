package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func sendRequest(db *database, function func(database, http.ResponseWriter, *http.Request), method string, path string) (rec *httptest.ResponseRecorder, req *http.Request) {
	rec = httptest.NewRecorder()
	req = httptest.NewRequest(method, path, nil)
	function(*db, rec, req)
	return
}

func TestDBCreateOK(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	rec, _ := sendRequest(&db, database.create, "POST", "/create?item=shirt&price=10.5")

	if rec.Code != http.StatusCreated {
		t.Fatalf("HTTP Status: expected: 201, got: %v(%v)", rec.Code, rec.Body)
	}

	if price, ok := db["shirt"]; !ok || price != 10.5 {
		t.Fatalf("Price of shirt: expected: 10.5, got: %f", float32(price))
	}
}

func TestDBUpdateOK(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	rec, _ := sendRequest(&db, database.update, "POST", "/update?item=shoes&price=40.5")

	if rec.Code != http.StatusOK {
		t.Fatalf("HTTP Status: expected: 200, got: %v(%v)", rec.Code, rec.Body)
	}

	if price, ok := db["shoes"]; !ok || price != 40.5 {
		t.Fatalf("Wrong price of shoes! expected: 40.5, got: %f", float32(price))
	}
}

func TestDBUpdateWithoutPrice(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	rec, _ := sendRequest(&db, database.update, "POST", "/update?item=shoes")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("HTTP Status: expected: 400, got: %v(%v)", rec.Code, rec.Body)
	}

	if price, ok := db["shoes"]; !ok || price != 50 {
		t.Fatalf("Wrond price of shoes! expected: 50, got: %f", float32(price))
	}
}

func TestDBDeleteOK(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	rec, _ := sendRequest(&db, database.delete, "POST", "/delete?item=shoes")

	if rec.Code != http.StatusOK {
		t.Fatalf("HTTP Status: expected: 200, got: %v(%v)", rec.Code, rec.Body)
	}

	if _, ok := db["shoes"]; ok {
		t.Fatalf("Shoes is not deleted")
	}
}
