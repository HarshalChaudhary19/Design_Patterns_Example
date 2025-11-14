package designpatterns2

import "fmt"

//This is the method that we need to use
type Driven interface {
	Drive()
}

//This is the struct that will use the drive method
type Carstruc struct{}

func (c *Carstruc) Drive() {
	fmt.Println("Car being driven")
}

//We will check the age of the Driver
type Driver struct {
	age int
}

//This is the proxy so we will use this struct
type CarProxy struct {
	car    Carstruc
	driver *Driver
}

//This way we will use drive via proxy and check if the age is above 18 or not
func (c *CarProxy) Drive() {
	if c.driver.age >= 18 {
		c.car.Drive()
	} else {
		fmt.Println("Driver is too young")
	}
}

//To create new proxy

func NewProxy(driver *Driver) *CarProxy {
	return &CarProxy{Carstruc{}, driver}
}

func Proxy() {
	car := NewProxy(&Driver{12})
	car.Drive() //This will print the else statement
}
