package goeureka

import "strings"

//File  : config.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/7

// define eureka config
var configStr = `{
  "instance": {
    "hostName":"${ipAddress}",
    "app":"${appName}",
    "ipAddr":"${ipAddress}",
    "vipAddress":"${appName}",
    "status":"UP",
    "port": {
      "$":${port},
      "@enabled": true
    },
    "securePort": {
      "$":${securePort},
      "@enabled": false
    },
    "homePageUrl" : "http://${ipAddress}:${port}/",
    "statusPageUrl": "http://${ipAddress}:${port}/info",
    "healthCheckUrl": "http://${ipAddress}:${port}/health",
    "dataCenterInfo" : {
      "@class":"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
      "name": "MyOwn"
    },
    "metadata": {
      "instanceId" : "${appName}${instanceId}:${port}"
    }
  }
}`

// newConfig load cfg from configStr
func newConfig(appName, port, securePort string)  string{
	// load config
	cfg := string(configStr)
	cfg = strings.Replace(cfg, "${ipAddress}", getLocalIP(), -1)
	cfg = strings.Replace(cfg, "${port}", port, -1)
	cfg = strings.Replace(cfg, "${securePort}", securePort, -1)
	cfg = strings.Replace(cfg, "${instanceId}", instanceId, -1)
	cfg = strings.Replace(cfg, "${appName}", appName, -1)
	return cfg
}