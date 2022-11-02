package main

import "fmt"

type Family interface {
	DoOwnThing()
}

type Stu struct {
}

func (this *Stu) DoOwnThing() {
	fmt.Println("开始学习...")
}

type Mom struct {
}

func (this *Mom) DoOwnThing() {
	fmt.Println("开始做饭...")
}

type Dad struct {
}

func (this *Dad) DoOwnThing() {
	fmt.Println("开始打扫卫生...")
}

func FamilyDoOwnThing(owner Family) {
	owner.DoOwnThing()
}

func main() {
	st := &Stu{}
	FamilyDoOwnThing(st)

	mom := &Mom{}
	FamilyDoOwnThing(mom)

	dad := &Dad{}
	FamilyDoOwnThing(dad)

	// 开闭原则定义: 在修改需求的时候，应该尽量通过扩展来实现变化，而不是通过修改已有代码来实现变化
	// 接口的最大的意义就是实现多态的思想 ！！！！
	// 我们可以根据interface类型来设计API接口，那么这种API接口的适应能力不仅能适应当下所实现的全部模块，也适应未来实现的模块来进行调用
}
