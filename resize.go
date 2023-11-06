package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// given a file name,s resize one image
func resizeImage(fileName string) {
	filePath := fmt.Sprintf("./tiny-imagenet-200/test/images/%s", fileName)
	outputPath := fmt.Sprintf("./output/resized_%s", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Load picture failed", fileName)
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Decode picture failed", fileName)
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(256, 256, img, resize.Lanczos3)

	out, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Write picture failed", fileName)
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
}
