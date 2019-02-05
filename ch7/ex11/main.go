package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	var m sync.RWMutex
	db := database{m, make(map[string]dollars)}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32
type database struct {
	sync.RWMutex
	items map[string]dollars
}

func (db database) print(w io.Writer) {
	for item, price := range db.items {
		fmt.Fprintf(w, "%v: %v\n", item, price)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	db.RLock()
	defer db.RUnlock()

	db.print(w)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	db.RLock()
	defer db.RUnlock()

	item := req.URL.Query().Get("item")
	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%v\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	defer db.print(w)
	if _, ok := db.items[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "already exist item: %v\n", item)
		return
	}
	pricef, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %v\n", price)
		return
	}
	db.items[item] = dollars(pricef)
	fmt.Fprintf(w, "success create item: %v, price: %v\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	defer db.print(w)
	if _, ok := db.items[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %v\n", item)
		return
	}
	pricef, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %v\n", price)
		return
	}
	db.items[item] = dollars(pricef)
	fmt.Fprintf(w, "success update item: %v, price: %v\n", item, price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()

	item := req.URL.Query().Get("item")
	defer db.print(w)
	if _, ok := db.items[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %v\n", item)
		return
	}
	delete(db.items, item)
	if _, ok := db.items[item]; ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed delete item: %v\n", item)
		return
	}
	fmt.Fprintf(w, "success delete item: %v\n", item)
}
