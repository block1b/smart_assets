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

2019年5月11日13点32分

用户查看账单，就是所有的转账记录s，只要根据sn，查metadata就可以了。
查看设备的，看的是asset+metadata,根据sn查最后一个output的asset+metadata。

14点30分 发布订阅 done

然后是 addrouter 这样就不用再 defaultpublish中手撸
router了。

发布订阅只要还是用在了与postServer交互了。

aaaaaaaaa，原来，订阅的时候可以直接指定handler
那样的话，addrouter就是对subscribe的补充了。
比如订阅统配主题，可以通过add添加handler。

<(￣︶￣)↗[GO!]

打开新大门

那样的话，要SubPub就有点没必要了 todo

对应了 event service 的模式 ！！！物模型相关

0 发布订阅模式是service client的角度定的，全称是
接收pub topic后，返回sub topic，发布订阅是一种模式，

实现方式，sub(subTopic,callback(pub(pubTopic,msg)))
是提供服务的service

主动的请求event，只有pub就好了。

16点59分

哇，原来defaultPubHandler是会被覆盖的；
吐槽，pub sub的主语一直在隐士切换，注意！！！

2019年5月12日14点51分

坑就是，向postServer转发后，请求响应模式会失去响应对象，
一个方法就是，向postServer发送的请求 payload中带响应topic，
smartServer收到响应后，转发到前端；

2019年5月13日16点56分
完了呀，balance和asset分不清

还是得用 assetId，
罢了，就用assetId了，不玩那么多的花了。

自动创建用户余额这个也是个坑 

2019年5月14日08点53分

使用自动生成的assetId，在创建的时候，需要响应回本地，
会有post追踪号丢失的问题，最好的解决方法还是，smart只是转发，
post对象标识使用请求者的，而不是smart

另外，把post放到handler里就可以了。

但是这样的话，query这种自动提交的部分就并不能自动执行了，理论上也不该自动执行。

11点33分

让 smart订阅所有+/post，然后转发给weChat，weChat更新本地数据；

post这次转发真的成了大坑了，两个create需要将assetId返回，
罢了 smart增加一个postHandler，解析响应成功的asset，发给wechat。

create 出来的id是第一次的transferId，会作为assetId，
所以第一次创建提交的结果返回值是个Transaction

14点56分

emmm transfaction 的fullfiled终于知道是干啥了，传入的fulfillment是公钥+签名算法，
用私钥签名了成为fulfillment string

16点08分
emmm果然给新建资源给坑了一把，一会儿的余额也会有，
简单的做法就是，创建后，先用一次。
todo，否则会一堆的异常处理

17点12分
完了完了完了，对资产的理解出现了偏差，所有某资产的事务都
记录在改资产下，代币这个东西就会造成一个问题，从创建开始，
他的amount就已经固定了，这就是一直有的违和感，
为什么设备和余额的处理不一样，就因为把余额和人混淆了，

在chain中任何物都是asset的意思就是，钱就是钱，不分是谁的钱。
按谁的钱这样定义，在系统角度总量是不变的，而我用token来代表一个人的余额。
造成的问题就是，该余额的余额总量不可变，只能变化拥有者，钱这个对象本身
总数是不会变化的，即不能无中生有。不好救啊。

改的话需要创建一个单独的余额，设备倒是可以随便搞，算了，把转账
跟设备同化好了，这是paln B，已经被方圆cut了，启用方案A。

使用大钱包，用户的余额都是一个账单而已，查个人余额的条件改成

publicKey -> outputs 该用户未消费的余额，
blanceAssetId只有一个，我甚至不需要改函数，fufu，
调用的时候，使用同样的assetId即可。

查个人账单需要用sn过滤，todo
应该是过滤不了，会和他的设备混在一起，按type应该能过滤掉，再说了。

现在是需要给admin一个超大的余额就好了。

2019年5月17日09点58分
emmm output 查不出来刚创建的asset output，命名在交易中根据assetId就可以查出来的
而且交易号就是资产号，但是outpu不能查到，日。
还是先做一套初始数据，把流程跑起来好了。

初始化管理员钱包和初始化普通用户余额不同，

先测试一下，。。。不对啊，命名能查到的,emmm竟然是因为公钥打错了。

重写钱包初始化，用户余额初始化，用户余额查询。
默认第一步已经执行。

余额，充值测试完成

19点36分
GetPersonBills 获取个人消费记录重新设计
不应该只是metadate，原始的模糊查询在业务中并没有什么卵用，key太难定义了。

还是应该依赖output的接口做查询，按公钥查某人相关的交易。
然后把每次交易的metadata提取出来。

2019年5月18日15点23分

啊啊啊啊啊啊啊啊啊啊，充值的时候余额计算错了，这里最好不要同时做合并操作。
但是这样的话，引发消费问题，日。。。

但是在转账的时候同时做合并时不对的，主要时每次的余额查询需要做合并，
嗯，merge post就不推送到 client了，但是依然会有问题，每次使用余额，
都需要先检查是否需要合并，然后才能用于交易。

这是一个同步问题，罢了，默认，每次交易后，都让他查余额，然后就触发合并了。

只有支付者做input，接收者分新代币，在查询的时候，让他合并。其他操作的获取只能获取到1个output，计时需要合并。
o(￣┰￣*)ゞ

14点39分2019年5月19日
todo 详细分析租用和归还的操作差异
租用设备的时候，需要检查的设备状态，跟归还时检查的个数不一样，
在查到租户余额，设备信息，设备拥有者后， 
租用操作需要检查的是 设备状态（可用）
归还操作需要检查的是 设备状态（使用中），使用者（归还者），余额。
所以非要写一个函数的话还是需要先判断操作类型。

差异，租用是没有py交易的，so，post空的py交易时不对的。

16点59分
还有一个修改设备状态的操作，虽然前端还没写应该是必要的。
修改需要检查两个动作（开启，关闭）
设备状态（禁用|可用），拥有者（拥有者）

状态通过状态机确认，只是许哟啊加一个用户权限检查;

2019年5月20日19点56分

hahaha第一次查重18%
就剩前端了
本地需要记录的资产,
钱包assetId，拥有设备的assetIds
{
    balance_asset_id:1234
    iot_asset_ids:[]
}

clientid,nicename这两做全局变量

看js 更新list Array() push,就ok了。
文件内容也缓存全局好了。

2019年5月21日09点25分
点击扫一扫，通过二维码获取deviceAssetID，
{"clientId":"WeChat","iot":{"nice_name":"block","private_key":"HwLCf9fbhm6BHTagY5aC1uVKR6sz57h7viuS8DUR9x34","public_key":"3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm","type":"iot","id":"clock0","asset_id":"027e9943c4f18f1873b00b08a02477d036ebe93cc98d8a551389ed252fe95ed1"}}
ClientId:"WeChat", done
PrivateKey:ADMIN_PRIVATE_KEY, 
PublicKey:ADMIN_PUBLIC_KEY, 
AssetId:ADMIN_IOT_ASSET_ID, done

果然再所有业务之前，还是需要先把密钥缓存一下。done

2019年5月21日22点00分
/balanceAssetId
/iotAssetId
前端监听这两个主题，收到后就更新，
本地文件一份，缓存一份，收到balalnce就全量更新，收到iot增量更新。
balance这个理论上都一样，不会改变，iot会由多个，得用[]
那就分两个文件好了。

2019年5月22日09点03分
主页面加载一次，收到flush主题后，保存后加载一次。
增加了资产刷新接口。但是觉得会是个坑。。。

10点47分 文件记录done，该做模板了。

模板直接在map上修改字段好了。
new mapTemp
set key=value
应该没什么问题，写个demo页面看一下。
done，没什么问题，修改数据模板中的字段即可。

还应该由一个初始化前端的脚本，直接手工初始化一个文件好了。
把钱包和设备先放一份，用于演示。done

数据模板，基本上几个数据拿到够，改map字段即可。

2019年5月25日14点18分

代币合并测试，如何展现，就把交易记录展示一下好了。

```json5
    {
        "inputs": [
            {
                "owners_before": [
                    "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"
                ],
                "fulfills": {
                    "transaction_id": "5286da424ad18021513a7449666ac818d951dd0b51d31fb5544d44dfb860213a",
                    "output_index": 0
                },
                "fulfillment": "pGSAIPVZtx4o_rOegIGfPcYb4EaFDHeiZeOMo33DyRobp6QIgUAHNh88PeSqANC_IWyCA-qWN37jjZRYF_HAv2O8I1qFXR-c_yy1aJ3xnls6O53GwpZ0sxmZnJYCTiPNTCCh-6IA"
            }
        ],
        "outputs": [
            {
                "public_keys": [
                    "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"
                ],
                "condition": {
                    "details": {
                        "type": "ed25519-sha-256",
                        "public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"
                    },
                    "uri": "ni:///sha-256;gq2QbZZEcAzcCuCv4xicDCYQ_0RmN09VLSH6t2H0pKQ?fpt=ed25519-sha-256&cost=131072"
                },
                "amount": "98"
            },
            {
                "public_keys": [
                    "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
                ],
                "condition": {
                    "details": {
                        "type": "ed25519-sha-256",
                        "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
                    },
                    "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"
                },
                "amount": "2"
            }
        ],
        "operation": "TRANSFER",
        "metadata": {
            "info": {
                "signer_nick_name": "alice",
                "signer_public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq",
                "recipient_nick_name": "admin",
                "recipient_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm",
                "reason": "payment",
                "cost": "2",
                "time": "2019-05-19 04:40:21"
            },
            "sn": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq.balance.main"
        },
        "asset": {
            "id": "d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"
        },
        "version": "2.0",
        "id": "d3513f4195792742c2434cbc18f835c63524f74bf19663731be4af04e63f265a"
    },
    {
        "inputs": [
            {
                "owners_before": [
                    "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
                ],
                "fulfills": {
                    "transaction_id": "00a37b6b3f0c6bae2d5aff197d6a71273ff82bbb457ab34dcae6ae92fcf92911",
                    "output_index": 0
                },
                "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBcU9PdezrYZxsrnj3-FZSPPH9DEbC4fN677GR75FRXjjq1fP8ff7mmtZOIGJi90e7OrbpU6bD-fVycVGwYyM4E"
            }
        ],
        "outputs": [
            {
                "public_keys": [
                    "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
                ],
                "condition": {
                    "details": {
                        "type": "ed25519-sha-256",
                        "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
                    },
                    "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"
                },
                "amount": "900"
            },
            {
                "public_keys": [
                    "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"
                ],
                "condition": {
                    "details": {
                        "type": "ed25519-sha-256",
                        "public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq"
                    },
                    "uri": "ni:///sha-256;gq2QbZZEcAzcCuCv4xicDCYQ_0RmN09VLSH6t2H0pKQ?fpt=ed25519-sha-256&cost=131072"
                },
                "amount": "50"
            }
        ],
        "operation": "TRANSFER",
        "metadata": {
            "info": {
                "signer_nick_name": "Admin",
                "signer_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm",
                "recipient_nick_name": "alice",
                "recipient_public_key": "HWkENox4DM4Tp3qSfYW8igndpog9GpKFzB7Tp7yXgpBq",
                "reason": "recharge",
                "cost": "50",
                "time": "2019-05-18 03:41:07"
            },
            "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.balance.main"
        },
        "asset": {
            "id": "d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"
        },
        "version": "2.0",
        "id": "c5aabe9fe7230ada547c0288ed08a7b61d8abaf421571b43033a640af3da09d1"
    },
    
```

```
    {
            "inputs": [
                {
                    "owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"],
                    "fulfills": {
                        "transaction_id": "c5aabe9fe7230ada547c0288ed08a7b61d8abaf421571b43033a640af3da09d1",
                        "output_index": 0 },
                    "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBLI5foDk7lr1OMVsgFL5dgDjMVjW-gGkF03UVub1XQzm1e80cPralrVk-4q6O1pnHy862A0RgSEFcm7w8AI6IN"
                },
                {
                    "owners_before": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"],
                    "fulfills": {
                        "transaction_id": "d3513f4195792742c2434cbc18f835c63524f74bf19663731be4af04e63f265a",
                        "output_index": 1 },
                    "fulfillment": "pGSAICNu3miMPbgnonQOrcbGgUknVoZB1S3vKt5lGBSitv90gUBv1O0Bt9Lu5HDJ4-lIanRSkGmnBN0L5ksnRoay7UETNAtjFn0p4ZH3P35H_Zag_vXrorsXC3m-EVBRbpDVEqQO"
                }
            ],
            "outputs": [{
                    "public_keys": ["3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"],
                    "condition": {
                        "details": {
                            "type": "ed25519-sha-256",
                            "public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm"
                        },
                        "uri": "ni:///sha-256;t1GM7Bud9-p2hvzkPDn8AN8FRSc8azG8u0BG_KLswiE?fpt=ed25519-sha-256&cost=131072"
                    },
                    "amount": "902"
                }],
            "operation": "TRANSFER",
            "metadata": {
                "info": {
                    "signer_nick_name": "admin",
                    "signer_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm",
                    "recipient_nick_name": "admin",
                    "recipient_public_key": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm",
                    "reason": "合并余额",
                    "cost": "Ά",
                    "time": "2019-05-19 04:43:42"
                },
                "sn": "3PKKhLTbaFSjpjdEtNYqPTSrgp17Vur25NwVjQNKK7Hm.balance.main"
            },
            "asset": {
                "id": "d6464d9f40ef5656c307a7750a2ac6d2dc76835f7c0fd188ff6d866bd12eb7de"
            },
            "version": "2.0",
            "id": "3da2381fcffad7d65bbf59177a5112cd21f437cdfecb89eae5180a5fb2d1f830"
    }
```

