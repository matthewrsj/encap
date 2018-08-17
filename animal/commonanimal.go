package animal

type CommonAnimal struct {
	name   string
	weight int
	age    int
}

func NewCommonAnimal(name string, weight, age int) CommonAnimal {
	return CommonAnimal{name: name, weight: weight, age: age}
}

func (a *CommonAnimal) Name() string {
	return a.name
}

func (a *CommonAnimal) Weight() int {
	return a.weight
}

func (a *CommonAnimal) Age() int {
	return a.age
}

func (a *CommonAnimal) Speak() string {
	return "hi"
}

func (a *CommonAnimal) Eat() string {
	return "chew"
}

func (a *CommonAnimal) Bite() string {
	return "chomp"
}
