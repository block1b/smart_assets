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

	NickForm `json:"nick_form"`
}

// 前端请求设备信息
type GetIotInfoForm struct {
	ClientId string `json:"clientId"`
	Iot NickForm `json:"iot"`
}

// 创建设备
type NewDeviceForm struct {
	ClientId string `json:"clientId"`
	DeviceForm DeviceForm `json:"device_form"`
}

// 使用设备
type UserIotForm struct {
	ClientId string `json:"clientId"`
	User NickForm `json:"user"`
	Iot DeviceForm `json:"iot"`
}

// 获取用户账单
type GetUerBillsForm struct {
	ClientId string `json:"clientId"`
	User NickForm `json:"user"`
}

// 获取用户余额
type GetUerBalanceForm struct {
	ClientId string `json:"clientId"`
	User NickForm `json:"user"`
}

// 充值提现
type UseMoneyForm struct {
	ClientId string `json:"clientId"`
	AUser UseBalanceForm `json:"a_user"`
	BUser NickForm `json:"b_user"`
}

// post请求
type ClientIdTransfer struct {
	ClientId string `json:"clientId"`
	TransferData TransferPrepare `json:"transferData"`
}

// 余额响应
type BalanceAmount struct {
	Amount string `json:"amount"`
} 