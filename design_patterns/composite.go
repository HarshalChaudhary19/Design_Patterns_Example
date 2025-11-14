package designpatterns

import "fmt"

type Component interface {
	Draw()
}

type Leaf struct {
	name string
}

func (l Leaf) Draw() {
	fmt.Println("Leaf", l.name)
}

type Composite struct {
	children []Component // Can hold leaves or other composites
	name     string
}

func (c *Composite) Add(child Component) {
	c.children = append(c.children, child)
}

func (c *Composite) Draw() {
	fmt.Println("Composite", c.name)
	for _, ch := range c.children {
		ch.Draw() // recursive call: each child draws itself
	}
}

func CompositeExample() {
	root := &Composite{name: "root"}
	root.Add(Leaf{"A"})

	child := &Composite{name: "sub"}
	child.Add(Leaf{"B"})

	root.Add(child)

	root.Draw()
}
