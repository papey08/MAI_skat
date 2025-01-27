package chich_obj

import (
	"hash/fnv"
	"strconv"
)

type String string

func NewString(s string) Object {
	return String(s)
}

func (s String) Type() Type {
	return StringType
}

func (s String) String() string {
	return string(s)
}

func (s String) Val() string {
	return string(s)
}

func (s String) Quoted() string {
	return strconv.Quote(string(s))
}

func (s String) KeyHash() KeyHash {
	var h = fnv.New64a()
	_, _ = h.Write([]byte(s))

	return KeyHash{Type: StringType, Value: h.Sum64()}
}
