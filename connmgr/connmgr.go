package connmgr

/*
import (
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/kingbuffalo/seelog"
	"github.com/valyala/gorpc"
	"sync"
	"time"
)

const (
	eRROR_DATA_LEN = 4
)

type WS struct {
	conn      *websocket.Conn
	player_id int32
	writeMsg  chan []byte
	closeChan chan bool
}

var playerMapConn map[int32](*WS)
var rwmutex = &sync.RWMutex{}

var rpcClient map[string](*gorpc.Client)

func addConn(ws *WS, player_id int32) {
	ws.player_id = player_id
	rwmutex.Lock()
	wsOld, ok := playerMapConn[player_id]
	if ok && ws == wsOld {
		rwmutex.Unlock()
		return
	}
	if ok {
		_ = wsOld.conn.Close()
	}
	playerMapConn[player_id] = ws
	rwmutex.Unlock()
}

func printOnlineCount() {
	rwmutex.RLock()
	seelog.Infof("--------------------online count=%d----------------------", len(playerMapConn))
	rwmutex.RUnlock()
}

func rmConn(ws *WS) {
	player_id := ws.player_id
	if player_id != 0 {
		rwmutex.Lock()
		defer rwmutex.Unlock()
		current_ws, ok := playerMapConn[player_id]
		if ok {
			if current_ws.conn == ws.conn {
				delete(playerMapConn, player_id)
			}
		}
	}
}

func sendMsg(player_id int32, data []byte) {
	rwmutex.RLock()
	ws, ok := playerMapConn[player_id]
	rwmutex.RUnlock()
	if !ok {
		return
	}
	ws.writeMsg <- data
}

func getLenId(message []byte) (int, int) {
	lenh, lenl := message[0], message[1]
	idh, idl := message[2], message[3]

	id := (int(idh)<<8 | int(idl)) & 0xffff
	dataLen := (int(lenh)<<8 | int(lenl)) & 0xffff
	return dataLen, id
}

func NewWs(conn *websocket.Conn) *WS {
	return &WS{
		conn:      conn,
		player_id: 0,
		writeMsg:  make(chan []byte, 10),
		closeChan: make(chan bool),
	}
}

func (ws *WS) readMsg_respond(dataLen, messageLen int, id int, data []byte) ([]byte, int, int32) {
	if dataLen != messageLen-2 {
		return nil, id, 102
	}
	protoName := common.GetProtoName(id)
	rpcProto := &common.RpcC2SProto{
		Len:  dataLen,
		Pid:  id,
		Data: data,
	}
	wMsg, errCode := ws.rpcReadMsg(protoName, rpcProto)
	return wMsg, id, errCode
}
func (ws *WS) readMsg_push(dataLen int, id int, data []byte) ([]byte, int, int32) {
	switch id {
	case enumErrCode.BATTLE_PUSH_PROTO_ID:
		var g2g king.G2G_PushBattle
		if err := proto.Unmarshal(data, &g2g); err != nil {
			seelog.WarnWE("proto.Unmarshal push err=", err)
		} else {
			sendMsg(*g2g.PlayerId, []byte(*g2g.Data))
		}
	case enumErrCode.BATTLE_PUSH_KEEP_ALIVE_ID:
		printOnlineCount()
	}
	return nil, 0, 0

}

func (ws *WS) readMsg() ([]byte, int, int32) {
	wc := ws.conn

	if err := wc.SetReadDeadline(time.Now().Add(time.Second * enumErrCode.WS_READ_TIMEOUT_SECOND)); err != nil {
		return nil, 0, 102
	}

	_, message, err := wc.ReadMessage()
	if err != nil {
		seelog.Trace(err)
		return nil, 0, 101
	}

	dataLen, id := getLenId(message)
	data := message[4:]
	seelog.Trace("message=", message)
	if id >= 100 && dataLen != 0 {
		messageLen := len(message)
		return ws.readMsg_respond(dataLen, messageLen, id, data)
	} else {
		return ws.readMsg_push(dataLen, id, data)
	}
}

func pushToOther(rpcRetVO *common.RpcS2CRet) {
	if rpcRetVO.PushArr == nil {
		return
	}
	for _, v := range rpcRetVO.PushArr {
		sendMsg(v.PlayerId, v.Data)
	}
}

func (ws *WS) rpcReadMsg(protoName string, rpcProto *common.RpcC2SProto) ([]byte, int32) {
	protoStr := protoName[enumErrCode.PROTO_PREFIX_LEN:]
	ipPortStr, bExist := startcfg.GetPort(protoStr)
	if !bExist {
		return nil, 105
	}
	cmdrpc := rpcClient[ipPortStr]
	var writeByteArr []byte
	resp, err := cmdrpc.Call(rpcProto)
	if err != nil {
		return nil, 103
	}
	rpcRetVO := resp.(common.RpcS2CRet)
	writeByteArr = rpcRetVO.Respond
	seelog.Trace("writeByteArr", writeByteArr)
	if len(writeByteArr) <= eRROR_DATA_LEN {
		errCode := enumErrCode.GetErrorCode(writeByteArr)
		return nil, errCode
	}
	_, protoId := getLenId(writeByteArr)
	tmp := common.GetProtoName(protoId)
	seelog.Trace("tmp=", tmp)
	sendProtoName := common.GetProtoName(protoId)
	if sendProtoName == "king.S2C_Login" || sendProtoName == "king.S2C_ReConnect" {
		var player_id int32
		if sendProtoName == "S2C_Login" {
			s2c_login := &king.S2C_Login{}
			s2c_loginData := writeByteArr[4:]
			err := proto.Unmarshal(s2c_loginData, s2c_login)
			if err != nil {
				return nil, 104
			}
			player_id = *s2c_login.PlayerId
		} else {
			s2c_reconnect := &king.S2C_ReConnect{}
			s2c_reconnectData := writeByteArr[4:]
			err := proto.Unmarshal(s2c_reconnectData, s2c_reconnect)
			if err != nil {
				return nil, 104
			}
			player_id = *s2c_reconnect.PlayerId
		}
		addConn(ws, player_id)
	}
	pushToOther(&rpcRetVO)
	return writeByteArr, 0
}

func (ws *WS) Read() {
	for {
		writeByteArr, id, errCode := ws.readMsg()
		if errCode == 0 && writeByteArr == nil {
			continue
		}
		if errCode == 101 {
			ws.closeChan <- true
			break
		}
		if errCode != 0 {
			s2c_errMsg := &king.S2C_ErrMsg{
				Err:     proto.Int32(errCode),
				ProtoId: proto.Int32(int32(id)),
			}
			data, err := proto.Marshal(s2c_errMsg)
			if err != nil {
				seelog.Trace("marshal error:", err)
				ws.closeChan <- true
				break
			}
			lenh, lenl := common.GetHLFromWord(len(data) + enumErrCode.DATA_LEN_EXTERN)
			id := common.GetProtoId("S2C_ErrMsg")
			idh, idl := common.GetHLFromWord(id)

			writeByteArr = []byte{lenh, lenl, idh, idl}
			writeByteArr = append(writeByteArr, data...)
		}
		ws.writeMsg <- writeByteArr
	}
}

func (ws *WS) Write() {
L:
	for {
		select {
		case writeByteArr := <-ws.writeMsg:
			err := ws.conn.WriteMessage(websocket.BinaryMessage, writeByteArr)
			if err != nil {
				break L
			}
		case _ = <-ws.closeChan:
			break L
		}
	}

	rmConn(ws)
	_ = ws.conn.Close()
}

func init() {
	portArr := startcfg.GetPortArr()
	rpcClient = make(map[string](*gorpc.Client), len(portArr))
	playerMapConn = make(map[int32](*WS))
	for _, ipPortStr := range portArr {
		rpcAddr := ipPortStr
		c := &gorpc.Client{
			Addr: rpcAddr,
		}
		c.Start()
		rpcClient[ipPortStr] = c
	}
}*/
