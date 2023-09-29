package calc_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Times struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewTimes(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Times{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (t Times) String() string {
	return fmt.Sprintf("(%v * %v)", t.l, t.r)
}

func (t Times) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = t.l.Compile(c); err != nil {
		return
	}
	if position, err = t.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpMul)
	c.Bookmark(t.pos)
	return
}

func (t Times) IsConstExpression() bool {
	return t.l.IsConstExpression() && t.r.IsConstExpression()
}
