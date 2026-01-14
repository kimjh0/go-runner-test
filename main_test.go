package main

import (
	"testing"

	"github.com/valyala/gozstd"
)

func TestCompressDecompress(t *testing.T) {
	original := "Hello, World!"

	compressed := gozstd.Compress(nil, []byte(original))
	if len(compressed) == 0 {
		t.Fatal("compressed data is empty")
	}

	decompressed, err := gozstd.Decompress(nil, compressed)
	if err != nil {
		t.Fatalf("decompress failed: %v", err)
	}

	if string(decompressed) != original {
		t.Fatalf("expected %q, got %q", original, string(decompressed))
	}
}
