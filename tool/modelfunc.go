package tool

import "fmt"

//bob_public_key='3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm'
//bob_private_key='HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34'
const (
	ADMIN_NICK_NAME   = "BLOCK.1.B"
	ADMIN_PUBLIC_KEY  = "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
	ADMIN_PRIVATE_KEY = "HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34"
)

func (self *Sn) String() string {
	return fmt.Sprintf("%v.%v.%v", self.PublicKey, self.Type, self.Id)
}

func (self Recipient) ToList() []interface{} {
	key := []string{self.PublicKey}
	result := []interface{}{key, self.Count}
	return result
}
