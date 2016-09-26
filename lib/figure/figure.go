package figure

import (
	"image"
	"image/color"
)

type Canvas struct {
	myRGBA *image.RGBA
}

type Point struct{
	X, Y int
}

func Coords() *Canvas {
	width, height := 201, 201
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.White)
		}
	}
	return &Canvas{myRGBA: img}
}

func (c *Canvas) Drawfunc(a, b int) {

	width, height := c.getWidthHeight()
	for y := -height / 2; y < height/2; y++ {
		for x := -width / 2; x < width/2; x++ {
			if a*x+b-1 <= y && y <= a*x+b+1 {
				c.SetColorAt(x, y, color.RGBA{0x24, 0x6D, 0xB7, 0xFF})
			}
		}
	}
}

func (c *Canvas) getWidthHeight() (int, int) {
	var i *image.RGBA = c.myRGBA
	var p image.Point = i.Rect.Size()
	return p.X, p.Y
}

func (c *Canvas) SetColorAt(x, y int, color color.RGBA) {
	width, height := c.getWidthHeight()
	var i *image.RGBA = c.myRGBA

	i.SetRGBA(x+width/2, -(y - height/2), color)
}

func (c * Canvas) SetPointAt(x, y int, color color.RGBA){
	width, height := c.getWidthHeight()
	var img *image.RGBA = c.myRGBA
	for j := -(y - height/2) -1; j < -(y - height/2) + 1; j++{
		for i := x+width/2 -1; i < x+width/2 + 1; i ++{
			if 0 <= j && j < height && 0 <= i && i < width{			 	
				img.SetRGBA(i, j, color)	
			}
		}
	}
}

func (c *Canvas) GetMyRGBA() *image.RGBA {
	return c.myRGBA
}
