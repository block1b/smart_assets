package tool

// 自定义结构体，资产信息和元数据
type Data interface {
}

type Asset struct {
	Id string             `json:"id"`
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

type Input struct {
	OwnersBefore []string `json:"owners_before"`
	Fulfills Fulfills     `json:"fulfills"`
	Fulfillment string    `json:"fulfillment"`
}

type Fulfills struct {
	TransactionId string  `json:"transaction_id"`
	outputIndex int       `json:"output_index"`
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
	TransactionId string `json:"transaction_id"`
	OutputIndex int `json:"output_index"`
}

type GetAssetResult struct {
	Data Data `json:"data"`
	Id string `json:"id"`  // AssetId
}

type GetMetadataResult struct {
	Metadata Data `json:"metadata"`
	Id string `json:"id"`  // TransactionID
}