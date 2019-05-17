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
	// {"balance_asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}
	// 余额1000
}

// 充值|提现
func TestTempUseMoneyForm(t *testing.T) {
	Byte, _ := json.Marshal(TempUseMoneyForm())
	fmt.Println(string(Byte))
	// smartServer/useBalance", UserBalancePubHandler
	// 用户A充值50
	// {"clientId":"WeChat","a_user":{"cost_type":"recharge","money":"50","nice_name":"alice","private_key":"88L2BJC9eNtSWhpPwWqqsLDRGz7aBPhuRNyfsWx4QxWR","public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"},"b_user":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}}
	// postServer/post
	//{"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": {"transaction_id": "c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3", "output_index": 0}, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUD-Hf0ZfAWimQnAHBx2Oub-D3wd9LUKT-pH73mxcNIPJD8Ncr3ZwF6032li-fw9MLpkx-RmidvCctNqD_H-eesE"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "950"}, {"public_keys": ["HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"}, "uri": "ni:///sha-256;gq2QbZZEcAzcCuCv4xicDCYQ_0RmN09VLSH6t2H0pKQ?fpt=ed25519-sha-256&cost=131072"}, "amount": "50"}], "operation": "TRANSFER", "metadata": {"info": {"signer_nick_name": "Admin", "signer_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "recipient_nick_name": "alice", "recipient_public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq", "reason": "recharge", "cost": "50", "time": "2019-05-17 02:55:02"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.balance.main"}, "asset": {"id": "c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}, "version": "2.0", "id": "0c049a852f4768618c8229e2d022b815337d504d68e32b9fcedd1c0ee61fea99"}
    // check 主钱包变成了 950 元
    // smartServer/balanceInfo
    // {"clientId":"WeChat","user":{"nice_name":"admin","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}}
    // check alice 有50元
	//{"clientId":"WeChat","user":{"nice_name":"admin","private_key":"88L2BJC9eNtSWhpPwWqqsLDRGz7aBPhuRNyfsWx4QxWR","public_key":"HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq","type":"balance","id":"main","asset_id":"c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"}}

}