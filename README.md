# lzo1z

[![Go Reference](https://pkg.go.dev/badge/github.com/rhnvrm/lzo1z.svg)](https://pkg.go.dev/github.com/rhnvrm/lzo1z)
[![CI](https://github.com/rhnvrm/lzo1z/actions/workflows/ci.yml/badge.svg)](https://github.com/rhnvrm/lzo1z/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rhnvrm/lzo1z)](https://goreportcard.com/report/github.com/rhnvrm/lzo1z)
[![Made at Zerodha Tech](https://zerodha.tech/static/images/github-badge.svg)](https://zerodha.tech)

Pure Go implementation of LZO1Z compression and decompression.

## Overview

LZO1Z is a variant of the LZO1X compression algorithm used in real-time data feeds and other applications requiring fast compression/decompression.

This package provides both compression and decompression, fully compatible with the [liblzo2](http://www.oberhumer.com/opensource/lzo/) library.

### Features

- Pure Go - no CGO, no external dependencies
- Zero allocations per call
- ~420 MB/s compression, ~1 GB/s decompression
- Compatible with liblzo2
- Cross-compilation friendly

## Installation

```bash
go get github.com/rhnvrm/lzo1z
```

## Usage

### Compression

```go
input := []byte("Hello, World! Hello, World! Hello, World!")

// Allocate buffer for compressed data
compressed := make([]byte, lzo1z.MaxCompressedSize(len(input)))

// Compress
n, err := lzo1z.Compress(input, compressed)
if err != nil {
    log.Fatal(err)
}
compressed = compressed[:n]

fmt.Printf("%d bytes -> %d bytes\n", len(input), n)
// Output: 41 bytes -> 21 bytes
```

### Decompression

```go
// Allocate buffer (must know or estimate decompressed size)
output := make([]byte, expectedSize)

// Decompress
n, err := lzo1z.Decompress(compressed, output)
if err != nil {
    log.Fatal(err)
}
result := output[:n]
```

### Buffer Sizing

Use `MaxCompressedSize` to allocate compression buffers:

```go
bufSize := lzo1z.MaxCompressedSize(inputLen)
// Returns: inputLen + inputLen/16 + 64 + 3
```

For decompression, you must know or estimate the output size. LZO does not store the decompressed size in the stream.

## Performance

Benchmarks on Intel i7-1355U:

```
BenchmarkCompress-12         114550    10700 ns/op    420 MB/s    0 B/op    0 allocs/op
BenchmarkDecompress-12     13102768       94 ns/op   1060 MB/s    0 B/op    0 allocs/op
```

### Compression Ratios

| Input | Size | Compressed | Ratio |
|-------|------|------------|-------|
| Repeated "A" | 40 B | 9 B | 4.4x |
| "ABCD" x 100 | 400 B | 16 B | 25x |
| English text | 17 KB | 322 B | 53x |
| Random bytes | 256 B | 261 B | 0.98x |

## LZO1Z vs LZO1X

LZO1Z differs from LZO1X in offset encoding:

| Aspect | LZO1Z | LZO1X |
|--------|-------|-------|
| Offset encoding | `(b0 << 6) + (b1 >> 2)` | `(b0 >> 2) + (b1 << 6)` |
| M2 offset reuse | Yes | No |
| M2_MAX_OFFSET | 1792 | 2048 |

These differences mean LZO1X and LZO1Z are **not** compatible.

## Limitations

- **Buffer sizing** - caller must provide appropriately sized buffers
- **No streaming** - data must fit in memory

## Testing

```bash
go test -v ./...           # Run tests
go test -bench=. -benchmem # Run benchmarks
```

Test vectors are verified against liblzo2 for both compression and decompression.

## Credits

Based on the [LZO algorithm](http://www.oberhumer.com/opensource/lzo/) by Markus Franz Xaver Johannes Oberhumer.

## License

MIT License - see [LICENSE](LICENSE).
