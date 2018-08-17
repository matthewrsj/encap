package animal

type Dog struct {
	CommonAnimal
}

func NewDog(name string, age, weight int) Animal {
	return &Dog{NewCommonAnimal(name, age, weight)}
}

func (d *Dog) Speak() string {
	return "bark bark"
}

func (d *Dog) Eat() string {
	return "munch munch"
}

func (d *Dog) Bite() string {
	return "chomp chomp"
}
