package main

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	out := os.Getenv("PORT")
	t.Error(out)
}
