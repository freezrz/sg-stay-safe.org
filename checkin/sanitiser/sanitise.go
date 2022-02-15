package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"sg-stay-safe.org/checkin/pkg/cache"
	"sg-stay-safe.org/checkin/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	log.Println("sanitiser invoked")
	if event.AnonymousId == "" || event.SiteId == "" {
		return protocol.GeneralResponse{Code: http.StatusNotAcceptable, Msg: "not accepted"}, errors.New("not accepted")
	}
	redisCli := cache.New()
	click, err := redisCli.Get(event.SiteId)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(click)
	return protocol.GeneralResponse{Code: http.StatusOK, Msg: "sanitise ok..."}, nil
}
