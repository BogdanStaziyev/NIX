package workDb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type comments []comment

type comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func WorkWithDb(i int) {
	var coo comments

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId="+"%d", i)
	res := findResults(r)
	err = json.Unmarshal(res, &coo)
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range coo {
		_, err := db.Query("INSERT INTO comments (post_id, id, name, email, body) VALUES (?,?,?,?,?)", val.PostId, val.Id, val.Name, val.Email, val.Body)
		if err != nil {
			log.Println(err)
		}
	}
}

func findResults(url string) []byte {
	var comm comments
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&comm); err != nil {
		log.Println(err)
	}

	out, err := json.MarshalIndent(comm, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return out
}
