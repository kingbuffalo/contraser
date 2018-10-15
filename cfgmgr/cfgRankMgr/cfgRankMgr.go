package cfgRankMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
	"sort"
)

type CfgRank struct {
	Id     int32     `json:"id"`
	Name   string    `json:"name"`
	Score  int32     `json:"score"`
	Reward [][]int32 `json:"reward"`
}

var cfgArr []*CfgRank

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_rank.json"

	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	cfgArr = make([]*CfgRank, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	sort.Slice(cfgArr, func(i, j int) bool {
		return cfgArr[i].Score > cfgArr[j].Score
	})
}

func GetCfgRankByScore(score int32) *CfgRank {
	for _, v := range cfgArr {
		if score >= v.Score {
			return v
		}
	}
	cfgLen := len(cfgArr)
	return cfgArr[cfgLen-1]
}
