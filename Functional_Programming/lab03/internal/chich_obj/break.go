package chich_obj

type Break struct{}

func (b Break) String() string {
	return "break"
}

func (b Break) Type() Type {
	return BreakType
}
