package main

import "fmt"

// 动物抽象层

type Animal interface {
	Show()
}

// 基础能力

type Monkey struct{}

func (m *Monkey) Show() {
	fmt.Println("我是猴子")
}

type Dog struct{}

func (m *Dog) Show() {
	fmt.Println("我是小狗")
}

type Cat struct{}

func (m *Cat) Show() {
	fmt.Println("我是小猫")
}

// 简单工厂模块

type AnimalFactory struct{}

func (a *AnimalFactory) ShowAnimal(animal string) *Animal {
	var animalBasic Animal
	switch animal {
	case "monkey":
		animalBasic = &Monkey{}
	case "dog":
		animalBasic = &Dog{}
	case "cat":
		animalBasic = &Cat{}
	}
	return &animalBasic
}

func main() {
	AnimalFactory := &AnimalFactory{}
	monkey := AnimalFactory.ShowAnimal("monkey")
	(*monkey).Show()

	dog := AnimalFactory.ShowAnimal("dog")
	(*dog).Show()

	cat := AnimalFactory.ShowAnimal("cat")
	(*cat).Show()
	/*
		优点：
			1. 实现了对象创建和使用的分离。
			2. 不需要记住具体类名，记住参数即可，减少使用者记忆量。

		缺点：
			1. 对工厂类职责过重，一旦不能工作，系统受到影响。
			2. 增加系统中类的个数，复杂度和理解度增加。
			3. 违反「开闭原则」，添加新产品需要修改工厂逻辑，工厂越来越复杂。

		适用场景：
			1. 工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
			2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。
	*/
}
