package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	//	"github.com/mohno007/practice-gopl/ch02/ex02/filenconv"
)

// Feet は における長さを保持する型です
// 1フィート(フット)の長さは国際フィートに基づきます
type Feet float64

// Meter は 国際メートル法における長さを保持する型です
// 国際メートル法の物理的な長さはSI単位系の定義に従います
type Meter float64

// FootM は、1フィート(フット) の 国際メートル法での長さ (単位: メートル)
const FootM Meter = Meter(0.3048)

func (f Feet) ToMeter() Meter   { return Meter(float64(f) * float64(FootM)) }
func (f Feet) ToString() string { return fmt.Sprintf("%g ft", f) }

func (m Meter) ToFeet() Feet     { return Feet(float64(m / FootM)) }
func (m Meter) ToString() string { return fmt.Sprintf("%g m", m) }

func main() {
	if len(os.Args) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, ": %v\n", err)
			os.Exit(1)
		}

		t, err := strconv.ParseFloat(string(b), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		printValues(t)
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		printValues(t)
	}
}

func printValues(t float64) {
	m := Meter(t)
	f := Feet(t)

	fmt.Printf("%f = %f, %f = %f\n", f, f.ToMeter(), m, m.ToFeet())
}
