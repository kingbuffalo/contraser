package cfgMonthSignMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgMonthSign struct {
	Id     int32 `json:id`
	ConfId int32 `json:confId`
}

const (
	lc_MAX_MONTH_DAYS = 32
	lc_MIN_MONTH_DAYS = 28
)

var monthIdMap_monthdaysMapCfgMonthSignGift map[int32](map[int32]*CfgMonthSignGift)
var default_monthdaysMapCfgMonthSignGift map[int32]*CfgMonthSignGift

func init() {
	cfgArr := make([]*CfgMonthSign, 0)

	path := startcfg.GetCfgPath()
	fn := path + "cfg_monthsign.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	monthIdMap_monthdaysMapCfgMonthSignGift = make(map[int32](map[int32]*CfgMonthSignGift), 0)
	for _, v := range cfgArr {
		monthDaysMapCfg, ok := monthIdMap_monthdaysMapCfgMonthSignGift[v.Id]
		if !ok {
			monthDaysMapCfg = make(map[int32]*CfgMonthSignGift, lc_MIN_MONTH_DAYS)
		}
		for i := 1; i < lc_MAX_MONTH_DAYS; i++ {
			cfgMonthSignGift := getCfgMonthSignGift(v.ConfId, int32(i))
			if cfgMonthSignGift == nil {
				break
			}
			monthDaysMapCfg[cfgMonthSignGift.MonthDays] = cfgMonthSignGift
		}
		monthIdMap_monthdaysMapCfgMonthSignGift[v.Id] = monthDaysMapCfg
	}

	var defaultConfId int32 = 0
	default_monthdaysMapCfgMonthSignGift = make(map[int32]*CfgMonthSignGift, lc_MIN_MONTH_DAYS)
	for i := 1; i < lc_MAX_MONTH_DAYS; i++ {
		cfgMonthSignGift := getCfgMonthSignGift(defaultConfId, int32(i))
		if cfgMonthSignGift == nil {
			break
		}
		default_monthdaysMapCfgMonthSignGift[cfgMonthSignGift.MonthDays] = cfgMonthSignGift
	}
}

func GetCfgMonthSignGift(monthId, monthDays int32) *CfgMonthSignGift {
	monthDaysMapCfg, ok := monthIdMap_monthdaysMapCfgMonthSignGift[monthId]
	if !ok {
		monthDaysMapCfg = default_monthdaysMapCfgMonthSignGift
	}
	cfg, ok := monthDaysMapCfg[monthDays]
	if !ok {
		return nil
	}
	return cfg
}
