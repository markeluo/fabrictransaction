package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
type IdType struct {
	idtype   string   `json:"idtype"`   //idtype
	idstring []string `json:"idstring"` //idstring
}

//用户
type User struct {
	ID                 string 					`json:"id"`                 //用户ID
	Name               string 					`json:"name"`               //用户名字
	IdentificationType int    					`json:"identificationType"` // 证件类型
	Identification     string 					`json:"identification"`     //证件号码
	Sex                int    					`json:"sex"`                //性别
	Birthday           string 					`json:"birthday"`           //生日
	BankCard           string 					`json:"bankcard"`           //银行卡号
	PhoneNumber        string 					`json:"phonoumber"`         //手机号
	Token              string 					`json:"token"`              //密钥
	ProductMap     map[string]Product     		`json:"productmap"`     //产品
	TransactionMap map[string]Transaction 		`json:"transactionmap"` //交易
}

//产品
type Product struct {
	ProductID      	string `json:"productid"`      	//产品id
	ProductName    	string `json:"productname"`    	//产品名称
	ProductType    	int    `json:"producttype"`    	//产品类型
	ProDuctDesc		string `json:"proDuctDesc"` 	//产品介绍
	SPID 			string `json:"sPID"` 			//产品所属服务商id
	Portion        	int    `json:"portion"`        	//产品份额
	Price          	int    `json:"price"`          	//单价

}

//服务商
type ServiceProvider struct {
	ServiceProviderID   string `json:"serviceProviderID"`   //服务商id
	ServiceProviderName string `json:"serviceProviderName"` //服务商名称
	ServiceProviderType int    `json:"serviceProviderType"` //服务商类型

}

//交易内容
type Transaction struct {
	TransId       		string 	`json:"id"`         		//交易id
	TransType     		int    	`json:"transtype"`  		//交易类型 0，在线交易 1，线下交易
	TransStatus			int		`json:"transStatus"`		//交易状态
	FromType      		int    	`json:"fromtype"`   		//发送方角色
	FromID        		string 	`json:"fromid"`     		//发送方 ID
	ToType        		int    	`json:"totype"`     		//接收方角色
	ToID          		string 	`json:"toid"`       		//接收方 ID
	ConfirmType			int    	`json:"confirmType"`		//确认方角色
    ConfirmID			string 	`json:"confirmID"`  		//确认方 ID
	TransDate       	string 	`json:"transDate"`  		//交易时间
	ProductID     		string 	`json:"productid"`  		//交易产品id
	Account       		int    	`json:"account"`    		//交易 份额
	Price         		int    	`json:"price"`      		//交易价格
	TransConfirmDate	string	`json:"transConfirmDate"`	//交易确认时间
	PayConfirmID		string	`json:"payConfirmID"`		//交易付款确认方ID
	PayConfirmDate		string	`json:"payConfirmDate"`		//付款确认时间
	TransSuccedDate		string	`json:"transSuccedDate"`	//交易完成时间
	ParentOrderNo 		string 	`json:parentOrderNo"` 		//父订单号
}

var err error

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}
	return shim.Success(nil)
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function != "invoke" {
		return shim.Error("Unknown function call")
	}
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting at least 2")
	}
	switch {

	case args[0] == "CreateUser":
		return t.CreateUser(stub, args)
	case args[0] == "UserLogin":
		return t.UserLogin(stub, args)	
	case args[0] == "CreateServiceProvider":
		return t.CreateServiceProvider(stub, args)
	case args[0] == "CreateProduct":
		return t.CreateProduct(stub, args)
	case args[0] == "GetTransactionByID":
		return t.GetTransactionByID(stub, args)
	case args[0] == "GetProduct":
		return t.GetProduct(stub, args)
	case args[0] == "getServiceProvider":
		return t.getServiceProvider(stub, args)
	case args[0] == "GetUser":
		return t.GetUser(stub, args)
	case args[0] == "WriteUser":
		return t.WriteUser(stub, args)
	case args[0] == "writeServiceProvider":
		return t.writeServiceProvider(stub, args)
	case args[0] == "WriteProduct":
		return t.WriteProduct(stub, args)
	case args[0] == "Transation":
		return t.Transation(stub, args)
	case args[0] == "GetUserAsset":
		return t.GetUserAsset(stub, args)
	case args[0] == "query":
		return t.query(stub, args)
	default:
		fmt.Printf("function is not exist\n")
	}

	return shim.Error("Unknown action,")
}

//用户登录验证
func (t *SimpleChaincode) UserLogin(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 UserLogin")

	var userid string   // 用户ID
	var username string //用户名称
	var token string    //用户密钥
	var user User

	if len(args) != 4 {
		return shim.Error("UserLogin :Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode

	userid = args[1]
	username = args[2]
	token = args[3]

	userinfo, err := stub.GetState(userid)
	if err != nil {
		return shim.Error(err.Error())
	}
	if userinfo == nil {
		return shim.Success(nil)
	} else {
		err = json.Unmarshal(userinfo, &user)
		if err != nil {
			return shim.Error(err.Error())
		} else if (username == user.Name) && (token == user.Token) {
			return shim.Success("success")
		}

	}
	return shim.Success(nil)
}

//用户查询当下的所有产品
func (t *SimpleChaincode) getUserProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var userid string //用户ID
	var user User
	var ProductMap map[string]Product

	if len(args) != 2 {
		return shim.Error("getUserProduct：Incorrect number of arguments. Expecting 2")
	}

	userid = args[1]
	UserInfo, err := stub.GetState(userid)
	if err != nil {
		return shim.Error(err.Error())
	}
	if UserInfo != nil {
		//将byte的结果转换成struct
		err = json.Unmarshal(UserInfo, &user)
		if err != nil {
			return shim.Error(err.Error())
		}
		ProductMap = user.ProductMap

		for key, value := range ProductMap {
			fmt.Printf("%s-%d\n", key, value)

			fmt.Printf("产品：", key, "产品内容：", value)

		}

		fmt.Printf(" CeateBank success \n")
		return shim.Success(ProductMap)

	}
	return shim.Success(nil)

}

//创建用户
func (t *SimpleChaincode) CreateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 CreateUser")
	//
	var ID string              //用户ID
	var Name string            //用户名字
	var IdentificationType int // 证件类型
	var Identification string  //证件号码
	var Sex int                //性别
	var Birthday string        //生日
	var BankCard string        //银行卡号
	var PhoneNumber string     //手机号
	var token string           //密钥

	var user User
	var idtype IdType

	if len(args) != 10 {
		return shim.Error("CreateUser：Incorrect number of arguments. Expecting 10")
	}

	ID = args[1]
	Name = args[2]
	IdentificationType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Identification = args[4]
	Sex, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Birthday = args[6]
	BankCard = args[7]
	PhoneNumber = args[8]
	token = args[9]

	user.ID = ID
	user.BankCard = BankCard
	user.Birthday = Birthday
	user.Identification = Identification
	user.IdentificationType = IdentificationType
	user.Name = Name
	user.PhoneNumber = PhoneNumber
	user.Sex = Sex
	user.Token = token
	map_pro := make(map[string]Product)
	map_transaction := make(map[string]Transaction)
	user.ProductMap = map_pro
	user.TransactionMap = map_transaction

	jsons_users, err := json.Marshal(user) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}

	//	IdtypeInfo, err := stub.GetState(1)
	//	if err != nil {
	//		return shim.Error(err.Error())
	//	}
	//	if IdtypeInfo == nil {
	//		idtype.idtype = 1
	//		idtype.idstring = append(user.ID)

	//	} else {
	//		err = json.Unmarshal(IdtypeInfo, &idtype)
	//		if err != nil {
	//			return shim.Error(err.Error())
	//		} else {
	//			idtype.idtype = 1
	//			idtype.idstring = append(user.ID)
	//		}
	//	}

	//jsons_idtype, err := json.Marshal(idtype) //转换成JSON返回的是byte[]
	//将byte的结果转换成struct
	//	if err != nil {
	//		return shim.Error(err.Error())
	//	}
	//	err = stub.PutState(1, jsons_idtype)
	//	if err != nil {
	//		return shim.Error(err.Error())
	//	}

	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_users)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(" CreateUser success \n", jsons_users)
	return shim.Success("success")
}

//用户查询某服务商的产品
func (t *SimpleChaincode) getUserProductogOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var userid string //用户ID
	var org_id string //用户ID
	var user User
	var ProductMap map[string]Product

	if len(args) != 3 {
		return shim.Error("getUserProductogOrg number of arguments. Expecting 3")
	}

	userid = args[1]
	org_id = args[2]
	UserInfo, err := stub.GetState(userid)
	if err != nil {
		return shim.Error(err.Error())
	}
	if UserInfo != nil {
		//将byte的结果转换成struct
		err = json.Unmarshal(UserInfo, &user)
		if err != nil {
			return shim.Error(err.Error())
		}
		ProductMap = user.ProductMap
		map_product_org := make(map[string]Product)
		for key, value := range ProductMap {
			fmt.Printf("%s-%d\n", key, value)

			fmt.Printf("产品：", key, "产品内容：", value)
			if value.OrganizationID == org_id {
				map_product_org[key] = value
			}

		}

		fmt.Printf(" CeateBank success \n")
		return shim.Success(map_product_org)

	}
	return shim.Success(nil)

}

//用户查询某机构下的交易情况
func (t *SimpleChaincode) getUserProductogOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var userid string //用户ID
	var org_id string //用户ID
	var user User
	var trans map[string]Transaction

	if len(args) != 3 {
		return shim.Error("getUserProductogOrg number of arguments. Expecting 3")
	}

	userid = args[1]
	org_id = args[2]
	UserInfo, err := stub.GetState(userid)
	if err != nil {
		return shim.Error(err.Error())
	}
	if UserInfo != nil {
		//将byte的结果转换成struct
		err = json.Unmarshal(UserInfo, &user)
		if err != nil {
			return shim.Error(err.Error())
		}
		trans = user.TransactionMap
		map_trans_org := make(map[string]Transaction)
		for key, value := range trans {
			fmt.Printf("%s-%d\n", key, value)

			fmt.Printf("ID：", key, "交易内容：", value)
			if value.OrganizationID == org_id {
				map_trans_org[key] = value
			}

		}

		fmt.Printf(" CeateBank success \n")
		return shim.Success(map_trans_org)

	}
	return shim.Success(nil)

}

//创建服务商
func (t *SimpleChaincode) CreateServiceProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 CreateServiceProvider")

	var ServiceProviderID string      //机构id
	var ServiceProviderName string    //机构名称
	var ServiceProviderType int       //机构类型
	var serviceprovider ServiceProvider //机构

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Initialize the chaincode
	ServiceProviderID = args[1]
	ServiceProviderName = args[2]

	ServiceProviderType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	serviceprovider.ServiceProviderID = ServiceProviderID
	serviceprovider.ServiceProviderName = ServiceProviderName
	serviceprovider.ServiceProviderType = ServiceProviderType

	jsons_ServiceProvider, err := json.Marshal(serviceprovider) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_ServiceProvider)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("CreateServiceProvider \n", jsons_ServiceProvider)

	return shim.Success(nil)
}

//创建产品
func (t *SimpleChaincode) CreateProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 CreateProduct")

	var ProductID string      //产品id
	var ProductName string    //产品名称
	var ProductType int       //产品类型
	var ProDuctDesc string 	  //产品介绍
	var SPID string 		  //产品所属服务商id
	var Portion int           //产品份额
	var Price int             //价格
	var product Product

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	ProductID = args[1]
	ProductName = args[2]
	ProductType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ProDuctDesc = args[4]
	SPID = args[5]
	Portion, err = strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	Price, err = strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	product.ProductID = ProductID
	product.ProductName = ProductName
	product.ProductType = ProductType
	product.ProDuctDesc = ProDuctDesc
	product.SPID = SPID
	product.Portion = Portion
	product.Price = Price

	jsons_product, err := json.Marshal(product) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_product)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf(" CreateProduct success \n", jsons_product)
	return shim.Success(nil)
}

//GetTransactionByID 获取某笔交易
func (t *SimpleChaincode) GetTransactionByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 GetTransactionByID")

	var Transactin_ID string //交易ID
	var transaction Transaction
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	Transactin_ID = args[1]

	TransactionInfo, err := stub.GetState(Transactin_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(TransactionInfo, &transaction)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  TransactionInfo  = %d  \n", TransactionInfo)

	return shim.Success(nil)
}

//GetProduct 获取产品信息
func (t *SimpleChaincode) GetProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 GetProduct")

	var Product_ID string //产品ID
	var product Product
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	Product_ID = args[1]

	ProductInfo, err := stub.GetState(Product_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(ProductInfo, &product)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  ProductInfo  = %d  \n", ProductInfo)
	return shim.Success(nil)
}

//getServiceProvider 获取服务商信息
func (t *SimpleChaincode) getServiceProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getServiceProvider")

	var ServiceProvider_ID string // 商业银行ID
	var serviceprovider ServiceProvider

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	ServiceProvider_ID = args[1]

	ServiceProviderInfo, err := stub.GetState(ServiceProvider_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(ServiceProviderInfo, &serviceprovider)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  OrganizationInfo  = %d  \n", ServiceProviderInfo)

	return shim.Success(nil)
}

//GetUser 获取用户信息
func (t *SimpleChaincode) GetUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 GetUser")

	var User_ID string // 用户ID
	var user User

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	User_ID = args[1]
	userinfo, err := stub.GetState(User_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(userinfo, &user)

	fmt.Printf("  userinfo  = %d  \n", userinfo)

	return shim.Success(nil)
}

//writeUser  修改用户信息
func (t *SimpleChaincode) WriteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var ID string              //用户ID
	var Name string            //用户名字
	var IdentificationType int // 证件类型
	var Identification string  //证件号码
	var Sex int                //性别
	var Birthday string        //生日
	var BankCard string        //银行卡号
	var PhoneNumber string     //手机号
	var user User

	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	ID = args[1]
	Name = args[2]
	IdentificationType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Identification = args[4]
	Sex, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Birthday = args[6]
	BankCard = args[7]
	PhoneNumber = args[8]

	user.ID = ID
	user.BankCard = BankCard
	user.Birthday = Birthday
	user.Identification = Identification
	user.IdentificationType = IdentificationType
	user.Name = Name
	user.PhoneNumber = PhoneNumber
	user.Sex = Sex

	jsons_users, err := json.Marshal(user) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}

	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_users)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(" CeateBank success \n")
	return shim.Success(nil)
}

//writeServiceProvider         服务商
func (t *SimpleChaincode) writeServiceProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 writeServiceProvider")

	var ServiceProviderID string      //服务商id
	var ServiceProviderName string    //服务商名称
	var ServiceProviderType int       //服务商类型
	var serviceprovider ServiceProvider //服务商

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Initialize the chaincode
	ServiceProviderID = args[1]
	ServiceProviderName = args[2]

	ServiceProviderType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	serviceprovider.OrganizationID = ServiceProviderID
	serviceprovider.OrganizationName = ServiceProviderName
	serviceprovider.OrganizationType = ServiceProviderType

	jsons_serviceprovider, err := json.Marshal(serviceprovider) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_serviceprovider)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("CreateServiceProvider \n")

	return shim.Success(nil)
}

//WriteProduct 修改产品
func (t *SimpleChaincode) WriteProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteProduct")

	var ProductID string     //产品id
	var ProductName string  //产品名称
	var ProductType int     //产品类型
	var ProDuctDesc string 	//产品简介
	var SPID string 		//产品所属机构id
	var Portion int         //产品份额
	var Price int           //价格
	var product Product

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	ProductID = args[1]
	ProductName =args[2]
	ProductType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ProDuctDesc =args[4]
	SPID = args[5]
	Portion, err = strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	Price, err = strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	product.ProductID = ProductID
	product.ProductName = ProductName
	product.ProductType = ProductType
	product.ProDuctDesc = ProDuctDesc
	product.SPID = SPID
	product.Portion = Portion
	product.Price = Price

	jsons_product, err := json.Marshal(product) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(ProductID, jsons_product)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf(" CreateProduct success \n")
	return shim.Success(nil)
}

//Transation交易
func (t *SimpleChaincode) Transation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 Transation交易")

	var TransId string          //交易id
	var TransType int           //交易类型
	var FromType int            //发送方角色
	var FromID string           //发送方 ID
	var ToType int              //接收方角色
	var ToID string             //接收方 ID
	var ConfirmType int			//确认方角色
    var ConfirmID string 		//确认方 ID
	var TransDate string        //交易时间
	var ProductID string        //交易产品id
	var TransConfirmDate string //交易确认时间
	var PayConfirmID string 	//交易付款确认方ID
	var PayConfirmDate string 	//付款确认时间
	var TransSuccedDate string  //交易完成时间
	var Account int             //交易 份额
	var price int               //价格
	var ParentOrderNo string    //父订单号

	var transaction Transaction //交易
	var product Product

	var user User
	var ProductMap map[string]Product
	if len(args) != 18 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	TransId = args[1]
	TransType, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	FromType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	FromID = args[4]
	ToType, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ToID = args[6]
	ConfirmType, err = strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ConfirmID = args[8]
	TransDate = args[9]
	ProductID = args[10]
	TransConfirmDate = args[11]
	PayConfirmID, err = strconv.Atoi(args[12])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	PayConfirmDate = args[13]
	TransSuccedDate = args[14]
	Account, err = strconv.Atoi(args[15])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	price, err = strconv.Atoi(args[16])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ParentOrderNo = args[17]

	transaction.TransId = TransId
	transaction.TransType = TransType
	transaction.FromType = FromType
	transaction.FromID = FromID
	transaction.ToType = ToType
	transaction.ToID = ToID
	transaction.ConfirmType =ConfirmType
    transaction.ConfirmID=ConfirmID
	transaction.TransDate = TransDate
	transaction.ProductID = ProductID
	transaction.TransConfirmDate = TransConfirmDate
	transaction.PayConfirmID = PayConfirmID
	transaction.PayConfirmDate = PayConfirmDate
	transaction.TransSuccedDate = TransSuccedDate
	transaction.Account = Account
	transaction.Price = price
	transaction.ParentOrderNo = ParentOrderNo

	jsons_transaction, err := json.Marshal(transaction) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(TransId, jsons_transaction)
	if err != nil {
		return shim.Error(err.Error())
	}

	FromInfo, err := stub.GetState(FromID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(FromInfo, &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	ProductMap = user.ProductMap
	product = ProductMap[ProductID]
	product.Portion = product.Portion + Account
	ProductMap[ProductID] = product

	user.ProductMap = ProductMap

	jsons_User, err := json.Marshal(user) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(FromID, jsons_User)

	fmt.Printf(" Transation success \n")
	return shim.Success(nil)
}

//GetUserAsset  查询用户资产
func (t *SimpleChaincode) GetUserAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var User_ID string //用户ID
	//	var Name string            //用户名字
	//	var IdentificationType int // 证件类型
	//	var Identification string  //证件号码
	//	var Sex int                //性别
	//	var Birthday string        //生日
	//	var BankCard string        //银行卡号
	//	var PhoneNumber string     //手机号
	//	var TransactionIDArray []string

	var user User
	var ProductMap map[string]Product

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	User_ID = args[1]
	UserInfo, err := stub.GetState(User_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(UserInfo, &user)
	if err != nil {
		return shim.Error(err.Error())
	}
	ProductMap = user.ProductMap

	for key, value := range ProductMap {
		fmt.Printf("%s-%d\n", key, value)

		fmt.Printf("产品：", key, "产品内容：", value)

	}

	fmt.Printf(" CeateBank success \n")
	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[1]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// Deletes an entity from state

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[1]

	// Get the state from the ledger
	Avalbytes, erro := stub.GetState(A)
	if erro != nil {
		return shim.Error(erro.Error())
	}
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
