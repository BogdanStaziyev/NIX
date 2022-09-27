package workDb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"nix_practice/Beginer/domain"
	"os"
	"time"
)

func WorkWithDb(i int) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	writeToDbPosts(i, db)
}

func writeToDbPosts(i int, db *sql.DB) {
	var posts []domain.Post

	r := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId="+"%d", i)
	res := findResultsPosts(r, posts)
	err := json.Unmarshal(res, &posts)
	if err != nil {
		log.Fatal(err)
	}
	for g := 0; g <= len(posts)-1; g++ {
		_, err = db.Query("INSERT INTO posts (user_id, id, title, body) VALUES (?,?,?,?)", posts[g].UserId, posts[g].Id, posts[g].Title, posts[g].Body)
		if err != nil {
			log.Println(err)
		}
		writeToDbComments(posts[g].Id, db)
	}
}

func findResultsPosts(url string, posts []domain.Post) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
		log.Println(err)
	}

	out, err := json.MarshalIndent(posts, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return out
}

func writeToDbComments(i int, db *sql.DB) {
	var comments []domain.Comment

	r := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId="+"%d", i)
	res := findResultsComments(r, comments)
	err := json.Unmarshal(res, &comments)
	if err != nil {
		log.Fatal(err)
	}
	for val := 0; val < len(comments); val++ {
		_, err = db.Query("INSERT INTO comments (post_id, id, name, email, body, time) VALUES (?,?,?,?,?,?)", comments[val].PostId, comments[val].Id, comments[val].Name, comments[val].Email, comments[val].Body, time.Now().Nanosecond())
		if err != nil {
			log.Println(err)
		}
	}
}

func findResultsComments(url string, comments []domain.Comment) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&comments); err != nil {
		log.Println(err)
	}

	out, err := json.MarshalIndent(comments, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return out
}
