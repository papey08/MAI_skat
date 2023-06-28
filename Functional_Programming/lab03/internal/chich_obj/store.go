package chich_obj

type Store map[string]Object

func (s Store) Get(n string) (Object, bool) {
	ret, ok := s[n]
	return ret, ok
}

func (s Store) Set(n string, o Object) Object {
	s[n] = o
	return o
}
