package main

import "fmt"

/*
将一个请求封装为一个对象，从而让我们可用不同的请求对客户进行参数化；对请求排队或者记录请求日志，以及支持可撤销的操作。命令模式是一种对象行为型模式，其别名为动作(Action)模式或事务(Transaction)模式。
命令模式可以将请求发送者和接收者完全解耦，发送者与接收者之间没有直接引用关系，发送请求的对象只需要知道如何发送请求，而不必知道如何完成请求。
*/

type Doctor struct{}

func (d *Doctor) DoctorTreatEye() {
	fmt.Println("治疗眼睛")
}
func (d *Doctor) DoctorTreatNose() {
	fmt.Println("治疗鼻子")
}

type Command interface {
	Treat()
}

type TreatEye struct {
	d *Doctor
}

func (t *TreatEye) Treat() {
	t.d.DoctorTreatEye()
}

type TreatNose struct {
	d *Doctor
}

func (t *TreatNose) Treat() {
	t.d.DoctorTreatNose()
}

type Nurse struct {
	cmdList []Command
}

func (n *Nurse) CallDoctor() {
	if len(n.cmdList) == 0 {
		fmt.Println("没人，贼清闲！")
		return
	}
	for _, val := range n.cmdList {
		val.Treat()
	}
}

func main() {
	// 找个护士上榜
	nurse := new(Nurse)
	// 打榜展示今天可以挂号的科室
	treatEye := &TreatEye{
		new(Doctor),
	}
	treatNose := &TreatNose{
		new(Doctor),
	}

	// 来了病人了。。。。 嗷嗷挂号
	nurse.cmdList = append(nurse.cmdList, treatEye)
	nurse.cmdList = append(nurse.cmdList, treatNose)
	nurse.CallDoctor()

	/*
		优点：
			(1) 降低系统的耦合度。由于请求者与接收者之间不存在直接引用，因此请求者与接收者之间实现完全解耦，相同的请求者可以对应不同的接收者，同样，相同的接收者也可以供不同的请求者使用，两者之间具有良好的独立性。
			(2) 新的命令可以很容易地加入到系统中。由于增加新的具体命令类不会影响到其他类，因此增加新的具体命令类很容易，无须修改原有系统源代码，甚至客户类代码，满足“开闭原则”的要求。
			(3) 可以比较容易地设计一个命令队列或宏命令（组合命令）。
		缺点：
			使用命令模式可能会导致某些系统有过多的具体命令类。因为针对每一个对请求接收者的调用操作都需要设计一个具体命令类，因此在某些系统中可能需要提供大量的具体命令类，这将影响命令模式的使用。

		使用场景：
			(1) 系统需要将请求调用者和请求接收者解耦，使得调用者和接收者不直接交互。请求调用者无须知道接收者的存在，也无须知道接收者是谁，接收者也无须关心何时被调用。
			(2) 系统需要在不同的时间指定请求、将请求排队和执行请求。一个命令对象和请求的初始调用者可以有不同的生命期，换言之，最初的请求发出者可能已经不在了，而命令对象本身仍然是活动的，可以通过该命令对象去调用请求接收者，而无须关心请求调用者的存在性，可以通过请求日志文件等机制来具体实现。
			(3) 系统需要将一组操作组合在一起形成宏命令。
	*/
}
