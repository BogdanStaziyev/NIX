package placeholder

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Post struct {
	UserId int    `json:"userId,omitempty"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

func JsonPlaceholder() []byte {
	var jsonRes *[]Post

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
	return out
}
