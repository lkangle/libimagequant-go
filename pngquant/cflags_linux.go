//go:build linux
// +build linux

package pngquant

/*
#cgo CFLAGS: -std=c99 -O3 -DNDEBUG -DUSE_SSE=1 -msse -fopenmp
#cgo LDFLAGS: -lm -fopenmp
*/
import "C"
