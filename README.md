# libimagequant-go

![GitHub release](https://img.shields.io/github/tag/lkangle/libimagequant-go.svg?label=release)

A golang wrapper for [libimagequant](https://github.com/ImageOptim/libimagequant).

## Install

```shell
go get github.com/lkangle/libimagequant-go
```

## Usage

### High Level API

```
// compress png iamge
func Compress(img image.Image, quality int, speed int, floyd float32) (image.Image, error)
```

### Low Level API

[GoDoc](https://godoc.org/github.com/lkangle/libimagequant-go/pngquant)

## Example

```golang
package main

import (
	"fmt"
	"github.com/lkangle/libimagequant-go/pngquant"
	"image/png"
	"os"
)

func main() {
	source, _ := os.OpenFile("testdata/1.png", os.O_RDONLY, 0444)
	img, _ := png.Decode(source)
	quality := 98
	speed := 5 // 1~10
	output, err := pngquant.Compress(img, quality, speed, 1.)
}
```
