package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Greater struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewGreater(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Greater{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (g Greater) String() string {
	return fmt.Sprintf("(%v > %v)", g.l, g.r)
}

func (g Greater) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = g.l.Compile(c); err != nil {
		return
	}
	if position, err = g.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpGreaterThan)
	c.Bookmark(g.pos)
	return
}

func (g Greater) IsConstExpression() bool {
	return g.l.IsConstExpression() && g.r.IsConstExpression()
}
