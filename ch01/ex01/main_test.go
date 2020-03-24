package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func Test(t *testing.T) {
	out, err := exec.Command("go", "run", "main.go", "1", "2", "3").Output()
	if err != nil {
		t.Fatal("Command execution failed")
	}

	contained := []byte("/main 1 2 3")
	if !bytes.Contains(out, contained) {
		t.Fatalf("Output should contain '%s', but not contained: %s", contained, out)
	}
}
