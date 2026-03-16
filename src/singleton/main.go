package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	number int
}

var once sync.Once
var instance *Singleton

func GetSingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetSingleton()
	s2 := GetSingleton()
	s1.number = 42
	fmt.Println(s2.number) // 42
}
