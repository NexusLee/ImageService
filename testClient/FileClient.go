package main

import (
	"net/http"
	"os"
	"fmt"
	"time"
	"io"
)

func main() {
	for i := 0; i < 200; i++ {
		go testing()
	}
	time.Sleep(time.Second * 15)

}

func testing() {
	file, err := os.Open("/tmp/image.png")
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("http://192.168.1.14:3000/","file/image",file)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	io.Copy(os.Stdout,response.Body)

}
