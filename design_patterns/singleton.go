package designpatterns

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

//Basically Ek he instance hoga kisi object ka aur wohi pass krenge hmm sb jagah
//And ek global Point hoga jahan se wo Data ko read krskte hain
//Just like we do it in Gin_Service_Whatever

type Config struct {
	Port int
	DB   *gorm.DB
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{Port: 8080}
	})
	return cfg
}

func SingletonCompare() {
	a := GetConfig()
	b := GetConfig()

	fmt.Println(a == b) //This will be true as there is only one instance for the Config that is getting passed
}
