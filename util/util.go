package util

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"math"
)

func AverageColor(img image.Image) color.RGBA {
	bounds := img.Bounds()

	var r, g, b, a, count uint64
	count = 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			count++
			t_r, t_g, t_b, t_a := img.At(x, y).RGBA()
			r += uint64(t_r)
			g += uint64(t_g)
			b += uint64(t_b)
			a += uint64(t_a)
		}
	}

	r = r / count
	g = g / count
	b = b / count
	a = a / count

	a_color := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}

	// fmt.Printf("r: %i g: %i b: %i a: %i",a_color.R,a_color.G,a_color.B,a_color.A)
	return a_color
}

// Nearest-neighbor interpolation Algorithm
func ResizeImage(dst draw.Image, src image.Image) {
	// TODO: comprobaciones tamaño (implementado solo reduccion)

	src_bounds := src.Bounds()
	dst_bounds := dst.Bounds()

	src_height := (src_bounds.Max.Y - src_bounds.Min.Y)
	src_widht := (src_bounds.Max.X - src_bounds.Min.X)
	dst_height := (dst_bounds.Max.Y - dst_bounds.Min.Y)
	dst_widht := (dst_bounds.Max.X - dst_bounds.Min.X)

	v_splits := src_height / dst_height
	h_splits := src_widht / dst_widht

	for y := 0; y < dst_height; y++ {
		for x := 0; x < dst_widht; x++ {
			color := src.At((x*h_splits)+src_bounds.Min.X, (y*v_splits)+src_bounds.Min.Y)
			dst.Set(x+dst_bounds.Min.X, y+dst_bounds.Min.Y, color)
		}
	}
}

func ColorDistance(c1, c2 color.Color) float64 {
	c1_r, c1_g, c1_b, _ := c1.RGBA()
	c2_r, c2_g, c2_b, _ := c2.RGBA()

	r := math.Pow(float64(c1_r-c2_r), 2)
	g := math.Pow(float64(c1_g-c2_g), 2)
	b := math.Pow(float64(c1_b-c2_b), 2)
	// a := math.Pow(float64(c1_a - c2_a), 2)

	return math.Sqrt(r + g + b)
}
