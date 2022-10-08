package Trainee

import (
	"fmt"
	"io"
	"net"
	"time"
)

func PracticeNet() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, fmt.Sprint("Hello world", time.Now(), "\n"))

		conn.Close()
	}
}
