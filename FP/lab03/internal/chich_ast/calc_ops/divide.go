package calc_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Divide struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewDivide(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Divide{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (d Divide) String() string {
	return fmt.Sprintf("(%v / %v)", d.l, d.r)
}

func (d Divide) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = d.l.Compile(c); err != nil {
		return
	}
	if position, err = d.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpDiv)
	c.Bookmark(d.pos)
	return
}

func (d Divide) IsConstExpression() bool {
	return d.l.IsConstExpression() && d.r.IsConstExpression()
}
