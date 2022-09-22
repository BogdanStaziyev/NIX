package beginer

import (
	"fmt"
	"io/ioutil"
	"log"
	"nix_practice/Beginer/2"
	"os"
)

func CreateFileNix(num int) {
	for i := 1; i <= num; i++ {
		str := fmt.Sprintf("./storage/posts/")
		adr := fmt.Sprintf(str+"%d.txt", i)
		res := beginer.ConvJsonToByte(i)
		if err := os.MkdirAll(str, 0777); err != nil {
			fmt.Println("MakeDir failed:", err)
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(adr, res, 0600); err != nil {
			log.Fatal(err)
		}
	}
}
