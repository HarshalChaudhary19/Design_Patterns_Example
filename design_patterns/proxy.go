package designpatterns

import "fmt"

//Basically kind of a middleware is proxy jo ki decide krta hai ki real object ka method ke pass request ko jane dena hai ki nhi
//Ya pehle authentication hogi ya fir kuch aur

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func (r *RealImage) Display() {
	fmt.Println("Displaying", r.filename)
}

type ProxyImage struct {
	filename string
	real     *RealImage
}

func (p *ProxyImage) Display() {
	if p.real == nil {
		p.real = &RealImage{filename: p.filename}
	}
	p.real.Display()
}

func Proxy() {
	img := &ProxyImage{filename: "pic.jpg"}
	img.Display() //first call -> creates Realimage and displays it
	img.Display() //second call -> reuses already created Realimage
}
