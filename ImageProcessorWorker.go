package ImageFileMicroservices

import (
	"image/png"
	"os"
	"fmt"
	"image"
	"image/color"
)

func startProcessor(workToDo chan string, finishedWorkMap *map[string]bool) {
	var workId string
	runningProcesses := 0
	finishedWorkCommunicator := make(chan string)
	for {
		select {
		case workId = <- finishedWorkCommunicator:
			(*finishedWorkMap)[workId] = true
			runningProcesses--
			fmt.Println("Process finished:", runningProcesses)
		case workId = <- workToDo:
			if(runningProcesses < 10) {
				runningProcesses++
				fmt.Println("New process:", runningProcesses)
				go processImage(workId, finishedWorkCommunicator)
			} else {
				fmt.Println("Process loopbacked.")
				go workIdLoopback(workId, workToDo)
			}
		}
	}
}
func workIdLoopback(workId string, workToDo chan string) {
	workToDo <- workId
}

func processImage(workId string, finishedWorkCommunicator chan string) {
	if modifyImage(workId) {
		finishedWorkCommunicator <- workId
	}
}

func modifyImage(workId string) bool {
	file, err := os.Open("/tmp/" + workId + ".png")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	myImage, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	file.Close()
	file, err = os.Create("/tmp/" + workId + ".png")
	if err != nil {
		fmt.Println(err)
		return false
	}
	m := image.NewRGBA(myImage.Bounds())
	for i := 0; i < m.Rect.Max.X; i++ {
		for j := 0; j < m.Rect.Max.Y; j++ {
			r, g, b, _ := myImage.At(i, j).RGBA()
			myColor := new(color.RGBA)
			myColor.R = uint8((g * g) / 255)
			myColor.G = uint8((r * r) / 255)
			myColor.B = uint8((b * b) / 255)
			myColor.A = uint8(255)
			m.Set(i, j, myColor)
		}
	}
	png.Encode(file, m)
	file.Close()
	return true
}