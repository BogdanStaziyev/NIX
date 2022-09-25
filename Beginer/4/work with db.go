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
	for _, val := range posts {
		writeToDbComments(val.Id, db)
		_, err = db.Query("INSERT INTO posts (user_id, id, title, body) VALUES (?,?,?,?)", val.UserId, val.Id, val.Title, val.Body)
		if err != nil {
			log.Println(err)
		}
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
	for _, val := range comments {
		_, err = db.Query("INSERT INTO comments (post_id, id, name, email, body) VALUES (?,?,?,?,?)", val.PostId, val.Id, val.Name, val.Email, val.Body)
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
