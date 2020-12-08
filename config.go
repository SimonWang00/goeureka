package goeureka

import (
	"strings"
	"time"
)

//File  : config.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/7

// define eureka config
var configStr = `
{
	'instanceId': '${ipAddress}:${appName}:${port}',
	'hostName': '${ipAddress}',
	'app': '${appName}',
	'ipAddr': '${ipAddress}',
	'port': {
		'$': ${port},
		'@enabled': 'true'
	},
	'securePort': {
		'$': ${securePort},
		'@enabled': 'false'
	},
	'countryId': 1,
	'dataCenterInfo': {
		'@class': 'com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo',
		'name': 'MyOwn'
	},
	'leaseInfo': {
		'renewalIntervalInSecs': 30,
		'durationInSecs': 90,
		'registrationTimestamp': 0,
		'lastRenewalTimestamp': 0,
		'evictionTimestamp': 0,
		'serviceUpTimestamp': 0
	},
	'metadata': {
		'management.port': '${port}'
	},
	'homePageUrl': 'http://${ipAddress}:${port}/',
	'statusPageUrl': 'http://${ipAddress}:${port}/info',
	'healthCheckUrl': 'http://${ipAddress}:${port}/health',
	'secureHealthCheckUrl': '',
	'vipAddress': '${appName}',
	'secureVipAddress': '${appName}',
	'isCoordinatingDiscoveryServer': 'false',
	'status': 'UP',
	'overriddenstatus': 'UNKNOWN',
	'lastUpdatedTimestamp': '${nano}',
	'lastDirtyTimestamp': '${nano}'
}
`


// newConfig load cfg from configStr
func newConfig(appName, port, securePort string)  string{
	// load config
	cfg := string(configStr)
	nano := string(time.Now().UnixNano()/1e6)
	cfg = strings.Replace(cfg, "${ipAddress}", getLocalIP(), -1)
	cfg = strings.Replace(cfg, "${port}", port, -1)
	cfg = strings.Replace(cfg, "${securePort}", securePort, -1)
	cfg = strings.Replace(cfg, "${instanceId}", instanceId, -1)
	cfg = strings.Replace(cfg, "${appName}", appName, -1)
	cfg = strings.Replace(cfg, "${nano}", nano, -1)

	return cfg
}