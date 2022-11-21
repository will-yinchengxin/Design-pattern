package main

import "fmt"

/*
观察者模式是用于建立一种对象与对象之间的依赖关系，一个对象发生改变时将自动通知其他对象，其他对象将相应作出反应。在观察者模式中，
发生改变的对象称为观察目标，而被通知的对象称为观察者，一个观察目标可以对应多个观察者，而且这些观察者之间可以没有任何相互联系，
可以根据需要增加和删除观察者，使得系统更易于扩展。
*/

//--------- 抽象层 --------

//抽象的观察者

type Listener interface {
	OnStealing() //观察者得到通知后要触发的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// ------ 实现层 --------
type XiaotouA struct {
	Staff string
}

func (x *XiaotouA) OnStealing() {
	fmt.Println("小偷A偷了: " + x.Staff)
}

type XiaotouB struct {
	Staff string
}

func (x *XiaotouB) OnStealing() {
	fmt.Println("小偷B偷了: " + x.Staff)
}

type Fangshao struct {
	Tonghuo []Listener
}

func (m *Fangshao) AddListener(listener Listener) {
	m.Tonghuo = append(m.Tonghuo, listener)
}

func (m *Fangshao) RemoveListener(listener Listener) {
	for index, l := range m.Tonghuo {
		//找到要删除的元素位置
		if listener == l {
			//将删除的点前后的元素链接起来
			m.Tonghuo = append(m.Tonghuo[:index], m.Tonghuo[index+1:]...)
			break
		}
	}
}

func (m *Fangshao) Notify() {
	for _, listener := range m.Tonghuo {
		//依次调用全部观察的具体动作
		listener.OnStealing()
	}
}

// ------- 业务层 --------
func main() {
	s1 := &XiaotouA{
		Staff: "珠宝",
	}
	s2 := &XiaotouB{
		Staff: "黄金",
	}
	fangshao := &Fangshao{}
	fangshao.Tonghuo = append(fangshao.Tonghuo, s1)
	fangshao.Tonghuo = append(fangshao.Tonghuo, s2)

	fmt.Println("开偷")
	s1.OnStealing()
	s2.OnStealing()

	fmt.Println("警察来了")

	fangshao.Notify()
}
