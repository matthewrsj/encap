package animal

type Gopher struct {
	CommonAnimal
}

func NewGopher(name string, age, weight int) Animal {
	return &Gopher{NewCommonAnimal(name, age, weight)}
}

func (g *Gopher) Speak() string {
	return "mpf mpf mpf"
}

func (g *Gopher) Eat() string {
	return "chomp chomp"
}

func (g *Gopher) Bite() string {
	return "chew"
}
