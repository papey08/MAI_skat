package chich_ast

import (
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Assign struct {
	l   Node
	r   Node
	pos int
}

func NewAssign(l, r Node, pos int) Node {
	return Assign{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (a Assign) String() string {
	return fmt.Sprintf("(%v = %v)", a.l, a.r)
}

func (a Assign) Compile(c *chich_compiler.Compiler) (p int, err error) {
	defer c.Bookmark(p)

	switch left := a.l.(type) {
	case Identifier:
		symbol := c.Define(left.String())
		if p, err = a.r.Compile(c); err != nil {
			return
		}

		if symbol.Scope == chich_compiler.GlobalScope {
			p = c.Emit(chich_code.OpSetGlobal, symbol.Index)
			c.Bookmark(a.pos)
			return
		} else {
			p = c.Emit(chich_code.OpSetLocal, symbol.Index)
			c.Bookmark(a.pos)
			return
		}

	default:
		return 0, fmt.Errorf("cannot assign to literal")
	}
}

func (a Assign) IsConstExpression() bool {
	return false
}
