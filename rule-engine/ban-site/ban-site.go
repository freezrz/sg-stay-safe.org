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

func handler(ctx context.Context, event protocol.BanSiteRequest) (protocol.GeneralResponse, error) {
	fmt.Println("ban-site invoked")

	redisCli := cache.New(config.BanCache)
	if event.ShouldBan {
		// 1 means banned
		if err := redisCli.Set(fmt.Sprintf(config.BanSiteFormat, event.SiteId), "1", 0); err != nil {
			return protocol.GeneralResponse{Code: config.CodeBanSiteUpdateCacheError, Msg: fmt.Sprintf("%s banned", event.SiteId)}, nil
		}
		return protocol.GeneralResponse{Code: config.CodeOK, Msg: fmt.Sprintf("%s banned", event.SiteId)}, nil
	} else {
		if err := redisCli.Del(fmt.Sprintf(config.BanSiteFormat, event.SiteId)); err != nil {
			return protocol.GeneralResponse{Code: config.CodeUnBanSiteUpdateCacheError, Msg: fmt.Sprintf("%s unbanned", event.SiteId)}, nil
		}
		return protocol.GeneralResponse{Code: config.CodeOK, Msg: fmt.Sprintf("%s unbanned", event.SiteId)}, nil
	}
}
