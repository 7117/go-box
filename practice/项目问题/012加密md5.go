package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	appid := "aaa"
	psa := "/ums/operation_monitoring_system/metadata-manager"
	ts := "1647331721"
	Secret := "bbb"
	str := fmt.Sprintf("%s%s%s%s", appid, psa, ts, Secret)
	h := md5.Sum([]byte(str))
	fmt.Println(hex.EncodeToString(h[:]))
}
