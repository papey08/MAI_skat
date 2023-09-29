package chich_obj

type Env struct {
	Outer *Env
	Store Store
	file  string
	dir   string
}

func (e *Env) Get(n string) (Object, bool) {
	ret, ok := e.Store[n]
	if !ok && e.Outer != nil {
		return e.Outer.Get(n)
	}
	return ret, ok
}

func (e *Env) Set(n string, o Object) Object {
	e.Store[n] = o
	return o
}
