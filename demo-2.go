package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type Page struct {
	Url string `json:"url"`
}

func main() {
	page1 := Page{"https://www.nairaland.com/7106890/software-engineering-journey-alx-africa"}
	page1json, err := json.Marshal(page1)
	logError(err)
	page, err := http.Post("http://localhost:8000/", "application/json", bytes.NewBuffer(page1json))
	logError(err)
	pagetext, err := ioutil.ReadAll(page.Body)
	logError(err)
	fmt.Printf("%s", pagetext)
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!", err)
	}
}
