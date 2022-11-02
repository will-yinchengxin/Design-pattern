package main

// 概念：
// 如果一个struct嵌套了另一个有名结构体，那么这个模式就叫「组合」。
// 如果一个struct嵌套了另一个匿名结构体（只有类型没有名字），那么这个结构可以直接访问匿名结构体的方法，从而实现了「继承」。
// 如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现了「多重继承」

// 如果使用「继承」，会导致父类的任何变换都可能影响到子类的行为。如果使用对象「组合」，就降低了这种依赖关系。
// 对于继承和组合，「优先使用组合」

import "fmt"

type Cat struct {
	food string
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

//给小猫添加一个 可以睡觉的方法 （使用继承来实现）
type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉1")
}

//给小猫添加一个 可以睡觉的方法 （使用组合的方式）
type CatC struct {
	C *Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉2")
}

func main() {
	//通过继承增加的功能，可以正常使用
	cb := new(CatB)
	cb.Eat()
	cb.Sleep()

	//通过组合增加的功能，可以正常使用
	cc := new(CatC)
	cc.C = new(Cat)
	cc.C.Eat()
	cc.Sleep()
}
