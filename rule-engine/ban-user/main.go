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

func Handler(ctx context.Context, event protocol.BanUserRequest) (protocol.BanUserResponse, error) {
	fmt.Println("ban-user invoked")
	if event.AnonymousId == "" {
		return protocol.BanUserResponse{Code: http.StatusInternalServerError, Msg: "internal error"}, errors.New("internal error")
	}
	redisCli := pkg.New(config.BanUserCache)
	if event.ShouldBan {
		// 1 means banned
		err := redisCli.Set(fmt.Sprintf("%s%s", config.BanUserPrefix, event.AnonymousId), "1", 0)
		if err != nil {
			log.Println(err.Error())
		}
		return protocol.BanUserResponse{Code: http.StatusOK, Msg: fmt.Sprintf("%s banned", event.AnonymousId)}, nil
	} else {
		err := redisCli.Del(fmt.Sprintf("%s%s", config.BanUserPrefix, event.AnonymousId))
		if err != nil {
			log.Println(err.Error())
		}
		return protocol.BanUserResponse{Code: http.StatusOK, Msg: fmt.Sprintf("%s unbanned", event.AnonymousId)}, nil
	}
}
