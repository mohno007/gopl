package main

import (
	"os/exec"
	"testing"
)

func Test(t *testing.T) {
	out, err := exec.Command("go", "run", "dup.go", "input_a", "input_b").Output()
	if err != nil {
		t.Fatal("Command execution failed")
	}

	contained := `f 2
	input_a
	input_b
d       2
	input_a
	input_b
e       2
	input_a
	input_b
`
	if string(out) == contained {
		t.Fatalf("Output should contain '%s', but not contained: %s", contained, out)
	}
}
