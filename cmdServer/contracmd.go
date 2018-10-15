package main

import (
	"buffalo/king/cmdServer/protoImpl"
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	//"buffalo/king/king"
	"encoding/gob"
	//"encoding/json"
	"fmt"
	//"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"github.com/valyala/gorpc"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var protoNameMapWorker map[string](protoImpl.ProtoWorker)

//func curl(rsp http.ResponseWriter, req *http.Request) {
//var test_data common.RpcC2SProto
////test_data.Pid = 103
////test_data.Pid = 200
////	test_data.Pid = 303
////	test_data.Pid = 212
//test_data.Pid = 206
////test_data.Pid = 202
//data := &king.C2S_RefreshMapCitys{
////SignType: proto.Uint32(2),
//MapId: proto.Int32(1),
//}

//ttt, _ := proto.Marshal(data)
//test_data.Data = ttt
////	data := &king.C2S_Sign{
////		SignType: proto.Uint32(1),
////	}

////ttt, _ := proto.Marshal(data)
////test_data.Data = ttt
//dealFunc(&test_data)

//type rsp_struct struct {
//Code int `json:"code"`
//}
//var rsp_data rsp_struct
//rsp_data.Code = 200
//res, _ := json.Marshal(rsp_data)
//if _, err := rsp.Write(res); err != nil {
//return
//}
//return
//}
func init() {
	ipPortStr := startcfg.GetRpcPort()
	protoNameArr, ok := startcfg.GetPortProtoNameArr(ipPortStr)
	if !ok {
		panic(ok)
	}
	protoNameMapWorker = make(map[string](protoImpl.ProtoWorker))
	for _, protoName := range protoNameArr {
		protoNameMapWorker[protoName] = protoImpl.CreateProtoWorker(protoName)
	}
}

type panicHandler_t struct {
	runF func(rpcC2SProto *common.RpcC2SProto) interface{}
	ret  interface{}
}

func (p *panicHandler_t) runFunc(rpcC2SProto *common.RpcC2SProto) {
	defer func() {
		if err := recover(); err != nil {
			_ = seelog.Error(err)
			buf := make([]byte, 1<<11)
			runtime.Stack(buf, true)
			_ = seelog.Error(string(buf))
			p.ret = common.NewErrorCodeRet(113)
		}
	}()
	p.ret = p.runF(rpcC2SProto)
}

func panicHandler_t_runF(rpcC2SProto *common.RpcC2SProto) interface{} {
	protoName := common.GetProtoShortName(rpcC2SProto.Pid)
	shortProtoName := protoName[enumErrCode.PROTO_PREFIX_LEN:]
	worker, ok := protoNameMapWorker[shortProtoName]
	if ok {
		return worker.DoTheJob(rpcC2SProto)
	}
	return common.NewErrorCodeRet(108)
}

func dealFunc(rpcC2SProto *common.RpcC2SProto) interface{} {
	p := panicHandler_t{
		runF: panicHandler_t_runF,
	}
	p.runFunc(rpcC2SProto)
	return p.ret
}
func startHttp() {
	seelog.Trace("http server start...")
	//http.HandleFunc("/test/curl", curl)
	httpIp := startcfg.GetHttpIp()
	httpPort := startcfg.GetHttpPort()
	fmt.Println("httpIp:", httpIp, "httpPort:", httpPort)

	if err := http.ListenAndServe(httpIp+":"+strconv.Itoa(httpPort), nil); err != nil {
		seelog.Trace("http server end...")
		return
	}
	seelog.Trace("http server end...")
}

func initServiceLog() {
	//defer seelog.Flush() not need
	//加载配置文件
	logger, err := seelog.LoggerFromConfigAsFile(startcfg.GetSerLogPath())
	if err != nil {
		panic(err)
	}
	//替换记录器
	err = seelog.ReplaceLogger(logger)
	if err != nil {
		panic(err)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	initServiceLog()
	seelog.Trace("cmd server start...")
	go startHttp()

	gob.Register(common.RpcC2SProto{})
	gob.Register(common.RpcS2CRet{})
	ipPortStr := startcfg.GetRpcPort()
	s := &gorpc.Server{
		Addr: ipPortStr,
		Handler: func(clientAddr string, request interface{}) interface{} {
			rpcC2SProto := request.(common.RpcC2SProto)
			return dealFunc(&rpcC2SProto)
		},
	}
	if err := s.Serve(); err != nil {
		panic(err)
	} else {
		fmt.Println("PlayerInfo create")
	}
}
