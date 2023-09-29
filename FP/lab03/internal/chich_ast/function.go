package chich_ast

import (
	"fmt"
	"strings"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
	"chich/internal/chich_obj"
)

type Function struct {
	body   Node
	Name   string
	pos    int
	params []Identifier
}

func NewFunction(params []Identifier, body Node, pos int) Node {
	return Function{
		params: params,
		body:   body,
		pos:    pos,
	}
}

func (f Function) String() string {
	var params []string

	for _, p := range f.params {
		params = append(params, p.String())
	}
	return fmt.Sprintf("fn(%s) { %v }", strings.Join(params, ", "), f.body)
}

func (f Function) Compile(c *chich_compiler.Compiler) (position int, err error) {
	c.EnterScope()

	if f.Name != "" {
		c.DefineFunctionName(f.Name)
	}

	for _, p := range f.params {
		c.Define(p.String())
	}

	if position, err = f.body.Compile(c); err != nil {
		return
	}

	if c.LastIs(chich_code.OpPop) {
		c.ReplaceLastPopWithReturn()
	}
	if !c.LastIs(chich_code.OpReturnValue) {
		c.Emit(chich_code.OpReturn)
	}

	freeSymbols := c.FreeSymbols
	nLocals := c.NumDefs
	ins, bookmarks := c.LeaveScope()

	for _, s := range freeSymbols {
		position = c.LoadSymbol(s)
	}

	fn := chich_obj.NewFunctionCompiled(ins, nLocals, len(f.params), bookmarks)
	position = c.Emit(chich_code.OpClosure, c.AddConstant(fn), len(freeSymbols))
	c.Bookmark(f.pos)
	return
}

func (f Function) IsConstExpression() bool {
	return false
}
