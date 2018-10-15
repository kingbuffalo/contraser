package mapCity_db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/cfgmgr/cfgMapCityMgr"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/commonVO"
)

type MapCityNpcVO struct {
	NpcIdx        int32                    `json:"npc_idx"`
	MapNpcHeroVOs []*commonVO.MapNpcHeroVO `json:"map_npc_heroVOs"`
	Res           *commonVO.RobAbleResVO   `json:"res"`
}

func getKey(npcIdx int32) string {
	return fmt.Sprintf("s:sg:MapCityNpcVO:%d")
}

func NewMapCityNpcVO(mapCityNpcId int32) *MapCityNpcVO {
	cfg := cfgMapCityMgr.GetCfgMapCityNpcVO(mapCityNpcId)
	if cfg != nil {
		npcIdx := gameutil.GIncr(enumErrCode.NPC_IDX_INC_KEY)
		return &MapCityNpcVO{
			NpcIdx:        int32(npcIdx),
			MapNpcHeroVOs: cfg.MapNpcHeroVOs,
			Res:           cfg.Res,
		}
	}
	return nil
}

func GetMapCityNpcVO(npcIdx int32) (*MapCityNpcVO, error) {
	seelog.Trace("npcIdx:", npcIdx)
	key := getKey(npcIdx)
	rep := gameutil.GGet(key)
	if rep == nil {
		return nil, errors.New("not found")
		//vo, err := getMapCityNpcVOFromMysql(npcIdx)
		//if err != nil {
		//return nil, err
		//}
		////set 到 redis中
		//err :=vo.SaveVO(false)
		//return vo, err
	}
	var vo MapCityNpcVO
	err := json.Unmarshal(rep, &vo)
	if err != nil {
		seelog.Trace(" Unmarshal error")
		return nil, err
	}
	return &vo, nil
}

func (vo *MapCityNpcVO) SaveVO( /*bSaveToMysql bool*/ ) error {
	key := getKey(vo.NpcIdx)
	b, err := json.Marshal(vo)
	if err != nil {
		return nil
	}
	gameutil.GSet(key, b)

	//if bSaveToMysql {
	//db := gameutil.GetDB()
	//db.Model(&vo).Updates(vo)
	//}
	return nil
}
