# goeureka
提供Go微服务客户端注册到Eureka中心。

## 功能特性

本项目封装了Eureka rest操作，实现了以下功能：

- register 
- sendheartbeat 
- deregister 
- instances 

调用前三个特性register 、heartbeat 和deregister基本上可以实现微服务注册到Eureka中心，其中appid实例是客户端的名称。

## 使用方法

工具包的导入方法：

```go
import "github.com/SimonWang00/goeureka"
```

如果您默认使用本地启动的Eureka Server，注册代码如下：

```go
goeureka.RegisterClient("http://127.0.0.1:8761","my-goserver", "8080", "9043")
```

register方法是通过心跳与Eureka服务端保持通信，当Eureka客户端和服务端注册成功后，则每30秒钟发送一次心跳。当您的微服务实例通过Sigterm或OS中断信号退出时，则本客户端会在关闭之前注销Eureka，以确保服务实例不会发生冲突。

## 接口函数

RegisterLocal和RegisterClient方法自动封装了注册、发送心跳和取消注册的功能，直接导入到客户端完成调用即可。

- RegisterLocal
- RegisterClient
- GetServiceInstances
- GetInfoWithappName
- GetServices
- Sendheartbeat

## 使用示例

**在http中使用：**

```go
import (
	"fmt"
	"github.com/SimonWang00/goeureka"
	"net/http"
)

func main()  {
	goeureka.RegisterClient("http://127.0.0.1:8761","myapp", "8080", "43")
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		resp := "hello goeureka!"
		_, _ = responseWriter.Write([]byte(resp))
	})
	// start server
	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		fmt.Println(err)
	}
}
```

**在gin框架中使用：**

```go
import (
	"github.com/SimonWang00/goeureka"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("hello", func(c *gin.Context) {
		c.String(200, "hello goeureka")
	})
	goeureka.RegisterClient("http://127.0.0.1:8761","myapp", "8080", "43")
	r.Run("127.0.0.1:8000")
}
```

**在beego中使用：**

待补充



## Test

 java端测试代码

[eurekaconsumer](!https://github.com/SimonWang00/eurekaconsumer.git)