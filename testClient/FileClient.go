package main

import (
	"net/http"
	"os"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 200; i++ {
		go testing()
	}
	time.Sleep(time.Second * 15)
}

func testing() {
	file, err := os.Open("C:/temp/image.png")
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("http://localhost:3000/","file/image",file)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	fmt.Println(response.Body)

}
