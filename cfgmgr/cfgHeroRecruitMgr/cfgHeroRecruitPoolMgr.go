package cfgHeroRecruitMgr

import (
	"buffalo/king/cfgmgr/cfgHeroMgr"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type CfgHeroRecruitPool struct {
	Id             int32     `json:"id"`
	LocIdCount     [][]int32 `json:"loc_id_count"`
	Weight         int32     `json:"weight"`
	AccWeightTimes int32     `json:"acc_weight_times"`
	AccWeightValue int32     `json:"acc_weight_value"`
}

var id_map_accsum map[int32]int32
var id_map_origsum map[int32]int32
var id_map_cfgArr map[int32]([]*CfgHeroRecruitPool)

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_hero_recruit_pool.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgHeroRecruitPool = make([]*CfgHeroRecruitPool, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	id_map_cfgArr = make(map[int32]([]*CfgHeroRecruitPool), 0)
	id_map_accsum = make(map[int32]int32, 0)
	id_map_origsum = make(map[int32]int32, 0)

	for _, v := range cfgArr {
		arr := id_map_cfgArr[v.Id]
		if arr == nil {
			arr = make([]*CfgHeroRecruitPool, 0)
		}
		arr = append(arr, v)
		id_map_cfgArr[v.Id] = arr

		accsum, ok := id_map_accsum[v.Id]
		if !ok {
			accsum = 0
		}
		accsum += v.AccWeightValue
		id_map_accsum[v.Id] = accsum

		origsum, ok := id_map_origsum[v.Id]
		if !ok {
			origsum = 0
		}
		origsum += v.Weight
		id_map_origsum[v.Id] = origsum
		if len(v.LocIdCount) != 1 {
			errMsg := "len of loc id count != 1"
			panic(errMsg)
		}
		for _, locidcountArr := range v.LocIdCount {
			heroId := locidcountArr[1]
			cfgHero := cfgHeroMgr.GetCfgHero(heroId)
			if cfgHero == nil {
				errMsg := fmt.Sprintf("hero not found in hero_recruit_pool cfg: heroId=%d", heroId)
				panic(errMsg)
			}
		}
	}
}

func (cfg *CfgHeroRecruitPool) getSelfWeightSum(accTimes int32) int32 {
	return accTimes*cfg.AccWeightValue + cfg.Weight
}
func getTotalSum(id, accTimes int32) int32 {
	sum := id_map_origsum[id]
	acc, ok := id_map_accsum[id]
	if ok {
		sum += acc * accTimes
	}
	return sum
}

func randHeroId(id, randTimes, accTimes, tenTimes int32, isAcc bool) ([][]int32, int32, int32) {
	ret := make([][]int32, randTimes)

	var i int32
	for i = 0; i < randTimes; i++ {
		poolId := id
		if isAcc {
			if tenTimes+1 == enumErrCode.HERO_RECRUIT_SPEC_TIMES {
				poolId = enumErrCode.HERO_RECRUIT_TENTH_POOL_ID
				tenTimes = 0
			} else {
				tenTimes++
			}
		}
		sum := getTotalSum(poolId, accTimes)
		randInt := int32(rand.Intn(int(sum)))
		arr := id_map_cfgArr[poolId]
		var chooseHeroId []int32
		var chooseSum int32 = 0
		for _, v := range arr {
			chooseSum += v.getSelfWeightSum(accTimes)
			if chooseSum >= randInt {
				chooseHeroId = v.LocIdCount[0]
				break
			}
		}
		if chooseHeroId == nil {
			chooseHeroId = arr[0].LocIdCount[0]
		}
		ret[i] = chooseHeroId
		hid := chooseHeroId[1]
		heroCfg := cfgHeroMgr.GetCfgHero(hid)
		if isAcc {
			if heroCfg.StarLv >= enumErrCode.HERO_RECRUIT_SPEC_STAR {
				accTimes = 0
			} else {
				accTimes++
			}
		}
	}

	return ret, accTimes, tenTimes
}
