package main

import "fmt"

type Animal struct {
}

func (a *Animal) Show(name string) {
	if name == "monkey" {
		fmt.Println("我是猴子")
	} else if name == "dog" {
		fmt.Println("我是小狗")
	} else if name == "cat" {
		fmt.Println("我是小猫")
	}
}

func NewAnimal(name string) *Animal {
	fruit := new(Animal)

	if name == "monkey" {
		//创建monkey逻辑
	} else if name == "dog" {
		//创建dog逻辑
	} else if name == "cat" {
		//创建cat逻辑
	}

	return fruit
}

func main() {
	monkey := NewAnimal("monkey")
	monkey.Show("monkey")

	dog := NewAnimal("dog")
	dog.Show("dog")

	cat := NewAnimal("cat")
	cat.Show("cat")
	/*
		问题：
			(1) 在Fruit类中包含很多 「if…else…」代码块，整个类的代码相当冗长，代码越长，阅读难度、维护难度和测试难度也越大；
				而且大量条件语句的存在还将影响系统的性能，程序在执行过程中需要做大量的条件判断。
			(2) Fruit类的职责过重，它负责初始化和显示所有的水果对象，将各种水果对象的初始化代码和显示代码集中在一个类中实现，
				违反了「单一职责原则」，不利于类的重用和维护；
			(3) 当需要增加新类型的水果时，必须修改Fruit类的构造函数NewFruit()和其他相关方法源代码，违反了「开闭原则」。

		工厂模式：
			之前的依赖关系：业务层 ---》 基础能力
			新的依赖关系：  业务层 ---》 工厂模式 ---》基础类模块
	*/
}
