package main

import "fmt"

// project

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

// student

type Will struct {
}

func (w *Will) LearnMath(math Math) {
	fmt.Println("will learn")
	math.HomeWork()
}
func (w *Will) LearnChinese(chinese Chinese) {
	fmt.Println("will learn")
	chinese.HomeWork()
}

type Yin struct {
}

func (w *Yin) LearnMath(math Math) {
	fmt.Println("yin learn")
	math.HomeWork()
}
func (w *Yin) LearnChinese(chinese Chinese) {
	fmt.Println("yin learn")
	chinese.HomeWork()
}

// 业务逻辑
func main() {
	var math Math
	math = Math{}
	var chin Chinese
	chin = Chinese{}
	var stu *Will
	stu = &Will{}
	stu.LearnMath(math)
	stu.LearnChinese(chin)

	fmt.Println("--------------------")

	var math1 Math
	math1 = Math{}
	var chin1 Chinese
	chin1 = Chinese{}
	var stu1 *Yin
	stu1 = &Yin{}
	stu1.LearnMath(math1)
	stu1.LearnChinese(chin1)

	// 每个模块之间的依赖关系，实际上并没有用到任何的interface接口层的代码
	// 这种设计的问题就在于，小规模没什么问题，但是一旦程序需要扩展，比如我现在要增加一个 学科英语 或者 学生张三， 那么模块和模块的依赖关系将成指数级递增，想蜘蛛网一样越来越难维护和捋顺。
}
