package calc_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Plus struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewPlus(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Plus{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (p Plus) String() string {
	return fmt.Sprintf("(%v + %v)", p.l, p.r)
}

func (p Plus) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = p.l.Compile(c); err != nil {
		return
	}
	if position, err = p.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpAdd)
	c.Bookmark(p.pos)
	return
}

func (p Plus) IsConstExpression() bool {
	return p.l.IsConstExpression() && p.r.IsConstExpression()
}
