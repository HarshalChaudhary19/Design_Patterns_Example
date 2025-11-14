package designpatterns

import "fmt"

//Basically Jaise ham Interfaces and methods wgera bnate hain to pass to next level in
//our Architecture
//Bascically Ek interface bna kr and structs bna kr wo struct se related methods bna diye and then method ka naam se function
//and then wo struct ke instance se method call hoga interface ka

type Notifier interface {
	Notify(msg string)
}

type Email struct{}

type SMS struct{}

func (Email) Notify(msg string) {
	fmt.Println("Email is this", msg)
}

func (SMS) Notify(msg string) {
	fmt.Println("SMS is this", msg)
}
