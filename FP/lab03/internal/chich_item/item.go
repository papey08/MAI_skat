package chich_item

import "fmt"

type Item struct {
	Val string
	Typ Type
	Pos int
}

func (i Item) Is(t Type) bool {
	return i.Typ == t
}

func (i Item) String() string {
	if i.Is(Error) {
		return i.Val
	}
	if len(i.Val) > 10 {
		return fmt.Sprintf("%.10q...", i.Val)
	}
	return fmt.Sprintf("%q", i.Val)
}
