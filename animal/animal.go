package animal

// An Animal is a living organism that feeds on organic matter, typically having
// specialized sense organs and nervous system and able to respond rapidly to
// stimuli.
type Animal interface {
	// Commands
	Speak() string
	Eat() string
	Bite() string

	// Reporting functions
	Name() string
	Weight() int
	Age() int
}
