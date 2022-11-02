package main

import "fmt"

type Student struct {
}

func (this *Student) Learn() {
	fmt.Println("开始学习...")
}

func (this *Student) Cook() {
	fmt.Println("开始做饭...")
}

func (this *Student) Clean() {
	fmt.Println("开始打扫卫生...")
}

func main() {
	st := &Student{}
	st.Learn()
	st.Cook()
	st.Clean()

	// Student 负责了 n 多的模块，随着时间推移，可能还要新增许多的新功能方法，
	// 这样的设计会导致，当我们去给 Student 添加新的职责的时候，会直接修改原有的 Student 代码，那么 Student 模块的功能会越来越多，出现问题的几率也就越来越大，
	// 假如此时 Student 已经有99个职责了，现在我们要添加第100个职责，可能由于一次的不小心，导致之前99个职责也一起崩溃，因为所有的业务都在一个 Student 类里（结构体），
	// 他们的耦合度太高，Student 的职责也不够单一，代码的维护成本随着业务的复杂正比成倍增大。

}
