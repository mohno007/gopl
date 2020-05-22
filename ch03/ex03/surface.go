package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 500, 320            // キャンバスの大きさ
	cells         = 100                 // 格子のマス目の数(縦横)
	xyrange       = 30.0                // x-y軸の範囲
	xyscale       = width / 2 / xyrange // x単位 および y単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y軸の角度(180/6 = 30)
	zMax          = +0.30
	zMin          = -0.00
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
	next:
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			for _, v := range []float64{ax, ay, bx, by, cx, cy, dx, dy} {
				if math.IsInf(v, 0) || math.IsNaN(v) {
					break next
				}
			}

			zavg := (az + bz + cz + dz) / 4

			c := zColor(zavg)

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill: %s;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, rgbaToCSSString(c))
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx float64, sy float64, z float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z = f(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	return math.Min(math.Abs(math.Sin(x/2)), math.Abs(math.Sin(y/2))) / 3
}

func zColor(z float64) color.RGBA {
	z = math.Max(z, zMin)
	z = math.Min(z, zMax)

	scale := zMax - zMin
	zInScale := (z - zMin) / scale

	red := uint8(0xFF * (2*math.Max(zInScale, 0.5) - 1))
	blue := uint8(0xFF * (2 * math.Max(0.5-zInScale, 0)))

	return color.RGBA{red, 0x22, blue, 0xFF}
}

func rgbaToCSSString(c color.RGBA) string {
	return fmt.Sprintf("rgba(%d, %d, %d, %d)", c.R, c.G, c.B, c.A)
}
