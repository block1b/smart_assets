package test

import (
	"fmt"
	. "smart_assets/tool"
)

const alice_private_key = "88L2BJC9eNtSWhpPwWqqsLDRGz7aBPhuRNyfsWx4QxWR"
const alice_public_key = "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"


func prepareData()  {
	fmt.Println()
	// 生成数据
	//var getIotInfoForm GetIotInfoForm
	//var newDeviceForm NewDeviceForm
	//var userIotForm UserIotForm
	//var getUerBillsForm GetUerBillsForm
	//var getUerBalanceForm GetUerBalanceForm
	//var useMoneyForm UseMoneyForm
}

func TempGetIotInfoForm() GetIotInfoForm {
	return GetIotInfoForm{
		//ClientId string `json:"clientId"`
		//Iot NickForm `json:"iot"`
			//NiceName string `json:"nice_name"`
			//PrivateKey string `json:"private_key"`
			//Sn
				//PublicKey string `json:"public_key"`
				//Type string `json:"type"`  // [balance|iot]
				//Id string `json:"id"`
		ClientId:"WeChat",
		Iot:NickForm{
			NiceName:"block",
			PrivateKey:ADMIN_PRIVATE_KEY,  // may null
			Sn:Sn{
				PublicKey:ADMIN_PUBLIC_KEY, // not null
				Type:"iot",
				Id:"clock0",
			},
		},
	}
}

//NewDeviceForm
func TempNewDeviceForm() NewDeviceForm {
	return NewDeviceForm{
		//ClientId string `json:"clientId"`
		//DeviceForm DeviceForm `json:"device_form"`
			//DeviceName string     `json:"device_name"`  // 设备名
			//DeviceInfo string     `json:"device_info"`  // 设备描述
			//Status string `json:"status"`  // 设备状态
			//Ruler string `json:"ruler"`  // 收费规则
			//NickForm
				//NiceName string `json:"nice_name"`
				//PrivateKey string `json:"private_key"`
				//Sn
					//PublicKey string `json:"public_key"`
					//Type string `json:"type"`  // [balance|iot]
					//Id string `json:"id"`
		ClientId:"WeChat",
		DeviceForm:DeviceForm{
			DeviceName:"clock0",
			DeviceInfo:"shareParking",
			Status:"CanUse",
			Ruler:"5",
			NickForm:NickForm{
				NiceName:"block",
				PrivateKey:ADMIN_PRIVATE_KEY,
				Sn:Sn{
					PublicKey:ADMIN_PUBLIC_KEY, // not null
					Type:"iot",
					Id:"clock0",
				},
			},
		},

	}

}

//UserIotForm
func TempUserIotForm() UserIotForm {
	return UserIotForm{
		//ClientId string `json:"clientId"`
		//User NickForm `json:"user"`
		//Iot DeviceForm `json:"iot"`
		ClientId:"WeChat",
		User:NickForm{
			NiceName:"block",
			PrivateKey:ADMIN_PRIVATE_KEY,
			Sn:Sn{
				PublicKey:ADMIN_PUBLIC_KEY, // not null
				Type:"balance",
				Id:"main",
			},
		},
		Iot:DeviceForm{
			DeviceName:"clock0",
			DeviceInfo:"shareParking",
			Status:"rent",
			Ruler:"5",
			NickForm:NickForm{
				NiceName:"block",
				PrivateKey:ADMIN_PRIVATE_KEY,
				Sn:Sn{
					PublicKey:ADMIN_PUBLIC_KEY, // not null
					Type:"iot",
					Id:"clock0",
				},
			},
		},
	}

}

//GetUerBillsForm
func TempGetUerBillsForm() GetUerBillsForm {
	return GetUerBillsForm{
		//ClientId string `json:"clientId"`
		//User NickForm `json:"user"`
		ClientId:"WeChat",
		User:NickForm{
			NiceName:"block",
			PrivateKey:ADMIN_PRIVATE_KEY,
			Sn:Sn{
				PublicKey:ADMIN_PUBLIC_KEY, // not null
				Type:"balance",
				Id:"main",
			},
		},
	}
}

//GetUerBalanceForm 同 GetUerBillsForm
func TempGetUerBalanceForm() GetUerBalanceForm  {
	return GetUerBalanceForm{
		//ClientId string `json:"clientId"`
		//User NickForm `json:"user"`
		ClientId:"WeChat",
		User:NickForm{
			NiceName:"admin",
			PrivateKey:alice_private_key,
			Sn:Sn{
				PublicKey:alice_public_key, // not null
				Type:"balance",
				Id:"main",

				AssetId:ADMIN_BALANCE_ASSET_ID,  // 创建后填写
			},
		},
	}
}

//UseMoneyForm
func TempUseMoneyForm() UseMoneyForm {
	return UseMoneyForm{
		//ClientId string `json:"clientId"`
		//AUser UseBalanceForm `json:"a_user"`
		//BUser NickForm `json:"b_user"`
		ClientId:"WeChat",
		AUser:UseBalanceForm{
			//CostMoney
			//NickForm
			CostMoney:CostMoney{
				//CostType string `json:"cost_type"`
				//Money string `json:"money"`
				CostType:"recharge",  // recharge|withdrawal
				Money:"50",
			},
			NickForm:NickForm{
				NiceName:"alice",
				PrivateKey:alice_private_key,
				Sn:Sn{
					PublicKey:alice_public_key, // not null
					Type:"balance",
					Id:"main",
					AssetId:ADMIN_BALANCE_ASSET_ID,
				},
			},
		},
		BUser:NickForm{
			NiceName:"admin",
			PrivateKey:ADMIN_PRIVATE_KEY,
			Sn:Sn{
				PublicKey:ADMIN_PUBLIC_KEY, // not null
				Type:"balance",
				Id:"main",

				AssetId:ADMIN_BALANCE_ASSET_ID,
			},
		},
	}
}

// return result
// DeviceForm
// []GetMetadataResult

//	Metadata Data `json:"metadata"`
//	Id string `json:"id"`  // TransactionID

func TempDeviceForm() DeviceForm {
	return DeviceForm{
		DeviceName:"clock0",
		DeviceInfo:"shareParking",
		Status:"CanUse",
		Ruler:"5",
		NickForm:NickForm{
			NiceName:"block",
			PrivateKey:ADMIN_PRIVATE_KEY,
			Sn:Sn{
				PublicKey:ADMIN_PUBLIC_KEY, // not null
				Type:"iot",
				Id:"clock0",
			},
		},

	}
}

//[]GetMetadataResult
func TempMetadataResult() []GetMetadataResult {
	return []GetMetadataResult{
		GetMetadataResult{
			//Metadata Data `json:"metadata"`
			//Id string `json:"id"`  // TransactionID
			Metadata:Data{
				Info:"",
				Sn:"",
			},
		},
	}
}

//TransferPrepare
func TempTransferPrepare() TransferPrepare {
	return TransferPrepare{
		//Operation string `json:"operation"`
		//Asset Asset      `json:"asset"`
		//Signers string   `json:"signers, omitempty"`
		//Inputs []Input   `json:"inputs, omitempty"`
		//Recipients []interface{} `json:"recipients"`
		//PrivateKeys []string `json:"private_keys"`
		//Metadata Data `json:"metadata"`
		Signers:"block",
	}
}