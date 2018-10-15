package gameutil

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

const RAND_CHAR_ARR_LEN = 62

var charArr [RAND_CHAR_ARR_LEN]byte

func init() {
	idx := 0
	for i := 0; i < 10; i++ {
		charArr[idx] = byte(48 + i)
		idx++
	}
	for i := 0; i < 26; i++ {
		charArr[idx] = byte(65 + i)
		idx++
	}
	for i := 0; i < 26; i++ {
		charArr[idx] = byte(97 + i)
		idx++
	}
}

func RandIdx(maxIdx, needIdx int) []int {
	randIdxs := make([]int, maxIdx)
	for i, _ := range randIdxs {
		randIdxs[i] = i
	}
	for i := 0; i < maxIdx; i++ {
		idx := rand.Intn(maxIdx)
		randIdxs[i], randIdxs[idx] = randIdxs[idx], randIdxs[i]
	}
	return randIdxs[:needIdx]
}

var md5tokenfmt = "buffalokingbuffaloxxx_%d"

func GenToken(playerId int32) string {
	str := fmt.Sprintf(md5tokenfmt, playerId)
	md5sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", md5sum)
}

func GenRandStr(strLen int) string {
	ret := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		rIdx := rand.Intn(RAND_CHAR_ARR_LEN)
		ret[i] = charArr[rIdx]
	}
	s := string(ret)
	return s
}

func BToday(timestamp int32) bool {
	ct := time.Now()
	cmpT := time.Unix(int64(timestamp), 0)
	cy, cm, cd := ct.Date()
	cmpy, cmpm, cmpd := cmpT.Date()
	return cy == cmpy && cm == cmpm && cd == cmpd
}

//type PanicHandler struct {
//runF func() interface{}
//Ret  interface{}
//}

//func (p *PanicHandler) RunFunc() {
//defer func() {
//if err := recover(); err != nil {
//_ = seelog.Error(err)
//buf := make([]byte, 1<<11)
//runtime.Stack(buf, true)
//_ = seelog.Error(string(buf))
//p.Ret = common.NewErrorCodeRet(113)
//}
//}()
//p.Ret = p.runF()
//}
