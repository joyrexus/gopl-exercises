// Server3 is ...
package main

import (
	"fmt"
	"log"
	"net/http"
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
	lissajous.Write(w)
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
		log.Print(err)
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
