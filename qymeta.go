package qymeta

import (
	"fmt"
	"github.com/qymeta/qymeta_go/util"
	"strconv"
	"time"
)

// 创建合约
func InitLedger(appid string, secrect string, name string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["name"] = strconv.QuoteToASCII(name)
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("initLedger", parameter)
}

// 创建用户账户
func CreateAccountAddress(appid string, secrect string, name string, mobile string, idc string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["name"] = strconv.QuoteToASCII(name)
	parameter["mobile"] = mobile
	parameter["idc"] = idc
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("createAccountAddress", parameter)
}

// 创建一个nft
func CreateNft(appid string, secrect string, author string, title string, series_name string, series_id string, url string, address string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["author"] = strconv.QuoteToASCII(author)
	parameter["title"] = strconv.QuoteToASCII(title)
	parameter["series_name"] = strconv.QuoteToASCII(series_name)
	parameter["series_id"] = series_id
	parameter["url"] = url
	parameter["address"] = address
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("createNft", parameter)
}

// 授权上链
func Save(appid string, secrect string, source_url string, order_sn string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["source_url"] = source_url
	parameter["order_sn"] = strconv.QuoteToASCII(order_sn)
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	fmt.Println(parameter)
	return util.PostJsonRequest("save", parameter)
}

// 获取nft实际持有人
func GetOwnerof(appid string, secrect string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("getOwnerof", parameter)
}

// 获取nft的url
func GetTokenUrl(appid string, secrect string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("getTokenUrl", parameter)
}

// 查询某个账户拥有的nft个数
func GetBalanceof(appid string, secrect string, address string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["address"] = address
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("getBalanceof", parameter)
}

// 查询哈希存证，获取上链的存证信息
func Query(appid string, secrect string, hash string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["hash"] = hash
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("query", parameter)
}

// 转移nft到另一个账户
func TransFrom(appid string, secrect string, address_from string, address_to string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["address_from"] = address_from
	parameter["address_to"] = address_to
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("transFrom", parameter)
}

// 获取某种类型的nft的基本信息
func GetTokenBasic(appid string, secrect string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("getTokenBasic", parameter)
}

// 获取某种类型的nft的名称
func GetName(appid string, secrect string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("getName", parameter)
}

// 销毁nft
func Destruction(appid string, secrect string, address_from string, token_id string) interface{} {
	parameter := make(map[string]string)
	parameter["appid"] = appid
	parameter["secret"] = secrect
	parameter["address_from"] = address_from
	parameter["token_id"] = token_id
	parameter["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	parameter["sign"] = util.GenerateApiSign(parameter)
	return util.PostJsonRequest("destruction", parameter)
}
