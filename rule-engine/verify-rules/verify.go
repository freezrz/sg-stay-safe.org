package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/cache"
	"sg-stay-safe.org/protocol"
	"strconv"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	fmt.Println("verify-rules invoked")

	// check ban
	banCli := cache.New(os.Getenv("BanCache"))

	isSiteBan, err := banCli.Get(fmt.Sprintf(config.BanSiteFormat, event.SiteId))
	if err != nil {
		return protocol.GeneralResponse{Code: config.CodeSiteBannedCacheError, Msg: fmt.Sprintf("site %s banned", event.SiteId)}, nil
	}
	if isSiteBan == "1" {
		return protocol.GeneralResponse{Code: config.CodeSiteIsBannedError, Msg: fmt.Sprintf("site %s banned", event.SiteId)}, nil
	}

	isUserBan, err := banCli.Get(fmt.Sprintf(config.BanUserFormat, event.AnonymousId))
	if err != nil {
		return protocol.GeneralResponse{Code: config.CodeUserBannedCacheError, Msg: fmt.Sprintf("user %s banned", event.AnonymousId)}, nil
	}
	if isUserBan == "1" {
		return protocol.GeneralResponse{Code: config.CodeUserIsBannedError, Msg: fmt.Sprintf("user %s banned", event.AnonymousId)}, nil
	}

	// check rule
	rule := protocol.Rule{}
	ruleCli := cache.New(os.Getenv("RuleCache"))

	maxDailyCacheValue, err := ruleCli.Get(config.RuleMaxDailyCheckin)
	if err != nil {
		log.Println(err.Error())
	}
	if err := json.Unmarshal([]byte(maxDailyCacheValue), &rule); err != nil {
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: config.CodeVerifyUserMaxCheckinCacheError, Msg: fmt.Sprintf("verify user max checkin json umarshal error for %s", event.AnonymousId)}, nil
	}
	if rule.IsEnabled {
		userCheckInCount, _ := banCli.Get(fmt.Sprintf(config.User24HoursCheckinCountFormat, event.AnonymousId))
		if userCheckInCount == "" {
			userCheckInCount = "0"
		}
		count, _ := strconv.Atoi(userCheckInCount)
		if count >= rule.Value {
			return protocol.GeneralResponse{Code: config.CodeUserExceedDailyMaxCheckinError, Msg: fmt.Sprintf("%s has checkin %d times in 24hours. Stay home.", event.AnonymousId, rule.Value)}, nil
		}
	}

	return protocol.GeneralResponse{Code: config.CodeOK, Msg: "verify ban and rules ok..."}, nil
}
