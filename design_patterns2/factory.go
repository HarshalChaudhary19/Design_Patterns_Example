package designpatterns2

import "fmt"

//Basically Jitne bhi products hain unke methods ko define n krke Concrete Products mein decide krenge
//For Example Product mein hai deliver and Concrete products we can add as much as we want...and every product will have the method deliver\
//But will behave very differently as per the usecase

//Products
//Representation of the Object
type Car interface {
	Drive() string
}

//Concrete Products
//Specific implementation of the Product interface
//Jo jo features hain Product ke
type Sedan struct{}
type SUV struct{}

func (c *Sedan) Drive() string {
	return "Driving Sedan"
}

func (c *SUV) Drive() string {
	return "Driving SUV"
}

//Factory
//This is responsible for creating objects
type CarFactory interface {
	CreateCar() Car
}

//Concrete Factory
//Implements Factory methods and produces specific instances of the product
//Product create krne ke liye
type SedanFactory struct{}
type SUVFactory struct{}

func (sf *SedanFactory) CreateCar() Car {
	return &Sedan{}
}

func (sf *SUVFactory) CreateCar() Car {
	return &SUV{}
}

func Factory() {
	//Firstly creating a Sedan Car using Sedan Factory
	sedanFactoryTemp := &SedanFactory{}
	sedan := sedanFactoryTemp.CreateCar()
	//Then invoking the method
	fmt.Println("Sedan does this", sedan.Drive())
}
