package main

import "fmt"

// 抽象类，一个制作工厂的完整流程步骤

type Beverage interface {
	BoilWater() //煮开水
	Brew()      //冲泡
	PourInCup() //倒入杯中
	AddThings() //添加酌料

	WantAddThings() bool //是否加入酌料Hook
}

// 模版，使得各个具体流程继承并实现
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

// 具体实现类，如煮咖啡
type coffee struct {
	template
}

func NewCoffee() *coffee {
	Caffe := new(coffee)
	// b 为 Beverage，是 coffee 的接口，这里需要给接口赋值，指向具体的子类对象
	// 来触发b全部接口方法的多态特性。
	Caffe.b = Caffe
	return Caffe
}

func (c *coffee) BoilWater() {
	fmt.Println("煮开水")
}
func (c *coffee) Brew() {
	fmt.Println("冲泡")
}
func (c *coffee) PourInCup() {
	fmt.Println("倒入杯中")
}

func (c *coffee) WantAddThings() bool {
	fmt.Println("要加料的～ 科技与狠活")
	return true
}

func (c *coffee) AddThings() {
	fmt.Println("给点 coffee 因")
}

// 具体实现类，如泡茶
type tea struct {
	template
}

func NewTea() *tea {
	// b 为 Beverage，是 coffee 的接口，这里需要给接口赋值，指向具体的子类对象
	// 来触发b全部接口方法的多态特性。
	teab := new(tea)
	teab.b = teab
	return teab
}

func (c *tea) BoilWater() {
	fmt.Println("煮开水")
}
func (c *tea) Brew() {
	fmt.Println("冲泡")
}
func (c *tea) PourInCup() {
	fmt.Println("倒入杯中")
}

func (c *tea) WantAddThings() bool {
	fmt.Println("不要加料的～ 拒绝科技与狠活")
	return true
}

func (c *tea) AddThings() {
	fmt.Println("纯天然无污染")
}

// 业务逻辑
func main() {
	tea := NewTea()
	tea.MakeBeverage()

	fmt.Println("--------------")

	coffee := NewCoffee()
	coffee.MakeBeverage()

	/*
		优点：
			(1) 在父类中形式化地定义一个算法，而由它的子类来实现细节的处理，在子类实现详细的处理算法时并不会改变算法中步骤的执行次序。
			(2) 模板方法模式是一种代码复用技术，它在类库设计中尤为重要，它提取了类库中的公共行为，将公共行为放在父类中，而通过其子类来实现不同的行为，它鼓励我们恰当使用继承来实现代码复用。
			(3) 可实现一种反向控制结构，通过子类覆盖父类的钩子方法来决定某一特定步骤是否需要执行。
			(4) 在模板方法模式中可以通过子类来覆盖父类的基本方法，不同的子类可以提供基本方法的不同实现，更换和增加新的子类很方便，符合单一职责原则和开闭原则。
		缺点：
			需要为每一个基本方法的不同实现提供一个子类，如果父类中可变的基本方法太多，将会导致类的个数增加，系统更加庞大，设计也更加抽象。

		使用场景：
			(1)具有统一的操作步骤或操作过程;
			(2) 具有不同的操作细节;
			(3) 存在多个具有同样操作步骤的应用场景，但某些具体的操作细节却各不相同;
				在抽象类中统一操作步骤，并规定好接口；让子类实现接口。这样可以把各个具体的子类和操作步骤解耦合
	*/
}
