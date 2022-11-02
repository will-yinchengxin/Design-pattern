package main

import "fmt"

// 工厂方法模式通过引入工厂等级结构，解决了简单工厂模式中工厂类职责太重的问题，但由于工厂方法模式中的每个工厂只生产一类产品，
// 可能会导致系统中存在大量的工厂类，势必会增加系统的开销。因此，可以考虑将一些相关的产品组成一个“产品族”，由同一个工厂来统一生产

/*
从工厂方法模式可以看出来：
（1）当添加一个新动物的时候，比如大象，虽然不用修改代码，但是需要添加大量的类，而且还需要添加相对的工厂。（系统开销，维护成本）
（2）如果使用同一地域的动物（中国小狗，中国小猫，中国猴子），那么需要分别创建具体的工厂，如果选择出现失误，将会造成混乱，虽然可以加一些约束，但是代码实现变得复杂。
	所以“抽象工厂方法模式”引出了「动物族」和「动物等级结构」概念，其目的是为了更加高效的生产同一个「动物族」产品。

（3）动物族：全部来自统一产地，如：中国
	动物等级：都是猫
*/

// 抽象层

type AbstractMonkey interface {
	ShowMonkey()
}

type AbstractCat interface {
	ShowCat()
}

type AbstractDog interface {
	ShowDog()
}

type AbstractAnimal interface {
	CreateMonkey() AbstractMonkey
	CreateCat() AbstractCat
	CreateDog() AbstractDog
}

// 实现层
/*made in china*/
type ChineseMonkey struct{}

func (c *ChineseMonkey) ShowMonkey() {
	fmt.Println("chinese monkey")
}

type ChineseCat struct{}

func (c *ChineseCat) ShowCat() {
	fmt.Println("chinese cat")
}

type ChineseDog struct{}

func (c *ChineseDog) ShowDog() {
	fmt.Println("chinese dog")
}

type ChineseFactory struct{}

func (c *ChineseFactory) CreateMonkey() AbstractMonkey {
	monkey := &ChineseMonkey{}
	return monkey
}

func (c *ChineseFactory) CreateCat() AbstractCat {
	cat := &ChineseCat{}
	return cat
}

func (c *ChineseFactory) CreateDog() AbstractDog {
	dog := &ChineseDog{}
	return dog
}

/*made in usa*/
type USAMonkey struct{}

func (u *USAMonkey) ShowMonkey() {
	fmt.Println("usa monkey")
}

type USACat struct{}

func (u *USACat) ShowCat() {
	fmt.Println("usa cat")
}

type USADog struct{}

func (u *USADog) ShowDog() {
	fmt.Println("usa dog")
}

type USAFactory struct{}

func (u *USAFactory) CreateMonkey() AbstractMonkey {
	monkey := &USAMonkey{}
	return monkey
}

func (u *USAFactory) CreateCat() AbstractCat {
	cat := &USACat{}
	return cat
}

func (u *USAFactory) CreateDog() AbstractDog {
	dog := &USADog{}
	return dog
}

func main() {
	chinaFactory := &ChineseFactory{}
	chinaFactory.CreateMonkey().ShowMonkey()
	chinaFactory.CreateDog().ShowDog()
	chinaFactory.CreateCat().ShowCat()

	usaFactory := &USAFactory{}
	usaFactory.CreateMonkey().ShowMonkey()
	usaFactory.CreateDog().ShowDog()
	usaFactory.CreateCat().ShowCat()
	/*
		优点：
			1. 拥有工厂方法模式的优点
			2. 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。
			3  增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。

		缺点：
			1. 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，甚至需要修改抽象层代码，这显然会带来较大的不便，违背了“开闭原则”。


	*/
}
