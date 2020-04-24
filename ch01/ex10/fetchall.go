package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	for i, url := range os.Args[1:] {
		go fetch(url, i, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, n int, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}

	path := strings.ReplaceAll(url, "/", "_")
	path = strings.ReplaceAll(path, "\\", "_")
	path = fmt.Sprintf("%s_%d", path, n)
	output, err := os.Create(path)

	if err != nil {
		ch <- fmt.Sprintf("failed to create %s: %v", path, err)
	}

	nbytes, err := io.Copy(output, resp.Body) // とりあえず読み込む。読込中にエラーがないかどうかも検証する。
	resp.Body.Close()                         // 資源をリークさせない
	output.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
