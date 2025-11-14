package designpatterns

import "fmt"

//Basically Hmm ek chiz ko call krenge and then wo multiple ko aage call kregi for complex systems
//And we don't need to know the internal working in the complex interface

//We have done this in the Gin_Event_App_Gorm_Arc....Check

//System 1
type AuthService struct{} //Iss Authservice struct ke assosiation mein Login method hai
func (AuthService) Login() {
	fmt.Println("User has logged in ")
}

//System 2
type PaymentService struct{} //Ab iss PaymentService ke struct se assosiated hai Charge method
func (PaymentService) Charge() {
	fmt.Println("User has been charged")
}

type ShopFacade struct {
	auth AuthService
	pay  PaymentService
}

func (s ShopFacade) Buy() {
	s.auth.Login() //Complex method se Login ko call kiya
	s.pay.Charge() //Complex method se Charge ko call kiya
	fmt.Println("Product is sold")
}

func Facade() {
	shop := ShopFacade{}
	shop.Buy()
}
