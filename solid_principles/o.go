package solidprinciples

import "math"

//Open/Close principle
//This states that the functions/methods should be open to extension but closed to modification
//Mtlb koi nye scenerio pr purane wale mein modifications n krni pde sirf code ko extend krwayen

//Niche wale mein agr circle wgera ke liye Area nikalna hua to Area function mein modifications krni pdengi
type Rectangle struct {
	Width  float64
	Height float64
}

func Area(rectangle *Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

//But niche wale example mein we can add more structs as shapes and add there implementations
//Jaise ye Circle
type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
