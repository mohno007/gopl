package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type lissajousParameters struct {
	cycles  int     // 発振器 x が完了する周回数
	res     float64 // 回転の分解数(resolution)
	size    int     // 画像の縦横幅の半分, 実際の画像はこの値の2倍+1になる
	nframes int     // アニメーションフレーム数
	delay   int     // 10ms単位でのフレーム間の遅延 (nframes 枚 * (delay * 10ms/枚) = ms)
}

var palette = []color.Color{
	color.Black,
	color.RGBA{0xE6, 0x00, 0x12, 0xff},
	color.RGBA{0xF3, 0x98, 0x00, 0xff},
	color.RGBA{0xFF, 0xF1, 0x00, 0xff},
	color.RGBA{0x00, 0x99, 0x44, 0xff},
	color.RGBA{0x00, 0x68, 0xB7, 0xff},
	color.RGBA{0x1D, 0x20, 0x88, 0xff},
	color.RGBA{0x92, 0x07, 0x83, 0xff},
}

const (
	blackIndex  = 0
	redIndex    = 1
	orangeIndex = 2
	yellowIndex = 3
	greenIndex  = 4
	blueIndex   = 5
	indigoIndex = 6
	purpleIndex = 7
)

func defaultLissajousParameters() lissajousParameters {
	return lissajousParameters{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
}

func main() {
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())

	param := defaultLissajousParameters()

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	log.Printf("query: %v", r.Form)

	for k, v := range r.Form {
		switch k {
		case "cycles":
			n, err := strconv.Atoi(v[len(v)-1])
			if err != nil || !(n >= 1 && n <= 100) {
				break
			}
			param.cycles = n
		case "res":
			n, err := strconv.ParseFloat(v[len(v)-1], 64)
			if err != nil || !(n > 0.0 && n <= 1.0) {
				break
			}
			param.res = n
		case "size":
			n, err := strconv.Atoi(v[len(v)-1])
			if err != nil || !(n >= 32 && n <= 1024) {
				break
			}
			param.size = n
		case "nframes":
			n, err := strconv.Atoi(v[len(v)-1])
			if err != nil || !(n >= 32 && n <= 1024) {
				break
			}
			param.nframes = n
		case "delay":
			n, err := strconv.Atoi(v[len(v)-1])
			if err != nil || !(n >= 1 && n <= 1024) {
				break
			}
			param.delay = n
		default:
			log.Printf("Unknown parameter: %s", k)
		}
	}

	lissajous(w, param)
}

// 頭回ってないので全部コメントつけていく
// lissajous はリサージュ曲線をGIFとして出力します
func lissajous(out io.Writer, param lissajousParameters) {
	var (
		cycles  = param.cycles
		res     = param.res
		size    = param.size
		nframes = param.nframes
		delay   = param.delay
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)                                    // -1 ~ 1
			y := math.Sin(t*freq + phase)                       // -1 ~ 1
			color := uint8(1 + 7*t/(float64(cycles)*2*math.Pi)) // 周期に基づく(なので、線の色が変わっていく感じになる)
			// color := uint8(1 + 7*math.Abs(x))           // xに基づく(なので、縞々になる)
			// color := uint8(1 + 7*float64(i)/nframes)    // フレームに基づく(なので、時間が立つと変わっていく)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), color)
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
