package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/lkangle/libimagequant-go/pngquant"
)

func TestMain(t *testing.T) {
	source, err := os.OpenFile("./testdata/head.png", os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stat, _ := source.Stat()
	img, err := png.Decode(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	output1, err := pngquant.Compress(img, 98, pngquant.SPEED_DEFAULT, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, output1); err != nil {
		fmt.Println(err.Error())
		return
	}
	outputSize1 := int64(len(buf.Bytes()))

	fmt.Printf("outputSize1 %d, origin size %d\n", outputSize1, stat.Size())

	os.WriteFile("./temp.png", buf.Bytes(), os.ModePerm)
}
