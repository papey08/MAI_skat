package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type LessEq struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewLessEq(l, r chich_ast.Node, pos int) chich_ast.Node {
	return LessEq{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (l LessEq) String() string {
	return fmt.Sprintf("(%v <= %v)", l.l, l.r)
}

func (l LessEq) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = l.r.Compile(c); err != nil {
		return
	}
	if position, err = l.l.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpGreaterThanEqual)
	c.Bookmark(l.pos)
	return
}

func (l LessEq) IsConstExpression() bool {
	return l.l.IsConstExpression() && l.r.IsConstExpression()
}
