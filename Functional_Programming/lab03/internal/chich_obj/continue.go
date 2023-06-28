package chich_obj

type Continue struct{}

func (c Continue) String() string {
	return "continue"
}

func (c Continue) Type() Type {
	return ContinueType
}
