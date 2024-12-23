package pngquant

import (
	"unsafe"

	"github.com/pkg/errors"
)

/*
#include "libimagequant.h"
*/
import "C"

type Image struct {
	p        *C.struct_liq_image
	w, h     int
	released bool
}

// Callers MUST call Release() on the returned object to free memory.
func NewImage(attr *Attributes, rgba32data []byte, width, height int, gamma float64) (*Image, error) {
	pImg := C.liq_image_create_rgba(attr.p, unsafe.Pointer(C.CString(string(rgba32data))), C.int(width), C.int(height), C.double(gamma))
	if pImg == nil {
		return nil, errors.New("Failed to create image (invalid argument)")
	}

	return &Image{
		p:        pImg,
		w:        width,
		h:        height,
		released: false,
	}, nil
}

// Free memory. Callers must not use i object after Release has been called.
func (i *Image) Release() {
	C.liq_image_destroy(i.p)
	i.released = true
}

func (i *Image) Quantize(attr *Attributes) (*Result, error) {
	res := Result{
		im: i,
	}
	r := C.liq_quantize_image(attr.p, i.p)
	if r == nil {
		return nil, ErrInvalidPointer
	}

	res.p = r
	return &res, nil
}
