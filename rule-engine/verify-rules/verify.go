package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/cache"
	"sg-stay-safe.org/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	fmt.Println("verify-rules invoked")

	redisCli := cache.New(config.BanCache)

	isSiteBan, err := redisCli.Get(fmt.Sprintf(config.BanSiteFormat, event.SiteId))
	if err != nil {
		return protocol.GeneralResponse{Code: config.CodeSiteBannedCacheError, Msg: fmt.Sprintf("site %s banned", event.SiteId)}, nil
	}
	if isSiteBan == "1" {
		return protocol.GeneralResponse{Code: config.CodeSiteIsBannedError, Msg: fmt.Sprintf("site %s banned", event.SiteId)}, nil
	}

	isUserBan, err := redisCli.Get(fmt.Sprintf(config.BanUserFormat, event.AnonymousId))
	if err != nil {
		return protocol.GeneralResponse{Code: config.CodeUserBannedCacheError, Msg: fmt.Sprintf("user %s banned", event.AnonymousId)}, nil
	}
	if isUserBan == "1" {
		return protocol.GeneralResponse{Code: config.CodeUserIsBannedError, Msg: fmt.Sprintf("user %s banned", event.AnonymousId)}, nil
	}

	return protocol.GeneralResponse{Code: config.CodeOK, Msg: "verify rules ok..."}, nil
}
