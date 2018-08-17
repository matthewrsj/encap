package animal

type Platypus struct {
	CommonAnimal
}

func NewPlatypus(name string, age, weight int) Animal {
	return &Platypus{NewCommonAnimal(name, age, weight)}
}

func (g *Platypus) Speak() string {
	return "quack quack"
}

func (g *Platypus) Eat() string {
	return "slurp slurp"
}

func (g *Platypus) Bite() string {
	return "smack smack"
}
