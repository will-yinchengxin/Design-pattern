package main

import "fmt"

type Work struct{}

func (w *Work) Coding() {
	fmt.Println("Work Coding")
}

type Eat struct{}

func (e *Eat) Enjoy() {
	fmt.Println("Eat Enjoy")
}

func main() {
	// 工作就是coding⌨️
	w := new(Work)
	w.Coding()

	// 吃饭就是享受😎
	e := new(Eat)
	e.Enjoy()
}
