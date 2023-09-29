package calc_ops

import (
	"chich/internal/chich_ast"
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Mod struct {
	l   chich_ast.Node
	r   chich_ast.Node
	pos int
}

func NewMod(l, r chich_ast.Node, pos int) chich_ast.Node {
	return Mod{
		l:   l,
		r:   r,
		pos: pos,
	}
}

func (m Mod) String() string {
	return fmt.Sprintf("(%v %% %v)", m.l, m.r)
}

func (m Mod) Compile(c *chich_compiler.Compiler) (position int, err error) {

	if position, err = m.l.Compile(c); err != nil {
		return
	}
	if position, err = m.r.Compile(c); err != nil {
		return
	}
	position = c.Emit(chich_code.OpMod)
	c.Bookmark(m.pos)
	return
}

func (m Mod) IsConstExpression() bool {
	return m.l.IsConstExpression() && m.r.IsConstExpression()
}
