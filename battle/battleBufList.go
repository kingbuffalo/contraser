package battle

/*
import (
	//"buffalo/kingv2Game/allCmdServerNeed/tblMgr/tblBattleBufferMgr"
	//"buffalo/utils/utilsFunc"
	"errors"
)

type batBuf_s struct {
	battleUnitPos     int
	bufferId          int
	tblBattleBufferVO *tblBattleBufferMgr.TblBattleBufferVO
	bufferValue       int
	endStep           int
}

type batBuf_node struct {
	v    *batBuf_s
	next *batBuf_node
}

type batBufList_s struct {
	head *batBuf_node
	tail *batBuf_node
	iter *batBuf_node
}

func createbatBuf_s(battleUnitPos, bufferId, bufferValue, endStep int) (*batBuf_s, error) {
	tblBattleBufferVO, msg := tblBattleBufferMgr.GetTblBattleBufferVO(bufferId)
	if msg != nil {
		utilsFunc.DevPanic(msg.ToString())
		return nil, errors.New(msg.ToString())
	}
	return &batBuf_s{
		battleUnitPos:     battleUnitPos,
		bufferId:          bufferId,
		bufferValue:       bufferValue,
		endStep:           endStep,
		tblBattleBufferVO: tblBattleBufferVO,
	}, nil
}

func (bl *batBufList_s) addUnitOne(battleUnitPos, bufferId, bufferValue, endStep int) error {
	if bl.head == nil {
		batBuf, err := createbatBuf_s(battleUnitPos, bufferId, bufferValue, endStep)
		if err != nil {
			return err
		}
		node := &batBuf_node{
			v: batBuf,
		}
		bl.head = node
		bl.tail = node
	} else {
		iter := bl.head
		bFound := false
		for iter != nil {
			if iter.v.battleUnitPos == battleUnitPos && iter.v.bufferId == bufferId {
				iter.v.bufferValue = bufferValue
				iter.v.endStep = endStep
				bFound = true
				break
			}
			iter = iter.next
		}
		if !bFound {
			batBuf, err := createbatBuf_s(battleUnitPos, bufferId, bufferValue, endStep)
			if err != nil {
				return err
			}
			node := &batBuf_node{
				v: batBuf,
			}
			bl.tail.next = node
			bl.tail = node
		}
	}
	return nil
}

func (bl *batBufList_s) clearTimeEndBuf(curStep int) {
	if bl.head != nil {
		iter := bl.head
		for iter != nil && iter.v.endStep <= curStep {
			iter = iter.next
		}
		bl.head = iter
		if iter != nil {
			for iter.next != nil {
				if iter.next.v.endStep <= curStep {
					iter.next = iter.next.next
				}
				iter = iter.next
			}
			bl.tail = iter
		} else {
			bl.tail = nil
		}
	}
}

func (bl *batBufList_s) iterMoveToFist() {
	bl.iter = bl.head
}
func (bl *batBufList_s) iterBatBuf() *batBuf_s {
	if bl.iter != nil {
		return bl.iter.v
	}
	return nil
}
func (bl *batBufList_s) iterNext() {
	bl.iter = bl.iter.next
}*/
