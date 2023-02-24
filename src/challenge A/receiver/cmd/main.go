package main

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/receiver/config"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/adapter/broker/redis"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/adapter/store/mysql"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/interactor/order"
	"log"
	"os"
)

const (
	redisDSN     = "REDIS_DSN"
	mySqlDSN     = "MYSQL_DSN"
	portErrorMsg = "port should be pass as argument"
)

func main() {
	appSetting := GetAppConfig()
	redisQ := redis.InitRedis(appSetting.BrokerDSN)
	mysqlStore := mysql.InitMySql(appSetting.MySQLDSN)
	orderInteractor := order.New(*redisQ, mysqlStore)

	log.Println("start consuming orders")
	ctx := context.Background()
	go orderInteractor.CreateOrder(ctx)

	forever := make(chan bool)
	go func() {
		for {

		}
	}()
	<-forever // Block main from exiting
}

func GetAppConfig() config.AppConfig {
	if len(os.Args) < 2 {
		log.Fatalf(portErrorMsg)
	}
	return config.AppConfig{
		InProduction: false,
		BrokerDSN:    os.Getenv(redisDSN),
		MySQLDSN:     os.Getenv(mySqlDSN),
		PortServer:   os.Args[1],
	}
}
