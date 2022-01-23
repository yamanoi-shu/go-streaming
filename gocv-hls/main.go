package main

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

const (
	VideoCaptureWidth  = 1280
	VideoCaptureHeight = 720
	VideoCaptureFPS    = 24
)

func main() {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatal(err)
	}
	defer webcam.Close()

	height := webcam.Get(gocv.VideoCaptureFrameHeight)
	width := webcam.Get(gocv.VideoCaptureFrameWidth)
	fps := webcam.Get(gocv.VideoCaptureFPS)
	if fps < 0.0 {
		fps = 3.0
	}

	fmt.Printf("fps -> %f, height -> %f, width -> %f\n", fps, height, width)

	window := gocv.NewWindow("hls-sample")
	defer window.Close()
	img := gocv.NewMat()
	defer img.Close()

	for {
		if webcam.Read(&img) {
			if !img.Empty() {
				window.IMShow(img)
				fmt.Println(string(img.ToBytes()))
			}
		}
		window.WaitKey(1)
	}
}
