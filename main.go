package main

import (
	"fmt"
	"log"
	placeholder "nix_practice/Beginer/1"
	goroutines "nix_practice/Beginer/2"
	file "nix_practice/Beginer/3"
	workDb "nix_practice/Beginer/4"
	"nix_practice/Trainee"
)

func main() {
	var numGoroutines, numFiles int
	fmt.Println(string(placeholder.JsonPlaceholder()))

	fmt.Println("Enter number from one to 100 too demonstrate goroutine power")
	if _, err := fmt.Scanln(&numGoroutines); err != nil {
		log.Fatal(err)
	}
	goroutines.CreateGoRoutines(numGoroutines)

	fmt.Println("Enter number from one to 100 too create files")
	if _, err := fmt.Scanln(&numFiles); err != nil {
		log.Fatal(err)
	}
	file.CreateFile(numFiles)

	workDb.WorkWithDb(5)
	Trainee.WorkWithDbGorm(6)
}
