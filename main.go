package main

import (
	"./lib/figure"
	"image/color"
	"image/png"
	"os"
)

func main() {

	c := figure.Coords()
	clr := color.RGBA{0x00, 0x00, 0x00, 0xFF}
	for i := 0; i < 100; i++ {
		c.SetColorAt(i, i, clr)
	}

	a := 2
	b := 3
	c.Drawfunc(a, b)

	png.Encode(os.Stdout, c.GetMyRGBA())
}
