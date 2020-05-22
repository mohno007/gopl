// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 2048, 2048
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			c := supersample(px, py, newton)
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func supersample(px int, py int, f func(complex128) color.Color) color.Color {
	grid := []float64{-0.5, +0.5}
	var ry, rcb, rcr uint32

	for _, sx := range grid {
		for _, sy := range grid {
			// Grid https://en.wikipedia.org/wiki/Supersampling
			x := (float64(px)+sx)/width*(xmax-xmin) + xmin
			y := (float64(py)+sy)/height*(ymax-ymin) + ymin
			z := complex(x, y)

			// RGBでは平均値の色の出方が微妙だったので、YCbCr色空間を使ってみた
			// https://xn--go-hh0g6u.com/pkg/image/color/#RGBToYCbCr
			// RGB と Y'CbCr の間の変換は損失が大きいらしく、2回変換してるので微妙に色の出方が違うかもしれない
			c := f(z)
			r, g, b, _ := c.RGBA()
			cy, ccb, ccr := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))

			ry += uint32(cy)
			rcb += uint32(ccb)
			rcr += uint32(ccr)
		}
	}

	d := uint32(len(grid) * len(grid))
	return color.YCbCr{uint8(ry / d), uint8(rcb / d), uint8(rcr / d)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrastCb = 7
	const contrastCr = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// 色温度使いたかった
			return color.YCbCr{128, 255 - contrastCb*n, 255 - contrastCr*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
