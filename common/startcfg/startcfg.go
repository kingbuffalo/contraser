package startcfg

import (
	"encoding/json"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"io/ioutil"
	"os"
	"strconv"
)

type rpcProc_t struct {
	Ip             string
	ProtoNameArr   []string
	RedisConnCount int
}

type startCfgJson_t struct {
	GateWayIpStr   string
	PortList       []rpcProc_t
	UnitTest       int
	LogLevel       int
	GRedisHost     string
	GRedisPort     int
	GRedisDBIdx    int
	GRedisPasswd   string
	PRedisHost     string
	PRedisPort     int
	PRedisDBIdx    int
	PRedisPasswd   string
	ProtoFilePath  string
	MysqlUser      string
	MysqlPasswd    string
	MysqlDbName    string
	HttpIp         string
	HttpPort       int
	SConfigLogPath string
	GConfigLogPATH string
	CfgPath        string
}

var cfg startCfgJson_t
var protoNameMapPort map[string]string
var portMapProtoNameArr map[string]([]string)
var portArr []string
var ip string

func init() {
	if len(os.Args) > 2 {
		ip = os.Args[2]
	}
	cfgFn := os.Args[1]
	cfgByteArr, err := ioutil.ReadFile(cfgFn)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(cfgByteArr, &cfg)
	if err != nil {
		panic(err)
	}

	protoNameMapPort = make(map[string]string)
	portMapProtoNameArr = make(map[string]([]string))
	portArr = make([]string, len(cfg.PortList))
	for i, rpcProcCfg := range cfg.PortList {
		for _, protoName := range rpcProcCfg.ProtoNameArr {
			seelog.Trace(protoName)
			protoNameMapPort[protoName] = rpcProcCfg.Ip
		}
		portMapProtoNameArr[rpcProcCfg.Ip] = rpcProcCfg.ProtoNameArr
		portArr[i] = rpcProcCfg.Ip
	}
}

func GetMysqlIp() string {
	ret := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", cfg.MysqlUser, cfg.MysqlPasswd, cfg.MysqlDbName)
	return ret
}
func GetPortArr() []string {
	return portArr
}

func GetCfgPath() string {
	return cfg.CfgPath
}

func GetPortProtoNameArr(ip string) ([]string, bool) {
	ret, ok := portMapProtoNameArr[ip]
	return ret, ok
}

func GetRpcPort() string {
	return ip
}

func GetHttpIp() string {
	fmt.Printf(cfg.HttpIp)
	return cfg.HttpIp
}

func GetHttpPort() int {
	return cfg.HttpPort
}

func GetProtoFilePath() string {
	return cfg.ProtoFilePath
}

func GetPort(protoName string) (string, bool) {
	seelog.Trace("protoName", protoName)
	ret, ok := protoNameMapPort[protoName]
	return ret, ok
}

func GetBUnitTest() bool {
	return cfg.UnitTest == 1
}
func GetGRedisHostPasswd() (string, string, int) {
	return cfg.GRedisHost + ":" + strconv.Itoa(cfg.GRedisPort), cfg.GRedisPasswd, cfg.GRedisDBIdx
}
func GetPRedisHostPasswd() (string, string, int) {
	return cfg.PRedisHost + ":" + strconv.Itoa(cfg.PRedisPort), cfg.PRedisPasswd, cfg.PRedisDBIdx
}

func GetSerLogPath() string {
	return cfg.SConfigLogPath
}

func GetGateWayIpStr() string {
	return cfg.GateWayIpStr
}
func GetGatwayLogPath() string {
	return cfg.GConfigLogPATH
}
