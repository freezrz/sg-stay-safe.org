package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"sg-stay-safe.org/rule-engine/config"
	"sg-stay-safe.org/rule-engine/pkg"
	"sg-stay-safe.org/rule-engine/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.VerifyRulesForCheckInResponse, error) {
	fmt.Println("verify-rules invoked")
	if event.SiteId == "" {
		return protocol.VerifyRulesForCheckInResponse{Code: http.StatusInternalServerError, Msg: "internal error"}, errors.New("internal error")
	}
	redisCli := pkg.New(config.BanSiteCache)

	isSiteBan, err := redisCli.Get(fmt.Sprintf("%s%s", config.BanSitePrefix, event.SiteId))
	if err != nil {
		log.Println(err.Error())
	}
	if isSiteBan == "1" {
		return protocol.VerifyRulesForCheckInResponse{Code: config.CodeSiteBanned, Msg: fmt.Sprintf("site %s banned", event.SiteId)}, nil
	}

	isUserBan, err := redisCli.Get(fmt.Sprintf("%s%s", config.BanUserPrefix, event.SiteId))
	if err != nil {
		log.Println(err.Error())
	}
	if isUserBan == "1" {
		return protocol.VerifyRulesForCheckInResponse{Code: config.CodeUserBanned, Msg: fmt.Sprintf("user %s banned", event.SiteId)}, nil
	}

	return protocol.VerifyRulesForCheckInResponse{Code: config.CodeOK, Msg: "ok"}, nil
}
