package chich_vm

import (
	"chich/internal/chich_code"
	"chich/internal/chich_obj"
)

type Frame struct {
	cl          *chich_obj.Closure
	ip          int
	basePointer int
}

func NewFrame(cl *chich_obj.Closure, basePointer int) *Frame {
	return &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}
}

func (f *Frame) Instructions() chich_code.Instructions {
	return f.cl.Fn.Instructions
}
