package common

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	//"github.com/kingbuffalo/seelog"
	"io/ioutil"
	"strconv"
	"strings"
)

var idMapName map[int]string
var idMapShortName map[int]string
var nameMapId map[string]int

type pjkv_t struct {
	Key   string
	Value string
}

func init() {
	protoFP := startcfg.GetProtoFilePath()
	data, err := ioutil.ReadFile(protoFP)
	if err != nil {
		panic(err)
	}
	var pjkvs []pjkv_t
	err = json.Unmarshal(data, &pjkvs)
	if err != nil {
		panic(err)
	}
	idLen := len(pjkvs) / 2
	idMapName = make(map[int]string, idLen)
	idMapShortName = make(map[int]string, idLen)
	nameMapId = make(map[string]int, idLen)
	for _, v := range pjkvs {
		keyInt, err := strconv.Atoi(v.Key)
		if err == nil {
			idMapName[keyInt] = v.Value
			idMapShortName[keyInt] = strings.TrimPrefix(v.Value, "king.S2C_")
		} else {
			valueInt, _ := strconv.Atoi(v.Value)
			nameMapId[v.Key] = valueInt
		}
	}
}

func GetProtoShortName(protoId int) string {
	return idMapShortName[protoId]
}

func GetProtoName(protoId int) string {
	return idMapName[protoId]
}

func GetProtoId(protoName string) int {
	s := "king." + protoName
	return nameMapId[s]
}

func GetHLFromWord(intValue int) (byte, byte) {
	h := (intValue >> 8) & 0xff
	l := intValue & 0xff
	return byte(h), byte(l)
}

func GetBytesFromInt(intValue int) (byte, byte, byte, byte) {
	h1 := (intValue >> 24) & 0xff
	h2 := (intValue >> 16) & 0xff
	h := (intValue >> 8) & 0xff
	l := intValue & 0xff
	return byte(h1), byte(h2), byte(h), byte(l)
}
