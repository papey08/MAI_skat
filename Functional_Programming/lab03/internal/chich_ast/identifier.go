package chich_ast

import (
	"chich/internal/chich_compiler"
)

type Identifier struct {
	name string
	pos  int
}

func NewIdentifier(name string, pos int) Identifier {
	return Identifier{
		name: name,
		pos:  pos,
	}
}

func (i Identifier) String() string {
	return i.name
}

func (i Identifier) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if symbol, ok := c.Resolve(i.name); ok {
		return c.LoadSymbol(symbol), nil
	}
	return 0, c.UnresolvedError(i.name, i.pos)
}

func (i Identifier) IsConstExpression() bool {
	return false
}
