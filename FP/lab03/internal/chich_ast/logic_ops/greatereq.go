package logic_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type GreaterEq struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewGreaterEq(l, r chich_ast.Node, pos int) chich_ast.Node {
	return GreaterEq{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (g GreaterEq) String() string {
	return fmt.Sprintf("(%v >= %v)", g.l, g.r)
}

func (g GreaterEq) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = g.l.Compile(c); err != nil {
		return
	}
	if position, err = g.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpGreaterThanEqual)
	c.Bookmark(g.pos)
	return
}

func (g GreaterEq) IsConstExpression() bool {
	return g.l.IsConstExpression() && g.r.IsConstExpression()
}
