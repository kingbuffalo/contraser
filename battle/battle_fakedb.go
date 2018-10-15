package battle

import (
	//"buffalo/king/common/gameutil"
	//"encoding/json"
	//"errors"
	//"fmt"
	"github.com/kingbuffalo/seelog"
	"sync"
)

func GetBattle(playerId int32) *Battle {
	seelog.Trace("playerId:", playerId)
	battleMutex.RLock()
	defer battleMutex.RUnlock()
	ret, exist := pidMapBattle[playerId]
	if exist {
		return ret
	}
	return nil
}

func (v *Battle) Save() {
	battleMutex.Lock()
	defer battleMutex.Unlock()
	pidMapBattle[v.playerId] = v
}

var pidMapBattle map[int32]*Battle
var battleMutex = &sync.RWMutex{}

func init() {
	pidMapBattle = make(map[int32]*Battle, 1024)
}
