package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const bigchaindb_addr  = "192.168.113.6:9984"

// http get 查询 bigchaindb 的查询接口

func GetAsset(search string)([]byte, error){
	params := url.Values{}
	Url, err := url.Parse("http://"+bigchaindb_addr+"/api/v1/assets")
	if err != nil {
		panic(err.Error())
	}
	params.Set("search", search)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("read resp.Body ",err)
		return nil,err
	}
	return s, nil
}

func GetOutputs(publicKey,spent string)([]byte, error){
	params := url.Values{}
	Url, err := url.Parse("http://"+bigchaindb_addr+"/api/v1/outputs")
	if err != nil {
		panic(err.Error())
	}
	params.Set("public_key", publicKey)
	params.Set("spent", spent)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("read resp.Body ",err)
		return nil,err
	}
	return s, nil
}

// get transfer by transfer_id
func GetTransactionById(transaction_id string) ([]byte,error) {
	params := url.Values{}
	Url, err := url.Parse("http://"+bigchaindb_addr+"/api/v1/transactions/"+transaction_id)
	if err != nil {
		panic(err.Error())
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("read resp.Body ",err)
		return nil,err
	}
	return s, nil
}

//# 查询余额
//当前端查看余额时，本用例开始执行；
//## 基本流
//1. 前端获取 sn：用户昵称+公钥+钱包编号 提交到ss（smartServer）；
//2. ss 查询余额，余额资产中按sn查询资产交易&未消耗的outputs，余额结果返回。
//## 可选流
//1. 新用户，无余额资产，管理员创建该用户的余额资产，初始化为0，重新查询。
//2. 余额分散，使用该用户账号合并余额资产，重新查询。

func AssetQuery(args NickForm) (string, error) {
	sn := args.Sn
	snStr := sn.String()
	fmt.Println("select sn ", snStr)
	assetsByte, err := GetAsset(snStr)
	if err != nil{
		fmt.Println("get asset ",err)
		return "", err
	}
	fmt.Println("assets", string(assetsByte))
	var assets []GetAssetResult
	err = json.Unmarshal(assetsByte, &assets)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return "", err
	}
	asset_id := assets[0].Id
	fmt.Println("余额资产id", asset_id)
	publicKey := sn.PublicKey
	outputsByte, err := GetOutputs(publicKey,"true")
	if err != nil{
		fmt.Println("get asset ",err)
		return "", err
	}
	var getOutPutResults []GetOutputResult
	err = json.Unmarshal(outputsByte, &getOutPutResults)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return "", err
	}
	var unSpentOutputs []Output
	// 包括balance和iot的未消耗outputs，过滤
	for _, getOutPutResult := range getOutPutResults{
		transactionByte, err := GetTransactionById(getOutPutResult.TransactionId)
		if err != nil{
			fmt.Println("get transaction",err)
			return "", err
		}
		var transaction Transaction
		err = json.Unmarshal(transactionByte, &transaction)
		if err != nil{
			fmt.Println("unmarshal ",err)
			return "", err
		}
		// 按asset_id过滤 还好struct自己解决了无Id的问题 CREATE 无id
		if transaction.Asset.Id == asset_id{
			fmt.Println("unspent output")
			unSpentOutputs = append(unSpentOutputs, transaction.Outputs[getOutPutResult.OutputIndex])
		}
	}
	unSpentOutputsLen := len(unSpentOutputs)
	switch unSpentOutputsLen {
	case 0:
		fmt.Println("该用户无余额")
		// todo 新建用户余额资产
		err = CreateBalanceAsset(args)
		if err != nil{
			fmt.Println("create ", err)
			return "", err
		}
		// 返回0值
		return "0", nil
	case 1:
		fmt.Println("该用户余额为", unSpentOutputs[0].Amount)
		return unSpentOutputs[0].Amount, nil
	default:
		fmt.Println("该用于余额token需要合并")
		// todo 合并用户资产
	}

	return "",nil
}

// 创建余额资产;
// 准备数据
//byte_data = {
//"operation":"TRANSFER",
//"asset":{"id":"123456"},
//"inputs":[{"input":"input_msg"}],
//"recipients": [[["public_key1"],2],[["public_key2",6]]],
//"private_keys": ["p1","p2"]
//}
func CreateBalanceAsset(args NickForm) error {
	// prepare data
	operation := "CREATE"
	balanceInfo := BalanceInfo{
		OwnerNickName: args.NiceName,
		OwnerPublicKey: args.PublicKey,
		Type: args.Type,
		Id: args.Id,
	}
	asset := Asset{Data:Data{Sn:args.String(),Info:balanceInfo}}  // not ID
	//input := Input{}  // TRANSFER
	//Inputs := []Input{input}
	recipients := []interface{}{Recipient{args.PublicKey,1}.ToList()}
	privateKeys := []string{ADMIN_PRIVATE_KEY}
	billInfo := BillInfo{
		//SignerNickName string       `json:"signer_nick_name"`  // 发起人
		//SignerPublicKey string      `json:"signer_public_key"`
		//RecipientNickName string    `json:"recipient_nick_name"`  // 收款人
		//RecipientPublicKey string   `json:"recipient_public_key"`
		//Reason string               `json:"reason"`  // 支付原因
		//Cost string                 `json:"cost"`  // 支付金额
		//Time string                 `json:"time"`  // 支付时间
		SignerNickName:     ADMIN_NICK_NAME,
		SignerPublicKey:    ADMIN_PUBLIC_KEY,
		RecipientNickName:  args.NiceName,
		RecipientPublicKey: args.PublicKey,
		Reason:"新用户余额初始化",
		Cost:"1",
		Time:time.Now().String(),
	}
	metadata := Data{Sn:args.Sn.String(), Info:billInfo}

	transferPrepare := TransferPrepare{
		Operation:operation,
		Asset:asset,
		Recipients:recipients,
		PrivateKeys:privateKeys,
		Metadata:metadata,
	}

	fmt.Println("post :", transferPrepare)
	// todo 提交给postServer,路由中添加接收处理，响应前端
	return nil
}

// 合并资产
func MergeBalanceAsset()  {
	// todo 余额分散时合并
}