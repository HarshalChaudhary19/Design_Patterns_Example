package designpatterns

import "fmt"

//Ads behavior to an object dynamically, without changing the object
//Add logging,metrics, caching around existing functionality

type Service interface {
	Do() string
}

type RealService struct{}

func (RealService) Do() string {
	return "done"
}

type LoggingDecorator struct {
	s Service
}

func (l LoggingDecorator) Do() string {
	fmt.Println("Before")
	res := l.s.Do()
	fmt.Println("after")
	return res
}

func Decorator() {
	var svc Service = RealService{}
	svc = LoggingDecorator{svc}
	fmt.Println("Svc now has this method", svc.Do())
}
