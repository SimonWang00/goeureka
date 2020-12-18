package goeureka

import (
	"strings"
)

//File  : config.go
//Author: Simon
//Describe: the config for eureka client
//Date  : 2020/12/7

// define eureka config
var configStr = `{
  "instance": {
	"instanceId" : "${ipAddress}:${appName}:${port}",
    "hostName":"${ipAddress}",
    "app":"${appName}",
    "ipAddr":"${ipAddress}",
    "vipAddress":"${appName}",
	"overriddenstatus": "UNKNOWN",
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
      "management.port" : "${port}"
    }
  }
}`



// newConfig load cfg from configStr
func newConfig(appName,localip, port, securePort string)  string{
	if localip == ""{
		localip = getLocalIP()
	}
	// load config
	cfg := string(configStr)
	cfg = strings.Replace(cfg, "${ipAddress}", localip, -1)
	cfg = strings.Replace(cfg, "${port}", port, -1)
	cfg = strings.Replace(cfg, "${securePort}", securePort, -1)
	cfg = strings.Replace(cfg, "${appName}", appName, -1)
	return cfg
}
