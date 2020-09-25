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
	"runtime"
	"sync"
)

type plot struct {
	px    int
	py    int
	color color.Color
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 4096, 4096
	)

	var wg sync.WaitGroup

	plots := make(chan plot, 1024)
	done := make(chan struct{})
	cpuCount := runtime.NumCPU()
	for cpu := 0; cpu < cpuCount; cpu++ {
		pyStart := (height / cpuCount) * cpu
		pyEnd := (height / cpuCount) * (cpu + 1)
		if (cpu + 1) == cpuCount {
			pyEnd = height
		}
		wg.Add(1)
		// 論理コア数ぴったりだと、若干遅くなってしまう模様
		//    2.00s user 0.31s system 238% cpu 0.971 total
		// 論理コア数 + 1 ほぼ変わらず
		// 論理コア数 * 2 早い
		//    2.22s user 0.34s system 286% cpu 0.896 total
		// 論理コア数 * 3 早い
		//    2.46s user 0.32s system 315% cpu 0.879 total
		// 論理コア数 * 4 早い
		//    2.59s user 0.32s system 334% cpu 0.869 total
		// 論理コア数 * 16 早い
		//    2.58s user 0.33s system 336% cpu 0.862 total
		// 疑問
		//    CPU論理コア数は8あるのに 330%程度で頭打ちになる
		//    めちゃくちゃ多くすれば400%を超えるが...(SMT(hyper-threading)が使われる)
		//    もし物理コア8のマシンで実行すれば700%程度使われるのだろうか？
		go func(pyStart, pyEnd int) {
			defer wg.Done()
			for py := pyStart; py < pyEnd; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				// MEMO ここで、goroutineを呼び出すと、2048個のgoroutineができる
				// 2.66s user 0.35s system 314% cpu 0.955 total
				for px := 0; px < width; px++ {
					// MEMO ここで goroutineを呼び出すと、2048*2048 = 4194304個のgoroutineができてしまう
					// そこまで多いと、むしろ遅くなってしまう模様
					// 8.27s user 1.65s system 438% cpu 2.264 total
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					plots <- plot{px, py, mandelbrot(z)}
				}
			}
		}(pyStart, pyEnd)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	go func() {
		// channel with buffer can receive all messages after close
		for p := range plots {
			img.Set(p.px, p.py, p.color)
		}
		close(done)
	}()

	wg.Wait()
	close(plots) // all goroutines ends means all goroutines sent message to plots
	<-done
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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
