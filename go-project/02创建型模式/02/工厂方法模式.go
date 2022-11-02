package main

import "fmt"

// 动物抽象层

type Animal interface {
	Show()
}

type Factory interface {
	CreateAnimal() *Animal
}

// 基础能力

type Monkey struct {
	Animal // 方便理解可以「显示继承」，可以删除掉
}

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

// 工厂模块

type MonkeyFactory struct {
	Factory
}

func (m *MonkeyFactory) CreateAnimal() *Animal {
	var monkey Animal
	monkey = &Monkey{}
	return &monkey
}

type CatFactory struct{}

func (m *CatFactory) CreateAnimal() *Animal {
	var monkey Animal
	monkey = &Cat{}
	return &monkey
}

type DogFactory struct{}

func (m *DogFactory) CreateAnimal() *Animal {
	var monkey Animal
	monkey = &Dog{}
	return &monkey
}

func main() {
	monkey := &MonkeyFactory{}
	(*monkey.CreateAnimal()).Show()

	cat := &CatFactory{}
	(*cat.CreateAnimal()).Show()

	dog := &DogFactory{}
	(*dog.CreateAnimal()).Show()
	/*
		优点：
			1. 不需要记住具体类名，甚至连具体参数都不用记忆。
			2. 实现了对象创建和使用的分离。
			3. 系统的可扩展性也就变得非常好，无需修改接口和原类。
			4.对于新产品的创建，符合开闭原则。

		缺点：
			1. 增加系统中类的个数，复杂度和理解度增加。
			2. 增加了系统的抽象性和理解难度。

		适用场景：
			1. 客户端不知道它所需要的对象的类。
			2. 抽象工厂类通过其子类来指定创建哪个对象。
	*/
}
