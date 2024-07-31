package qiniu

import (
	"step2money-server/common/config"
	"testing"
)

func TestUpload(t *testing.T) {
	config.InitConfig()

	Init()

	Upload([]byte("hello, this is qiniu cloud"), "abc")
}
