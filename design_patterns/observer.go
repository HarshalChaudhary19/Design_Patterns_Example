package designpatterns

import "fmt"

//Also Called Publish or Subscribe

//Many Listner reacts when something happens
//Basically notifications wgera bhejne ke liye use hota hai

//Kuch bhi change hota hai to notification sbke pass jati hai

type ObserverItf interface {
	Update(string)
}

type Subject struct {
	observers []ObserverItf
}

func (s *Subject) Register(o ObserverItf) {
	s.observers = append(s.observers, o)
}
func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type Logger struct {
	name string
}

func (l *Logger) Update(msg string) {
	fmt.Println(l.name, "received the msg...", msg)
}

func Observer() {
	sub := &Subject{}
	sub.Register(&Logger{"Harshal"})
	sub.Register(&Logger{"Antier"})
	sub.Notify("Event Happened")
}
