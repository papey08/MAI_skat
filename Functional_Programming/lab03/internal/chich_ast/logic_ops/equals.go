package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Equals struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewEquals(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Equals{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (e Equals) String() string {
	return fmt.Sprintf("(%v == %v)", e.l, e.r)
}

func (e Equals) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = e.l.Compile(c); err != nil {
		return
	}
	if position, err = e.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpEqual)
	c.Bookmark(e.pos)
	return
}

func (e Equals) IsConstExpression() bool {
	return e.l.IsConstExpression() && e.r.IsConstExpression()
}
