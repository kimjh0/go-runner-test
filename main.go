package main

import (
	"fmt"

	"github.com/valyala/gozstd"
)

func main() {
	message := "Hello, World!"
	fmt.Println(message)

	compressed := gozstd.Compress(nil, []byte(message))
	fmt.Printf("Original:   %d bytes\n", len(message))
	fmt.Printf("Compressed: %d bytes\n", len(compressed))
	fmt.Printf("Compressed bytes: %v\n", compressed)

	decompressed, err := gozstd.Decompress(nil, compressed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decompressed: %s\n", string(decompressed))
}
