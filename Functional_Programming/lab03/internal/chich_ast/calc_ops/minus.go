package calc_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Minus struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewMinus(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Minus{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (m Minus) String() string {
	return fmt.Sprintf("(%v - %v)", m.l, m.r)
}

func (m Minus) Compile(c *chich_compiler.Compiler) (position int, err error) {

	if position, err = m.l.Compile(c); err != nil {
		return
	}
	if position, err = m.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpSub)
	c.Bookmark(m.pos)
	return
}

func (m Minus) IsConstExpression() bool {
	return m.l.IsConstExpression() && m.r.IsConstExpression()
}
