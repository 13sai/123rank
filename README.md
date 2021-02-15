# 123rank

有趣的排行榜，不限于自然、电影、小说、财经。

```
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl

goctl api new api

goctl model mysql datasource -url="root:111111@tcp(127.0.0.1:3306)/rank" -table="items" -dir ./model

go run rank.go -f etc/rank-api.yaml
```
