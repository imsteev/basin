package main

import (
	"testing"
)

const haiku = `as i drink some tea
i breathe in summer air
tinted with AC`

func TestBase16Encode(t *testing.T) {
	equals(t, "61732069206472696e6b20736f6d65207465610a69206272656174686520696e2073756d6d6572206169720a74696e7465642077697468204143", base16Encode([]byte(haiku)))
}

func equals(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
