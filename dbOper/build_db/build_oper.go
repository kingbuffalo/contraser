package build_db

import (
	"github.com/kingbuffalo/seelog"
)

func NewPlayerBuildingVO(playerId int32) *BuildingVO {
	vo := &BuildingVO{
		PlayerId: playerId,
	}
	if err := vo.SaveNewVO(); err != nil {
		seelog.WarnWE("building save err=", err)
	}
	return vo
}
