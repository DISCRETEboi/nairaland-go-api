package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Page struct {
	Url string //`json:"url"`
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	page1 := Page{}
	rbody, err := ioutil.ReadAll(r.Body)
	logError(err)
	err = json.Unmarshal(rbody, &page1)
	logError(err)
	w.Write([]byte(page1.Url))
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!", err)
	}
}
