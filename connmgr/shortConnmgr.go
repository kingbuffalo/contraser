package connmgr

import (
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/kingbuffalo/seelog"
	"github.com/valyala/gorpc"
	//"time"
)

type SWS struct {
	conn *websocket.Conn
}

const (
	eRROR_DATA_LEN = 4
)

func NewSWS(conn *websocket.Conn) *SWS {
	return &SWS{
		conn: conn,
	}
}

func (s *SWS) rpcReadMsg(protoName string, rpcProto *common.RpcC2SProto) ([]byte, int32) {
	protoStr := protoName[enumErrCode.PROTO_PREFIX_LEN:]
	seelog.Trace("readMsg,", protoStr)
	seelog.Trace("readMsg,", protoName)
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
	return writeByteArr, 0
}

func (s *SWS) readMsg_respond(dataLen, messageLen, id int, pid int32, data []byte) ([]byte, int, int32) {
	if dataLen != messageLen-6 {
		return nil, id, 102
	}
	protoName := common.GetProtoName(id)
	rpcProto := &common.RpcC2SProto{
		PlayerId: pid,
		Len:      dataLen,
		Pid:      id,
		Data:     data,
	}
	wMsg, errCode := s.rpcReadMsg(protoName, rpcProto)
	return wMsg, id, errCode
}

func (s *SWS) readMsg() ([]byte, int, int32) {
	wc := s.conn

	_, message, err := wc.ReadMessage()
	if err != nil {
		seelog.Trace(err)
		return nil, 0, 101
	}

	dataLen, id := getLenId(message)
	pid := getPlayerId(message)
	data := message[8:]
	seelog.Trace("message=", message)
	messageLen := len(message)
	return s.readMsg_respond(dataLen, messageLen, id, pid, data)
}

func (s *SWS) Respond() {
	writeByteArr, id, errCode := s.readMsg()
	if errCode == 0 && writeByteArr == nil {
		return
	}
	if errCode == 101 {
		return
	}
	if errCode != 0 {
		s2c_errMsg := &king.S2C_ErrMsg{
			Err:     proto.Int32(errCode),
			ProtoId: proto.Int32(int32(id)),
		}
		data, err := proto.Marshal(s2c_errMsg)
		if err != nil {
			seelog.Trace("marshal error:", err)
			return
		}
		lenh, lenl := common.GetHLFromWord(len(data) + enumErrCode.DATA_LEN_EXTERN)
		id := common.GetProtoId("S2C_ErrMsg")
		idh, idl := common.GetHLFromWord(id)

		writeByteArr = []byte{lenh, lenl, idh, idl}
		writeByteArr = append(writeByteArr, data...)
	}
	if err := s.conn.WriteMessage(websocket.BinaryMessage, writeByteArr); err != nil {
		return
	}
	_ = s.conn.Close()
}

var rpcClient map[string](*gorpc.Client)

func init() {
	portArr := startcfg.GetPortArr()
	rpcClient = make(map[string](*gorpc.Client), len(portArr))
	for _, ipPortStr := range portArr {
		rpcAddr := ipPortStr
		c := &gorpc.Client{
			Addr: rpcAddr,
		}
		c.Start()
		rpcClient[ipPortStr] = c
	}
}

func getPlayerId(message []byte) int32 {
	pid1, pid2, pid3, pid4 := message[4], message[5], message[6], message[7]
	pid := (int(pid1) << 24) | (int(pid2) << 16) | (int(pid3) << 8) | int(pid4)
	return int32(pid)
}
func getLenId(message []byte) (int, int) {
	lenh, lenl := message[0], message[1]
	idh, idl := message[2], message[3]

	id := (int(idh)<<8 | int(idl)) & 0xffff
	dataLen := (int(lenh)<<8 | int(lenl)) & 0xffff
	return dataLen, id
}
