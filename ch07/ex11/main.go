package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// DB
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

// API
func (db database) list(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item / price params are required")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong price")
		return
	}

	// update
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item %s is already created", item)
	} else {
		w.WriteHeader(http.StatusCreated)
		db[item] = dollars(price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" && req.Method != "PATCH" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item / price params are required")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong price")
		return
	}

	// update
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusOK)
		db[item] = dollars(price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %s is not found", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" && req.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	item := req.URL.Query().Get("item")

	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item params is required")
		return
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusOK)
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %s is not found", item)
	}
}
