package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	page, err := http.Get("http://localhost:8000/")
	logError(err)
	pagetext, err := ioutil.ReadAll(page.Body)
	logError(err)
	//fmt.Printf("%s", pagetext)
	fmt.Println(pagetext)
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!", err)
	}
}
