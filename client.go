package main

import (
	//"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"encoding/base64"
)

type Page struct {
	Url string `json:"url"`
}

func main() {
	page1 := Page{"https://www.nairaland.com/7106890/software-engineering-journey-alx-africa"}
	page1json, err := json.Marshal(page1)
	logError(err)
	//fmt.Printf("%s\n", page1json)
	page, err := http.Post("http://localhost:2069/", "application/json", bytes.NewBuffer(page1json))
	logError(err)
	pagetext, err := ioutil.ReadAll(page.Body)
	logError(err)
	pagetext_b64, err := base64.StdEncoding.DecodeString(string(pagetext))
	logError(err)
	//fmt.Printf("%s\n", pagetext_b64)
	err = ioutil.WriteFile("server-sent-npage.pdf", []byte(pagetext_b64), 0644)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered! :-", err)
	}
}
