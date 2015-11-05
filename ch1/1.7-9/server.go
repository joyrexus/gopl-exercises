// Server3 is ...
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hi)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi!")
}
