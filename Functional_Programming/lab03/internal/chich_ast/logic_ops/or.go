package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Or struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewOr(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Or{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (o Or) String() string {
	return fmt.Sprintf("(%v || %v)", o.l, o.r)
}

func (o Or) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = o.l.Compile(c); err != nil {
		return
	}
	if position, err = o.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpOr)
	c.Bookmark(o.pos)
	return
}

func (o Or) IsConstExpression() bool {
	return o.l.IsConstExpression() && o.r.IsConstExpression()
}
