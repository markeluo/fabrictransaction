
# servcie transactions
产品/服务交易智能合约，服务商、用户之间的产品服务交易。

##主要函数

CreateUser #创建用户

CreateServiceProvider #创建服务商

CreateProduct #创建服务／产品

Transaction # 交易 

getTransaction #获取所有交易 

getTransactionByID #获取某笔交易 

getProduct #获取服务／产品信息 

getServiceProvider #获取机构信息 

getUser #获取用户信息 

writeUser #修改用户信息 

writeServiceProvider #修改机构信息 

writeProduct #修改服务/产品信息 

getUserAsset #查询用户资产

##数据结构设计

###用户
```
ID：用户ID
Name: 姓名 
IdentificationType: 证件类型 
Identification: 证件号码
Sex: 性别 
Birthday：生日 
BankCard:银行卡号 
PhonoNumber:手机号 
Key: 秘钥
```
###服务/产品类
```
ID : 服务/产品编号 
ProductName: 服务/产品名称 
ProductType: 服务/产品类型
ProDuctDesc: 服务/产品描述
SPID:产品所属服务商ID 
Portion:产品份额
```
###服务商类
```
ServiceProviderID：服务提供商ID 
ServiceProviderName:服务商名称 
ServiceProviderType:服务商类型
```
###交易内容
```
ID：交易ID 
Trans_type:交易类型 
TransStatus:交易状态 
FromType:交易发起方类型 
FromID：交易发起方ID 
ToType:交易接收方类型 
ToID:交易接收方ID
ConfirmType:交易确认方类型
ConfirmID:交易确认方ID
ProductID：交易产品ID 
Account:份额 
TransDate:交易时间
TransConfirmDate:交易确认时间
PayConfirmID:交易付款确认方ID
PayConfirmDate:付款确认时间
TransSuccedDate:交易完成时间
ParentOrderNo:父订单ID
```
###入链协议类
```
SID：业务系统ID 
ReceiverSID:下游系统ID 
OriginSID：来源系统ID 
RequestSerial:来源请求流水号 
NextRequestSerial:下游请求流水号 
Time:入链时间
```

##接口设计
```
CreateUser #创建用户 
request 参数: 
    args[0]：用户ID 
    args[1]: 姓名 
    args[2]: 证件类型 
    args[3]: 证件号码 
    args[4]: 性别 
    args[5]：生日 
    args[6]:银行卡号 
    args[7]:手机号 
    args[8]: 秘钥 
response 参数: 
{ 
    "ID":"XXX", 
    "Name":"XXX", 
    "Identification_type":"XXX", 
    "Identification":"XXX", 
    "Sex":"XXX", 
    "Birthday":"XXX", 
    "BankCard":"XXX", 
    "PhonoNumber":"XXX", 
    "Key":"XXX"
}
```
```
CreateServiceProvider #创建服务商
request 参数: 
    args[0]:服务商ID 
    args[1]:服务商名称 
    args[2]: 服务商类型 
response 参数: 
{
    "ServiceProviderID":"XXX",
    "ServiceProviderName":"xxx",
    "ServiceProviderType":"xxx"
}
```
```
CreateProduct #创建产品 
request 参数: 
    args[0]:产品ID 
    args[1]:产品名称 
    args[2]:产品类型 
    args[3]:产品所属服务商ID
    args[4]:产品描述
    args[5]:产品份额 
response 参数: 
{
    "ID":"XXX",
    "ProductName":"xxx",
    "ProductType":"xxx",
    "SPID"："xxx"，
    "ProDuctDesc":"xxx",
    "Portion"："xxx" 
}
```
```
Transaction # 交易 
request 参数 
    args[0]：交易ID 
    args[1] :交易类型 0，在线交易 1，线下交易
    args[2]:交易状态 
    args[3]:交易发起方类型 
    args[4]：交易发起方ID 
    args[5]:交易接收方类型 
    args[6]:交易接收方ID
    args[7]：交易确认方类型
    args[8]:交易确认方ID 
    args[9]：交易产品ID
    args[10]:份额 
    args[11]:交易时间
    args[12]:交易确认时间
    args[13]:付款确认方ID
    args[14]:付款确认时间
    args[15]:交易完成时间
response 参数： 
{ 
    "ID":"XXX", 
    "Trans_type":"XXX"," 
    "TransStatus":"XXX", 
    "FromType":"XXX"," 
    "FromID":"XXX", 
    "ToType":"XXX", 
    "ToID":"XXX",
    "ConfirmType":"XXX",
    "ConfirmID":"XXX",
    "ProductID":"XXX", 
    "Account":"XXX", 
    "TransDate":"XXX",
    "TransConfirmDate":"XXX",
    "PayConfirmID":"XXX",
    "PayConfirmDate":"XXX",
    "TransSuccedDate":"XXX"
}
```
```
getTransaction #获取所有交易 
request 参数
    null
response 参数： 
{ 
    "ID":"XXX", 
    "Trans_type":"XXX"," 
    TransStatus":"XXX", 
    "FromType":"XXX"," 
    "FromID":"XXX", 
    "ToType":"XXX", 
    "ToID":"XXX",
    "ConfirmType":"XXX",
    "ConfirmID":"XXX",
    "ProductID":"XXX", 
    "Account":"XXX", 
    "TransDate":"XXX",
    "TransConfirmDate":"XXX",
    "PayConfirmID":"XXX",
    "PayConfirmDate":"XXX",
    "TransSuccedDate":"XXX"
}
```
```
getTransactionByID #获取某笔交易 
request 参数 
    args[0]：交易ID 
response 参数： 
{ 
"ID":"XXX", 
    "Trans_type":"XXX"," 
    TransStatus":"XXX", 
    "FromType":"XXX"," 
    "FromID":"XXX", 
    "ToType":"XXX", 
    "ToID":"XXX",
    "ConfirmType":"XXX",
    "ConfirmID":"XXX",
    "ProductID":"XXX", 
    "Account":"XXX", 
    "TransDate":"XXX",
    "TransConfirmDate":"XXX",
    "PayConfirmID":"XXX",
    "PayConfirmDate":"XXX",
    "TransSuccedDate":"XXX"
}
```
```
getProduct #获取产品信息 
request 参数: 
    args[0] :产品ID
response 参数: 
{
    "ID":"XXX",
    "ProductName":"XXX",
    "ProductType":"XXX",
    ProDuctDesc:"XXX",
    SPID:"XXX, 
    "Portion":"XXX"
}
```
```
getServiceProvider #获取服务商信息 
request 参数: 
    args[0] :机构ID
response 参数: 
{
    "ID":"XXX",
    "ServiceProviderName":"XXX",
    "ServiceProviderType":"XXX"
}
```
```
getUser #获取用户信息
request 参数: 
    args[0]：用户ID
response 参数: 
{ 
    "ID":"XXX", 
    "Name":"XXX", 
    "Identification_type":"XXX", 
    "Identification":"XXX", 
    "Sex":"XXX", 
    "Birthday":"XXX", 
    "BankCard":"XXX", 
    "PhonoNumber":"XXX" 
}
```
```

writeUser #修改用户信息 
request 参数: 
    args[0]：用户ID 
    args[1]: 姓名 
    args[2]: 证件类型 
    args[3]: 证件号码 
    args[4]: 性别 
    args[5]：生日 
    args[6]:银行卡号 
    args[7]:手机号 
    args[8]: 秘钥 
response 参数: 
{ 
    "ID":"XXX", 
    "Name":"XXX", 
    "Identification_type":"XXX", 
    "Identification":"XXX", 
    "Sex":"XXX", 
    "Birthday":"XXX", 
    "BankCard":"XXX", 
    "PhonoNumber":"XXX", 
    "Key":"XXX"
}
```
```
writeServiceProvider #修改服务商信息 
request 参数: 
    args[0]:机构ID 
    args[1]:服务商名称 
    args[2]: 服务商类型
response 参数: 
{
    "ID":"XXX",
    "ServiceProviderName":"XXX",
    "ServiceProviderType":"XXX"
}
```
```
writeProduct #修改产品信息 
request 参数: 
    args[0]:产品ID 
    args[1]:产品名称 
    args[2]: 产品类型 
    args[3]:产品所属机构 
    args[4]:产品份额 
response 参数: 
{
    "ID":"XXX",
    "ProductName":"XXX",
    "ProductType":"XXX" ，
    ProDuctDesc:"XXX",
    SPID:"XXX",
    "Portion"："XXX" 
}
```
```
getUserAsset #查询用户产品
request 参数 
    args[0] 用户ID 
response 参数： 
{
    "ID":"XXX", 
    "Name":"XXX", 
    "Identification_type":"XXX", 
    "Identification":"XXX", 
    "Sex":"XXX", 
    "Birthday":"XXX", 
    "BankCard":"XXX", 
    "PhonoNumber":"XXX", 
    "PID":"XXX", 
    "ProductName":"XXX", 
    "ProductType":"XXX" ， 
    "ProDuctDesc":"XXX",
    "SPID":"XXX",
    "Portion"："XXX"
}
```
