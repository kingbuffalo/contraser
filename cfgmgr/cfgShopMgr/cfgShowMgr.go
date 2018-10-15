package cfgShopMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgShop struct {
	ShopId          int32 `json:"shop_id"`
	LocIdCount      int32 `json:"loc_id_count"`
	BuyType         int32 `json:"buy_type"`
	Init            int32 `json:"init"`
	AddGold         int32 `json:"add_gold"`
	Discount        int32 `json:"discount"`
	PriceChangeType int32 `json:"price_change_type"`
	MaxTimes        int32 `json:"max_times"`
	MaxGold         int32 `json:"max_gold"`
}

var shopIdMapCfg map[int32]*CfgShop

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_shop.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgShop = make([]*CfgShop, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	shopIdMapCfg := make(map[int32]*CfgShop, len(cfgArr))
	for _, v := range cfgArr {
		shopIdMapCfg[v.ShopId] = v
	}
}
