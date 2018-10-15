package cfgWeekSignMgr

import (
	"encoding/json"
	//"errors"
	//"fmt"
	//"github.com/kingbuffalo/seelog"
	"io/ioutil"
	//"strconv"
	//"time"
	"buffalo/king/common/startcfg"
)

const (
	lc_MAX_WEEK_DAYS = 7
)

type CfgWeekSign struct {
	Id     int32 `json:id`
	ConfId int32 `json:confId`
}

var weekIdMap_weekdaysMapCfgWeekSignGift map[int32](map[int32]*CfgWeekSignGift)
var default_weekdaysMapCfgWeekSignGift map[int32]*CfgWeekSignGift

func init() {
	cfgArr := make([]*CfgWeekSign, 0)

	path := startcfg.GetCfgPath()
	fn := path + "cfg_weeksign.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	weekIdMap_weekdaysMapCfgWeekSignGift = make(map[int32](map[int32]*CfgWeekSignGift), 0)
	for _, v := range cfgArr {
		weekDaysMapCfg, ok := weekIdMap_weekdaysMapCfgWeekSignGift[v.Id]
		if !ok {
			weekDaysMapCfg = make(map[int32]*CfgWeekSignGift, lc_MAX_WEEK_DAYS)
		}
		for i := 1; i < lc_MAX_WEEK_DAYS; i++ {
			cfgWeekSignGift := getCfgWeekSignGift(v.ConfId, int32(i))
			if cfgWeekSignGift == nil {
				break
			}
			weekDaysMapCfg[cfgWeekSignGift.WeekDays] = cfgWeekSignGift
		}
		weekIdMap_weekdaysMapCfgWeekSignGift[v.Id] = weekDaysMapCfg
	}

	var defaultConfId int32 = 0
	default_weekdaysMapCfgWeekSignGift = make(map[int32]*CfgWeekSignGift, lc_MAX_WEEK_DAYS)
	for i := 1; i < lc_MAX_WEEK_DAYS; i++ {
		cfgWeekSignGift := getCfgWeekSignGift(defaultConfId, int32(i))
		if cfgWeekSignGift == nil {
			break
		}
		default_weekdaysMapCfgWeekSignGift[cfgWeekSignGift.WeekDays] = cfgWeekSignGift
	}
}

func GetCfgWeekSignGift(weekId, weekDays int32) *CfgWeekSignGift {
	weekDaysMapCfg, ok := weekIdMap_weekdaysMapCfgWeekSignGift[weekId]
	if !ok {
		weekDaysMapCfg = default_weekdaysMapCfgWeekSignGift
	}
	cfg, ok := weekDaysMapCfg[weekDays]
	if !ok {
		return nil
	}
	return cfg
}
