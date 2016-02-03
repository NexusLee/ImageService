package main

import (
	"net/http"
	"fmt"
	"image/png"
	"image/color"
	"image"
)

func main() {
	counter := 0;
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if(r.Method == "POST") {
			counter++
			current := counter;
			fmt.Println("Started:", current)
			processImage(w,r)
			fmt.Println("Finished:", current)
		} else {
			fmt.Fprintln(w,"ERROR: Only POST accepted.")
		}
	})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func processImage(w http.ResponseWriter, r *http.Request) {
	myImage, err := png.Decode(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	m := image.NewRGBA(myImage.Bounds())
	for i := 0; i < m.Rect.Max.X; i++ {
		for j := 0; j < m.Rect.Max.Y; j++ {
			r, g, b, _ := myImage.At(i, j).RGBA()
			myColor := new(color.RGBA)
			myColor.R = uint8((g * r) / 255)
			myColor.G = uint8((g * r) / 255)
			myColor.B = uint8((b * b) / 255)
			myColor.A = uint8(255)
			m.Set(i, j, myColor)
		}
	}
	png.Encode(w, m)
}