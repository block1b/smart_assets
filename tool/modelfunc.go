package tool

import "fmt"

//bob_public_key='3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm'
//bob_private_key='HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34'
const (
	ADMIN_NICK_NAME   = "Admin"
	ADMIN_PUBLIC_KEY  = "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
	ADMIN_PRIVATE_KEY = "HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34"
	DEFAULT_ACOUNT    = 1000  // 默认的钱包容量
	ADMIN_BALANCE_ASSET_ID = "c279f15ce6414a8c6e6e07313f93cf5c124caeeb30bf5a4ab8564c3fcdc626e3"
)

func (self *Sn) String() string {
	return fmt.Sprintf("%v.%v.%v", self.PublicKey, self.Type, self.Id)
}

func (self Recipient) ToList() []interface{} {
	key := []string{self.PublicKey}
	result := []interface{}{key, self.Count}
	return result
}
