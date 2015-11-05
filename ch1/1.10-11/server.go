// Server is an http server used for testing fetchall.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hi)
	http.HandleFunc("/noresponse", noresponse)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// hi responds with a little hello.
func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi!")
}

// noresonse never responds.
func noresponse(w http.ResponseWriter, r *http.Request) {
	for {
	}
}
