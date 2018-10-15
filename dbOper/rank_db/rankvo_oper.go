package rank_db

import (
	"time"
	"buffalo/king/cfgmgr/cfgRankMgr"
	"buffalo/king/common/enumErrCode"
)

var monthMapQuarter []int = []int{0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}

func bTheSameRankQuarter(timestamp int32) (bool, int32) {
	ct := time.Now()
	oldT := time.Unix(int64(timestamp), 0)
	_, oldM, _ := oldT.Date()
	_, cM, _ := ct.Date()
	oldMInt := int(oldM)
	cmInt := int(cM)
	return monthMapQuarter[cmInt] == monthMapQuarter[oldMInt], int32(ct.Unix())
}

func (vo *RankVO) AddScore(score int32) {
	vo.Score += score
	updateGlobalScore(vo.PlayerId, vo.Score)
}

func (vo *RankVO) UpdateAndGetPrevCfg() *cfgRankMgr.CfgRank {
	//TODO ADD gift by email
	bTheSame, ct := bTheSameRankQuarter(vo.RankTimestamp)
	if !bTheSame {
		cfg := cfgRankMgr.GetCfgRankByScore(vo.Score)
		vo.Score = enumErrCode.RANK_SCORE_INIT
		vo.RankTimestamp = ct
		return cfg
	}
	return nil
}
