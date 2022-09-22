package main

import (
	"fmt"
	"log"
	j "nix_practice/Beginer/1"
	g "nix_practice/Beginer/2"
	f "nix_practice/Beginer/3"
	"time"
)

func main() {
	var num int
	fmt.Println(string(j.JsonPlaceholder()))
	time.Sleep(2 * time.Second)
	fmt.Println("Enter number from one to 100 too demonstrate goroutine power")
	if _, err := fmt.Scanln(&num); err != nil {
		log.Fatal(err)
	}
	g.GoRoutines(num)
	time.Sleep(2 * time.Second)
	f.CreateFileNix(num)
}
