package designpatterns2

import "fmt"

//Product
type Computer struct {
	CPU     string
	RAM     string
	Storage string
}

func (c *Computer) DisplayInfo() {
	fmt.Println("CPU specs", c.CPU)
	fmt.Println("RAM available", c.RAM)
	fmt.Println("Storage available", c.Storage)
}

//Builder Interface
//Defines the methods for creating Building the Product
type Builder interface {
	BuildRAM()
	BuildCPU()
	BuildStorage()
	GetResult() *Computer
}

//ConcreteBuilder
//Implements the methods from Builder Interface and create the computer
type GamingComputerBuilder struct {
	computer Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{computer: Computer{}}
}

func (b *GamingComputerBuilder) BuildCPU() {
	b.computer.CPU = "Some CPU"
}

func (b *GamingComputerBuilder) BuildRAM() {
	b.computer.RAM = "8 GB"
}

func (b *GamingComputerBuilder) BuildStorage() {
	b.computer.Storage = "512 SSD"
}

func (b *GamingComputerBuilder) GetResult() *Computer {
	return &b.computer
}

//Director
//Basically this manages the Building process and defines the order
type ComputerDirector struct{}

func (c *ComputerDirector) Construct(builder Builder) {
	builder.BuildCPU()
	builder.BuildRAM()
	builder.BuildStorage()

}

//And ye niche hai client
func BuilderExample2() {
	builder := NewGamingComputerBuilder()
	director := &ComputerDirector{}
	director.Construct(builder)
	computer := builder.GetResult()
	computer.DisplayInfo()
}
