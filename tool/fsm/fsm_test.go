package fsm

import (
	"fmt"
	"testing"
)

func TestFSM_Call(t *testing.T) {
	efan := Init(FSMState("CanUse"))
	efan.Call(FSMEvent("Close"))  // 其实该用event 不是state
	newStatus := string(efan.GetState())
	fmt.Println(newStatus)
}
