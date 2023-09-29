package chich_vm

import (
	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
	"chich/internal/chich_err"
	"chich/internal/chich_obj"
	"fmt"
	"path/filepath"
)

type State struct {
	Symbols *chich_compiler.SymbolTable
	Consts  []chich_obj.Object
	Globals []chich_obj.Object
}

func NewState() *State {
	st := chich_compiler.NewSymbolTable()
	for i, builtin := range chich_obj.Builtins {
		st.DefineBuiltin(i, builtin.Name)
	}

	return &State{
		Consts:  []chich_obj.Object{},
		Globals: make([]chich_obj.Object, GlobalSize),
		Symbols: st,
	}
}

type VM struct {
	*State
	dir        string
	file       string
	stack      []chich_obj.Object
	frames     []*Frame
	localTable []bool
	sp         int
	frameIndex int
}

const (
	StackSize  = 2048
	GlobalSize = 65536
	MaxFrames  = 1024
)

var (
	True  = chich_obj.True
	False = chich_obj.False
	Null  = chich_obj.NullObj
)

func New(file string, bytecode *chich_compiler.Bytecode) *VM {
	vm := &VM{
		stack:      make([]chich_obj.Object, StackSize),
		frames:     make([]*Frame, MaxFrames),
		frameIndex: 1,
		localTable: make([]bool, GlobalSize),
		State:      NewState(),
	}

	vm.dir, vm.file = filepath.Split(file)
	vm.Consts = bytecode.Constants
	fn := &chich_obj.CompiledFunction{
		Instructions: bytecode.Instructions,
		Bookmarks:    bytecode.Bookmarks,
	}
	vm.frames[0] = NewFrame(&chich_obj.Closure{Fn: fn}, 0)
	return vm
}

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.frameIndex-1]
}

func (vm *VM) pushFrame(f *Frame) {
	vm.frames[vm.frameIndex] = f
	vm.frameIndex++
}

func (vm *VM) popFrame() *Frame {
	vm.frameIndex--
	return vm.frames[vm.frameIndex]
}

func (vm *VM) bookmark() chich_err.Bookmark {
	var (
		frame     = vm.currentFrame()
		offset    = frame.ip
		bookmarks = frame.cl.Fn.Bookmarks
	)

	if len(bookmarks) == 0 {
		return chich_err.Bookmark{}
	}

	prev := bookmarks[0]
	for _, cur := range bookmarks[1:] {
		if offset < prev.Offset {
			return prev
		} else if offset > prev.Offset && offset <= cur.Offset {
			return cur
		}
		prev = cur
	}
	return prev
}

func (vm *VM) errorf(s string, a ...any) error {
	return chich_err.NewFromBookmark(
		filepath.Join(vm.dir, vm.file),
		vm.bookmark(),
		s,
		a...,
	)
}

func (vm *VM) execAdd() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	switch {
	case chich_obj.AssertTypes(left, chich_obj.IntType) && chich_obj.AssertTypes(right, chich_obj.IntType):
		l := left.(chich_obj.Integer)
		r := right.(chich_obj.Integer)
		return vm.push(l + r)

	case chich_obj.AssertTypes(left, chich_obj.StringType) && chich_obj.AssertTypes(right, chich_obj.StringType):
		l := left.(chich_obj.String)
		r := right.(chich_obj.String)
		return vm.push(l + r)

	default:
		return vm.errorf("unsupported operator '+' for types %v and %v", left.Type(), right.Type())
	}
}

func (vm *VM) execSub() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	switch {
	case chich_obj.AssertTypes(left, chich_obj.IntType) && chich_obj.AssertTypes(right, chich_obj.IntType):
		l := left.(chich_obj.Integer)
		r := right.(chich_obj.Integer)
		return vm.push(l - r)

	default:
		return vm.errorf("unsupported operator '-' for types %v and %v", left.Type(), right.Type())
	}
}

func (vm *VM) execMul() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	switch {
	case chich_obj.AssertTypes(left, chich_obj.IntType) && chich_obj.AssertTypes(right, chich_obj.IntType):
		l := left.(chich_obj.Integer)
		r := right.(chich_obj.Integer)
		return vm.push(l * r)

	default:
		return vm.errorf("unsupported operator '*' for types %v and %v", left.Type(), right.Type())
	}
}

func (vm *VM) execDiv() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	if !chich_obj.AssertTypes(left, chich_obj.IntType) || !chich_obj.AssertTypes(right, chich_obj.IntType) {
		return fmt.Errorf("unsupported operator '/' for types %v and %v", left.Type(), right.Type())
	}

	l := left.(chich_obj.Integer)
	r := right.(chich_obj.Integer)
	return vm.push(l / r)
}

func (vm *VM) execMod() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	if !chich_obj.AssertTypes(left, chich_obj.IntType) || !chich_obj.AssertTypes(right, chich_obj.IntType) {
		return fmt.Errorf("unsupported operator '%%' for types %v and %v", left.Type(), right.Type())
	}

	l := left.(chich_obj.Integer)
	r := right.(chich_obj.Integer)

	if r == 0 {
		return vm.errorf("can't divide by 0")
	}
	return vm.push(l % r)
}

func (vm *VM) execEqual() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	switch {
	case chich_obj.AssertTypes(left, chich_obj.BoolType, chich_obj.NullType) || chich_obj.AssertTypes(right, chich_obj.BoolType, chich_obj.NullType):
		return vm.push(chich_obj.ParseBool(left == right))

	case chich_obj.AssertTypes(left, chich_obj.StringType) && chich_obj.AssertTypes(right, chich_obj.StringType):
		l := left.(chich_obj.String)
		r := right.(chich_obj.String)
		return vm.push(chich_obj.ParseBool(l == r))

	case chich_obj.AssertTypes(left, chich_obj.IntType) && chich_obj.AssertTypes(right, chich_obj.IntType):
		l := left.(chich_obj.Integer)
		r := right.(chich_obj.Integer)
		return vm.push(chich_obj.ParseBool(l == r))

	default:
		return vm.push(False)
	}
}

func (vm *VM) execNotEqual() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	switch {
	case chich_obj.AssertTypes(left, chich_obj.BoolType, chich_obj.NullType) || chich_obj.AssertTypes(right, chich_obj.BoolType, chich_obj.NullType):
		return vm.push(chich_obj.ParseBool(left != right))

	case chich_obj.AssertTypes(left, chich_obj.StringType) && chich_obj.AssertTypes(right, chich_obj.StringType):
		l := left.(chich_obj.String)
		r := right.(chich_obj.String)
		return vm.push(chich_obj.ParseBool(l != r))

	case chich_obj.AssertTypes(left, chich_obj.IntType) && chich_obj.AssertTypes(right, chich_obj.IntType):
		l := left.(chich_obj.Integer)
		r := right.(chich_obj.Integer)
		return vm.push(chich_obj.ParseBool(l != r))

	default:
		return vm.push(True)
	}
}

func (vm *VM) execAnd() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	return vm.push(chich_obj.ParseBool(chich_obj.IsTruthy(left) && chich_obj.IsTruthy(right)))
}

func (vm *VM) execOr() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	return vm.push(chich_obj.ParseBool(chich_obj.IsTruthy(left) || chich_obj.IsTruthy(right)))
}

func (vm *VM) execGreaterThan() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	switch {
	case chich_obj.AssertTypes(left, chich_obj.IntType) && chich_obj.AssertTypes(right, chich_obj.IntType):
		l := left.(chich_obj.Integer)
		r := right.(chich_obj.Integer)
		return vm.push(chich_obj.ParseBool(l > r))

	case chich_obj.AssertTypes(left, chich_obj.StringType) && chich_obj.AssertTypes(right, chich_obj.StringType):
		l := left.(chich_obj.String)
		r := right.(chich_obj.String)
		return vm.push(chich_obj.ParseBool(l > r))

	default:
		return vm.errorf("unsupported operator '>' for types %v and %v", left.Type(), right.Type())
	}
}

func (vm *VM) execGreaterThanEqual() error {
	var (
		right = chich_obj.Unwrap(vm.pop())
		left  = chich_obj.Unwrap(vm.pop())
	)

	switch {
	case chich_obj.AssertTypes(left, chich_obj.IntType) && chich_obj.AssertTypes(right, chich_obj.IntType):
		l := left.(chich_obj.Integer)
		r := right.(chich_obj.Integer)
		return vm.push(chich_obj.ParseBool(l >= r))

	case chich_obj.AssertTypes(left, chich_obj.StringType) && chich_obj.AssertTypes(right, chich_obj.StringType):
		l := left.(chich_obj.String)
		r := right.(chich_obj.String)
		return vm.push(chich_obj.ParseBool(l >= r))

	default:
		return vm.errorf("unsupported operator '>=' for types %v and %v", left.Type(), right.Type())
	}
}

func (vm *VM) execReturnValue() error {
	retVal := chich_obj.Unwrap(vm.pop())
	frame := vm.popFrame()
	vm.sp = frame.basePointer - 1

	return vm.push(retVal)
}

func (vm *VM) call(o chich_obj.Object, numArgs int) error {
	switch fn := chich_obj.Unwrap(o).(type) {
	case *chich_obj.Closure:
		return vm.callClosure(fn, numArgs)
	case chich_obj.Builtin:
		return vm.callBuiltin(fn, numArgs)
	default:
		return vm.errorf("calling non-function")
	}
}

func (vm *VM) execCall(numArgs int) error {
	return vm.call(vm.stack[vm.sp-1-numArgs], numArgs)
}

func (vm *VM) buildList(start, end int) chich_obj.Object {
	var elements = make([]chich_obj.Object, end-start)

	for i := start; i < end; i++ {
		elements[i-start] = vm.stack[i]
	}
	return chich_obj.NewList(elements...)
}

func (vm *VM) callClosure(cl *chich_obj.Closure, nargs int) error {
	if nargs != cl.Fn.NumParams {
		return vm.errorf("wrong number of arguments: expected %d, got %d", cl.Fn.NumParams, nargs)
	}

	frame := NewFrame(cl, vm.sp-nargs)
	vm.pushFrame(frame)
	vm.sp = frame.basePointer + cl.Fn.NumLocals
	return nil
}

func (vm *VM) callBuiltin(fn chich_obj.Builtin, nargs int) error {
	args := vm.stack[vm.sp-nargs : vm.sp]
	res := fn(args...)
	vm.sp = vm.sp - nargs - 1

	if res == nil {
		return vm.push(Null)
	}
	return vm.push(res)
}

func (vm *VM) pushClosure(constIdx, numFree int) error {
	constant := vm.Consts[constIdx]
	fn, ok := constant.(*chich_obj.CompiledFunction)
	if !ok {
		return vm.errorf("not a function: %+v", constant)
	}

	free := make([]chich_obj.Object, numFree)
	for i := 0; i < numFree; i++ {
		free[i] = vm.stack[vm.sp-numFree+i]
	}
	vm.sp = vm.sp - numFree
	return vm.push(&chich_obj.Closure{Fn: fn, Free: free})
}

func (vm *VM) Run() (err error) {
	var (
		ip  int
		ins chich_code.Instructions
		op  chich_code.Opcode
	)

	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	for vm.currentFrame().ip < len(vm.currentFrame().Instructions())-1 && err == nil {
		vm.currentFrame().ip++

		ip = vm.currentFrame().ip
		ins = vm.currentFrame().Instructions()
		op = chich_code.Opcode(ins[ip])

		switch op {
		case chich_code.OpConstant:
			constIndex := chich_code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2
			err = vm.push(vm.Consts[constIndex])

		case chich_code.OpJump:
			pos := int(chich_code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip = pos - 1

		case chich_code.OpJumpNotTruthy:
			pos := int(chich_code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip += 2

			if cond := chich_obj.Unwrap(vm.pop()); !chich_obj.IsTruthy(cond) {
				vm.currentFrame().ip = pos - 1
			}

		case chich_code.OpSetGlobal:
			globalIndex := chich_code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2
			vm.localTable[globalIndex] = true
			vm.Globals[globalIndex] = vm.peek()

		case chich_code.OpGetGlobal:
			globalIndex := chich_code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2
			err = vm.push(vm.Globals[globalIndex])

		case chich_code.OpGetLocal:
			localIndex := chich_code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			frame := vm.currentFrame()
			err = vm.push(vm.stack[frame.basePointer+int(localIndex)])

		case chich_code.OpList:
			nElements := int(chich_code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip += 2

			list := vm.buildList(vm.sp-nElements, vm.sp)
			vm.sp = vm.sp - nElements
			err = vm.push(list)

		case chich_code.OpCall:
			numArgs := chich_code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1
			err = vm.execCall(int(numArgs))

		case chich_code.OpGetBuiltin:
			idx := chich_code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1
			def := chich_obj.Builtins[idx]
			err = vm.push(def.Builtin)

		case chich_code.OpClosure:
			constIdx := chich_code.ReadUint16(ins[ip+1:])
			numFree := chich_code.ReadUint8(ins[ip+3:])
			vm.currentFrame().ip += 3
			err = vm.pushClosure(int(constIdx), int(numFree))

		case chich_code.OpReturnValue:
			err = vm.execReturnValue()

		case chich_code.OpNull:
			err = vm.push(Null)

		case chich_code.OpAdd:
			err = vm.execAdd()

		case chich_code.OpSub:
			err = vm.execSub()

		case chich_code.OpMul:
			err = vm.execMul()

		case chich_code.OpDiv:
			err = vm.execDiv()

		case chich_code.OpMod:
			err = vm.execMod()

		case chich_code.OpEqual:
			err = vm.execEqual()

		case chich_code.OpNotEqual:
			err = vm.execNotEqual()

		case chich_code.OpGreaterThan:
			err = vm.execGreaterThan()

		case chich_code.OpGreaterThanEqual:
			err = vm.execGreaterThanEqual()

		case chich_code.OpAnd:
			err = vm.execAnd()

		case chich_code.OpOr:
			err = vm.execOr()

		case chich_code.OpPop:
			vm.pop()
		}

	}
	return
}

func (vm *VM) push(o chich_obj.Object) error {
	if vm.sp >= StackSize {
		return vm.errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

func (vm *VM) pop() chich_obj.Object {
	vm.sp--
	o := vm.stack[vm.sp]
	return o
}

func (vm *VM) peek() chich_obj.Object {
	return vm.stack[vm.sp-1]
}
