package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type NotEquals struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewNotEquals(l, r chich_ast.Node, pos int) chich_ast.Node {
	return NotEquals{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (n NotEquals) String() string {
	return fmt.Sprintf("(%v != %v)", n.l, n.r)
}

func (n NotEquals) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = n.l.Compile(c); err != nil {
		return
	}
	if position, err = n.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpNotEqual)
	c.Bookmark(n.pos)
	return
}

func (n NotEquals) IsConstExpression() bool {
	return n.l.IsConstExpression() && n.r.IsConstExpression()
}
