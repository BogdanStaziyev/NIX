package Trainee

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"nix_practice/Beginer/domain"
	"os"
	"sync"
)

func WorkWithDbGorm(i int) {
	dsn := "root:@tcp(127.0.0.1:3306)/testdb"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	err = db.AutoMigrate(&domain.Post{}, &domain.Comment{})
	if err != nil {
		log.Println(err)
	}

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
					db.Create(&domain.Comment{
						PostId: comments[val].PostId,
						Id:     comments[val].Id,
						Name:   comments[val].Name,
						Email:  comments[val].Email,
						Body:   comments[val].Body,
					})
				}
				wg.Done()
			}()
		}()

		db.Create(&domain.Post{
			UserId: posts[g].UserId,
			Id:     posts[g].Id,
			Title:  posts[g].Title,
			Body:   posts[g].Body,
		})
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
