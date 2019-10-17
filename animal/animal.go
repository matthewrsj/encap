package animal

// Animal is a living organism that feeds on organic matter, typically having
// specialized sense organs and nervous system and able to respond rapidly to
// stimuli.
type Animal interface {
	// Commands
	// Speak returns the sound the animal makes
	Speak() string
	// Eat returns the sound the animal makes while eating and may modify weight
	Eat() string
	// Bite returns the sound the animal makes while biting
	Bite() string

	// Reporting functions
	// Name returns the name of the animal
	Name() string
	// Weight returns the weight (in lbs) of the animal
	Weight() int
	// Age returns the age (in years) of the animal
	Age() int
}
