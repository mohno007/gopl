package main

import (
	"testing"
)

func TestCommaBuffer(t *testing.T) {
	tc := [...]struct {
		input    string
		expected string
	}{
		{"", ""},
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1234"},
		{"+1234", "+1234"},
		{"-12.12", "-12.12"},
		{"-123.123", "-123.123"},
		{"-1234.1234", "-1,234.1234"},
		{"-12345.12345", "-12,345.12345"},
		{"-123456.123456", "-123,456.123456"},
		{"-1234567.123456", "-1,234,567.123456"},
		{"-12345678.123456", "-12,345,678.123456"},
	}

	for _, c := range tc {
		result := commaBuffer(c.input)
		if result != c.expected {
			t.Fatalf("%s (expected %s)", result, c.expected)
		}
	}
}
