module github.com/goees/yago

go 1.12

require (
	github.com/Shopify/sarama v1.19.0
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/bwmarrin/snowflake v0.3.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/garyburd/redigo v1.6.0
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-contrib/pprof v1.2.1
	github.com/gin-gonic/gin v1.4.0
	github.com/goees/go-mysql v1.5.1
	github.com/go-xorm/xorm v0.7.6
	github.com/golang/protobuf v1.3.1
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/levigross/grequests v0.0.0-20190130132859-37c80f76a0da
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.4.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.1
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65
	google.golang.org/grpc v1.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	xorm.io/core v0.7.0
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
