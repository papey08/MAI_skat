package chich_compiler

import (
	"fmt"

	"chich/internal/chich_code"
	"chich/internal/chich_err"
	"chich/internal/chich_obj"
)

type Compilable interface {
	Compile(c *Compiler) (int, error)
	IsConstExpression() bool
}

type EmittedInst struct {
	Opcode   chich_code.Opcode
	Position int
}

type CompilationScope struct {
	instructions chich_code.Instructions
	lastInst     EmittedInst
	prevInst     EmittedInst
	bookmarks    []chich_err.Bookmark
}

type Compiler struct {
	constants   *[]chich_obj.Object
	scopes      []CompilationScope
	scopeIndex  int
	fileName    string
	fileContent string
	*SymbolTable
}

type Bytecode struct {
	Instructions chich_code.Instructions
	Constants    []chich_obj.Object
	Bookmarks    []chich_err.Bookmark
}

func New() *Compiler {
	var st = NewSymbolTable()

	for i, b := range chich_obj.Builtins {
		st.DefineBuiltin(i, b.Name)
	}

	return &Compiler{
		SymbolTable: st,
		scopes:      []CompilationScope{{}},
		constants:   &[]chich_obj.Object{},
	}
}

func (c *Compiler) AddConstant(o chich_obj.Object) int {
	*c.constants = append(*c.constants, o)
	return len(*c.constants) - 1
}

func (c *Compiler) AddInstruction(ins []byte) int {
	var posNewInstruction = len(c.scopes[c.scopeIndex].instructions)

	c.scopes[c.scopeIndex].instructions = append(c.scopes[c.scopeIndex].instructions, ins...)
	return posNewInstruction
}

func (c *Compiler) setLastInstruction(op chich_code.Opcode, pos int) {
	prev := c.scopes[c.scopeIndex].lastInst
	last := EmittedInst{op, pos}
	c.scopes[c.scopeIndex].prevInst = prev
	c.scopes[c.scopeIndex].lastInst = last
}

func (c *Compiler) Emit(opcode chich_code.Opcode, operands ...int) int {
	ins := chich_code.Make(opcode, operands...)
	pos := c.AddInstruction(ins)
	c.setLastInstruction(opcode, pos)
	return pos
}

func (c *Compiler) LastIs(op chich_code.Opcode) bool {
	if len(c.scopes[c.scopeIndex].instructions) == 0 {
		return false
	}
	return c.scopes[c.scopeIndex].lastInst.Opcode == op
}

func (c *Compiler) RemoveLast() {
	last := c.scopes[c.scopeIndex].lastInst
	prev := c.scopes[c.scopeIndex].prevInst

	old := c.scopes[c.scopeIndex].instructions
	next := old[:last.Position]

	c.scopes[c.scopeIndex].instructions = next
	c.scopes[c.scopeIndex].lastInst = prev
}

func (c *Compiler) replaceInstruction(pos int, newInst []byte) {
	ins := c.scopes[c.scopeIndex].instructions

	for i := 0; i < len(newInst); i++ {
		ins[pos+i] = newInst[i]
	}
}

func (c *Compiler) ReplaceOperand(opPos, operand int) {
	op := chich_code.Opcode(c.scopes[c.scopeIndex].instructions[opPos])
	newInst := chich_code.Make(op, operand)
	c.replaceInstruction(opPos, newInst)
}

func (c *Compiler) ReplaceLastPopWithReturn() {
	lastPos := c.scopes[c.scopeIndex].lastInst.Position
	c.replaceInstruction(lastPos, chich_code.Make(chich_code.OpReturnValue))
	c.scopes[c.scopeIndex].lastInst.Opcode = chich_code.OpReturnValue
}

func (c *Compiler) EnterScope() {
	c.scopes = append(c.scopes, CompilationScope{})
	c.scopeIndex++
	c.SymbolTable = NewEnclosedSymbolTable(c.SymbolTable)
}

func (c *Compiler) LeaveScope() (chich_code.Instructions, []chich_err.Bookmark) {
	ins := c.scopes[c.scopeIndex].instructions
	bookmarks := c.scopes[c.scopeIndex].bookmarks
	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--
	c.SymbolTable = c.SymbolTable.outer

	return ins, bookmarks
}

func (c *Compiler) Pos() int {
	return len(c.scopes[c.scopeIndex].instructions)
}

func (c *Compiler) Bookmark(pos int) {
	if c.fileContent == "" {
		return
	}

	b := chich_err.NewBookmark(c.fileContent, pos, c.Pos())
	c.scopes[c.scopeIndex].bookmarks = append(c.scopes[c.scopeIndex].bookmarks, b)
}

func (c *Compiler) UnresolvedError(name string, pos int) error {
	if c.fileName == "" || c.fileContent == "" {
		return fmt.Errorf("undefined variable %s", name)
	}

	return chich_err.New(c.fileName, c.fileContent, pos, "undefined variable %s", name)
}

func (c *Compiler) Compile(node Compilable) error {
	_, err := node.Compile(c)
	return err
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.scopes[c.scopeIndex].instructions,
		Constants:    *c.constants,
		Bookmarks:    c.scopes[c.scopeIndex].bookmarks,
	}
}

func (c *Compiler) SetFileInfo(name, content string) {
	c.fileName = name
	c.fileContent = content
}

func (c *Compiler) LoadSymbol(s Symbol) int {
	switch s.Scope {
	case GlobalScope:
		return c.Emit(chich_code.OpGetGlobal, s.Index)
	case LocalScope:
		return c.Emit(chich_code.OpGetLocal, s.Index)
	case BuiltinScope:
		return c.Emit(chich_code.OpGetBuiltin, s.Index)
	case FreeScope:
		return c.Emit(chich_code.OpGetFree, s.Index)
	case FunctionScope:
		return c.Emit(chich_code.OpCurrentClosure)
	default:
		return 0
	}
}
