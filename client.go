package main

import (
	"fmt"
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
	var link string
	fmt.Println("Enter the link to be processed below (to process a default link, just press Enter) ")
	fmt.Print("Link >>> ")
	fmt.Scanf("%s", &link)
	page1 := Page{link}
	page1json, err := json.Marshal(page1)
	logError(err)
	//fmt.Printf("%s\n", page1json)
	fmt.Println("Making HTTP request ...")
	//page, err := http.Post("http://etp4ma.octamile.com:9602/", "application/json", bytes.NewBuffer(page1json))
	page, err := http.Post("http://localhost:9602/", "application/json", bytes.NewBuffer(page1json))
	logError(err)
	fmt.Println("HTTP request successful!")
	fmt.Println("Reading and decoding HTTP response content ...")
	pagetext, err := ioutil.ReadAll(page.Body)
	logError(err)
	pagetext_b64, err := base64.StdEncoding.DecodeString(string(pagetext))
	logError(err)
	//fmt.Printf("%s\n", pagetext_b64)
	err = ioutil.WriteFile("client-files/server-sent-npage.pdf", []byte(pagetext_b64), 0644)
	logError(err)
	fmt.Println("The pdf file is saved in the current working directory as 'server-sent-npage.pdf'!")
	var ent string
	fmt.Print("Press Enter to exit ...")
	fmt.Scanf("%s", &ent)
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered! :- ", err)
	}
}
