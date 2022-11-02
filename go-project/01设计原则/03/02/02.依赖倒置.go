package main

import "fmt"

// 抽象类

type Student interface {
	Learn(Project)
}

type Project interface {
	HomeWork()
}

// 实现类

// 学生
type Will struct {
}

func (w *Will) Learn(pro Project) {
	fmt.Println("will learn")
	pro.HomeWork()
}

type Yin struct {
}

func (y *Yin) Learn(pro Project) {
	fmt.Println("yin learn")
	pro.HomeWork()
}

// 学科

type Math struct {
}

func (this *Math) HomeWork() {
	fmt.Println("math HomeWork")
}

type Chinese struct {
}

func (this *Chinese) HomeWork() {
	fmt.Println("chinese HomeWork")
}

// 业务逻辑
func main() {
	var project1 *Math
	project1 = &Math{}
	var stu *Will
	stu = &Will{}
	stu.Learn(project1)

	fmt.Println("--------------------")

	var project2 *Chinese
	project2 = &Chinese{}
	var stu1 *Yin
	stu1 = &Yin{}
	stu1.Learn(project2)

	// 那么，我们首先将抽象层的模块和接口定义出来，这里就需要了interface接口的设计，然后我们依照抽象层，依次实现每个实现层的模块，在我们写实现层代码的时候，
	// 实际上我们只需要参考对应的抽象层实现就好了，实现每个模块，也和其他的实现的模块没有关系，符合 「开闭原则」

	// 抽象层暴露出来的接口就是我们业务层可以使用的方法，然后可以通过多态的线下，
	// 接口指针指向哪个实现模块，调用了就是具体的实现方法，这样我们业务逻辑层也是依赖抽象成编程
	// 这种的设计原则叫做依赖倒转原则
}
