package chich_obj

type Null struct{}

func NewNull() Object {
	return new(Null)
}

func (n Null) String() string {
	return "null"
}

func (n Null) Type() Type {
	return NullType
}
