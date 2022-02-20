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

	checkinCli := cache.New(config.CheckInSiteCache)
	t := convert.CleanTime(config.Region, time.Now().Unix(), time.Minute*config.CacheDuration)
	if err := checkinCli.Incr(fmt.Sprintf(config.SiteCountFormat, t, event.SiteId), config.CacheDuration); err != nil {
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: config.CodeRecordCacheEventError, Msg: "record checkin fail..."}, err

	}

	antiFraudCli := cache.New(config.AntiFraudCache)
	t = convert.CleanTime(config.Region, time.Now().Unix(), time.Minute*config.UserVisitSiteIntervalTimeDuration)
	// if 1, means the user visits this site in 5mins
	if err := antiFraudCli.Set(fmt.Sprintf(config.UserVisitSiteHistoryFormat, event.AnonymousId, event.SiteId), "1", config.UserVisitSiteIntervalTimeDuration); err != nil {
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: config.CodeRecordUserVisitSiteEventError, Msg: "record checkin fail..."}, err
	}

	banCli := cache.New(config.BanCache)
	if err := banCli.Incr(fmt.Sprintf(config.User24HoursCheckinCountFormat, event.AnonymousId), 24*60); err != nil { // 24hours
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: config.CodeIncrUser24hrsCheckinError, Msg: "increase user 24hrs checkin number fail..."}, err
	}

	return protocol.GeneralResponse{Code: config.CodeOK, Msg: "record checkin ok..."}, nil
}
