package main

import "fmt"

// 将一个类的接口转换成客户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

//适配的目标
type V100 interface {
	Use100V()
}

type JiayongDain struct {
	v V100
}

func (j *JiayongDain) Use100V() {
	j.v.Use100V()
	fmt.Println("使用的是 100v 电")
}

//被适配的角色，适配者
func NewJiayongDain(v100 V100) *JiayongDain {
	return &JiayongDain{v100}
}

type GaoyaDian struct{}

func (g *GaoyaDian) Use1000v() {
	fmt.Println("1000v 高压电")
}

// 适配器
type Adapter struct {
	v1000 *GaoyaDian
}

func (a *Adapter) Use100V() {
	a.v1000.Use1000v()
	fmt.Println("开始适配")
}

func NewAdapter(v1000 *GaoyaDian) *Adapter {
	return &Adapter{v1000: v1000}
}

// 业务裸机
func main() {
	jiayong := NewJiayongDain(NewAdapter(&GaoyaDian{}))
	jiayong.Use100V()
	/*
		优点：
			(1) 将目标类和适配者类解耦，通过引入一个适配器类来重用现有的适配者类，无须修改原有结构。
			(2) 增加了类的透明性和复用性，将具体的业务实现过程封装在适配者类中，对于客户端类而言是透明的，而且提高了适配者的复用性，同一个适配者类可以在多个不同的系统中复用。
			(3) 灵活性和扩展性都非常好，可以很方便地更换适配器，也可以在不修改原有代码的基础上增加新的适配器类，完全符合“开闭原则”。
		缺点:
			适配器中置换适配者类的某些方法比较麻烦。

		适应场景
			(1) 系统需要使用一些现有的类，而这些类的接口（如方法名）不符合系统的需要，甚至没有这些类的源代码。
			(2) 想创建一个可以重复使用的类，用于与一些彼此之间没有太大关联的一些类，包括一些可能在将来引进的类一起工作。
	*/
}
