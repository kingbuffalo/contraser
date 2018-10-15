package protoImpl

import (
	"buffalo/king/common"
	"buffalo/king/dbOper/pveLevel_db"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
)

type GetPVELevelId struct {
	proto_t
}

func (p *GetPVELevelId) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_GetPVELevelId
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(7051)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	playerId := rpcC2SProto.PlayerId
	//TODO ADD YOUR CODE HERE
	pveLevelVO, err := pveLevel_db.GetPVELevelVO(playerId)
	if err != nil {
		return common.NewErrorCodeRet(7052)
	}

	s2c := &king.S2C_GetPVELevelId{
		LevelId: proto.Int32(pveLevelVO.LevelId),
	}
	return getProtoS2CDataWithoutPush(playerId, "S2C_GetPVELevelId", s2c)
}
