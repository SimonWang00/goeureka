package goeureka

//File  : models.go
//Author: Simon
//Describe: Defines all structs for eureka
//Date  : 2020-12-03

/**
  Defines all structs for eureka

  Request Struct:
		HttpAction

  Rsponse Struct:
		EurekaServiceResponse
		EurekaApplicationsRootResponse
		EurekaApplicationsResponse
		EurekaApplication
		EurekaInstance
		EurekaPort

  Example:
		{
			"application": {
				"name": "MYMICROSERVICE2",
				"instance": [{
					"instanceId": "10.8.30.55:myMicroservice2:338440300596826112",
					"hostName": "10.8.30.55",
					"app": "MYMICROSERVICE2",
					"ipAddr": "10.8.30.55",
					"status": "UP",
					"overriddenStatus": "UNKNOWN",
					"port": {
						"$": 8080,
						"@enabled": "true"
					},
					"securePort": {
						"$": 8443,
						"@enabled": "true"
					},
					"countryId": 1,
					"dataCenterInfo": {
						"@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
						"name": "MyOwn"
					},
					"leaseInfo": {
						"renewalIntervalInSecs": 30,
						"durationInSecs": 90,
						"registrationTimestamp": 1606975540476,
						"lastRenewalTimestamp": 1606975540476,
						"evictionTimestamp": 0,
						"serviceUpTimestamp": 1606975539963
					},
					"metadata": {
						"instanceId": "myMicroservice2:338440300596826112"
					},
					"homePageUrl": "http://10.8.30.55:8080/",
					"statusPageUrl": "http://10.8.30.55:8080/info",
					"healthCheckUrl": "http://10.8.30.55:8080/health",
					"vipAddress": "MYMICROSERVICE2",
					"isCoordinatingDiscoveryServer": "false",
					"lastUpdatedTimestamp": "1606975540476",
					"lastDirtyTimestamp": "1606975539962",
					"actionType": "ADDED"
				}]
			}
		}
*/


// http action for request
type RequestAction struct {
	Method      string `yaml:"method"`
	Url         string `yaml:"url"`
	Body        string `yaml:"body"`
	Template    string `yaml:"template"`
	Accept      string `yaml:"accept"`
	ContentType string `yaml:"contentType"`
	Title       string `yaml:"title"`
	StoreCookie string `yaml:"storeCookie"`
}

// Response for /eureka/apps/{appName}
type ServiceResponse struct {
	Application Application `json:"application"`
}

// Response for /eureka/apps
type ApplicationsRootResponse struct {
	Resp ApplicationsResponse `json:"applications"`
}

type ApplicationsResponse struct {
	Version      string        `json:"versions__delta"`
	AppsHashcode string        `json:"versions__delta"`
	Applications []Application `json:"application"`
}

type Application struct {
	Name     string     `json:"name"`
	Instance []Instance `json:"instance"`
}

// Instance
//eg. [{},]
type Instance struct {
	InstanceId string `json:"instanceId"`
	HostName string `json:"hostName"`
	Port     Port   `json:"port"`
	App 	string	`json:"app"`
	IpAddr	string	`json:"ipAddr"`
	Status	string	`json:"status"`
	SecurePort SecurePort	`json:"securePort"`
	LastDirtyTimestamp string `json:"lastDirtyTimestamp"`
}

// pare port
// eg."port": {
//				"$": 8080,
//				"@enabled": "true"
//			}
type Port struct {
	Port int `json:"$"`
}

// pare SecurePort
// eg."securePort": {
//				"$": 8443,
//				"@enabled": "true"
//			}
type SecurePort struct {
	SecurePort int `json:"$"`
}