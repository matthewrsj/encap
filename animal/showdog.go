package animal

type ShowDog struct {
	Dog
}

func NewShowDog(name string, age, weight int) Animal {
	return &ShowDog{
		Dog{NewCommonAnimal(name, age, weight)},
	}
}

func (sd *ShowDog) Bite() string {
	return ""
}

func (sd *ShowDog) Eat() string {
	sd.weight += 2
	return "munch munch"
}
