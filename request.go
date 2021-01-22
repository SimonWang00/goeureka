package goeureka

//File  : request.go
//Author: Simon
//Describe: Defines all request for client request
//Date  : 2020-12-03 11:12:23

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// executeQuery request for eureka server
func executeQuery(requestAction RequestAction) ([]byte, error) {
	request := newHttpRequest(requestAction)

	var DefaultTransport http.RoundTripper = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*10)    //设置建立连接超时
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Second * 10))    //设置发送接受数据超时
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * 10,
	}

	resp, err := DefaultTransport.RoundTrip(request)
	if err != nil {
		return []byte(nil), err
	} else {
		defer resp.Body.Close()
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []byte(nil), err
		}
		return responseBody, nil
	}
}

// isDoHttpRequest return request eureka server is ok
func isDoHttpRequest(requestAction RequestAction) bool {
	request := newHttpRequest(requestAction)
	var DefaultTransport http.RoundTripper = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp, err := DefaultTransport.RoundTrip(request)
	if resp != nil && resp.StatusCode > 299 {
		defer resp.Body.Close()
		log.Printf("HTTP request failed with status (%d)", resp.StatusCode)
		return false
	} else if err != nil {
		log.Printf("HTTP request failed with error (%s)", err.Error())
		return false
	} else {
		return true
		defer resp.Body.Close()
	}
	return false
}


// newHttpRequest build request for eureka
func newHttpRequest(requestAction RequestAction) *http.Request {
	var (
		err error
		request *http.Request
	)
	//log.Printf("DoHttpRequest URL(%v)",requestAction.Url)
	// load body and template for request
	if requestAction.Body != "" {					// add body
		reader := strings.NewReader(requestAction.Body)
		request, err = http.NewRequest(requestAction.Method, requestAction.Url, reader)
	} else if requestAction.Template != "" {		// add template
		reader := strings.NewReader(requestAction.Template)
		request, err = http.NewRequest(requestAction.Method, requestAction.Url, reader)
	} else {
		request, err = http.NewRequest(requestAction.Method, requestAction.Url, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
	// Add headers for request
	request.Header = map[string][]string{
		"Accept":       {requestAction.Accept},
		"Content-Type": {requestAction.ContentType},
	}
	// Add auth username and password
	if username != ""{
		request.SetBasicAuth(username,password)
	}
	return request
}
