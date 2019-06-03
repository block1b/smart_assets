package tool

import "fmt"

//bob_public_key='3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm'
//bob_private_key='HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34'
const (
	ADMIN_NICK_NAME   = "Admin"
	ADMIN_PUBLIC_KEY  = "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
	ADMIN_PRIVATE_KEY = "HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34"
	DEFAULT_ACOUNT    = 1000  // 默认的钱包容量
	//ADMIN_BALANCE_ASSET_ID = "c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"
	ADMIN_BALANCE_ASSET_ID = "d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"
	ADMIN_IOT_ASSET_ID = "027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"

	CLIENTID  = "smartServer"
	MQTT_BROKER_ADDR = "192.168.0.111:1883"
    bigchaindb_addr  = "192.168.0.111:9984"
)

func (self *Sn) String() string {
	return fmt.Sprintf("%v.%v.%v", self.PublicKey, self.Type, self.Id)
}

func (self Recipient) ToList() []interface{} {
	key := []string{self.PublicKey}
	result := []interface{}{key, self.Count}
	return result
}
