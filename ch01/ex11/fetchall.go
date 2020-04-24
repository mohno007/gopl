package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // とりあえず読み込む。読込中にエラーがないかどうかも検証する。
	resp.Body.Close()                                 // 資源をリークさせない

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
