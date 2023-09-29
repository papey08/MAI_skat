package chich_ast

import (
	"chich/internal/chich_compiler"
)

type parseFn func(string, string) (Node, []error)

type Node interface {
	String() string
	chich_compiler.Compilable
}
