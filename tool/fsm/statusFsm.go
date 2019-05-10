package fsm

import "fmt"

//State:
//可用，占用，禁用
//Event:
//租用，归还，关闭
//Handler:

var (
	CanUse        = FSMState("CanUse")
	InUse         = FSMState("InUse")
	UnUse         = FSMState("UnUse")
	Err           = FSMState("Err")

	RentEvent     = FSMEvent("Rent")
	ReturnEvent   = FSMEvent("Return")
	CloseEvent    = FSMEvent("Close")
	OpenEvent     = FSMEvent("Open")

	RentHandler = FSMHandler(func() FSMState {
		fmt.Println("设备已出租")
		return InUse
	})
	ReturnHandler = FSMHandler(func() FSMState {
		fmt.Println("设备已归还")
		return CanUse
	})
	CloseHandler = FSMHandler(func() FSMState {
		fmt.Println("设备禁用")
		return UnUse
	})
	OpenHandler = FSMHandler(func() FSMState {
		fmt.Println("设备启用")
		return CanUse
	})
	ErrHandler = FSMHandler(func() FSMState {
		fmt.Println("不可操作")
		return Err
	})
)

// 设备
type ElectricFan struct {
	*FSM
}

// 实例化设备
func NewElectricFan(initState FSMState) *ElectricFan {
	return &ElectricFan{
		FSM: NewFSM(initState),
	}
}

// 入口函数
func Init(initState FSMState) *ElectricFan{
	//RentEvent     = FSMEvent("租用")
	//ReturnEvent   = FSMEvent("归还")
	//CloseEvent    = FSMEvent("关闭")
	//OpenEvent     = FSMEvent("开启")

	efan := NewElectricFan(initState) // 初始状态是可用的
	// 可用状态
	efan.AddHandler(CanUse, RentEvent, RentHandler)
	efan.AddHandler(CanUse, ReturnEvent, ErrHandler)
	efan.AddHandler(CanUse, CloseEvent, CloseHandler)
	efan.AddHandler(CanUse, OpenEvent, ErrHandler)
	// 占用状态
	efan.AddHandler(InUse, RentEvent, ErrHandler)
	efan.AddHandler(InUse, ReturnEvent, ReturnHandler)
	efan.AddHandler(InUse, CloseEvent, ErrHandler)
	efan.AddHandler(InUse, OpenEvent, ErrHandler)
	// 禁用状态
	efan.AddHandler(UnUse, RentEvent, ErrHandler)
	efan.AddHandler(UnUse, ReturnEvent, ErrHandler)
	efan.AddHandler(UnUse, CloseEvent, ErrHandler)
	efan.AddHandler(UnUse, OpenEvent, OpenHandler)

	// 开始测试状态变化
	//efan.Call(CanUseEvent)   // 按下关闭按钮
	return efan
}
