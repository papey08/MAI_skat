package chich_obj

import "reflect"

type NativeStruct struct {
	s any
}

func NewNativeStruct(s any) Object {
	return &NativeStruct{s}
}

func (n *NativeStruct) Get(name string) (Object, bool) {
	return Builtin(func(a ...Object) (o Object) {
		defer func() {
			if err := recover(); err != nil {
				o = NewError("%v", err)
			}
		}()

		f := reflect.ValueOf(n.s).MethodByName(name)
		arguments, err := args(f.Type(), a...)
		if err != nil {
			return NewError(err.Error())
		}

		return multiplex(f.Call(arguments))
	}), true
}

func (n *NativeStruct) Set(string, Object) Object {
	return NewError("cannot assign to native struct")
}

func (n *NativeStruct) String() string {
	return "<native struct>"
}

func (n *NativeStruct) Type() Type {
	return ObjectType
}
