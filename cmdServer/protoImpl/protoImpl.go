package protoImpl

import (
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	//"buffalo/king/dbOper/playerInfo_db"
)

type ProtoWorker interface {
	DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet
	Close() error
}

type proto_t struct {
	bClose bool
}

func (p *proto_t) Close() error {
	if !p.bClose {
		p.bClose = true
		if err := gameutil.CloseRedis(); err != nil {
			return err
		}
		if err := gameutil.CloseMysql(); err != nil {
			return err
		}
	}
	return nil
}

var nameMapWorker map[string]ProtoWorker = map[string]ProtoWorker{
	"Login":         &Login{},
	"HeroRecruit":   &HeroRecruit{},
	"GetInitGame":   &GetInitGame{},
	"FightPVE":      &FightPVE{},
	"EnQueueSingle": &EnQueueSingle{},
	"ContinueFight": &ContinueFight{},
	/*
		"AttackCity":            &AttackCity{},
		"BuyHero":               &BuyHero{},
		"EnQueue":               &EnQueue{},
		"GetAllBattleResult":    &GetAllBattleResult{},
		"GetBattleResultDetail": &GetBattleResultDetail{},
		"GetMapCitys":           &GetMapCitys{},
		"GetRedPoint":           &GetRedPoint{},
		"GetSignInfo":           &GetSignInfo{},
		"Heart":                 &Heart{},
		"Ping":                  &Ping{},
		"ReConnect":             &ReConnect{},
		"RefreshMapCitys":       &RefreshMapCitys{},
		"SetDefArmy":            &SetDefArmy{},
		"Sign":                  &Sign{},
		"TimerRefreshMapCitys":  &TimerRefreshMapCitys{},
		"EnQueueSingle":         &EnQueueSingle{},
	*/
}

func CreateProtoWorker(protoName string) ProtoWorker {
	worker, ok := nameMapWorker[protoName]
	if !ok {
		panic("proto not found")
	}
	return worker
}

//func getPhpLoginKey(sessionId string) string {
//return "s:pl:sessionid" + sessionId
//}

type phpLoginJson_t struct {
	OpenId     string `json:"openid"`
	AppId      string `json:"appid"`
	SessionKey string `json:"session_key"`
}

////////////login php backend  end

func getProtoS2CDataWithoutPush(playerId int32, protoName string, pb proto.Message) *common.RpcS2CRet {
	if playerId != 0 {
		seelog.Info("send:playerId=", playerId, ",msg=", pb.String())
	}
	data, err := proto.Marshal(pb)
	if err != nil {
		return common.NewErrorCodeRet(111)
	}
	lenh, lenl := common.GetHLFromWord(len(data) + enumErrCode.DATA_LEN_EXTERN)
	id := common.GetProtoId(protoName)
	idh, idl := common.GetHLFromWord(id)

	dataWithLen := []byte{lenh, lenl, idh, idl}
	dataWithLen = append(dataWithLen, data...)
	//return dataWithLen
	return &common.RpcS2CRet{
		Respond: dataWithLen,
	}
}

/*
func getProtoS2CDataWithPush1(protoName string, pb proto.Message, push *common.PushToOtherMsg) *common.RpcS2CRet {
	seelog.Info("send ", pb.String())
	data, err := proto.Marshal(pb)
	if err != nil {
		return common.NewErrorCodeRet(112)
	}
	lenh, lenl := common.GetHLFromWord(len(data) + enumErrCode.DATA_LEN_EXTERN)
	id := common.GetProtoId(protoName)
	idh, idl := common.GetHLFromWord(id)

	dataWithLen := []byte{lenh, lenl, idh, idl}
	dataWithLen = append(dataWithLen, data...)

	p, err := push.ToPushToOthers()
	if err == nil {
		pushArr := []common.PushToOthers{*p}
		return &common.RpcS2CRet{
			Respond: dataWithLen,
			PushArr: pushArr,
		}
	} else {
		return &common.RpcS2CRet{
			Respond: dataWithLen,
		}
	}

}
*/
