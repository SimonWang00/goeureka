# goeureka
提供Go微服务客户端注册到Eureka中心。

## Eureka服务端API

以下是非java应用程序使用Eureka的REST操作。特殊的参数说明如下：

- **appid**是应用程序的名称；
- **instanceid**是与实例关联的唯一id。

在AWS云中，instanceID是实例的**实例id**，在其他数据中心，是实例的**主机名**。对于JSON/XML，提供的内容类型必须是**application/XML**或者**application/JSON**。

| 描述                                                         | 接口名称                                                     | 输入输出                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Register new application instance                            | POST /eureka/apps/**appID**                                  | Input: JSON/XML payload HTTP Code: 204 on success            |
| De-register application instance                             | DELETE /eureka/apps/**appID**/**instanceID**                 | HTTP Code: 200 on success                                    |
| Send application instance heartbeat                          | PUT /eureka/apps/**appID**/**instanceID**                    | HTTP Code: 200 on success ; 404 if **instanceID** doesn’t exist |
| Query for all instances                                      | GET /eureka/apps                                             | HTTP Code: 200 on success Output: JSON/XML                   |
| Query for all **appID** instances                            | GET /eureka/apps/**appID**                                   | HTTP Code: 200 on success Output: JSON/XML                   |
| Query for a specific **appID**/**instanceID**                | GET /eureka/apps/**appID**/**instanceID**                    | HTTP Code: 200 on success Output: JSON/XML                   |
| Query for a specific **instanceID**                          | GET /eureka/instances/**instanceID**                         | HTTP Code: 200 on success Output: JSON/XML                   |
| Take instance out of service                                 | PUT /eureka/apps/**appID**/**instanceID**/status?value=OUT_OF_SERVICE | HTTP Code: 200 on success ;500 on failure                    |
| Move instance back into service (remove override)            | DELETE /eureka/apps/**appID**/**instanceID**/status?value=UP (The value=UP is optional, it is used as a suggestion for the fallback status due to removal of the override) | HTTP Code: * 200 on success * 500 on failure                 |
| Update metadata                                              | PUT /eureka/apps/**appID**/**instanceID**/metadata?key=value | HTTP Code: * 200 on success * 500 on failure                 |
| Query for all instances under a particular **vip address**   | GET /eureka/vips/**vipAddress**                              | * HTTP Code: 200 on success Output: JSON/XML * 404 if the **vipAddress** does not exist. |
| Query for all instances under a particular **secure vip address** | GET /eureka/svips/**svipAddress**                            | * HTTP Code: 200 on success Output: JSON/XML * 404 if the **svipAddress** does not exist. |

## 功能特性

本项目封装了Eureka rest操作，实现了以下功能：

- register appid
- sendheartbeat appid
- deregister appid
- instances of appid

调用前三个特性register 、heartbeat 和deregister基本上可以实现微服务注册到Eureka中心，其中appid实例是客户端的名称。

## 使用方法

工具包的导入方法：

```go
import "github.com/SimonWang00/goeureka/eureka"
```

如果您默认使用本地启动的Eureka Server，注册代码如下：

```go
eureka.RegisterLocal("my-goserver", "8080", "8443")
```

或者这样：
```go
eureka.RegisterClient("http://127.0.0.1:8761","my-goserver", "8080", "8443")
```

register方法是通过心跳与Eureka服务端保持通信，当Eureka客户端和服务端注册成功后，则每30秒钟发送一次心跳。当您的微服务实例通过Sigterm或OS中断信号退出时，则本客户端会在关闭之前注销Eureka，以确保服务实例不会发生冲突。

## 接口函数

- RegisterLocal-RegisterClient register this app at the Eureka server
- RegisterClient-RegisterClient register this app at the Eureka server
- GetServiceInstances- GetServiceInstances is a function query all instances by appName
- GetServiceInstanceIdWithappName-GetServiceInstanceIdWithappName : in this function, we can get InstanceId by appName
- GetServices-GetServices :get all services for eureka
- Sendheartbeat-Sendheartbeat is a test case for heartbeat

这些register方法会自动处理重试、心跳和取消注册。

## 使用示例

在http中使用：

待补充

在gin框架中使用：

待补充

在beego中使用：

待补充