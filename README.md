# uncompr

A Go [io.ReadCloser](https://pkg.go.dev/io#ReadCloser) and [io.WriteCloser](https://pkg.go.dev/io#WriteCloser)
that automatically detects and compresses/decompresses a wide variety of compression formats.

## Supported Formats

* bzip2 (decompression only)
* gzip
* lz4
* xz
* zstd

## Usage

```go
package main

import (
  "fmt" 
  "io"
  "log"
  "os"

  "github.com/dpeckett/uncompr"
)

func main() {
  f, err := os.Open("hello.gz")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  r, err := uncompr.NewReader(f)
  if err != nil {
    log.Fatal(err)
  }
  defer r.Close()

  contents, err := io.ReadAll(r)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(contents))
}
```
