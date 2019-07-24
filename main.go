package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Prova struct {
	Id int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var p Prova

		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Println(r.Body)
		fmt.Println(p.Id)
	})

	http.ListenAndServe(":3300", nil)
}
