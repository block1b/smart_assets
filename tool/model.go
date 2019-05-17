package tool

// 资产唯一标识
type Sn struct {
	PublicKey string `json:"public_key"`
	Type string `json:"type"`  // [balance|iot]
	Id string `json:"id"`

	AssetId string `json:"asset_id"`
}

// 自定义结构体，资产信息和元数据
type Info interface {
}

// 余额资产信息
type BalanceInfo struct {
	OwnerNickName string  `json:"owner_nick_name"`
	OwnerPublicKey string `json:"owner_public_key"`
	Type string           `json:"type"`  // default：balance
	Id string             `json:"id"`    // 钱包编号
}

// 设备资产信息
type IotInfo struct {
	OwnerNickName string  `json:"owner_nick_name"`
	OwnerPublicKey string `json:"owner_public_key"`
	Type string           `json:"type"`  // default：iot
	Id string             `json:"id"`    // 设备编号
	DeviceName string     `json:"device_name"`  // 设备名
	DeviceInfo string     `json:"device_info"`  // 设备描述
}

// 转账记录
type BillInfo struct {
	SignerNickName string       `json:"signer_nick_name"`  // 发起人
	SignerPublicKey string      `json:"signer_public_key"`
	RecipientNickName string    `json:"recipient_nick_name"`  // 收款人
	RecipientPublicKey string   `json:"recipient_public_key"`
	Reason string               `json:"reason"`  // 支付原因
	Cost string                 `json:"cost"`  // 支付金额
	Time string                 `json:"time"`  // 支付时间
}

// 租用记录
type RentInfo struct {
	DeviceId string `json:"device_id"`  // 设备号
	OwnerNickName string  `json:"owner_nick_name"`
	OwnerPublicKey string `json:"owner_public_key"`
	UserNickName string `json:"user_nick_name"`
	UserPublicKey string `json:"user_public_key"`
	Status string `json:"status"`  // 设备状态
	Ruler string `json:"ruler"`  // 收费规则
	StartTime string `json:"start_time"` // 开始租用时间
	CostTime string `json:"cost_time"`  // 租用时间
}

type Data struct {
	Info Info          `json:"info"`
	Sn string          `json:"sn"`  // asset,metadata的唯一标识 public_key+type+id
}

type Asset struct {
	Id string             `json:"id, omitempty"`  // asset_id
	Data Data             `json:"data, omitempty"`
}

type Transaction struct {
	Inputs []Input        `json:"inputs"`
	Outputs []Output      `json:"outputs"`
	Operation string      `json:"operation"`
	Metadata Data         `json:"metadata"`
	Asset Asset           `json:"asset"`  // Tran只有id；Create还有data
	Version string        `json:"version"`
	Id string             `json:"id"`  // TransactionId

}

type TilfilledTransaction struct {
	Inputs []TilfilledInput  `json:"inputs"`
	Outputs []Output      `json:"outputs"`
	Operation string      `json:"operation"`
	Metadata Data         `json:"metadata"`
	Asset Asset           `json:"asset"`  // Tran只有id；Create还有data
	Version string        `json:"version"`
	Id string             `json:"id"`  // TransactionId
}

type TilfilledInput struct {
	OwnersBefore []string `json:"owners_before"`
	Fulfills Fulfills     `json:"fulfills"`
	Fulfillment string    `json:"fulfillment"`
}

type Input struct {
	OwnersBefore []string `json:"owners_before"`
	Fulfills Fulfills     `json:"fulfills"`
	Fulfillment Details    `json:"fulfillment"`
}

type Fulfills struct {
	TransactionId string  `json:"transaction_id"`
	OutputIndex int       `json:"output_index"`
}

type Output struct {
	PublicKeys []string   `json:"public_keys"`
	Condition Condition   `json:"condition"`
	Amount string         `json:"amount"`
}

type Condition struct {
	Details Details       `json:"details"`
	Uri string            `json:"uri"`
}

type Details struct {
	Type string           `json:"type"`
	PublicKey string      `json:"public_key"`
}

type GetOutputResult struct {
	Fulfills
}

type GetAssetResult struct {
	Asset
}

type GetMetadataResult struct {
	Metadata Data `json:"metadata"`
	Id string `json:"id"`  // TransactionID
}

// 用于传输的Transfer prepare数据结构
//"operation":"TRANSFER",
//"asset":{"id":"123456"},
//"inputs":[{"input":"input_msg"}],
//"recipients": [[["public_key1"],2],[["public_key2",6]]],
//"private_keys": ["p1","p2"]

//operation='CREATE',
//signers=alice.public_key,
//recipients=[([bob.public_key], 10)],
//asset=game_boy_token
//inputs=transfer_input,

type TransferPrepare struct {
	Operation string `json:"operation"`
	Asset Asset      `json:"asset"`
	Signers string   `json:"signers, omitempty"`
	Inputs []Input   `json:"inputs, omitempty"`
	Recipients []interface{} `json:"recipients"`
	PrivateKeys []string `json:"private_keys"`
	Metadata Data `json:"metadata"`
}

//[["public_key1"],2]
type Recipient struct {
	PublicKey string
	Count int
}
