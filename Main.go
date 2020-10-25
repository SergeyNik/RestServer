package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct  {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Poster struct  {
	Name string
}

func main() {
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			user := User { Id: 1, Name: "John Doe" }
			_ = json.NewEncoder(w).Encode(user)
		}

		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			var t Poster
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			log.Println(t.Name)

			w.Header().Set("Content-Type", "application/json")
			user := User { Id: 1, Name: t.Name }
			_ = json.NewEncoder(w).Encode(user)
		}
	})

	_ = http.ListenAndServe(":3000", nil)
}

