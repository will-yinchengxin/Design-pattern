package main

import (
	"fmt"
	"sync"
)

// 线程安全

var once sync.Once

type singelton struct{}

var instance *singelton

func GetInstance() *singelton {

	once.Do(func() {
		instance = new(singelton)
	})

	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("线程安全-单例")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
