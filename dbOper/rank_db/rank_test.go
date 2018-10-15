package rank_db

import (
	"testing"
)

func Test_Rank(t *testing.T) {
	//rvo, err := GetRankVO(1)
	_, err := GetRankVO(1)
	if err != nil {
		t.Error(err)
	}

	//rvo.Score = 101
	//_ = rvo.Save(true)
}
