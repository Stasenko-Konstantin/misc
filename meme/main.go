// 135
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/png"
	"os"
)

var table [][]color.Color

func readTable(img image.Image) {
	bounds := img.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		table = append(table, nil)
		for y := 0; y < bounds.Max.Y; y++ {
			table[x] = append(table[x], img.At(x, y))
		}
	}
}

func checkX(x, i, lenTable int) int {
	var a int
	if x > lenTable/2 {
		a = x + i
		if a > lenTable {
			return x
		}
	} else {
		a = x - i
		if a < 0 {
			return x
		}
	}
	return a
}

func main() {
	const (
		nframes = 64
		delay   = 8
	)
	args := os.Args
	if len(args) < 2 {
		panic("Need image and gif as argument: ./meme in.png out.gif")
	}
	path := args[1]
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	anim := gif.GIF{LoopCount: nframes, BackgroundIndex: 0}
	bounds := img.Bounds()
	readTable(img)
	fmt.Println("start rendering")
	percent := (bounds.Max.X / 2) / 10
	var (
		edges    []int
		turn     bool
		lenTable = len(table)
	)
	fmt.Print("[")
	for i := 1; i <= bounds.Max.X/2; i++ {
		edges = append(edges, i, bounds.Max.X-i)
		palleted := image.NewPaletted(bounds, palette.Plan9)
		draw.Draw(palleted, bounds, img, image.Point{}, draw.Src)
		for x := 0; x < bounds.Max.X; x++ {
			for y := 0; y < bounds.Max.Y; y++ {
				if contains(edges, x) {
					palleted.Set(x, y, color.White)
				} else {
					if turn {
						palleted.Set(x, y, table[checkX(x, i, lenTable)][y])
					} else {
						palleted.Set(x, y, table[x][y])
					}
				}
			}
		}
		turn = !turn
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, palleted)
		if i%percent == 0 {
			fmt.Print("|")
		}
	}
	f, err = os.Create(args[2])
	if err != nil {
		panic(err)
	}
	err = gif.EncodeAll(f, &anim)
	if err != nil {
		panic(err)
	}
	fmt.Println("]")
}

func contains[T comparable](s []T, e T) bool {
	for _, c := range s {
		if c == e {
			return true
		}
	}
	return false
}
