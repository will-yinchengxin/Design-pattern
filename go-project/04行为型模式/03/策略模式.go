package main

import "fmt"

// 接口类
type Weapon interface {
	UseIt()
}

type Ak47 struct{}

func (a *Ak47) UseIt() {
	fmt.Println("Ak47 就是吐吐～～～")
}

type DaGai struct{}

func (a *DaGai) UseIt() {
	fmt.Println("38 大盖也挺好使～～～")
}

type Soldier struct {
	w Weapon
}

func (s *Soldier) SetWeapon(w Weapon) {
	s.w = w
}

func (s *Soldier) Fight() {
	s.w.UseIt()
}

func main() {
	soldier := &Soldier{}

	soldier.SetWeapon(new(Ak47))
	soldier.Fight()

	soldier.SetWeapon(new(DaGai))
	soldier.Fight()
	/*
		优点：
			(1) 策略模式提供了对“开闭原则”的完美支持，用户可以在不修改原有系统的基础上选择算法或行为，也可以灵活地增加新的算法或行为。
			(2) 使用策略模式可以避免多重条件选择语句。多重条件选择语句不易维护，它把采取哪一种算法或行为的逻辑与算法或行为本身的实现逻辑混合在一起，将它们全部硬编码(Hard Coding)在一个庞大的多重条件选择语句中，比直接继承环境类的办法还要原始和落后。
			(3) 策略模式提供了一种算法的复用机制。由于将算法单独提取出来封装在策略类中，因此不同的环境类可以方便地复用这些策略类。
		缺点：
			(1) 客户端必须知道所有的策略类，并自行决定使用哪一个策略类。这就意味着客户端必须理解这些算法的区别，以便适时选择恰当的算法。换言之，策略模式只适用于客户端知道所有的算法或行为的情况。
			(2) 策略模式将造成系统产生很多具体策略类，任何细小的变化都将导致系统要增加一个新的具体策略类。

		适用场景
			准备一组算法，并将每一个算法封装起来，使得它们可以互换
	*/
}
