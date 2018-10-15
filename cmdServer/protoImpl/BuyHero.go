package protoImpl

/*
import (
	"buffalo/king/cfgmgr/cfgHeroMgr"
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/hero_db"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
)

type BuyHero struct {
	proto_t
}

func (b *BuyHero) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	var c2s_buyHero king.C2S_BuyHero
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s_buyHero); err != nil {
		return common.NewErrorCodeRet(4011)
	}
	seelog.Info(c2s_buyHero.PlayerId, ",", c2s_buyHero.String())
	heroId := cfgHeroMgr.FakeBuyHero()
	heroVO, err := hero_db.GetHeroVO(rpcC2SProto.PlayerId, heroId)
	var chip int32 = 0
	var chipHeroIdArr []int32 = make([]int32, 0)
	if err != nil {
		return common.NewErrorCodeRet(4012)
	}
	//随机产生一个heroId
	cfgHero := cfgHeroMgr.GetCfgHero(heroId)
	cityId := enumErrCode.DEFAULT_CITY_ID
	if heroVO == nil {
		heroVO = hero_db.NewHero(cfgHero, cityId, rpcC2SProto.PlayerId)
		if err := heroVO.Save(true); err != nil {
			seelog.WarnWE(err.Error())
			return common.NewErrorCodeRet(4015)
		}
	} else {
		chip = cfgHero.Chip
		chipHeroIdArr = append(chipHeroIdArr, chip)
		playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(rpcC2SProto.PlayerId)
		if err != nil {
			seelog.WarnWE(err.Error())
			return common.NewErrorCodeRet(4013)
		}
		playerInfoVO.Chip += chip
		if err := playerInfoVO.Save(); err != nil {
			seelog.WarnWE(err.Error())
			return common.NewErrorCodeRet(4014)
		}
	}

	s2c_buyHero := king.S2C_BuyHero{
		HeroIds:       []int32{heroId},
		Chip:          proto.Int32(chip), //TODO not the final Chip
		ToChipHeroIds: chipHeroIdArr,
	}
	playerId := rpcC2SProto.PlayerId
	return getProtoS2CDataWithoutPush(playerId, "S2C_BuyHero", &s2c_buyHero)
}*/
