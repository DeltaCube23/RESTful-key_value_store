package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

const (
	length = 101
)

type store struct {
	db [length]*linkedlist
	mu sync.RWMutex
}

//test handler
func (kv *store) kvstoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/kvstore" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "KEY VALUE STORE !!!")
}

//handles all post requests
func (kv *store) postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	kv.mu.Lock()
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	kv.Push(name, address)
	kv.mu.Unlock()
}

//handles all get requests
func (kv *store) getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	kv.mu.RLock()
	fmt.Fprintf(w, "GET request successful\n")
	key := r.URL.Query().Get("name")
	addr := kv.Get(key)
	fmt.Fprintf(w, "Address = %s\n", addr)
	kv.mu.RUnlock()
}

//handles all put requests
func (kv *store) putHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	kv.mu.Lock()
	fmt.Fprintf(w, "PUT request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	ok := kv.Put(name, address)

	if ok == true {
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	} else {
		fmt.Fprintf(w, "Invalid key value pair\n")
	}
	kv.mu.Unlock()
}

//handles all delete requests
func (kv *store) deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	kv.mu.Lock()
	fmt.Fprintf(w, "DELETE request successful\n")
	key := r.URL.Path[8:]
	ok := kv.Delete(key)

	if ok == true {
		fmt.Fprintf(w, "Removed Name = %s\n", key)
	} else {
		fmt.Fprintf(w, "Invalid key value pair\n")
	}
	kv.mu.Unlock()
}

//creates a new instance of key value store
func newStore() *store {
	return &store{}
}

func main() {
	kv := newStore()
	http.HandleFunc("/kvstore", kv.kvstoreHandler)
	http.HandleFunc("/create", kv.postHandler)
	http.HandleFunc("/read/", kv.getHandler)
	http.HandleFunc("/update", kv.putHandler)
	http.HandleFunc("/delete/", kv.deleteHandler)

	//Start the server and listen for requests
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
