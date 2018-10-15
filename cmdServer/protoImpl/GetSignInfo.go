package protoImpl

/*

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/dbOper/signInfo_db"
	"buffalo/king/king"
)

type GetSignInfo struct {
	proto_t
}

func (S *GetSignInfo) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if S.bClose {
		return common.NewErrorCodeRet(107)
	}
	c2s_getSignInfo := &king.C2S_GetSignInfo{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s_getSignInfo)
	if err != nil {
		seelog.Trace("proto Unmarshal error!")
		return common.NewErrorCodeRet(3011)
	}

	seelog.Info(c2s_getSignInfo.String())
	signInfoT, err := signInfo_db.GetSignInfoById(rpcC2SProto.PlayerId)
	if err != nil {
		//拉取签到信息失败
		seelog.Trace("GetSignInfoById error")
		return common.NewErrorCodeRet(3012)
	}

	bW, bM := signInfoT.BSignToday()
	s2c_getSignInfo := &king.S2C_GetSignInfo{
		WeekSignDay:     proto.Int32(signInfoT.WeekSignDays),
		MonthSignDay:    proto.Int32(signInfoT.MonthSignDays),
		BWeekSignToday:  proto.Bool(bW),
		BMonthSignToday: proto.Bool(bM),
		WeekId:          proto.Int32(signInfoT.WeekId),
	}
	seelog.Trace(" ok! LastWeekSignTime,LastMonthSignTime, weeksignDays, monthsigndays:", signInfoT.LastWeekSignTime, signInfoT.LastMonthSignTime, signInfoT.WeekSignDays, signInfoT.MonthSignDays)
	playerId := rpcC2SProto.PlayerId
	return getProtoS2CDataWithoutPush(playerId, "S2C_GetSignInfo", s2c_getSignInfo)
}*/
