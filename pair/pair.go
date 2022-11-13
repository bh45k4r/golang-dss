package pair

type Pair[K, V any] struct {
	key   K
	value V
}

func New[K, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{key, value}
}

func (p Pair[K, V]) Key() K {
	return p.key
}

func (p Pair[K, V]) Value() V {
	return p.value
}
