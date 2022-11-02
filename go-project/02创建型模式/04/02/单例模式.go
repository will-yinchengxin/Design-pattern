package main

import "fmt"

// 懒汉式

type singelton struct{}

var instance *singelton

func GetInstance() *singelton {
	if instance == nil {
		instance = new(singelton)
		return instance
	}

	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("懒汉式-单例")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
