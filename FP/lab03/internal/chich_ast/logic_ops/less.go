package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Less struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewLess(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Less{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (l Less) String() string {
	return fmt.Sprintf("(%v < %v)", l.l, l.r)
}

func (l Less) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = l.r.Compile(c); err != nil {
		return
	}
	if position, err = l.l.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpGreaterThan)
	c.Bookmark(l.pos)
	return
}

func (l Less) IsConstExpression() bool {
	return l.l.IsConstExpression() && l.r.IsConstExpression()
}
