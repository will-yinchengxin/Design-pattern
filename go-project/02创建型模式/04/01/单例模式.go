package main

import "fmt"

// 保证一个类、只有一个实例存在，同时提供能对该实例加以访问的全局访问方法。
// 单例模式要解决的问题是：保证一个类永远只能有一个对象，且该对象的功能依然能被其他模块使用。

// 饿汉式：
// 述是单例模式中的一种，属于“饿汉式”。含义是，在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存。饿汉式的好处是，不会出现线程并发创建，
// 导致多个单例的出现，但是缺点是如果这个单例对象在业务逻辑没有被使用，也会客观的创建一块内存对象

type singelton struct{}

var instance *singelton = new(singelton)

func GetInstance() *singelton {
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("饿汉式-单例")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
