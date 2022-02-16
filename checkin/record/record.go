package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/cache"
	convert "sg-stay-safe.org/pkg/time"
	"sg-stay-safe.org/protocol"
	"time"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	log.Println("record checkin invoked")

	redisCli := cache.New(config.CheckInSiteCache)
	t := convert.CleanTime(config.Region, time.Now().Unix(), time.Minute*config.CacheDuration)
	err := redisCli.Incr(fmt.Sprintf(config.SiteCountFormat, t, event.SiteId))
	if err != nil {
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: config.CodeRecordCacheEventError, Msg: "record checkin fail..."}, err
	}

	return protocol.GeneralResponse{Code: config.CodeOK, Msg: "record checkin ok..."}, nil
}
