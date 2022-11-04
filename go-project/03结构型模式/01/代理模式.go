package main

import "fmt"

type Goods struct {
	Kind string //商品种类
	Fact bool   //商品真伪
}

// =========== 抽象层 ===========

type Shopping interface {
	Buy(goods *Goods)
}

// =========== 实现层 ===========

type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物, 买了 ", goods.Kind)
}

// OverseasProxy 代理
type OverseasProxy struct {
	shopping Shopping //代理某个主题，这里是抽象类型
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1. 先验货
	if op.distinguish(goods) == true {
		//2. 进行购买
		op.shopping.Buy(goods) //调用原被代理的具体主题任务
		//3 海关安检
		op.check(goods)
	}
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

// 业务逻辑
func main() {
	g1 := Goods{
		Kind: "NBA门票",
		Fact: true,
	}

	g2 := Goods{
		Kind: "德云社门票",
		Fact: false,
	}

	shopping := new(KoreaShopping) //具体购买主题

	fmt.Println("---------------以下是 未使用 代理模式-------")
	if g1.Fact == true {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪.")
		shopping.Buy(&g1)
		fmt.Println("对[", g1.Kind, "] 进行了海关检查， 成功的带回祖国")
	}

	fmt.Println("---------------以下是 使用 代理模式-------")
	overseasProxy := NewProxy(shopping)
	overseasProxy.Buy(&g1)
	overseasProxy.Buy(&g2)

	/*
		优点：
			(1) 能够协调调用者和被调用者，在一定程度上降低了系统的耦合度。
			(2) 客户端可以针对抽象主题角色进行编程，增加和更换代理类无须修改源代码，符合开闭原则，系统具有较好的灵活性和可扩展性。

		缺点：
			(1) 代理实现较为复杂。

		适用场景
			为其他对象提供一种代理以控制对这个对象的访问。
	*/
}
