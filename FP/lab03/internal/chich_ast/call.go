package chich_ast

import (
	"fmt"
	"strings"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Call struct {
	Fn   Node
	Args []Node
	pos  int
}

func NewCall(fn Node, args []Node, pos int) Node {
	return Call{
		Fn:   fn,
		Args: args,
		pos:  pos,
	}
}

func (c Call) String() string {
	var args []string

	for _, a := range c.Args {
		args = append(args, a.String())
	}
	return fmt.Sprintf("%v(%s)", c.Fn, strings.Join(args, ", "))
}

func (c Call) Compile(comp *chich_compiler.Compiler) (p int, err error) {
	if p, err = c.Fn.Compile(comp); err != nil {
		return
	}

	for _, a := range c.Args {
		if p, err = a.Compile(comp); err != nil {
			return
		}
	}

	p = comp.Emit(chich_code.OpCall, len(c.Args))
	comp.Bookmark(c.pos)
	return
}

func (c Call) IsConstExpression() bool {
	return false
}
