package placeholder

import (
	"encoding/json"
	"log"
	"net/http"
	"nix_practice/Beginer/domain"
	"os"
)

func JsonPlaceholder() []byte {
	var jsonRes *[]domain.Post

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
