package animal

type Chihuahua struct {
	Dog
}

func NewChihuahua(name string, age, weight int) Animal {
	return &Chihuahua{Dog{NewCommonAnimal(name, age, weight)}}
}

func (c *Chihuahua) Speak() string {
	c.weight--
	return "bark bark bark bark bark bark bark bark bark bark bark bark!"
}
