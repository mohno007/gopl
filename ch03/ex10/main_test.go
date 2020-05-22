package main

import (
	"fmt"
	"testing"
)

func TestCommaBuffer(t *testing.T) {
	v := fmt.Sprintf("%d", 12345678)
	result := commaBuffer(v)
	if result != "12,345,678" {
		t.Fatalf("%s", result)
	}
}
