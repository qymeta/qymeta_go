package qymeta

import (
	"fmt"
	"testing"
)

func TestInitLedger(t *testing.T) {
	res := InitLedger("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "qychan_test")

	fmt.Println("res:", res)

}

func TestCreateAccountAddress(t *testing.T) {
	res := CreateAccountAddress("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "zhangsi", "17756183275", "42070419110257855")
	fmt.Println("res:", res)
}

func TestCreateNft(t *testing.T) {
	res := CreateNft("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "柯镇恶", "神雕", "神雕", "123", "http://www.qingyuan.com/logon.jpg", "0x5578b9413f381640dad4d296e968dd64c2bafa5c")
	fmt.Println(res)
}

func TestSave(t *testing.T) {
	res := Save("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "http://openqkl.newmin.cn/resouce/IsNjP1wBB9MYH3vXuG2spHoZKHkXOW", "1213522")
	fmt.Println(res)
}

func TestGetOwnerof(t *testing.T) {
	res := GetOwnerof("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "5c6e7cb451528317dd20752d2af8764d8d67e97ca00dbff4372bd592e05f5308")
	fmt.Println(res)
}

func TestGetTokenUrl(t *testing.T) {
	res := GetTokenUrl("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "5c6e7cb451528317dd20752d2af8764d8d67e97ca00dbff4372bd592e05f5308")
	fmt.Println(res)
}

func TestGetBalanceof(t *testing.T) {
	res := GetBalanceof("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "0x5578b9413f381640dad4d296e968dd64c2bafa5c")
	fmt.Println(res)
}

func TestQuery(t *testing.T) {
	res := Query("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "5668abb067aaaa56129b2f9f524011c3af74953c9f3fc4595e3ab95c39f1d034")
	fmt.Println(res)
}

func TestTransFrom(t *testing.T) {
	res := TransFrom("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "0x5578b9413f381640dad4d296e968dd64c2bafa5c", "0x1c0a67251993cf4f4202ba5c5c5256c27c94986", "5c6e7cb451528317dd20752d2af8764d8d67e97ca00dbff4372bd592e05f5308")
	fmt.Println(res)
}

func TestGetTokenBasic(t *testing.T) {
	res := GetTokenBasic("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk")
	fmt.Println(res)
}

func TestGetName(t *testing.T) {
	res := GetName("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk")
	fmt.Println(res)
}

func TestDestruction(t *testing.T) {
	res := Destruction("GL0lFgy0AP1I", "ON4oLiIXNFIe54WCdCZk", "0x1c0a67251993cf4f4202ba5c5c5256c27c94986", "5c6e7cb451528317dd20752d2af8764d8d67e97ca00dbff4372bd592e05f5308")
	fmt.Println(res)
}
