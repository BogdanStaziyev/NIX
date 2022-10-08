package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"nix_practice/Beginer/2"
	"os"
)

func CreateFile(num int) {
	for i := 1; i <= num; i++ {
		fileLocation := "./storage/posts/"
		createFileName := fmt.Sprintf(fileLocation+"%d.txt", i)
		//use ConvJsonByte from goroutines package not to duplicate code
		res := goroutines.ConvJsonToByte(i)
		if err := os.MkdirAll(fileLocation, 0777); err != nil {
			log.Println("MakeDir failed:", err)
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(createFileName, res, 0600); err != nil {
			log.Fatal(err)
		}
	}
}
