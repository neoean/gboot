package types

import (
	"encoding/json"
	"step2money-server/common/logs"
	"strconv"
)

func ObjToJson(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func UInt64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		logs.Debug("StringToInt Err: %v", err)
	}

	return num
}

func StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logs.Debug("StringToInt Err: %v", err)
	}

	return num
}

func Booltonumber(b bool) int {
	result := 0
	if b {
		result = 1
	}
	return result
}
