package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"sg-stay-safe.org/checkin-sanitiser/pkg"
	"sg-stay-safe.org/checkin-sanitiser/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.SanitiserResponse, error) {
	fmt.Println("checkin-service event sanitise handler invoked")
	if event.AnonymousId == "" || event.PlaceId == "" {
		return protocol.SanitiserResponse{Code: http.StatusNotAcceptable, Msg: "not accepted"}, errors.New("not accepted")
	}
	redisCli := pkg.New()
	click, err := redisCli.Get(event.PlaceId)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(click)
	return protocol.SanitiserResponse{Code: http.StatusOK, Msg: "sanitise ok..."}, nil
}
