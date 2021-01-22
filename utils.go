package goeureka

//File  : utils.go
//Author: Simon
//Describe: Defines all util for eureka
//Date  : 2020-12-03 11:12:23

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	epoch int64 = 1526285084373
	numWorkerBits = 10
	numSequenceBits = 12
	MaxWorkId = -1 ^ (-1 << numWorkerBits)
	MaxSequence = -1 ^ (-1 << numSequenceBits)
)

type SnowFlake struct {
	lastTimestamp uint64
	sequence      uint32
	workerId      uint32
	lock          sync.Mutex
}

// getLocalIP get loacal host ip address
// if not panic("Unable to get the local IP address")
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to get the local IP address")
}

// getIpFromDocker 从docker中的环境变量获取宿主的IP地址，dockerfile中需要传入IP到环境变量
func getIpFromDocker() string {
	environ := os.Environ()
	for i := range environ {
		env_param := environ[i]
		param := strings.Split(env_param, "=")
		if len(param) ==2 && param[0] == "SystemRoot"{
			fmt.Println(param)
			return param[1]
		}
	}
	return ""
}

// GetLocalIPis getLocalIP for test
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	panic("Unable to get the local IP address")
}

// getUuid use snowflake compute uuid
func getUuid() string {
	sf, err := newSnowFlake(1)
	if err != nil {
		panic(err)
	}

	// Generate a snowflake ID.
	uuid, _:= sf.Generate()
	return uuid
}

// GetUuid getUuid for test
func GetUuid() (string, error){
	sf, err := newSnowFlake(1)
	if err != nil {
		panic(err)
	}

	// Generate a snowflake ID.
	uuid, err:= sf.Generate()
	return uuid, err
}


// trimChar trim first and last byte
// unuse
func trimChar(s string, r byte) string {
	sz := len(s)

	if sz > 0 && s[sz-1] == r {
		s = s[:sz-1]
	}
	sz = len(s)
	if sz > 0 && s[0] == r {
		s = s[1:sz]
	}
	return s
}

// pack program compute uuid
func (sf *SnowFlake) pack() string {
	uuid := (sf.lastTimestamp << (numWorkerBits + numSequenceBits)) | (uint64(sf.workerId) << numSequenceBits) | (uint64(sf.sequence))
	return fmt.Sprintf("%v", uuid)
}

// newSnowFlake returns a new snowflake node that can be used to generate snowflake
func newSnowFlake(workerId uint32) (*SnowFlake, error) {
	if workerId < 0 || workerId > MaxWorkId {
		return nil, errors.New("invalid worker Id")
	}
	return &SnowFlake{workerId: workerId}, nil
}

// Next creates and returns a unique snowflake ID
func (sf *SnowFlake) Generate() (string, error) {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	ts := timestamp()
	if ts == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & MaxSequence
		if sf.sequence == 0 {
			ts = sf.waitNextMilli(ts)
		}
	} else {
		sf.sequence = 0
	}

	if ts < sf.lastTimestamp {
		return "", errors.New("invalid system clock")
	}

	sf.lastTimestamp = ts
	return sf.pack(), nil
}

// waitNextMilli if that microsecond is full
// wait for the next microsecond
func (sf *SnowFlake) waitNextMilli(ts uint64) uint64 {
	for ts == sf.lastTimestamp {
		time.Sleep(100 * time.Microsecond)
		ts = timestamp()
	}
	return ts
}

// timestamp
func timestamp() uint64 {
	return uint64(time.Now().UnixNano()/int64(1000000) - epoch)
}