package tool

import (
	"encoding/json"
	"fmt"
	"github.com/influxdata/influxdb/kit/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"smart_assets/tool/fsm"
	"strconv"
	"time"
)

const bigchaindb_addr  = "192.168.113.6:9984"

// http get 查询 bigchaindb 的查询接口

func GetMetadata(search string)([]byte, error){
	params := url.Values{}
	Url, err := url.Parse("http://"+bigchaindb_addr+"/api/v1/metadata")
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
// 根据sn查询 unspet balance|iot
func OutputQuery(args NickForm) (Output,GetOutputResult, error) {
	sn := args.Sn
	snStr := sn.String()
	fmt.Println("select sn ", snStr)
	assetsByte, err := GetAsset(snStr)
	if err != nil{
		fmt.Println("get asset ",err)
		return Output{},GetOutputResult{}, err
	}
	fmt.Println("assets", string(assetsByte))
	var assets []GetAssetResult
	err = json.Unmarshal(assetsByte, &assets)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return Output{},GetOutputResult{}, err
	}
	asset_id := assets[0].Id  // 无用
	asset_id = args.Sn.AssetId  // 老老实实用唯一标识就是了

	fmt.Println("余额资产id", asset_id)
	publicKey := sn.PublicKey
	outputsByte, err := GetOutputs(publicKey,"true")
	if err != nil{
		fmt.Println("get asset ",err)
		return Output{},GetOutputResult{}, err
	}
	if outputsByte == nil{
		// 无数据
		if args.Sn.Type == "balance"{
			// 新建资产
			err = CreateBalanceAsset(args)
			if err != nil{
				fmt.Println("create ", err)
				return Output{},GetOutputResult{}, err
			}
			// 返回amount = 1 不能初始化0 needs to be greater than zero
			// 再次查询
			time.Sleep(time.Second*1)
			return OutputQuery(args)
		}else {
			return Output{},GetOutputResult{}, err
		}

	}
	var getOutPutResults []GetOutputResult
	err = json.Unmarshal(outputsByte, &getOutPutResults)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return Output{},GetOutputResult{}, err
	}
	var unSpentOutputResults []GetOutputResult
	var amount string  // 坑，amount输出是string，输入的时候是int
	var outPut Output  // 单例返回
	// 包括balance和iot的未消耗outputs，过滤
	for _, getOutPutResult := range getOutPutResults{
		transactionByte, err := GetTransactionById(getOutPutResult.TransactionId)
		if err != nil{
			fmt.Println("get transaction",err)
			return Output{},GetOutputResult{}, err
		}
		var transaction Transaction
		err = json.Unmarshal(transactionByte, &transaction)
		if err != nil{
			fmt.Println("unmarshal ",err)
			return Output{},GetOutputResult{}, err
		}
		// 按asset_id过滤 还好struct自己解决了无Id的问题 CREATE 无id
		if transaction.Asset.Id == asset_id{
			fmt.Println("unspent output")
			unSpentOutputResults = append(unSpentOutputResults, getOutPutResult)
			amount = transaction.Outputs[getOutPutResult.OutputIndex].Amount  // 只在 unSpentOutputResults len=1 有效
			outPut = transaction.Outputs[getOutPutResult.OutputIndex]
		}
	}
	unSpentOutputsLen := len(unSpentOutputResults)
	switch unSpentOutputsLen {
	case 0:
		fmt.Println("无记录")
		// 如果是设备查询，直接返回空
		if args.Type == "iot"{
			return Output{},GetOutputResult{}, err
		}

		// 新建资产
		err = CreateBalanceAsset(args)
		if err != nil{
			fmt.Println("create ", err)
			return Output{},GetOutputResult{}, err
		}
		// 返回amount = 1 不能初始化0 needs to be greater than zero
		// 再次查询
		time.Sleep(time.Second*1)
		return OutputQuery(args)
	case 1:
		fmt.Println("该用户资产可用数量为", amount ,outPut.Amount)
		return outPut,unSpentOutputResults[0], nil
	default:
		if args.Type == "balance"{

			fmt.Println("该用于余额token需要合并")
			// 合并用户资产
			err = MergeBalanceAsset(args, unSpentOutputResults)
			if err != nil{
				fmt.Println("merge ", err)
				return Output{},GetOutputResult{},nil
			}
			// 再次查询
			time.Sleep(time.Second*1)
			return OutputQuery(args)
		}else {
			// 理论上只有balance需要合并
			fmt.Println("bad request")
			return Output{},GetOutputResult{},nil
		}
	}
	//return Output{},nil
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
		Time:time.Now().Format("2006-01-02 03:04:05"),
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
	// 提交给postServer,路由中添加接收处理，响应前端

	err := PostWork(transferPrepare)
	if err != nil {
		return err
	}
	return nil
}

// 合并资产
func MergeBalanceAsset(args NickForm,outPutResults []GetOutputResult) error{
	// 余额分散时合并
	// prepare data
	operation := "TRANSFER"
	Inputs := []Input{}
	amount := 0
	for _, outPutResult := range outPutResults{
		transactionByte, err := GetTransactionById(outPutResult.TransactionId)
		if err != nil{
			fmt.Println("get transaction",err)
			return err
		}
		var transaction Transaction
		err = json.Unmarshal(transactionByte, &transaction)
		if err != nil{
			fmt.Println("unmarshal ",err)
			return err
		}
		// prepare inputs
		output := transaction.Outputs[outPutResult.OutputIndex]
		amountInt,err:=strconv.Atoi(output.Amount)
		if err != nil{
			fmt.Println("Atoi ",err)
			return err
		}
		amount += amountInt
		input := Input{
			//OwnersBefore []string `json:"owners_before"`
			//Fulfills Fulfills     `json:"fulfills"`
			//Fulfillment string    `json:"fulfillment"`
			OwnersBefore:output.PublicKeys,
			Fulfillment:output.Condition.Details,
			Fulfills:Fulfills{
				TransactionId: outPutResult.TransactionId,
				OutputIndex:   outPutResult.OutputIndex,
			},
		}
		Inputs = append(Inputs, input)
	}
	// 合并代币
	recipients := []interface{}{Recipient{args.PublicKey,amount}.ToList()}
	privateKeys := []string{ADMIN_PRIVATE_KEY}
	billInfo := BillInfo{
		//SignerNickName string       `json:"signer_nick_name"`  // 发起人
		//SignerPublicKey string      `json:"signer_public_key"`
		//RecipientNickName string    `json:"recipient_nick_name"`  // 收款人
		//RecipientPublicKey string   `json:"recipient_public_key"`
		//Reason string               `json:"reason"`  // 支付原因
		//Cost string                 `json:"cost"`  // 支付金额
		//Time string                 `json:"time"`  // 支付时间
		SignerNickName:     args.NiceName,
		SignerPublicKey:    args.PublicKey,
		RecipientNickName:  args.NiceName,
		RecipientPublicKey: args.PublicKey,
		Reason:"合并余额",
		Cost:string(amount),
		Time:time.Now().Format("2006-01-02 03:04:05"),
	}
	metadata := Data{Sn:args.Sn.String(), Info:billInfo}

	transferPrepare := TransferPrepare{
		Operation:operation,
		Inputs:Inputs,
		Recipients:recipients,
		PrivateKeys:privateKeys,
		Metadata:metadata,
	}

	fmt.Println("post :", transferPrepare)
	// 提交给postServer,路由中添加接收处理，响应前端

	err := PostWork(transferPrepare)
	if err != nil {
		return err
	}
	return nil
}

// 充值提现
//# 充值/提现
//当用户充值/提现的时候，本用用例开始执行
//## 基本流
//1. 前端获取 sn，cost_type, money 提交到ss；
//1.2 . 查unspent output
//2. 充值，管理员向该用户转移代币
//3. 提现，用户向管理员转移代币
//4. 查询余额
//## 可选流
//1. 执行失败，提示原因，用例结束

func UseBalance(args UseBalanceForm, bUser NickForm) error {
	cost := args.CostMoney
	var a_user NickForm
	var b_user NickForm
	// 充值|提现
	if args.CostType == "recharge"{
		// 充值，由admin sign
		a_user = NickForm{
			NiceName:ADMIN_NICK_NAME,
			Sn:Sn{
				PublicKey:ADMIN_PUBLIC_KEY,
				Type:"balance",
				Id:"main",  // 主钱包
			},
		}
		b_user = args.NickForm
	}
	if args.CostType == "withdrawal"{
		// 提现，由用户 sign
		b_user = NickForm{
			NiceName:ADMIN_NICK_NAME,
			Sn:Sn{
				PublicKey:ADMIN_PUBLIC_KEY,
				Type:"balance",
				Id:"main",  // 主钱包
			},
		}
		a_user = args.NickForm
	}
	if args.CostType == "payment"{
		// 用户a支付给b
		a_user = args.NickForm
		b_user = bUser
	}
	err := BalanceTransfer(a_user, b_user, cost)
	if err != nil{
		return err
	}
	return nil
}

// BalanceTransfer(A,B,CostMoney)
func BalanceTransfer(A, B NickForm, cost CostMoney) error {
	//Operation string `json:"operation"`
	//Inputs []Input   `json:"inputs, omitempty"`
	//Recipients []interface{} `json:"recipients"`
	//PrivateKeys []string `json:"private_keys"`
	//Metadata Data `json:"metadata"`
	operation := "TRANSFER"
	var err error
	var a_unspentOutput Output
	var a_unspentOutputResult GetOutputResult
	var b_unspentOutput Output
	//var b_unspentOutputResult GetOutputResult
	var inputs []Input
	var recipients []interface{}
	var privateKeys []string

	a_unspentOutput,a_unspentOutputResult, err = OutputQuery(A)  // 支付方
	if err != nil{
		fmt.Println("asset query ", err)
		return err
	}
	b_unspentOutput,_, err = OutputQuery(A)  // 收款方
	if err != nil{
		fmt.Println("asset query ", err)
		return err
	}
	input := Input{
		//OwnersBefore []string `json:"owners_before"`
		//Fulfills Fulfills     `json:"fulfills"`
		//Fulfillment string    `json:"fulfillment"`
		OwnersBefore:a_unspentOutput.PublicKeys,
		Fulfillment:a_unspentOutput.Condition.Details,
		Fulfills:Fulfills{
			TransactionId: a_unspentOutputResult.TransactionId,
			OutputIndex:   a_unspentOutputResult.OutputIndex,
		},
	}
	inputs = []Input{input}
	// 重新分配代币
	costMoney, err := strconv.Atoi(cost.Money)
	aMoney, err := strconv.Atoi(a_unspentOutput.Amount)
	bMoney, err := strconv.Atoi(b_unspentOutput.Amount)
	a_amount := aMoney-costMoney
	b_amount := bMoney+costMoney

	// check amount 虽然不必要
	if a_amount < 0 || b_amount <0{
		return errors.New("amount less : not enough")
	}
	recipients = []interface{}{
		Recipient{A.PublicKey,a_amount}.ToList(),
		Recipient{B.PublicKey,b_amount}.ToList(),
	}
	privateKeys = []string{A.PrivateKey}

	billInfo := BillInfo{
		SignerNickName:     A.NiceName,
		SignerPublicKey:    A.PublicKey,
		RecipientNickName:  B.NiceName,
		RecipientPublicKey: B.PublicKey,
		Reason:cost.CostType,
		Cost:cost.Money,
		Time:time.Now().Format("2006-01-02 03:04:05"),
	}
	metadata := Data{
		Sn:A.Sn.String(),
		Info:billInfo,
	}

	transferPrepare := TransferPrepare{
		Operation:operation,
		Inputs:inputs,
		Recipients:recipients,
		PrivateKeys:privateKeys,
		Metadata:metadata,
	}

	fmt.Println("post :", transferPrepare)
	// 提交给postServer,路由中添加接收处理，响应前端
	err = PostWork(transferPrepare)
	if err != nil {
		return err
	}
	return nil
}

// 查看单设备信息，同资产查询

// 查看某人的设备信息，不知道支不支持模糊查询
// 不做了，自己的设备号，保存在本地，退化为多次查看单设备信息

//# 创建设备
//当设备拥有者要注册设备的时候，本用例开始执行。
//## 基本流
//1. 前端 填写设备基本信息，sn 提交到 ss；
//2. 管理员创建该设备资产，生成sn；
//3. 前端生成该设备sn的二维码；
//## 可选流
//1. 提交失败，提示原因，用例结束
func CreateDevice(deviceForm DeviceForm) error {
	// prepare data
	args := deviceForm.NickForm
	operation := "CREATE"
	iotInfo := IotInfo{
		//OwnerNickName string  `json:"owner_nick_name"`
		//OwnerPublicKey string `json:"owner_public_key"`
		//Type string           `json:"type"`  // default：iot
		//Id string             `json:"id"`    // 设备编号
		//DeviceName string     `json:"device_name"`  // 设备名
		//DeviceInfo string     `json:"device_info"`  // 设备描述
		OwnerNickName: args.NiceName,
		OwnerPublicKey: args.PublicKey,
		Type: args.Type,  // iot
		Id: args.Id,  // 设备编号
		DeviceName:deviceForm.DeviceName,
		DeviceInfo:deviceForm.DeviceInfo,
	}
	asset := Asset{Data:Data{Sn:args.String(),Info:iotInfo}}
	//input := Input{}  // TRANSFER
	//Inputs := []Input{input}
	recipients := []interface{}{Recipient{args.PublicKey,1}.ToList()}
	privateKeys := []string{args.PrivateKey}
	rentInfo := RentInfo{
		//DeviceId string `json:"device_id"`  // 设备号
		//OwnerNickName string  `json:"owner_nick_name"`
		//OwnerPublicKey string `json:"owner_public_key"`
		//UserNickName string `json:"user_nick_name"`
		//UserPublicKey string `json:"user_public_key"`
		//Status string `json:"status"`  // 设备状态
		//Ruler string `json:"ruler"`  // 收费规则
		//StartTime string `json:"start_time"` // 开始租用时间
		//CostTime string `json:"cost_time"`  // 租用时间
		DeviceId:args.Id,
		OwnerNickName:args.NiceName,
		OwnerPublicKey:args.PublicKey,
		UserNickName:"null",
		UserPublicKey:"null",
		Status:deviceForm.Status,
		Ruler:deviceForm.Ruler,
		StartTime:time.Now().Format("2006-01-02 03:04:05"),
		CostTime:"0",
	}
	metadata := Data{Sn:args.Sn.String(), Info:rentInfo}

	transferPrepare := TransferPrepare{
		Operation:operation,
		Asset:asset,
		Signers:ADMIN_PUBLIC_KEY,
		Recipients:recipients,
		PrivateKeys:privateKeys,
		Metadata:metadata,
	}

	fmt.Println("post :", transferPrepare)
	// 提交给postServer,路由中添加接收处理，响应前端

	err := PostWork(transferPrepare)
	if err != nil {
		return err
	}
	return nil
}

// 租用/归还设备
//# 租用/归还
//当使用者租用/归还设备时本用例开始执行；
//## 基本流
//1. 前端通过扫一扫获取sn，选择type；
//2. 租用，根据sn查询到设备资产id，通过公钥查询未使用的outputs；
//3. 取交集，获取到未使用的设备output,
//4. 检查metadata，生成事务的metadata；
//5. 归还，计算支付金额；租用设备；
//## 可选流
//1. 余额不足，先充值，本用例结束；
//2. 设备状态判断不通过，提示原因，用例结束；

func UseIot(user NickForm, iotForm DeviceForm) error {
	iot := iotForm.NickForm
	// 检查设备,获取output
	//OutputQuery(args NickForm) (Output,GetOutputResult, error)
	unspentOutput, unspentOutputResult, err := OutputQuery(iot)
	if err!= nil{
		return errors.New("bad device : device un define")
	}
	operation := "TRANSFER"
	var inputs []Input
	var recipients []interface{}
	var privateKeys []string

	input := Input{
		//OwnersBefore []string `json:"owners_before"`
		//Fulfills Fulfills     `json:"fulfills"`
		//Fulfillment string    `json:"fulfillment"`
		OwnersBefore:unspentOutput.PublicKeys,
		Fulfillment:unspentOutput.Condition.Details,
		Fulfills:Fulfills{
			TransactionId: unspentOutputResult.TransactionId,
			OutputIndex:   unspentOutputResult.OutputIndex,
		},
	}
	inputs = []Input{input}
	recipients = []interface{}{
		Recipient{unspentOutput.PublicKeys[0],1}.ToList(),
	}
	privateKeys = []string{iot.PrivateKey}  // 坑啊，罢了，直接用管理员的好了

	// 获取当前设备状态
	transactionByte, err := GetTransactionById(unspentOutputResult.TransactionId)
	if err != nil{
		fmt.Println("get transaction",err)
		return err
	}
	var transaction Transaction
	err = json.Unmarshal(transactionByte, &transaction)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return err
	}

	var oldrentInfo RentInfo
	rentInfoByte, err := json.Marshal(transaction.Metadata.Info)
	if err != nil{
		fmt.Println("marshal ",err)
		return err
	}
	err = json.Unmarshal(rentInfoByte, &oldrentInfo)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return err
	}
	if oldrentInfo.UserPublicKey != user.PublicKey{
		// 非授权用户操作；本来应该直接把设备转给user的，就没有这种问题了 todo
		return errors.New("bad user : permission denied")
	}
	// 判断status转换逻辑
	efan := fsm.Init(fsm.FSMState(oldrentInfo.Status))
	efan.Call(fsm.FSMEvent(iotForm.Status))  // 其实该用event 不是state
	newStatus := string(efan.GetState())
	if newStatus == "Err"{
		return errors.New("bad use : status not allow")
	}
	costTime := "0"  // 初始化为0
	userNiceName := user.NiceName
	userPublicKey := user.PublicKey
	// 归还设备操作计算时间
	if iotForm.Status == "Return"{
		// 计算花费时间
		startTime,err := time.Parse("2006-01-02 03:04:05", oldrentInfo.StartTime)
		if err != nil{
			return err
		}
		costTime = string(time.Now().Unix()-startTime.Unix())  // 时间单位 s
		// 支付
		a_user := user
		b_user := NickForm{
			//NiceName string `json:"nice_name"`
			//PrivateKey string `json:"private_key"`
			//Sn
			NiceName: iot.NiceName,
			PrivateKey: iot.PrivateKey,
			Sn:Sn{
				PublicKey: iot.Sn.PublicKey,
				Type: "balance",
				Id: "main",
			},
		}
		// 暂定消费金额为 costTime*Ruler
		//money := costTime * oldrentInfo.Ruler
		c,err:=strconv.Atoi(costTime)
		r,err:=strconv.Atoi(oldrentInfo.Ruler)
		if err != nil{
			fmt.Println("Atoi ",err)
			return err
		}
		money := string(c*r)
		cost := CostMoney{
			//CostType string `json:"cost_type"`
			//Money string `json:"money"`
			CostType: "payment",
			Money:money,
		}
		err = BalanceTransfer(a_user, b_user, cost)
		if err != nil {
			return err
		}
		// 支付完成
		userNiceName = iotForm.NiceName
		userPublicKey = iotForm.PublicKey
	}  
	
	rentInfo := RentInfo{
		//DeviceId string `json:"device_id"`  // 设备号
		//OwnerNickName string  `json:"owner_nick_name"`
		//OwnerPublicKey string `json:"owner_public_key"`
		//UserNickName string `json:"user_nick_name"`
		//UserPublicKey string `json:"user_public_key"`
		//Status string `json:"status"`  // 设备状态
		//Ruler string `json:"ruler"`  // 收费规则
		//StartTime string `json:"start_time"` // 开始租用时间
		//CostTime string `json:"cost_time"`  // 租用时间
		DeviceId:iot.Id,
		OwnerNickName:iot.NiceName,
		OwnerPublicKey:iot.PublicKey,
		UserNickName:userNiceName,
		UserPublicKey:userPublicKey,
		Status:newStatus,
		Ruler:iotForm.Ruler,
		StartTime:time.Now().Format("2006-01-02 03:04:05"),
		CostTime:costTime,
	}
	metadata := Data{
		Sn:iot.Sn.String(),
		Info:rentInfo,
	}

	transferPrepare := TransferPrepare{
		Operation:operation,
		Inputs:inputs,
		Recipients:recipients,
		PrivateKeys:privateKeys,
		Metadata:metadata,
	}

	fmt.Println("post :", transferPrepare)
	// 提交给postServer,路由中添加接收处理，响应前端
	err = PostWork(transferPrepare)
	if err != nil {
		return err
	}
	
	return nil
}

// 获取个人历史账单 balanceSn
// 同getMetadata() param:Sn return [{metadataresult}]
func GetPersonBills(args NickForm)([]byte, error){
	sn := args.Sn.String()
	return GetMetadata(sn)
}

// 查看设备信息
// OutputQuery -> outputResult(transferId)
// transferById -> assetInfo，metadataInfo
func GetIotInfo(args NickForm) (DeviceForm, error) {
	_, unspentOutputResult, err := OutputQuery(args)
	if err!= nil{
		return DeviceForm{}, errors.New("bad device : device un define")
	}
	transactionByte, err := GetTransactionById(unspentOutputResult.TransactionId)
	if err!= nil{
		return DeviceForm{}, errors.New("GetTransactionById : no result")
	}
	var transaction Transaction
	err = json.Unmarshal(transactionByte, &transaction)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return DeviceForm{}, err
	}

	// 抽取 asset，metadata中的info 部分字段用于展示
	iotAssetInfo := transaction.Asset.Data.Info
	iotmetadataInfo := transaction.Metadata.Info
	iab, err := json.Marshal(iotAssetInfo)
	imb, err := json.Marshal(iotmetadataInfo)
	var iotInfo IotInfo
	var rentInfo RentInfo
	err = json.Unmarshal(iab, &iotInfo)
	err = json.Unmarshal(imb, & rentInfo)
	if err != nil{
		fmt.Println("unmarshal ",err)
		return DeviceForm{}, err
	}

	deviceForm := DeviceForm{
		//DeviceName string     `json:"device_name"`  // 设备名
		//DeviceInfo string     `json:"device_info"`  // 设备描述
		//Status string `json:"status"`  // 设备状态
		//Ruler string `json:"ruler"`  // 收费规则
		DeviceName: iotInfo.DeviceName,
		DeviceInfo: iotInfo.DeviceInfo,
		Status: rentInfo.Status,
		Ruler:rentInfo.Ruler,
	}
	return deviceForm,err
}