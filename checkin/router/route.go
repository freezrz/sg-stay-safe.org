package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sg-stay-safe.org/checkin/config"
	"sg-stay-safe.org/checkin/protocol"
	"time"
)

func main() {
	log.Println("checkin in router service...")
	router := gin.Default()
	router.POST("/", checkIn)

	router.Run(fmt.Sprintf(":%s", config.AWSForwardPort))
}

func checkIn(c *gin.Context) {
	var event protocol.CheckInEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		clientError(c, err)
		return
	}
	// use server timestamp
	event.Timestamp = time.Now().Unix()

	err := record(event)
	if err != nil {
		log.Println(err.Error())
		serverError(c, err)
		return
	}
	success(c)
}

func record(event protocol.CheckInEvent) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambdaService.New(sess, &aws.Config{Region: aws.String(config.AWSRegion)})

	request := protocol.CheckInEvent{AnonymousId: event.AnonymousId, SiteId: event.SiteId}

	payload, err := json.Marshal(request)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	result, err := client.Invoke(&lambdaService.InvokeInput{FunctionName: aws.String(config.RecordCheckinLambda), Payload: payload})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// debug purpose
	log.Println(result.GoString())

	var resp protocol.SanitiserResponse

	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println(resp.Msg)
	return nil
}

func serverError(c *gin.Context, err error) {
	resp := protocol.CheckInResponse{
		Code: 50001,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func clientError(c *gin.Context, err error) {
	resp := protocol.CheckInResponse{
		Code: 40001,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusBadRequest, resp)
}

func success(c *gin.Context) {
	resp := protocol.CheckInResponse{
		Code: 0,
		Msg:  "check in successfully",
	}
	c.JSON(http.StatusOK, resp)
}
