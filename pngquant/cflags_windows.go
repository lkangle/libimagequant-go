//go:build windows
// +build windows

package pngquant

/*
#cgo CFLAGS: -std=c99 -O3 -DNDEBUG -DUSE_SSE=1 -msse -fopenmp
#cgo LDFLAGS: -lm -fopenmp
*/
import "C"
