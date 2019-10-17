package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/fogleman/gorgb"
)

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	im, _, err := image.Decode(file)
	return im, err
}

func savePNG(path string, im image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: gorgb input.png output.png")
		return
	}

	im, err := loadImage(args[0])
	if err != nil {
		log.Fatal(err)
	}

	im = gorgb.Convert(im)

	// if err := gorgb.Verify(im); err != nil {
	// 	log.Fatal(err)
	// }

	err = savePNG(args[1], im)
	if err != nil {
		log.Fatal(err)
	}
}
