package chich_obj

type Object interface {
	Type() Type
	String() string
}

type Getter interface {
	Object() Object
}

type Setter interface {
	Set(Object) Object
}

type GetSetter interface {
	Object
	Getter
	Setter
}

type MapGetSetter interface {
	Get(string) (Object, bool)
	Set(string, Object) Object
}

type KeyHash struct {
	Type  Type
	Value uint64
}

type Hashable interface {
	KeyHash() KeyHash
}

type Type int

const (
	NullType Type = iota
	ErrorType
	IntType
	BoolType
	StringType
	ObjectType
	ReturnType
	FunctionType
	ClosureType
	BuiltinType
	ListType
	ContinueType
	BreakType
)

var (
	NullObj = NewNull()
	True    = NewBoolean(true)
	False   = NewBoolean(false)
)

func ParseBool(b bool) Object {
	if b {
		return True
	}
	return False
}

func Unwrap(o Object) Object {
	if g, ok := o.(Getter); ok {
		return g.Object()
	}
	return o
}

func AssertTypes(o Object, types ...Type) bool {
	for _, t := range types {
		if t == o.Type() {
			return true
		}
	}
	return false
}

func IsTruthy(o Object) bool {
	switch val := o.(type) {
	case *Boolean:
		return o == True
	case Integer:
		return val != 0
	case *Null:
		return false
	default:
		return true
	}
}
