package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	webcam.Set(gocv.VideoCaptureFPS, 24)
	window := gocv.NewWindow("hls-sample")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
		fmt.Println(string(img.ToBytes()))
	}
}
