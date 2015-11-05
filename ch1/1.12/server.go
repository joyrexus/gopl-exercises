// Server3 is ...
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/joyrexus/lissajous"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", render)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// render renders a lissajous figure.
func render(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	cycles, err := strconv.Atoi(r.FormValue("cycles"))
	if err != nil {
		cycles = 1
	}
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		size = 100
	}
	lissajous.Write(w, cycles, size)
}

// echo echoes the HTTP request.
func echo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Path = %q\n", r.URL.Path)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
