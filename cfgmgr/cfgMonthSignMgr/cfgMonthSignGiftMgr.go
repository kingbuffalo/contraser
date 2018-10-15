package cfgMonthSignMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgMonthSignGift struct {
	RewardId  int32     `json:rewardId`
	MonthDays int32     `json:monthDays`
	Gift      [][]int32 `json:gift`
	ExtGift   [][]int32 `json:extGift`
}

var cfgArr []*CfgMonthSignGift

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_monthsigngift.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	cfgArr = make([]*CfgMonthSignGift, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
}

func getCfgMonthSignGift(rewardId, monthDays int32) *CfgMonthSignGift {
	for _, v := range cfgArr {
		if v.RewardId == rewardId && v.MonthDays == monthDays {
			return v
		}
	}
	return nil
}
