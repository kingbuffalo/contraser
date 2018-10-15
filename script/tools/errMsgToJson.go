package main

import (
	"buffalo/king/common/enumErrCode"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

func main() {
	var errCode = enumErrCode.ErrCode
	var strMapStr map[string]string = make(map[string]string, 0)
	for k, v := range errCode {
		strK := strconv.Itoa(k)
		strMapStr[strK] = v
	}
	b, err := json.Marshal(strMapStr)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("/tmp/errCode.json", b, 0644)

	if err != nil {
		panic(err)
	}
}
