package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	for i := 0; i < 400; i++ {
		go testing()
	}
	time.Sleep(time.Second * 15)

}

func testing() {
	file, err := os.Open("/tmp/image.png")
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("http://localhost:3000/", "file/image", file)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	io.Copy(os.Stdout, response.Body)

}
