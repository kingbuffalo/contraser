package playerInfo_db

import (
	"buffalo/king/common/gameutil"
	"encoding/json"
)

func (v *PlayerInfoVO) Clone() *PlayerInfoVO {
	return &PlayerInfoVO{
		PlayerId: v.PlayerId,
		Name:     v.Name,
		Url:      v.Url,
		OpenId:   v.OpenId,
		Coin:     v.Coin,
		Diamond:  v.Diamond,
		Wood:     v.Wood,
		Mineral:  v.Mineral,
		Level:    v.Level,
		FiveStarHeroRecruitTimes: v.FiveStarHeroRecruitTimes,
		TenTimesHeroRecuritTimes: v.TenTimesHeroRecuritTimes,
		FreeHeroRecruitTimestamp: v.FreeHeroRecruitTimestamp,
	}
}

func (v *PlayerInfoVO) getRedisKey() string {
	return getPlayerInfoKey(v.PlayerId)
}

func (v *PlayerInfoVO) DiffSave(oldv *PlayerInfoVO) error {
	key := getPlayerInfoKey(v.PlayerId)
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	gameutil.GSet(key, b)

	diff := v.getDiff(oldv)
	db := gameutil.GetDB()
	db.Model(v).Updates(diff)
	return nil
}

/*
func (v *PlayerInfoVO) GetUpdatePush(oldv *PlayerInfoVO) *king.S2C_UpdatePush {
	updatePushArr := make([]*king.UpdatePush, 0)
	if v.Coin != oldv.Coin {
		up := &king.UpdatePush{
			Type:  proto.Int32(enumErrCode.SPEC_RESCOURCE_COIN),
			Value: proto.Int32(v.Coin),
		}
		updatePushArr = append(updatePushArr, up)
	}
	if v.Diamond != oldv.Diamond {
		up := &king.UpdatePush{
			Type:  proto.Int32(enumErrCode.SPEC_RESCOURCE_DIAMOND),
			Value: proto.Int32(v.Diamond),
		}
		updatePushArr = append(updatePushArr, up)
	}
	if v.Wood != oldv.Wood {
		up := &king.UpdatePush{
			Type:  proto.Int32(enumErrCode.SPEC_RESCOURCE_WOOD),
			Value: proto.Int32(v.Wood),
		}
		updatePushArr = append(updatePushArr, up)
	}
	if v.Mineral != oldv.Mineral {
		up := &king.UpdatePush{
			Type:  proto.Int32(enumErrCode.SPEC_RESCOURCE_MINERAL),
			Value: proto.Int32(v.Mineral),
		}
		updatePushArr = append(updatePushArr, up)
	}
	if v.Chip != oldv.Chip {
		up := &king.UpdatePush{
			Type:  proto.Int32(enumErrCode.SPEC_RESCOURCE_MINERAL),
			Value: proto.Int32(v.Chip),
		}
		updatePushArr = append(updatePushArr, up)
	}

	return &king.S2C_UpdatePush{
		UpdatePushs: updatePushArr,
	}
}
*/

func (v *PlayerInfoVO) getDiff(oldv *PlayerInfoVO) *PlayerInfoVO {
	var ret PlayerInfoVO
	if v.Name != oldv.Name {
		ret.Name = v.Name
	}
	if v.Url != oldv.Url {
		ret.Url = v.Url
	}
	if v.OpenId != oldv.OpenId {
		ret.OpenId = v.OpenId
	}
	if v.Coin != oldv.Coin {
		ret.Coin = v.Coin
	}
	if v.Diamond != oldv.Diamond {
		ret.Diamond = v.Diamond
	}
	if v.Wood != oldv.Wood {
		ret.Wood = v.Wood
	}
	if v.Mineral != oldv.Mineral {
		ret.Mineral = v.Mineral
	}
	if v.Level != oldv.Level {
		ret.Level = v.Level
	}

	if v.FiveStarHeroRecruitTimes != oldv.FiveStarHeroRecruitTimes {
		ret.FiveStarHeroRecruitTimes = v.FiveStarHeroRecruitTimes
	}
	if v.TenTimesHeroRecuritTimes != oldv.TenTimesHeroRecuritTimes {
		ret.TenTimesHeroRecuritTimes = v.TenTimesHeroRecuritTimes
	}

	if v.FreeHeroRecruitTimestamp != oldv.FreeHeroRecruitTimestamp {
		ret.FreeHeroRecruitTimestamp = v.FreeHeroRecruitTimestamp
	}

	return &ret
}
