package weApp

import (
	"crypto/tls"
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/logger"
	"log"
	"net/http"
	"step2money-server/common/config"
	"step2money-server/common/logs"
	"time"
)

func InitWeApp() {
	weApp := config.Config.WeApp

	httpCli := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			// 跳过安全校验
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	lgr := logger.NewLogger(log.New(logs.GetWriter("/log/step2money_wx.log"), "\r\n", log.LstdFlags), logger.Info, true)

	config.WeAppClient = weapp.NewClient(
		weApp.AppId, weApp.Secret,
		weapp.WithHttpClient(httpCli),
		weapp.WithLogger(lgr),
		weapp.WithCache(new(WeCache)),
	)
}
