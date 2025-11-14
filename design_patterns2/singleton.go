package designpatterns2

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

//We can create singleton design pattern using Mutex so that if the Variable is nil then we can initialize the value again
//Or else use the one we were using before

type Config2 struct {
	Port int
	DB   *gorm.DB
}

var (
	cfg *Config2
	mu  sync.Mutex
)

func GetConfig() *Config2 {
	mu.Lock()
	defer mu.Unlock()
	if cfg == nil {
		cfg = &Config2{Port: 8080}
		fmt.Println("Initializing the cfg")
	}
	return cfg
}

func SingletonEx2() {
	var1 := GetConfig()
	fmt.Println("var1 cfg", var1.Port)
	cfg = nil
	var2 := GetConfig() //This will reinitialize the cfg
	fmt.Println("var2 cfg", var2.Port)
	var3 := GetConfig() //This will return the already stored value instead of initiating it again
	fmt.Println("var3 cfg", var3.Port)
}
