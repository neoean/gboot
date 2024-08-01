package qiniu

import (
	"testing"
	"{{.Package}}/common/config"
)

func TestUpload(t *testing.T) {
	config.InitConfig()

	Init()

	Upload([]byte("hello, this is qiniu cloud"), "abc")
}
