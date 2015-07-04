package main

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/ww24/go-pseudo-color/pixel"
	"github.com/ww24/go-pseudo-color/pseudo"
	"gopkg.in/qml.v1"
)

func main() {
	err := qml.Run(run)
	fmt.Fprintln(os.Stderr, err)
}

func run() (e error) {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("file.qml")
	if err != nil {
		return err
	}

	file, err := os.Open("material/line.png")
	if err != nil {
		return err
	}
	defer file.Close()

	pix := pixel.NewPixel(file)

	images := []*pixel.Pixel{pix,
		pix.Map(pseudo.ConvLinear),
		pix.Map(pseudo.ConvSigmoid),
		pix.Map(pseudo.ConvSin)}

	engine.AddImageProvider("imgprv", func(imgId string, width, height int) (img image.Image) {
		index, _ := strconv.Atoi(imgId)
		img = images[index].Image
		return
	})

	win := component.CreateWindow(nil)
	win.Show()
	win.Wait()
	return
}
