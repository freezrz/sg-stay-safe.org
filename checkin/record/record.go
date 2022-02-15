package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"sg-stay-safe.org/checkin/config"
	"sg-stay-safe.org/checkin/pkg/cache"
	convert "sg-stay-safe.org/checkin/pkg/time"
	"sg-stay-safe.org/checkin/protocol"
	"time"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	log.Println("record checkin invoked")
	if event.AnonymousId == "" || event.SiteId == "" {
		return protocol.GeneralResponse{Code: http.StatusNotAcceptable, Msg: "not accepted"}, errors.New("invalid request")
	}

	redisCli := cache.New(config.CheckInSiteCache)
	t := convert.CleanTime(config.Region, time.Now().Unix(), time.Minute*config.CacheDuration)
	err := redisCli.Incr(fmt.Sprintf(config.SiteCountFormat, t, event.SiteId))
	if err != nil {
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: http.StatusInternalServerError, Msg: "record checkin fail..."}, nil
	}

	return protocol.GeneralResponse{Code: http.StatusOK, Msg: "record checkin ok..."}, nil
}
