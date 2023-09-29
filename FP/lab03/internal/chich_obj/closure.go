package chich_obj

import "fmt"

type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

func (c *Closure) String() string {
	return fmt.Sprintf("closure[%p]", c)
}

func (c *Closure) Type() Type {
	return ClosureType
}
