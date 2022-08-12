package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
)

type Page struct {
	Url string //`json:"url"`
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on localhost:2069")
	log.Fatal(http.ListenAndServe("localhost:2069", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	page1 := Page{}
	rbody, err := ioutil.ReadAll(r.Body)
	logError(err)
	err = json.Unmarshal(rbody, &page1)
	logError(err)
	page1_b64 := base64.StdEncoding.EncodeToString([]byte(page1.Url))
	w.Write([]byte(page1_b64))
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!", err)
	}
}
