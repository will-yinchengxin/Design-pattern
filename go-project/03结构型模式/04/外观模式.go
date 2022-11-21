package main

import "fmt"

// 根据迪米特法则，如果两个类不必彼此直接通信，那么这两个类就不应当发生直接的相互作用

type ActivityOne struct{}

func (one *ActivityOne) DetailOne() {
	fmt.Println("this is DetailOne")
}

type ActivityTwo struct{}

func (two *ActivityTwo) DetailTwo() {
	fmt.Println("this is DetailTwo")
}

type ActivityThree struct{}

func (there *ActivityThree) DetailThree() {
	fmt.Println("this is DetailThree")
}

type ActivityFour struct{}

func (four *ActivityFour) DetailFour() {
	fmt.Println("this is DetailFour")
}

// 外观模式
type Facade struct {
	one   *ActivityOne
	two   *ActivityTwo
	there *ActivityThree
	four  *ActivityFour
}

func (f *Facade) MethodOne() {
	f.one.DetailOne()
	f.two.DetailTwo()
}

func (f *Facade) MethodTwo() {
	f.there.DetailThree()
	f.four.DetailFour()
}

// ..... 还可以进行多种组合

func main() {
	//如果不用 外观模式 实现 MethodOne() 和 MethodTwo()
	one := new(ActivityOne)
	one.DetailOne()
	two := new(ActivityTwo)
	two.DetailTwo()

	fmt.Println("-----------")
	//使用外观模式
	f := Facade{
		one:   new(ActivityOne),
		two:   new(ActivityTwo),
		there: new(ActivityThree),
		four:  new(ActivityFour),
	}
	//调用外观包裹方法
	f.MethodOne()
}
