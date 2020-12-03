package gotest

import (
	goeureka "github.com/SimonWang00/goeureka/eureka"
	"log"
	"testing"
)

//File  : goeureka_test.go
//Author: Simon
//Describe: unit test for eureka
//Date  : 2020/12/3



//UuidTest getuuid test
//=== RUN   TestUuid
//    goeureka_test.go:19: uuid: 338439481084350464
//--- PASS: TestUuid (0.00s)
//PASS
func TestUuid(t *testing.T)  {
	uuid, err:= goeureka.GetUuid()
	if err!= nil{
		t.Errorf("getuuid error(%v)", err.Error())
	}
	t.Log("uuid:", uuid)
}


//TestGetLocalHost getLocalHost test
//=== RUN   TestGetLocalHost
//    goeureka_test.go:35: host: 10.8.30.55
//--- PASS: TestGetLocalHost (0.01s)
//PASS
func TestGetLocalHost(t *testing.T)  {
	ip, err := goeureka.GetLocalIP()
	if err != nil{
		t.Errorf("GetLocalHost error(%v)", err.Error())
	}
	t.Log("host:", ip)
}

//TestRegisterLocal test RegisterDefault for client
//=== RUN   TestRegisterDefault
//2020/12/03 14:02:49 Registration Sucess!
//--- PASS: TestRegisterDefault (0.04s)
//PASS
func TestRegisterLocal(t *testing.T) {
	goeureka.RegisterLocal("go-service1", "8080", "8443")
}

//TestRegisterDefault test TestRegisterClient for client
//=== RUN   TestRegisterClient
//2020/12/03 13:59:27 Registration Sucess!
//--- PASS: TestRegisterClient (0.04s)
//PASS
func TestRegisterClient(t *testing.T) {
	goeureka.RegisterClient("http://127.0.0.1:8761/","go-service2", "8080", "8443")
}

//TestGetServiceInstances get appName instance
//=== RUN   TestGetServiceInstances
//2020/12/03 14:29:00 Querying eureka for instances of myMicroservice2 at: http://127.0.0.1:8761/eureka/apps/myMicroservice2
//2020/12/03 14:29:00 Doing queryAction using URL: http://127.0.0.1:8761/eureka/apps/myMicroservice2
//2020/12/03 14:29:00 Got instances response from Eureka:
//{"application":{"name":"MYMICROSERVICE2","instance":[{"instanceId":"10.8.30.55:myMicroservice2:338445840785870848","hostName":"10.8.30.55","app":"MYMICROSERVICE2","ipAddr":"10.8.30.55","status":"UP","overriddenStatus":"UNKNOWN","port":{"$":8080,"@enabled":"true"},"securePort":{"$":8443,"@enabled":"true"},"countryId":1,"dataCenterInfo":{"@class":"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo","name":"MyOwn"},"leaseInfo":{"renewalIntervalInSecs":30,"durationInSecs":90,"registrationTimestamp":1606976861357,"lastRenewalTimestamp":1606976861357,"evictionTimestamp":0,"serviceUpTimestamp":1606976860846},"metadata":{"instanceId":"myMicroservice2:338445840785870848"},"homePageUrl":"http://10.8.30.55:8080/","statusPageUrl":"http://10.8.30.55:8080/info","healthCheckUrl":"http://10.8.30.55:8080/health","vipAddress":"MYMICROSERVICE2","isCoordinatingDiscoveryServer":"false","lastUpdatedTimestamp":"1606976861357","lastDirtyTimestamp":"1606976860846","actionType":"ADDED"}]}}
//2020/12/03 14:30:33 instanct host:10.8.30.55, port:{8080}, SecurePort:{8443}
//--- PASS: TestGetServiceInstances (0.03s)
//PASS
func TestGetServiceInstances(t *testing.T){
	appName := "myMicroservice2"
	eus, _ :=goeureka.GetServiceInstances(appName)
	for _, e := range eus{
		log.Printf("instanct host:%v, port:%v, SecurePort:%v", e.HostName, e.Port, e.SecurePort)
	}
}

//TestGetServices get all services for eureka
//=== RUN   TestGetServices
//2020/12/03 14:34:39 Query for all instanceshttp://127.0.0.1:8761/eureka/apps
//2020/12/03 14:34:39 Doing queryAction using URL: http://127.0.0.1:8761/eureka/apps
//2020/12/03 14:34:39 Got services from Eureka:
//{"applications":{"versions__delta":"1","apps__hashcode":"UP_1_","application":[{"name":"EUREKACONSUMER","instance":[{"instanceId":"DESKTOP-SR3A9FI:eurekaconsumer:9002","hostName":"10.8.30.55","app":"EUREKACONSUMER","ipAddr":"10.8.30.55","status":"UP","overriddenStatus":"UNKNOWN","port":{"$":9002,"@enabled":"true"},"securePort":{"$":443,"@enabled":"false"},"countryId":1,"dataCenterInfo":{"@class":"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo","name":"MyOwn"},"leaseInfo":{"renewalIntervalInSecs":30,"durationInSecs":90,"registrationTimestamp":1606972241104,"lastRenewalTimestamp":1606977221655,"evictionTimestamp":0,"serviceUpTimestamp":1606972240279},"metadata":{"management.port":"9002"},"homePageUrl":"http://10.8.30.55:9002/","statusPageUrl":"http://10.8.30.55:9002/actuator/info","healthCheckUrl":"http://10.8.30.55:9002/actuator/health","vipAddress":"eurekaconsumer","secureVipAddress":"eurekaconsumer","isCoordinatingDiscoveryServer":"false","lastUpdatedTimestamp":"1606972241104","lastDirtyTimestamp":"1606972240190","actionType":"ADDED"}]}]}}
//2020/12/03 14:34:39 app.Name(EUREKACONSUMER),app.Instance([{10.8.30.55 {9002} EUREKACONSUMER 10.8.30.55 UP {443}}])
//--- PASS: TestGetServices (0.03s)
//PASS
func TestGetServices(t *testing.T)  {
	apps, err := goeureka.GetServices()
	if err != nil{
		t.Errorf("test GetServices error(%v)", err)
	}
	for _, app := range apps{
		log.Printf("app.Name(%v),app.Instance(%v)", app.Name, app.Instance)
	}
}

// TestSendheartbeat test send heart beat
func TestSendheartbeat(t *testing.T)  {
	appName := "GOLANG-SERVER"
	goeureka.Sendheartbeat(appName)
}