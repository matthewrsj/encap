package animal

type Animal interface {
	Speak() string
	Eat() string
	Bite() string
	Name() string
	Weight() int
	Age() int
}
