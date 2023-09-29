package chich_ast

import (
	"strconv"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
	"chich/internal/chich_obj"
)

type Integer int64

func NewInteger(i int64) Node {
	return Integer(i)
}

func (i Integer) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Integer) Compile(c *chich_compiler.Compiler) (position int, err error) {
	return c.Emit(chich_code.OpConstant, c.AddConstant(chich_obj.Integer(i))), nil
}

func (i Integer) IsConstExpression() bool {
	return true
}
