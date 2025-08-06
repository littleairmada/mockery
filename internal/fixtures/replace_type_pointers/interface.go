package replace_type_pointers

type Foo int
type Bar int

type InterfaceWithPointers interface {
	FooFunc(*Foo) *Foo
	BarFunc(Bar) Bar
}
