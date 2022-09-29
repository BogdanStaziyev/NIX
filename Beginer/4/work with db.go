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
	"sync"
)

func WorkWithDb(i int) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var posts []domain.Post
	var wg sync.WaitGroup

	r := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId="+"%d", i)
	res := findResultsPosts(r, posts)
	err = json.Unmarshal(res, &posts)
	if err != nil {
		log.Fatal(err)
	}
	for g := range posts {
		wg.Add(1)
		var comments []domain.Comment
		go func() {
			urlComments := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId="+"%d", posts[g].Id)
			sliceByteComments := findResultsComments(urlComments, comments)
			err = json.Unmarshal(sliceByteComments, &comments)
			if err != nil {
				log.Fatal(err)
			}
			go func() {
				for val := range comments {
					_, err = db.Query("INSERT INTO comments (post_id, id, name, email, body) VALUES (?,?,?,?,?)", comments[val].PostId, comments[val].Id, comments[val].Name, comments[val].Email, comments[val].Body)
					log.Println(comments[val].Id)
					if err != nil {
						log.Println(err)
					}
				}
				wg.Done()
			}()
		}()

		_, err = db.Query("INSERT INTO posts (user_id, id, title, body) VALUES (?,?,?,?)", posts[g].UserId, posts[g].Id, posts[g].Title, posts[g].Body)
		if err != nil {
			log.Println(err)
		}
	}
	wg.Wait()
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
