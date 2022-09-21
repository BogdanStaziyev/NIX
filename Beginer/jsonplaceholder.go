package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type post struct {
	UserId int    `json:"userId,omitempty"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

func main() {
	var jsonRes *[]post
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&jsonRes); err != nil {
		log.Println(err)
	}
	out, err := json.MarshalIndent(jsonRes, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
