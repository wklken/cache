package key

import "fmt"

type ExampleKey struct {
	Field1 string
	Field2 int64
}

func NewExampleKey(field1 string, field2 int64) *ExampleKey {
	return &ExampleKey{
		Field1: field1,
		Field2: field2,
	}
}

func (k ExampleKey) Key() string {
	return fmt.Sprintf("%s:%d", k.Field1, k.Field2)
}
