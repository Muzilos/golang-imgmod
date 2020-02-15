package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

type pixel struct {
	r, g, b, a uint8
}
type convert func(uint8) uint8

func main() {
	var infile string
	flag.StringVar(&infile, "i", "./img/painting.jpg", "jpeg format image file to read")
	flag.StringVar(&infile, "input", "./img/painting.jpg", "jpeg format image file to read")
	var outfile string
	flag.StringVar(&outfile, "o", "./img/modded.jpg", "jpeg format filename for output")
	flag.StringVar(&outfile, "output", "./img/modded.jpg", "jpeg format filename for output")
	flag.Parse()

	fmt.Println("Reading: ", infile)
	fmt.Println("Writing: ", outfile)
	img := loadImage(infile)
	imgPixels := getPixels(img)

	var moddedImage = modImage(imgPixels, img.Bounds().Dx(), img.Bounds().Dy(), pixelMod)
	writeImage(outfile, moddedImage)
	// for i, pixel := range getPixels(moddedImage) {
	// 	fmt.Println("Pixel", i, "\t r g b a:", pixel)
	// }
}

func pixelMod(p uint8) uint8 {
	return uint8(float64(p) / 1.2)
}

func loadImage(filename string) image.Image {
	f, err := os.Open((filename))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func writeImage(filename string, img image.Image) {
	outfile, err := os.Create((filename))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	var opt jpeg.Options
	opt.Quality = 100
	// ok, write out the data into the new JPEG file

	err = jpeg.Encode(outfile, img, &opt) // put quality to 80%
	if err != nil {
		log.Fatal(err)
	}
}

func getPixels(img image.Image) []pixel {
	bounds := img.Bounds()
	Dx := bounds.Dx()
	Dy := bounds.Dy()
	fmt.Println(Dx, " x ", Dy) // debugging
	pixels := make([]pixel, Dx*Dy)

	for i := 0; i < Dx*Dy; i++ {
		x := i % Dx
		y := i / Dx
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = uint8(r)
		pixels[i].g = uint8(g)
		pixels[i].b = uint8(b)
		pixels[i].a = uint8(a)
	}
	return pixels
}

func modImage(pixels []pixel, Dx int, Dy int, fn convert) image.Image {
	newRect := image.Rectangle{image.Point{0, 0}, image.Point{Dx, Dy}}
	newRGBA := image.NewRGBA(newRect)
	modded := make([]pixel, len(pixels))
	for i, pixel := range pixels {
		modded[i].r = fn(pixel.r)
		modded[i].g = fn(pixel.g)
		modded[i].b = fn(pixel.b)
		modded[i].a = fn(pixel.a)

		var loc = modded[i]
		rgbaColor := color.RGBA{loc.r, loc.g, loc.b, loc.a}
		newRGBA.SetRGBA(i%Dx, i/Dx, rgbaColor)
	}
	return newRGBA
}
