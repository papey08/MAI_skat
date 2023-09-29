package chich_ast

import (
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type IfExpr struct {
	cond   Node
	body   Node
	altern Node
	pos    int
}

func NewIfExpr(cond, body, alt Node, pos int) Node {
	return IfExpr{
		cond:   cond,
		body:   body,
		altern: alt,
		pos:    pos,
	}
}

func (i IfExpr) String() string {
	if i.altern != nil {
		return fmt.Sprintf("if %v { %v } else { %v }", i.cond, i.body, i.altern)
	}
	return fmt.Sprintf("if %v { %v }", i.cond, i.body)
}

func (i IfExpr) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if position, err = i.cond.Compile(c); err != nil {
		return
	}
	jumpNotTruthyPos := c.Emit(chich_code.OpJumpNotTruthy, 9999)
	if position, err = i.body.Compile(c); err != nil {
		return
	}

	if c.LastIs(chich_code.OpPop) {
		c.RemoveLast()
	}

	jumpPos := c.Emit(chich_code.OpJump, 9999)
	c.ReplaceOperand(jumpNotTruthyPos, c.Pos())

	if i.altern == nil {
		c.Emit(chich_code.OpNull)
	} else {
		if position, err = i.altern.Compile(c); err != nil {
			return
		}

		if c.LastIs(chich_code.OpPop) {
			c.RemoveLast()
		}
	}

	c.ReplaceOperand(jumpPos, c.Pos())
	return c.Pos(), nil
}

func (i IfExpr) IsConstExpression() bool {
	return false
}
