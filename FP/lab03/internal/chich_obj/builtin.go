package chich_obj

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	Stdout io.Writer = os.Stdout
)

type Builtin func(args ...Object) Object

func (b Builtin) Type() Type {
	return BuiltinType
}

func (b Builtin) String() string {
	return "<builtin function>"
}

type BuiltinImpl struct {
	Builtin Builtin
	Name    string
}

var Builtins = []BuiltinImpl{
	{
		Name: "len",
		Builtin: func(args ...Object) Object {
			if l := len(args); l != 1 {
				return NewError("len: wrong number of arguments, expected 1, got %d", l)
			}

			switch o := Unwrap(args[0]).(type) {
			case List:
				return Integer(len(o))
			case String:
				return Integer(len(o))
			default:
				return NewError("len: object of type %q has no length", o.Type())
			}
		},
	},
	{
		Name: "println",
		Builtin: func(args ...Object) Object {
			_, _ = fmt.Fprintln(Stdout, toAnySlice(args)...)
			return NullObj
		},
	},
	{
		Name: "input",
		Builtin: func(args ...Object) Object {
			var tmp string

			args = UnwrapAll(args)
			switch l := len(args); l {
			case 0:
				_, _ = fmt.Scanln(&tmp)

			case 1:
				fmt.Print(args[0])
				_, _ = fmt.Scanln(&tmp)

			default:
				return NewError("input: wrong number of arguments, expected 1, got %d", l)
			}
			return NewString(tmp)
		},
	},
	{
		Name: "string",
		Builtin: func(args ...Object) Object {
			if len(args) == 0 {
				return NewError("string: no argument provided")
			}

			args = UnwrapAll(args)
			return NewString(fmt.Sprint(toAnySlice(args)...))
		},
	},
	{
		Name: "error",
		Builtin: func(args ...Object) Object {
			return NewError(fmt.Sprint(toAnySlice(args)...))
		},
	},
	{
		Name: "int",
		Builtin: func(args ...Object) Object {
			if l := len(args); l != 1 {
				return NewError("int: wrong number of arguments, expected 1, got %d", l)
			}

			args = UnwrapAll(args)
			switch o := args[0].(type) {
			case Integer:
				return o

			case String:
				if a, err := strconv.ParseInt(string(o), 10, 64); err == nil {
					return Integer(a)
				}
				return NewError("%v is not a number", args[0])

			default:
				return NewError("%v is not a number", args[0])
			}
		},
	},
	{
		Name: "append",
		Builtin: func(args ...Object) Object {
			if len(args) == 0 {
				return NewError("append: no argument provided")
			}

			args = UnwrapAll(args)
			lst, ok := args[0].(List)
			if !ok {
				return NewError("append: first argument must be a list")
			}

			if len(args) > 1 {
				return append(lst, args[1:]...)
			}
			return lst
		},
	},
	{
		Name: "push",
		Builtin: func(args ...Object) Object {
			if len(args) == 0 {
				return NewError("push: no argument provided")
			}

			args = UnwrapAll(args)
			lst, ok := args[0].(List)
			if !ok {
				return NewError("push: first argument must be a list")
			}

			if len(args) > 1 {
				var tmp List

				for i := len(args) - 1; i > 0; i-- {
					tmp = append(tmp, args[i])
				}

				return append(tmp, lst...)
			}
			return lst
		},
	},
}

func UnwrapAll(a []Object) []Object {
	for i, o := range a {
		a[i] = Unwrap(o)
	}
	return a
}

func toAnySlice(args []Object) []any {
	var ret = make([]any, len(args))
	for i, a := range args {
		ret[i] = a
	}
	return ret
}
