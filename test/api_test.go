package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestTempNewDeviceForm(t *testing.T) {
	Byte, _ := json.Marshal(TempNewDeviceForm())
	fmt.Println(string(Byte))

	//req: smartServer/newIot
	//rep: smartServerPost WeChat/newIot
	//{"clientId":"WeChat","device_form":{"device_name":"clock0","device_info":"shareParking","status":"CanUse","ruler":"5","nick_form":{"nice_name":"block","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0"}}}
	// return
	// {"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": null, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBo0MskbRZIdJX1Bo0jzpIIbBmzEhVLW5LuscVCyq9P3FrSTvfQjGonZd9tDdAU2WE4gNFl1T7gUf75EbdgGiYE"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "1"}], "operation": "CREATE", "metadata": {"info": {"device_id": "clock0", "owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "user_nick_name": "null", "user_public_key": "null", "status": "CanUse", "ruler": "5", "start_time": "2019-05-13 04:23:40", "cost_time": "0"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "asset": {"data": {"info": {"owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "type": "iot", "id": "clock0", "device_name": "clock0", "device_info": "shareParking"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}}, "version": "2.0", "id": "74dcb60d222ba502519d40a09c2a0e26ee7784269d80d15eda9ee0cddcab1d03"}
	// check
	// http://192.168.113.6:9984/api/v1/assets/?search=3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0
	// body:[{"data": {"info": {"owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "type": "iot", "id": "clock0", "device_name": "clock0", "device_info": "shareParking"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "id": "74dcb60d222ba502519d40a09c2a0e26ee7784269d80d15eda9ee0cddcab1d03"}]

	err := errors.New("unexit:bbb")
	key := fmt.Sprint(err)
	if  key == "unexit:bbb"{
		fmt.Println("get unexit err")
	}

	topic := "WeChat/post"
	//sn := "1234.balance,main"
	ts := strings.Split(topic, "/")
	fmt.Println(ts)
	ts = strings.SplitN(topic, "/",0)
	fmt.Println(ts)
	ts = strings.SplitAfter(topic, "/")
	fmt.Println(ts)
	ts = strings.SplitAfterN(topic, "/", 0)
	fmt.Println(ts)

}
// 查看余额
func TestTempGetUerBalanceForm(t *testing.T) {
	Byte, _ := json.Marshal(TempGetUerBalanceForm())
	fmt.Println(string(Byte))
	//req: smartServer/balanceInfo GetUerBalanceForm
	//rep: WeChat/balanceInfo
	// admin
	//{"clientId":"WeChat","user":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}}
	// rep WeChat/post
	//{"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": null, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUCPN4bSg5xGzm_0c8CX88jqoy2qmXwEytgNiFWA0B_yxvs5dXQ71XihdZ675Fi6CJ2_NZezZOB1BMli0L4AJDoM"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "1000"}], "operation": "CREATE", "metadata": {"info": {"signer_nick_name": "Admin", "signer_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "recipient_nick_name": "Admin", "recipient_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "reason": "\u4e3b\u94b1\u5305\u521d\u59cb\u5316", "cost": "\u03e8", "time": "2019-05-17 02:36:07"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.balance.main"}, "asset": {"data": {"info": {"owner_nick_name": "Admin", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "type": "balance", "id": "main"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.balance.main"}}, "version": "2.0", "id": "c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}
	// turn pub WeChat/balanceAssetId
	// {"balance_asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"} // 113.6-2
	// {"balance_asset_id":"d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"} // 0.111
	// 余额1000
	// 测试合并资产
	// alice unmerge -> merge 100
	// done
}

// 充值|提现
func TestTempUseMoneyForm(t *testing.T) {
	Byte, _ := json.Marshal(TempUseMoneyForm())
	fmt.Println(string(Byte))
	// smartServer/useBalance, UserBalancePubHandler
	// 用户A充值50
	// {"clientId":"WeChat","a_user":{"cost_type":"recharge","money":"50","nice_name":"alice","private_key":"88L2BJC9eNtSWhpPwWqqsLDRGz7aBPhuRNyfsWx4QxWR","public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"},"b_user":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}}
	// postServer/post
	//{"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": {"transaction_id": "c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3", "output_index": 0}, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUD-Hf0ZfAWimQnAHBx2Oub-D3wd9LUKT-pH73mxcNIPJD8Ncr3ZwF6032li-fw9MLpkx-RmidvCctNqD_H-eesE"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "950"}, {"public_keys": ["HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"}, "uri": "ni:///sha-256;gq2QbZZEcAzcCuCv4xicDCYQ_0RmN09VLSH6t2H0pKQ?fpt=ed25519-sha-256&cost=131072"}, "amount": "50"}], "operation": "TRANSFER", "metadata": {"info": {"signer_nick_name": "Admin", "signer_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "recipient_nick_name": "alice", "recipient_public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq", "reason": "recharge", "cost": "50", "time": "2019-05-17 02:55:02"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.balance.main"}, "asset": {"id": "c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}, "version": "2.0", "id": "0c049a852f4768618c8229e2d022b815337d504d68e32b9fcedd1c0ee61fea99"}
    // check 主钱包变成了 950 元
    // smartServer/balanceInfo
    // {"clientId":"WeChat","user":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}}
    // check alice 有50元
	//{"clientId":"WeChat","user":{"nice_name":"admin","private_key":"88L2BJC9eNtSWhpPwWqqsLDRGz7aBPhuRNyfsWx4QxWR","public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}}
	// 继续充值50后，admin 900，alice 50+50
	// done
}

// 创建设备
func TestTempDeviceForm(t *testing.T) {
	Byte, _ := json.Marshal(TempNewDeviceForm())
	fmt.Println(string(Byte))
	// smartServer/newIot
	//{"clientId":"WeChat","device_form":{"device_name":"clock0","device_info":"shareParking","status":"CanUse","ruler":"5","nick_form":{"nice_name":"block","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":""}}}
	// rep
	// {"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": null, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUD9xM5swt9MUK9pvhRGiFI_3GvkEg8l05V1aW2CEPiJ6COTgcRBdN0AXlRwL2Spv_isokwO05SIB9nm8xgSXukP"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "1"}], "operation": "CREATE", "metadata": {"info": {"device_id": "clock0", "owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "user_nick_name": "null", "user_public_key": "null", "status": "CanUse", "ruler": "5", "start_time": "2019-05-19 01:26:15", "cost_time": "0"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "asset": {"data": {"info": {"owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "type": "iot", "id": "clock0", "device_name": "clock0", "device_info": "shareParking"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}}, "version": "2.0", "id": "027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}
	// admin 注册了该设备
	// {"iot_asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}
	// done
}


// 查看设备
func TestTempGetIotInfoForm(t *testing.T) {
	Byte, _ := json.Marshal(TempGetIotInfoForm())
	fmt.Println(string(Byte))
	// smartServer/iotInfo
	//{"clientId":"WeChat","iot":{"nice_name":"block","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}}
	// rep 设备名，设备信息，状态，收费规则
	//{"device_name":"clock0","device_info":"shareParking","status":"CanUse","ruler":"5","nick_form":{"nice_name":"","private_key":"","public_key":"","type":"","id":"","asset_id":""}}
	// done
	// 前端调试
	//{"clientId":"WeChat","iot":{"nice_name":"block","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}}
}
// 租用
func TestTempUserIotForm(t *testing.T) {
	Byte, _ := json.Marshal(TempUserIotForm())
	fmt.Println(string(Byte))
	// smartServer/rentIot
	// 租用
	//{"clientId":"WeChat","user":{"nice_name":"alice","private_key":"88L2BJC9eNtSWhpPwWqqsLDRGz7aBPhuRNyfsWx4QxWR","public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","type":"balance","id":"main","asset_id":"d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"},"iot":{"device_name":"clock0","device_info":"shareParking","status":"Rent","ruler":"5","nick_form":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}}}
	// rep
	//{"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": {"transaction_id": "027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1", "output_index": 0}, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBQALbf22TxHss1L9C_pMvpgdEnI5x5VekIsY-11TfIghFAVhjQPqeU4YtNkmVSX1FEdWdboFQznz6wuvZ2g4cI"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "1"}], "operation": "TRANSFER", "metadata": {"info": {"device_id": "clock0", "owner_nick_name": "admin", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "user_nick_name": "alice", "user_public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq", "status": "InUse", "ruler": "5", "start_time": "2019-05-19 04:39:11", "cost_time": "0"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "asset": {"id": "027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}, "version": "2.0", "id": "d29804fc6e1a3acfcfc0fd86d799feddbcb27eec0e5dd40d138eeddd592ad925"}

	// smartServer/returnIot
	// 归还
	//{"clientId":"WeChat","user":{"nice_name":"alice","private_key":"88L2BJC9eNtSWhpPwWqqsLDRGz7aBPhuRNyfsWx4QxWR","public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","type":"balance","id":"main","asset_id":"d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"},"iot":{"device_name":"clock0","device_info":"shareParking","status":"Return","ruler":"5","nick_form":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}}}
	//rep
	//{"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": {"transaction_id": "d29804fc6e1a3acfcfc0fd86d799feddbcb27eec0e5dd40d138eeddd592ad925", "output_index": 0}, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBzzX0F7vvtOhrTLED4_FqrYx-R_LoaC-2vz5roTKw6J2UkZT-gm8gNGBu-GCCkBYS2gu3AfVpkKzTcH4fWmoYB"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "1"}], "operation": "TRANSFER", "metadata": {"info": {"device_id": "clock0", "owner_nick_name": "admin", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "user_nick_name": "admin", "user_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "status": "CanUse", "ruler": "5", "start_time": "2019-05-19 04:40:21", "cost_time": "\u3886"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "asset": {"id": "027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}, "version": "2.0", "id": "57ea18c4d9d8d7815afcdd043c1073677e8dd1a34b00ead87b435b40ac3ef2bb"}
	// check : alice 98 done；admin merge 902 done；device status : CanUse
	// done
}
// 归还 done

// 查看账单
func TestTempGetUerBillsForm(t *testing.T) {
	Byte, _ := json.Marshal(TempGetUerBillsForm())
	fmt.Println(string(Byte))
	// smartServer/billInfo
	//{"clientId":"WeChat","user":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"}}
	// rep
	//[{"signer_nick_name":"Admin","signer_public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","recipient_nick_name":"Admin","recipient_public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","reason":"主钱包初始化","cost":"Ϩ","time":"2019-05-18 02:43:54"},{"signer_nick_name":"Admin","signer_public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","recipient_nick_name":"alice","recipient_public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","reason":"recharge","cost":"50","time":"2019-05-18 03:15:22"},{"signer_nick_name":"Admin","signer_public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","recipient_nick_name":"alice","recipient_public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","reason":"recharge","cost":"50","time":"2019-05-18 03:41:07"},{"signer_nick_name":"alice","signer_public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","recipient_nick_name":"admin","recipient_public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","reason":"payment","cost":"2","time":"2019-05-19 04:40:21"},{"signer_nick_name":"admin","signer_public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","recipient_nick_name":"admin","recipient_public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","reason":"合并余额","cost":"Ά","time":"2019-05-19 04:43:42"}]
	// done
}

// 修改设备共享状态
func TestTempUpdateIot(t *testing.T)  {
	Byte, _ := json.Marshal(TempUpdateIot())
	fmt.Println(string(Byte))
	// smartServer/updateIot
	// 关闭
	// {"clientId":"WeChat","user":{"nice_name":"admin","private_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"},"iot":{"device_name":"clock0","device_info":"shareParking","status":"Close","ruler":"5","nick_form":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}}}
	// rep
	//{"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": {"transaction_id": "485419c676e64c779b96d70be20675079e2cc96c42e2f048b27367f80c3c0ae3", "output_index": 0}, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBZgEl-y5uTVlFsxFi3knmo6hS0gB4yMIBWMcUGw3z6W06Wu6n9d7b2bXa5-BgsdvKTaQm4XstOaL6eOARrU9AO"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "1"}], "operation": "TRANSFER", "metadata": {"info": {"device_id": "clock0", "owner_nick_name": "admin", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "user_nick_name": "admin", "user_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "status": "UnUse", "ruler": "5", "start_time": "2019-05-19 05:37:11", "cost_time": "0"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "asset": {"id": "027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}, "version": "2.0", "id": "6b2499ea9283d0112cc6e89d6e01a33fd54209e37f037df74687306b8bf41ad5"}
	// check 设备状态 Unuse done
	// 开启
	// {"clientId":"WeChat","user":{"nice_name":"admin","private_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"},"iot":{"device_name":"clock0","device_info":"shareParking","status":"Open","ruler":"5","nick_form":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}}}
	// check ：CanUse done
}
