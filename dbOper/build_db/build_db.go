package build_db

import (
	"buffalo/king/dbOper"
)

const (
	lc_KEY_FMT = "s:sg:BuildingVO:%d"
)

type BuildingVO struct {
	PlayerId int32 `json:"player_id"`

	Farm1Lv int32 `json:"farm1_lv"`
	Farm2Lv int32 `json:"farm2_lv"`
}

func (BuildingVO) TableName() string {
	return "building"
}

func newBuildingVO() *BuildingVO {
	ret := BuildingVO{}
	return &ret
}

func (vo *BuildingVO) SetPlayerId(playerId int32) {
	vo.PlayerId = playerId
}

func (vo *BuildingVO) BExistInMysql() bool {
	return vo.PlayerId != 0
}

func (vo *BuildingVO) GetPlayerId() int32 {
	return vo.PlayerId
}

func GetBuildingVO(playerId int32) (*BuildingVO, error) {
	vo := newBuildingVO()
	err := dbOper.GetOneKeyVO(lc_KEY_FMT, playerId, vo)
	if err != nil {
		return nil, err
	}
	return vo, nil
}

func (vo *BuildingVO) Save(bSaveToMysql bool) error {
	return dbOper.SaveVO(vo, lc_KEY_FMT, bSaveToMysql)
}

func (vo *BuildingVO) SaveNewVO() error {
	return dbOper.SaveNewVO(vo, lc_KEY_FMT)
}
