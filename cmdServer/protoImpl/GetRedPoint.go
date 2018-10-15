package protoImpl

/*

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/signInfo_db"
	"buffalo/king/king"
)

type GetRedPoint struct {
	proto_t
}

func (S *GetRedPoint) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	seelog.Trace("[GetRedPoint::DoTheJob] start")
	if S.bClose {
		return common.NewErrorCodeRet(107)
	}
	c2s_getRedPoint := &king.C2S_GetRedPoint{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s_getRedPoint)
	if err != nil {
		seelog.Trace("[GetRedPoint::DoTheJob] proto Unmarshal error!")
		return common.NewErrorCodeRet(3021)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s_getRedPoint.String())
	signInfoT, err := signInfo_db.GetSignInfoById(rpcC2SProto.PlayerId)
	if err != nil {
		//拉取签到信息失败
		seelog.Trace("[GetRedPoint::DoTheJob] GetSignInfoById error")
		return common.NewErrorCodeRet(3022)
	}
	weekSignStatus := gameutil.BToday(signInfoT.LastWeekSignTime)
	monthsignStatus := gameutil.BToday(signInfoT.LastMonthSignTime)
	s2c_getRedPoint := &king.S2C_GetRedPoint{
	}
	seelog.Trace("[GetRedPoint::DoTheJob] ok! WeekSign, Monthsign:", weekSignStatus, monthsignStatus)
	playerId := rpcC2SProto.PlayerId
	return getProtoS2CDataWithoutPush(playerId, "S2C_GetRedPoint", s2c_getRedPoint)
}*/
