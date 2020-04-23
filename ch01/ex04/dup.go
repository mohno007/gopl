package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// リファレンスになるので、書き換え可能
	// 配列やマップは参照になる, makeで作られるのは参照
	// C++のnewみたいなやつだと思っておけば良さそう
	// { "": [ "", "", "" ] }
	counts := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, files := range counts {
		n := len(files)
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
			for _, f := range files {
				fmt.Printf("\t%s\n", f)
			}
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	// bufio.Scannerだけでなく、ioutil.ReadFile(ファイルのバイト列を読み込む)を使うという手もある
	input := bufio.NewScanner(f)
	for input.Scan() {
		files := counts[input.Text()]
		// ゼロ値はnil。nilに対してもappend関数は利用できる
		counts[input.Text()] = append(files, f.Name())
	}
	// 注意: input.Err() からのエラーの可能性を無視している
	// -> input.Scan()がfalseのとき、input.Err() にエラー情報が格納される
}
