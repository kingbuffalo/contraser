package rank_db

import (
	"time"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper"
)

type RankVO struct {
	PlayerId      int32 `json:"player_id" gorm:"primary_key"`
	Score         int32 `json:"score"`
	RankTimestamp int32 `json:"rank_timestamp"`
}

func (RankVO) TableName() string {
	return "rank"
}

const (
	lc_KEY_FMT = "s:sg:RankVO:%d"
)

func newRankVO() *RankVO {
	ct := int32(time.Now().Unix())
	ret := RankVO{
		Score:         enumErrCode.RANK_SCORE_INIT,
		RankTimestamp: ct,
	}
	return &ret
}

func (vo *RankVO) SetPlayerId(playerId int32) {
	vo.PlayerId = playerId
}

func (vo *RankVO) BExistInMysql() bool {
	return vo.PlayerId != 0
}

func (vo *RankVO) GetPlayerId() int32 {
	return vo.PlayerId
}

func GetRankVO(playerId int32) (*RankVO, error) {
	vo := newRankVO()
	err := dbOper.GetOneKeyVO(lc_KEY_FMT, playerId, vo)
	if err != nil {
		return nil, err
	}
	return vo, nil
}

func (vo *RankVO) Save(bSaveToMysql bool) error {
	return dbOper.SaveVO(vo, lc_KEY_FMT, bSaveToMysql)
}
