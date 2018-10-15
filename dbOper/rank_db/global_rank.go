package rank_db

import (
	"buffalo/king/common/gameutil"
)

const (
	lc_GLOBAL_ZSET_KEY = "zset:sg:globalrank"
)

func updateGlobalScore(playerId int32, score int32) {
	s := float64(score)
	gameutil.ZAdd(lc_GLOBAL_ZSET_KEY, s, playerId)
}

func init() {
	db := gameutil.GetDB()
	rows, err := db.Raw("select player_id,score from rank order by score desc limit 50").Rows()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	for rows.Next() {
		var pid, score int
		if err := rows.Scan(&pid, &score); err != nil {
			panic(err)
		}
		s := float64(score)
		gameutil.ZAdd(lc_GLOBAL_ZSET_KEY, s, pid)
	}
}
