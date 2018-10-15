package protoImpl

/*

import ( //"encoding/json"
	//"errors"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/cfgmgr/cfgMonthSignMgr"
	"buffalo/king/cfgmgr/cfgWeekSignMgr"
	"buffalo/king/cmdServer/protoImpl/comFunc"
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/dbOper/signInfo_db"
	"buffalo/king/king"
)

type Sign struct {
	proto_t
}

func (S *Sign) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if S.bClose {
		return common.NewErrorCodeRet(107)
	}
	c2s := &king.C2S_Sign{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s)
	if err != nil {
		seelog.Trace("[Sign::DoTheJob] proto Unmarshal error!")
		return common.NewErrorCodeRet(3001)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	signType := *c2s.SignType

	if !(signType == enumErrCode.SIGN_TYPE_WEEK || signType == enumErrCode.SIGN_TYPE_MONTH) {
		seelog.Trace("[Sign::DoTheJob] signType is not define", rpcC2SProto.PlayerId)
		return common.NewErrorCodeRet(3005)
	}

	signInfo, err := signInfo_db.GetSignInfoById(rpcC2SProto.PlayerId)
	if err != nil {
		seelog.Trace("[Sign::DoTheJob] GetSignInfoById error", rpcC2SProto.PlayerId)
		return common.NewErrorCodeRet(3002)
	}
	seelog.Trace("signInfo: ", signInfo)
	lastSignTime := signInfo.LastWeekSignTime
	if signType == enumErrCode.SIGN_TYPE_MONTH {
		lastSignTime = signInfo.LastMonthSignTime
	}

	bTodaySign := gameutil.BToday(lastSignTime)
	if bTodaySign {
		return common.NewErrorCodeRet(3003)
	}
	playerId := rpcC2SProto.PlayerId
	playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(playerId)
	if err != nil {
		seelog.WarnWE("GetPlayerInfoVOBy error=", err)
	}

	var giftIdCoArr [][]int32
	var extGiftIdCoArr [][]int32
	if signType == enumErrCode.SIGN_TYPE_MONTH {
		cfg := cfgMonthSignMgr.GetCfgMonthSignGift(signInfo.MonthId, signInfo.MonthSignDays+1)
		if cfg == nil {
			return common.NewErrorCodeRet(3006)
		}
		giftIdCoArr = cfg.Gift
		extGiftIdCoArr = cfg.ExtGift
	} else {
		cfg := cfgWeekSignMgr.GetCfgWeekSignGift(signInfo.WeekId, signInfo.WeekSignDays+1)
		if cfg == nil {
			return common.NewErrorCodeRet(3007)
		}
		giftIdCoArr = cfg.Gift
		extGiftIdCoArr = nil
	}

	comFunc.AddPlayerGift(playerInfoVO, giftIdCoArr)
	if extGiftIdCoArr != nil {
		comFunc.AddPlayerGift(playerInfoVO, extGiftIdCoArr)
	}
	if err := playerInfoVO.Save(); err != nil {
		return common.NewErrorCodeRet(3008)
	}

	signInfo.SignToday(signType)
	if err := signInfo.SaveVO(true); err != nil {
		return common.NewErrorCodeRet(3009)
	}

	reward := comFunc.GiftIdNumArrToGiftInfo(giftIdCoArr)
	var extReward []*king.GiftInfo
	if extGiftIdCoArr != nil {
		extReward = comFunc.GiftIdNumArrToGiftInfo(extGiftIdCoArr)
	}

	s2c := &king.S2C_Sign{
		SignType:  proto.Int32(*c2s.SignType),
		Reward:    reward,
		ExtReward: extReward,
	}
	return getProtoS2CDataWithoutPush(playerId, "S2C_Sign", s2c)
}*/
