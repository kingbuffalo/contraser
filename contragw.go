package main

import (
	"buffalo/king/common"
	"buffalo/king/common/startcfg"
	"buffalo/king/connmgr"
	"buffalo/king/dbOper/playerInfo_db"
	"encoding/gob"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/kingbuffalo/seelog"
	"math/rand"
	"net/http"
	"time"
)

var host = startcfg.GetRpcPort()

var addr = flag.String("addr", host, "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echoEx(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		seelog.Trace("upgrade:", err)
		return
	}

	s := connmgr.NewSWS(c)
	s.Respond()
}

//数据和redis初始化完成后, 初始化一些全局数据
func InitDbData() {
	playerInfo_db.SetMaxPlayId()
	seelog.Trace("init dbdata ...")
}

func init() {
	//加载配置文件
	logger, err := seelog.LoggerFromConfigAsFile(startcfg.GetGatwayLogPath())
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
	seelog.Trace("main server start...")
	rand.Seed(time.Now().Unix())

	InitDbData()
	gob.Register(common.RpcC2SProto{})
	gob.Register(common.RpcS2CRet{})

	flag.Parse()
	http.HandleFunc("/", echoEx)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		panic(err)
	}
}
