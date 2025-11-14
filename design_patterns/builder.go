package designpatterns

import "fmt"

//Step by Step creation of a complex Object
//Basically Chain create krni hai methods ki to create a Complex Object
//Here we are creating Server using ServerBuilder and all the methods like Host,port,TLS and all
//Ek ServerBulder hoga struct jiske andar main struct hoga and then uss struct ke aage methods assosiated honge

type Server struct {
	Host string
	Port int
	TLS  string
}

type ServerBuilder struct {
	s Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{Server{Port: 8080, Host: "localhost"}}
}

func (b *ServerBuilder) Host(h string) *ServerBuilder {
	b.s.Host = h
	return b
}

func (b *ServerBuilder) Port(p int) *ServerBuilder {
	b.s.Port = p
	return b
}

func (b *ServerBuilder) TLS(t string) *ServerBuilder {
	b.s.TLS = t
	return b
}

func (b *ServerBuilder) Build() Server {
	return b.s
}

func Builder() {
	s := NewServerBuilder().Host("example.com").Port(8080).TLS("fabric/ca-cert.pem").Build()
	fmt.Println("Server is this", s)
}
