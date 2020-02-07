package key

type StringKey struct {
	key string
}

func NewStringKey(key string) *StringKey {
	return &StringKey{
		key: key,
	}
}

func (s StringKey) Key() string {
	return s.key
}
