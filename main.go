package main

import (
	"flag"
	"fmt"
	"golang.org/x/image/bmp"
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
	flag.StringVar(&infile, "i", "./img/painting.jpg", "image file to read")
	flag.StringVar(&infile, "input", "./img/painting.jpg", "image file to read")
	var outfile string
	flag.StringVar(&outfile, "o", "./img/modded.jpg", "filename for output")
	flag.StringVar(&outfile, "output", "./img/modded.jpg", "filename for output")
	var imgtype string
	flag.StringVar(&imgtype, "t", "jpeg", "file format (jpeg, bmp)")
	flag.StringVar(&imgtype, "type", "jpeg", "file format (jpeg, bmp)")

	flag.Parse()

	fmt.Println("Reading: ", infile)
	fmt.Println("Writing: ", outfile)
	img := loadImage(infile, imgtype)
	imgPixels := getPixels(img)

	var moddedImage = modImage(imgPixels, img.Bounds().Dx(), img.Bounds().Dy(), pixelMod)
	writeImage(outfile, moddedImage, imgtype)
	// for i, pixel := range getPixels(moddedImage) {
	// 	fmt.Println("Pixel", i, "\t r g b a:", pixel)
	// }
}

func pixelMod(p uint8) uint8 {
	return uint8(float64(p) * 3.0)
}

func loadImage(filename string, imgtype string) image.Image {
	f, err := os.Open((filename))
	var img image.Image
	defer f.Close()
	if imgtype == "bmp" {
		img, err = bmp.Decode(f)
	} else {
		img, err = jpeg.Decode(f)
	}
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func writeImage(filename string, img image.Image, imgtype string) {
	outfile, err := os.Create((filename))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	if imgtype == "jpeg" {
		var opt jpeg.Options
		opt.Quality = 100
		// ok, write out the data into the new JPEG file
		err = jpeg.Encode(outfile, img, &opt) // put quality to 80%
	} else if imgtype == "bmp" {
		err = bmp.Encode(outfile, img)
	}
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
