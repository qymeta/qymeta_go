# qychan_go_sdk


清元多链qymeta_go是针对国内多家联盟链提供的一套完整的web3、数字藏品等解决方案的go sdk版。开发者无需深入了解各联盟链的相关技术知识，可通过qymeta_go 快速搭建自己的web3、数字藏品应用。

* 清元多链 [官网](http://openqkl.newmin.cn/)

* 清元链 NFT SDK [文档](https://github.com/qymeta/qymeta_go/blob/main/doc/qymeta_go.md)


### Installation
```
go get github.com/qymeta/qymeta_go
```



### qychan_go_sdk Start

```
import (
	"fmt"
	qychan "github.com/qymeta/qymeta_go"
)
func CreateAccountAddress() {
	res := qychan.CreateAccountAddress("...", "...", "...", "...", "...")
	fmt.Println("res:", res)
}

func main() {
	CreateAccountAddress()
}
```

### Change Logs

#### 1.0.2 2022/11/02

* 清元链 NFT 相关接口发布

