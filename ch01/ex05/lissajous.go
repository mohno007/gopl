package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

// 頭回ってないので全部コメントつけていく
// lissajous はリサージュ曲線をGIFとして出力します
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発振器 x が完了する周回数
		res     = 0.001 // 回転の分解数
		size    = 100   // 画像の縦横幅の半分, 実際の画像はこの値の2倍+1になる
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム間の遅延 (nframes 枚 * (delay * 10ms/枚) = ms)
	)

	// 乱数で周波数を決める
	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差

	// 毎フレームごとに実行
	for i := 0; i < nframes; i++ {
		//
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 二色の画像
		img := image.NewPaletted(rect, palette)
		// 2 Pi rad = 360 deg, 0.001 rad ずつ増やして、cycles周回分描画していく
		// 周回するのは
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)              // -1 ~ 1
			y := math.Sin(t*freq + phase) // -1 ~ 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) // エンコードエラーを無視
	if err != nil {
		panic(err)
	}
}
