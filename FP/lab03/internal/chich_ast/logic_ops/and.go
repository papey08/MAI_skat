package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type And struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewAnd(l, r chich_ast.Node, pos int) chich_ast.Node {
	return And{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (a And) String() string {
	return fmt.Sprintf("(%v && %v)", a.l, a.r)
}

func (a And) Compile(c *chich_compiler.Compiler) (p int, err error) {
	if p, err = a.l.Compile(c); err != nil {
		return
	}
	if p, err = a.r.Compile(c); err != nil {
		return
	}
	p = c.Emit(chich_code.OpAnd)
	c.Bookmark(a.pos)
	return
}

func (a And) IsConstExpression() bool {
	return a.l.IsConstExpression() && a.r.IsConstExpression()
}
