package playerInfo_db

import (
	"buffalo/king/common/enumErrCode"
)

func (p *PlayerInfoVO) BHasEnought(resType, resValue int32) bool {
	switch resType {
	case enumErrCode.SPEC_RESCOURCE_DIAMOND:
		return p.Diamond >= resValue
	case enumErrCode.SPEC_RESCOURCE_COIN:
		return p.Coin >= resValue
	case enumErrCode.SPEC_RESCOURCE_WOOD:
		return p.Wood >= resValue
	case enumErrCode.SPEC_RESCOURCE_MINERAL:
		return p.Mineral >= resValue
	case enumErrCode.SPEC_RESCOURCE_GRAIN:
		return p.Grain >= resValue
	}
	return false
}

func (p *PlayerInfoVO) AddRes(resType, resValue int32) {
	switch resType {
	case enumErrCode.SPEC_RESCOURCE_DIAMOND:
		p.Diamond += resValue
	case enumErrCode.SPEC_RESCOURCE_COIN:
		p.Coin += resValue
	case enumErrCode.SPEC_RESCOURCE_WOOD:
		p.Wood += resValue
	case enumErrCode.SPEC_RESCOURCE_MINERAL:
		p.Mineral += resValue
	case enumErrCode.SPEC_RESCOURCE_GRAIN:
		p.Grain += resValue
	}

}

func (vo *PlayerInfoVO) CheckPlayerGift(giftIdNumArr [][]int32) bool {
	for _, giftIdNum := range giftIdNumArr {
		if !vo.BHasEnought(giftIdNum[0], -giftIdNum[1]) {
			return false
		}
	}
	return true
}

func (vo *PlayerInfoVO) SubPlayerGift(giftIdNumArr [][]int32) {
	for _, giftIdNum := range giftIdNumArr {
		vo.AddRes(giftIdNum[0], -giftIdNum[1])
	}
}

func (vo *PlayerInfoVO) AddPlayerGift(giftIdNumArr [][]int32) {
	for _, giftIdNum := range giftIdNumArr {
		vo.AddRes(giftIdNum[0], giftIdNum[1])
	}
}
