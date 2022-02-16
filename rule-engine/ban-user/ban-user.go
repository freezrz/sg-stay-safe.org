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
	lambda.Start(handler)
}

func handler(ctx context.Context, event protocol.BanUserRequest) (protocol.GeneralResponse, error) {
	fmt.Println("ban-user invoked")

	redisCli := cache.New(config.BanSiteCache)
	if event.ShouldBan {
		// 1 means banned
		if err := redisCli.Set(fmt.Sprintf(config.BanUserFormat, event.AnonymousId), "1", 0); err != nil {
			return protocol.GeneralResponse{Code: config.CodeBanUserUpdateCacheError, Msg: fmt.Sprintf("%s banned", event.AnonymousId)}, err
		}
		return protocol.GeneralResponse{Code: config.CodeOK, Msg: fmt.Sprintf("%s banned", event.AnonymousId)}, nil
	} else {
		if err := redisCli.Del(fmt.Sprintf(config.BanUserFormat, event.AnonymousId)); err != nil {
			return protocol.GeneralResponse{Code: config.CodeUnBanUserUpdateCacheError, Msg: fmt.Sprintf("%s unbanned", event.AnonymousId)}, err
		}
		return protocol.GeneralResponse{Code: config.CodeOK, Msg: fmt.Sprintf("%s unbanned", event.AnonymousId)}, nil
	}
}
