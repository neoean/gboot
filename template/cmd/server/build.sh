GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o server -v

scp -P 22 server root@47.122.4.251:/www/wwwroot/step2money/cmd/server/
ssh root@47.122.4.251 "cd /www/wwwroot/step2money/cmd/server/; sh stop.sh"