package chich_obj

import (
	"fmt"
	"strings"

	"chich/internal/chich_code"
	"chich/internal/chich_err"
)

type Function struct {
	Body   any
	Env    *Env
	Params []string
}

func (f Function) Type() Type {
	return FunctionType
}

func (f Function) String() string {
	return fmt.Sprintf("fn(%s) { %v }", strings.Join(f.Params, ", "), f.Body)
}

type CompiledFunction struct {
	Instructions chich_code.Instructions
	NumLocals    int
	NumParams    int
	Bookmarks    []chich_err.Bookmark
}

func NewFunctionCompiled(i chich_code.Instructions, nLocals, nParams int, bookmarks []chich_err.Bookmark) Object {
	return &CompiledFunction{
		Instructions: i,
		NumLocals:    nLocals,
		NumParams:    nParams,
		Bookmarks:    bookmarks,
	}
}

func (c CompiledFunction) Type() Type {
	return FunctionType
}

func (c CompiledFunction) String() string {
	return "<compiled function>"
}
