package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPrepareData(t *testing.T)  {

	Byte, _ := json.Marshal(TempGetUerBalanceForm())
	fmt.Println(string(Byte))

	//req: smartServer/newInfo TempTransferPrepare()
	//rep: smartServerPost WeChat/newInfo
	//{"clientId":"WeChat","device_form":{"device_name":"clock0","device_info":"shareParking","status":"CanUse","ruler":"5","nick_form":{"nice_name":"block","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0"}}}
	// return
	// {"inputs": [{"owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "fulfills": null, "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBo0MskbRZIdJX1Bo0jzpIIbBmzEhVLW5LuscVCyq9P3FrSTvfQjGonZd9tDdAU2WE4gNFl1T7gUf75EbdgGiYE"}], "outputs": [{"public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"], "condition": {"details": {"type": "ed25519-sha-256", "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"}, "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"}, "amount": "1"}], "operation": "CREATE", "metadata": {"info": {"device_id": "clock0", "owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "user_nick_name": "null", "user_public_key": "null", "status": "CanUse", "ruler": "5", "start_time": "2019-05-13 04:23:40", "cost_time": "0"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "asset": {"data": {"info": {"owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "type": "iot", "id": "clock0", "device_name": "clock0", "device_info": "shareParking"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}}, "version": "2.0", "id": "74dcb60d222ba502519d40a09c2a0e26ee7784269d80d15eda9ee0cddcab1d03"}
	// check
	// http://192.168.113.6:9984/api/v1/assets/?search=3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0
	// body:[{"data": {"info": {"owner_nick_name": "block", "owner_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm", "type": "iot", "id": "clock0", "device_name": "clock0", "device_info": "shareParking"}, "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.iot.clock0"}, "id": "74dcb60d222ba502519d40a09c2a0e26ee7784269d80d15eda9ee0cddcab1d03"}]

	//req: smartServer/balanceInfo GetUerBalanceForm
	//rep: WeChat/balanceInfo
	//{"clientId":"WeChat","user":{"nice_name":"block","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"balance","id":"main"}}

}
