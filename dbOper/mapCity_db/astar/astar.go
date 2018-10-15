package astar

import (
	goAstar "github.com/kingbuffalo/go-astar"
	"math"
	"buffalo/king/cfgmgr/cfgMapCityMgr"
	"buffalo/king/common/enumErrCode"
)

var edw goAstar.EDWorld

func init() {
	mapDataInTileMapCfg := cfgMapCityMgr.GetData()
	w := cfgMapCityMgr.GetWidth()
	h := cfgMapCityMgr.GetHeight()
	idx := 0
	edw = goAstar.NewEDWorld()
	var i, j int32
	for j = 0; j < h; j++ {
		for i = 0; i < w; i++ {
			v := mapDataInTileMapCfg[idx]
			idx++
			edw.SetTile(v != 0, int(i), int(j))
		}
	}
}

func Distance(fromx, fromy, tox, toy int32) (float64, bool) {
	e := cfgMapCityMgr.GetNodeEdge()
	d, bFound := edw.Distance(int(fromx), int(fromy), int(tox), int(toy))
	if bFound {
		d = d * float64(e)
	}
	return d, bFound
}

func NeedTime(distance float64) int32 {
	return int32(math.Ceil(distance / float64(enumErrCode.ARMY_SPEED_PER_SECOND)))
}
