package main

import (
	"./lib/dl"
	"./lib/figure"
	"image/color"
	"image/png"
	"os"
)

func main() {

	c := figure.Coords()
	clr := color.RGBA{0x00, 0x00, 0x00, 0xFF}

	for i := -100; i < 100; i++ {

		point := dl.Point{
			X: int(rand.Float64()*201) - 101,
			Y: int(rand.Float64()*201) - 101,
		}

		c.SetColorAt(point.X, point.Y, clr)
	}

	a := 2
	b := 3
	c.Drawfunc(a, b)

	png.Encode(os.Stdout, c.GetMyRGBA())
}
