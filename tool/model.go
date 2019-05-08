package tool

type Asset struct {
	Data interface{} `json:"data"`
	Ns string `json:"ns"`  // 唯一标识
}

