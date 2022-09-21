package beginer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	beginer "nix_practice/Beginer/1"
	"os"
	"sync"
)

func GoRoutines(num int) {
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 1; i <= num; i++ {
		count := i
		go func() {
			coll(count)
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func coll(i int) {
	var jsonRes *beginer.Post
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
	fmt.Println(string(out))
}