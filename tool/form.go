package tool

// 接收字前端的表单
// 用户唯一标识结构
type NickForm struct {
	NiceName string `json:"nice_name"`
	PrivateKey string `json:"private_key"`
	Sn
}
// 用户提交的充值表单
type UseBalanceForm struct {
	CostMoney
	NickForm
}
// 前端选择的花费金额
type CostMoney struct {
	CostType string `json:"cost_type"`
	Money string `json:"money"`
}

type DeviceForm struct {
	DeviceName string     `json:"device_name"`  // 设备名
	DeviceInfo string     `json:"device_info"`  // 设备描述

	Status string `json:"status"`  // 设备状态
	Ruler string `json:"ruler"`  // 收费规则

	NickForm
}

//