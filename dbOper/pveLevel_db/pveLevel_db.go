package pveLevel_db

import (
	"buffalo/king/dbOper"
)

type PVELevelVO struct {
	PlayerId int32 `json:"player_id" gorm:"primary_key"`
	LevelId  int32 `json:"level_id"`
}

func (PVELevelVO) TableName() string {
	return "pve_level"
}

const (
	lc_KEY_FMT = "s:sg:PVELevelVO:%d"
)

func newPVELevelVO() *PVELevelVO {
	ret := PVELevelVO{}
	return &ret
}

func (vo *PVELevelVO) SetPlayerId(playerId int32) {
	vo.PlayerId = playerId
}

func (vo *PVELevelVO) BExistInMysql() bool {
	return vo.PlayerId != 0
}

func (vo *PVELevelVO) GetPlayerId() int32 {
	return vo.PlayerId
}

func GetPVELevelVO(playerId int32) (*PVELevelVO, error) {
	vo := newPVELevelVO()
	err := dbOper.GetOneKeyVO(lc_KEY_FMT, playerId, vo)
	if err != nil {
		return nil, err
	}
	return vo, nil
}

func (vo *PVELevelVO) Save(bSaveToMysql bool) error {
	return dbOper.SaveVO(vo, lc_KEY_FMT, bSaveToMysql)
}
