package goroutines

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nix_practice/Beginer/domain"
	"os"
	"sync"
)

func CreateGoRoutines(num int) {
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 1; i <= num; i++ {
		count := i
		go func() {
			b := ConvJsonToByte(count)
			fmt.Println(string(b))
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func ConvJsonToByte(i int) []byte {
	var jsonRes *domain.Post
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i)
	res, err := http.Get(url)
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
