package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/cache"
	"sg-stay-safe.org/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	log.Println("anti-fraud checkin invoked")

	redisCli := cache.New(config.AntiFraudCache)
	visited, err := redisCli.Get(fmt.Sprintf(config.UserVisitShopHistoryFormat, event.AnonymousId, event.SiteId))
	if visited != "" {
		return protocol.GeneralResponse{Code: config.CodeAntiFraudEventError, Msg: fmt.Sprintf("you have checked in this site in %d min(s)...", config.UserVisitSiteIntervalTimeDuration)}, err
	}
	if err != nil {
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: config.CodeAntiFraudEventError, Msg: "anti fraud checkin fail..."}, err
	}

	return protocol.GeneralResponse{Code: config.CodeOK, Msg: "anti fraud checkin ok..."}, nil
}
