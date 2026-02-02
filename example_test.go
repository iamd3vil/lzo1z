package lzo1z_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/rhnvrm/lzo1z"
)

func Example() {
	// Compress some data
	input := []byte("Hello, World! Hello, World! Hello, World!")

	compressed := make([]byte, lzo1z.MaxCompressedSize(len(input)))
	compLen, err := lzo1z.Compress(input, compressed)
	if err != nil {
		log.Fatal(err)
	}
	compressed = compressed[:compLen]

	// Decompress it back
	output := make([]byte, len(input)+100)
	decompLen, err := lzo1z.Decompress(compressed, output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Original: %d bytes\n", len(input))
	fmt.Printf("Compressed: %d bytes\n", compLen)
	fmt.Printf("Decompressed: %s\n", string(output[:decompLen]))
	// Output:
	// Original: 41 bytes
	// Compressed: 21 bytes
	// Decompressed: Hello, World! Hello, World! Hello, World!
}

func ExampleDecompress() {
	// Compressed "AAAAAAAAAA" (10 repeated 'A' characters)
	compressed := []byte{0x1b, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x11, 0x00, 0x00}

	output := make([]byte, 20)
	n, err := lzo1z.Decompress(compressed, output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decompressed %d bytes: %s\n", n, string(output[:n]))
	// Output: Decompressed 10 bytes: AAAAAAAAAA
}

func ExampleDecompress_repeated() {
	// Highly compressed data: 40 repeated 'A' characters
	// Demonstrates LZO1Z's match copying
	compressed := []byte{0x12, 0x41, 0x20, 0x06, 0x00, 0x00, 0x11, 0x00, 0x00}

	output := make([]byte, 50)
	n, err := lzo1z.Decompress(compressed, output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Compressed: %d bytes -> Decompressed: %d bytes\n", len(compressed), n)
	fmt.Printf("Ratio: %.1fx\n", float64(n)/float64(len(compressed)))
	// Output:
	// Compressed: 9 bytes -> Decompressed: 40 bytes
	// Ratio: 4.4x
}

func ExampleCompress() {
	// Compress repetitive data
	input := bytes.Repeat([]byte("ABCD"), 100) // 400 bytes

	compressed := make([]byte, lzo1z.MaxCompressedSize(len(input)))
	n, err := lzo1z.Compress(input, compressed)
	if err != nil {
		log.Fatal(err)
	}

	ratio := float64(len(input)) / float64(n)
	fmt.Printf("Input: %d bytes, Compressed: %d bytes, Ratio: %.1fx\n", len(input), n, ratio)
	// Output: Input: 400 bytes, Compressed: 16 bytes, Ratio: 25.0x
}

func ExampleMaxCompressedSize() {
	inputSize := 1000

	// Always allocate enough space for worst-case compression
	bufferSize := lzo1z.MaxCompressedSize(inputSize)
	fmt.Printf("For %d byte input, allocate %d byte buffer\n", inputSize, bufferSize)
	// Output: For 1000 byte input, allocate 1129 byte buffer
}
