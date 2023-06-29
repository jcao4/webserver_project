package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}

func goodbye(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Goodbye")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/goodbye", goodbye)

	http.ListenAndServe(":8080", nil)
}
