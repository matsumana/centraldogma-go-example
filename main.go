package main

import (
	"fmt"
	"net/http"

	"github.com/matsumana/centraldogma-go-example/dogma"
)

var greeting string

func main() {
	// dogma.NewClient()
	// fetched, err := dogma.Fetch()
	// if err != nil {
	// 	panic(err)
	// }
	// greeting = fetched

	err := dogma.Watch(func(data string) {
		greeting = data
	})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s, %s!", greeting, r.URL.Path[1:])
	})

	http.ListenAndServe(":8080", nil)
}
