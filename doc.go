// Package lzo1z implements LZO1Z compression and decompression in pure Go.
//
// LZO1Z is a variant of the LZO1X compression algorithm with different
// offset encoding, used primarily for real-time data compression.
// This package is compatible with the liblzo2 library.
//
// # Compression
//
// Compress data using the Compress function:
//
//	input := []byte("Hello, World!")
//	compressed := make([]byte, lzo1z.MaxCompressedSize(len(input)))
//
//	n, err := lzo1z.Compress(input, compressed)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	compressed = compressed[:n]
//
// Use MaxCompressedSize to determine the required buffer size for worst-case
// compression (incompressible data).
//
// # Decompression
//
// Decompress LZO1Z data:
//
//	output := make([]byte, expectedSize)
//
//	n, err := lzo1z.Decompress(compressed, output)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	result := output[:n]
//
// # Buffer Sizing
//
// The caller must provide appropriately sized buffers:
//   - For compression: use MaxCompressedSize(inputLen)
//   - For decompression: you must know or estimate the output size
//
// LZO does not store the decompressed size in the compressed stream,
// so the caller must track this separately.
//
// # Thread Safety
//
// Both Compress and Decompress are safe for concurrent use - they have
// no global state and perform zero allocations.
//
// # Performance
//
// The implementation achieves approximately 420 MB/s compression and
// 1 GB/s decompression on modern hardware with zero allocations.
package lzo1z
