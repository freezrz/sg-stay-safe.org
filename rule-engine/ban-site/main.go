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

func Handler(ctx context.Context, event protocol.BanSiteRequest) (protocol.BanSiteResponse, error) {
	fmt.Println("ban-site invoked")
	if event.SiteId == "" {
		return protocol.BanSiteResponse{Code: http.StatusInternalServerError, Msg: "internal error"}, errors.New("internal error")
	}
	redisCli := pkg.New(config.BanSiteCache)
	if event.ShouldBan {
		// 1 means banned
		err := redisCli.Set(fmt.Sprintf("%s%s", config.BanSitePrefix, event.SiteId), "1", 0)
		if err != nil {
			log.Println(err.Error())
		}
		return protocol.BanSiteResponse{Code: http.StatusOK, Msg: fmt.Sprintf("%s banned", event.SiteId)}, nil
	} else {
		err := redisCli.Del(fmt.Sprintf("%s%s", config.BanSitePrefix, event.SiteId))
		if err != nil {
			log.Println(err.Error())
		}
		return protocol.BanSiteResponse{Code: http.StatusOK, Msg: fmt.Sprintf("%s unbanned", event.SiteId)}, nil
	}
}
