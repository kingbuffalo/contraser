package cfgWeekSignMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgWeekSignGift struct {
	RewardId int32     `json:rewardId`
	WeekDays int32     `json:weekDays`
	Gift     [][]int32 `json:gift`
}

var cfgArr []*CfgWeekSignGift

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_weeksigngift.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	cfgArr = make([]*CfgWeekSignGift, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
}

func getCfgWeekSignGift(rewardId, weekDays int32) *CfgWeekSignGift {
	for _, v := range cfgArr {
		if v.RewardId == rewardId && v.WeekDays == weekDays {
			return v
		}
	}
	return nil
}
