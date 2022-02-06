package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"sg-stay-safe.org/checkin-cache/pkg"
)

func main() {
	lambda.Start(Handler)
}

type CheckInEvent struct {
	AnonymousId string `json:"anonymous_id"`
	PlaceId     string `json:"place_id"`
}

func Handler(ctx context.Context, event CheckInEvent) (string, error) {
	fmt.Println("checkin-service event sanitise handler invoked")
	if event.AnonymousId == "" || event.PlaceId == "" {
		return "unable to process your request", errors.New("invalid request")
	}

	cli := pkg.New()
	value, err := cli.Get(event.PlaceId)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println(value)
	return value, nil
}
