package solidprinciples

// Liskov Substitution Principle (LSP)
//This means the objects of the derived class should be capable of replacing objects of base class

//In the below example the method Fly doesn't apply to Penguin but it is a bird....
type Bird interface {
	Fly() string
}

type Pigeon struct{}

func (p *Pigeon) Fly() string {
	return "Pigeon is flying."
}

type Penguin struct{}

func (p *Penguin) Fly() string {
	return "Penguin is flying."
}

//And ab ye niche wale example mein we have segregated things and Now every bird can makesound but Only Flying bird can fly
//(And also make sounds)

type Birdy interface {
	MakeSound() string
}

type FlyingBirdy interface {
	Bird
	FlyB() string
}

type Pigeonn struct{}

func (p *Pigeonn) MakeSound() string {
	return "Coo"
}

func (p *Pigeon) FlyB() string {
	return "Pigeon is flying."
}

type Penguinn struct{}

func (p *Penguinn) MakeSound() string {
	return "Squawk"
}
