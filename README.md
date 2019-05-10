# 智能合约

本质，mqtt client

提供server服务；属于应用角色；
同 apiservergo get github.com/goiiot/libmqtt

mqttV5

2019年5月7日20点04分
v5的改动略大，弃之

2019年5月8日10点42分


postServer done

但是smartServer中的这个事务数据准备，需要好几个查询准备数据，需要根据业务进行；

还许需要解决topic的定义问题，

方案1：
服务提供者/服务名

坑就是，post需要转发一下，所以，smart转发的时候，把自己的clientid加个turn好了；

老方案中的/replay会很丑

# 初始化用户余额
当用后首次使用时，检查本地无密钥对，及新用户时，本用例开始执行；
## 基本流
1. 前端获取用户昵称作为clientId，提交到postServer获取密钥对，保存到本地；
2. smartServer准备数据 admin公私钥，用户公钥，用户余额0；
3. 提交到postServer，等待响应成功后，响应前端。
## 可选流
1. 不能获取用户昵称；不能分配密钥对；不能初始化；提示失败，用例结束；

11点13分 看看吃饭前能不能搞定，保存密钥到本地。

16点14分 保存密钥到本地 done

初始化 用户余额信息。
先查chain，用户余额，若无创建之。
这一步需要查询，按用户昵称查，资产，http get chaindb

后边的业务都是 查，改，增的操作，主要是 查和改的组合。
尽量不定义结构体。

需要几个模块 http 查询，数据组装，没什么了，底层逻辑由db做；

# 用例：查看余额
当用户需要查看自己的余额时，本用力开始执行；
## 基本流
1. 前端获取 用户昵称+公钥
2. 在chain中查询该用户的余额
3. 返回给前端，前端展示
## 可选流
1. 若该用户为新用户，为其初始化余额0；
2. 若在chain中该用户的资产分散了，合并后展示；

20190508 22：01
先定义一下 余额信息，设备信息，余额转移元数据，设备使用元数据 的字段
要支撑后面的迹象业务，怎么简单怎么来吧。

2019年5月9日08点57分

其实事务的签名应该交到用户端，不应该让用户发送私钥给智能合约中心。
设备端还好说交给了网管代理，应用端其实应该在用户本地签名（fulfill）后的数据，
交给设备侧，无中心就时说的这个。

但是这样势必需要有一次交互，暂时不做了，都交给网关做。

|blockchain|http|智能合约|go|用户端|
|:---:|:---:|:---:|:---:|:---:|
|-|-|v|<-|info|
|-|-|prepare| -> |v|
|-|-|v| <- |fulfill|
|v|<-|commit|-|-|
|result|->|check|->|show|

 时序图如上
 
 但是这个玩意需要2次交互，所以简化成下图
 反正没人看得出来，[]~(￣▽￣)~*
 
 |blockchain|http|智能合约|go|用户端|
 |:---:|:---:|:---:|:---:|:---:|
 |-|-|v|<-|info|
 |-|-|prepare|-|-|
 |-|-|fulfill|-|-|
 |v|<-|commit|-|-|
 |result|->|check|->|show|
 
 区别就是要不要把用户的私钥发送到smartServer用于签名。
 等有需要的时候，用户端的交给js做好了，暂时放到go端。
 
 09点20分
 ok 开始初始化用户余额，定义余额结构体
 
 查看余额
 用户查，通过 用户公钥 在资产 中查，
 但是这样会跟设备资产重合，所以有【余额，设备】的标识，
 
 公钥+type做资产唯一标识，这样会有个坑，没人可能有多台设备，
 所以还应加设备号，为了结构上统一，余额的加个钱包号0好了；
 
 公钥+type+id
 人+物+号
 
 asset_data{
    info：{
        余额描述|设备描述
    }
    ns:public_key+asset_type+id
 }
 
 用户自己，通过 公钥+"balance"+钱包号 查到 余额资产id
 然后用 资产id 去交易中查未消耗的output，
 若 output 个数 > 1 ，执行合并操作。
 
 为了支撑账单查询，metadata 中需要有
 A,B,cost,
 ns
 
 metadata{
   info{
       账单|租用
    }
    ns:public_key+asset_type+id
 }
 
 balance_info{
    owner:昵称
    public_key:公钥
    type:balance
    id:钱包号
 }
 
 iot_info{
    owner:昵称
    public_key:公钥
    type:device
    id:设备号
    device_name:设备名
    device_info:设备描述
 }
 
 bill_info{
    signer:发起人昵称
    signer_public_key:发起人公钥
    recipient:收款人昵称
    recipient_public_key:收款人公钥
    reason:支付原因
    cost:支付金额
    time:支付时间
 }
 
 rent_info{
    device_id:设备号
    owner:设备拥有者昵称
    owner_public_key:拥有者公钥
    user:设备租用者昵称
    user_public_key:使用者公钥
    status:设备状态
    start_time:开始时间
    end_time:结束时间
 }
 
 # 查询余额
 当前端查看余额时，本用例开始执行；
 ## 基本流
 1. 前端获取 sn：公钥+type+钱包编号 提交到ss（smartServer）；
 2. ss 查询余额，余额资产中按sn查询资产交易&未消耗的outputs，余额结果返回。
 ## 可选流
 1. 新用户，无余额资产，管理员创建该用户的余额资产，初始化为0，重新查询。
 2. 余额分散，使用该用户账号合并余额资产，重新查询。
 
 # 创建余额资产
 由管理员端，根据用户信息，创建余额资产。

 
 # 充值/提现
 当用户充值/提现的时候，本用用例开始执行
 ## 基本流
 1. 前端获取 sn，cost_type, money 提交到ss；
 1. 查unspent output
 2. 充值，管理员向该用户转移代币
 3. 提现，用户向管理员转移代币
 4. 查询余额
 ## 可选流
 1. 执行失败，提示原因，用例结束
 
 # 创建设备
 当设备拥有者要注册设备的时候，本用例开始执行。
 ## 基本流
 1. 前端 填写设备基本信息，sn 提交到 ss；
 2. 管理员创建该设备资产，生成sn；
 3. 前端生成该设备sn的二维码；
 ## 可选流
  1. 提交失败，提示原因，用例结束
  
 # 查看设备
 当设备拥有者要查看自己的设备时，本用例开始执行；
 ## 基本流
 1. 前端获取 sn 提交到ss；
 2. 在资产中 按公钥和设备 查询；
 3. 返回查询结果
 ## 可选流
 1. 无结果，提示无设备，用例结束
 
 # 租用/归还
 当使用者租用/归还设备时本用例开始执行；
 ## 基本流
 1. 前端通过扫一扫获取sn，选择type；
 2. 租用，根据sn查询到设备资产id，通过公钥查询未使用的outputs；
 3. 取交集，获取到未使用的设备output,
 4. 检查metadata，生成事务的metadata；
 5. 归还，计算支付金额；租用设备；
 ## 可选流
 1. 余额不足，先充值，本用例结束；
 2. 设备状态判断不通过，提示原因，用例结束；
 
 # 查看单设备信息
 根据扫码的设备sn，查询设备信息，transfer&output设备状态。
 
 15点25分
 
 头疼，实现 todo
 突然不想用go做server了，不过，这个时候换不合适。就这样吧。
 
坑就是， 通过tranfer只通过 asset_id,这样会查出来 余额的所有交易
output通过 public_key+spent查某人的的可用余额和设备的transfer都会出来；

所以要 某人的代币 使用 未使用的tranferid遍历transfer，然后通过asset_id,过滤。

21点53分 明天测试 创建0代币 

不能 needs to be greater than zero

2019年5月10日10点03分

merge done

mqtt client 果然还是应该独立初始化

10点29分 done

资产查询的资产类型识别待添加。

设备状体变化，使用有限状态机FSM

State:
空闲，使用中，禁用
Event:
租用，归还，禁用
Handler:

状态机少一个出错处理，操作名字定的也不好。
17点31分
done 优雅

时间格式化 todo 年月日时分秒
21点09分 done

// todo 学习mqtt clien api
// SubscribeMultiple(filters map[string]byte, callback MessageHandler) Token
// AddRoute(topic string, callback MessageHandler)

addRouter 好理解，直接为某个topic注册处理函数

SubscribeMultiple 

SubscribeMultiple starts a new subscription for multiple topics.
Provide a MessageHandler to be executed 
when a message is published on one of the topics provided.

这个应该是批量订阅

还是没有发布订阅模式
实现之
 
