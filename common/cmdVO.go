package common

import (
	//"buffalo/king/common/enumErrCode"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"strconv"
)

type CmdVO struct {
	Cmd int    `json:"cmd"`
	Pid int    `json:"pid"`
	P1  int    `json:"p1"`
	P2  int    `json:"p2"`
	P3  int    `json:"p3"`
	P4  string `json:"p4"`
	P5  string `json:"p5"`
}

type RpcC2SProto struct {
	PlayerId int32  `json:"playerId"`
	Len      int    `json:"len"`
	Pid      int    `json:"pid"`
	Data     []byte `json:"data"`
}

func Pb2ByteArr(len, id int, pb proto.Message) ([]byte, error) {
	data, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}
	lenh, lenl := GetHLFromWord(len)
	idh, idl := GetHLFromWord(id)
	dataWithLen := []byte{lenh, lenl, idh, idl}
	dataWithLen = append(dataWithLen, data...)
	return dataWithLen, nil
}

func Pb2ByteArrClient(playerId, id int, pb proto.Message) ([]byte, error) {
	data, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}
	len := len(data) + 2
	lenh, lenl := GetHLFromWord(len)
	idh, idl := GetHLFromWord(id)
	dataWithLen := []byte{lenh, lenl, idh, idl, 0, 0, 0, 0}
	p1, p2, p3, p4 := GetBytesFromInt(playerId)
	dataWithLen[4] = p1
	dataWithLen[5] = p2
	dataWithLen[6] = p3
	dataWithLen[7] = p4
	dataWithLen = append(dataWithLen, data...)
	return dataWithLen, nil
}
func Pb2ByteArrBy(id int, pb proto.Message) ([]byte, error) {
	data, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}
	len := len(data) + 2
	lenh, lenl := GetHLFromWord(len)
	idh, idl := GetHLFromWord(id)
	dataWithLen := []byte{lenh, lenl, idh, idl}
	dataWithLen = append(dataWithLen, data...)
	return dataWithLen, nil
}

type RpcS2CRet struct {
	Respond []byte `json:"respond"`
}

func createErrorCode(errCode int) []byte {
	int3 := errCode >> 24
	int2 := (errCode >> 16) & 0xff
	int1 := (errCode >> 8) & 0xff
	int0 := errCode & 0xff
	return []byte{byte(int3), byte(int2), byte(int1), byte(int0)}
}

func NewErrorCodeRet(errCode int) *RpcS2CRet {
	seelog.Tracef("error code = %d", errCode)
	res := createErrorCode(errCode)
	return &RpcS2CRet{
		Respond: res,
	}
}

func (r *RpcC2SProto) String() string {
	return "len=" + strconv.Itoa(r.Len) + ",protoId=" + strconv.Itoa(r.Pid)
}
