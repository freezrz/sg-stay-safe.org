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
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/lambda"
	"sg-stay-safe.org/protocol"
	"time"
)

func main() {
	log.Println("checkin in router service...")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.POST("/", route)
	router.GET("/welcome", welcome)

	router.Run(fmt.Sprintf(":%s", config.AWSForwardPort))
}

func route(c *gin.Context) {
	var event protocol.CheckInEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		clientError(c, config.CodeMarshalError, err)
		return
	}

	code, err := checkin(event)
	if err != nil {
		log.Println(err.Error())
		serverError(c, code, err)
		return
	}
	success(c)
}

func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Welcome": "Safe Check-in Platform",
	})
}

func checkin(event protocol.CheckInEvent) (int, error) {
	// use server timestamp
	event.Timestamp = time.Now().Unix()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambdaService.New(sess, &aws.Config{Region: aws.String(config.AWSRegion)})

	payload, err := json.Marshal(event)
	if err != nil {
		log.Println(err.Error())
		return config.CodeMarshalError, err
	}

	lambdaChain := []string{
		config.SanitiseCheckinLambda,
		config.AntiFraudCheckinLambda,
		config.VerifyRulesLambda,
		config.RecordCheckinLambda,
		config.ProduceCheckinMsgLambda,
	}
	for _, cmd := range lambdaChain {
		if resp, _, err := lambdaPkg.CallLambdaFunc(client, cmd, payload); err != nil {
			sendViolationMsg(client, event, resp) // if err, no need to return to customer
			return resp.Code, err
		}
	}

	return config.CodeOK, nil
}

func sendViolationMsg(client *lambdaService.Lambda, checkin protocol.CheckInEvent, resp protocol.GeneralResponse) {
	msg := protocol.ViolationEvent{
		CheckInEvent:    checkin,
		GeneralResponse: resp,
	}
	var cmd string
	switch resp.Code {
	case config.CodeSiteIsBannedError:
		cmd = config.ProduceSiteViolationMsgLambda
	case config.CodeUserIsBannedError, config.CodeUserExceedDailyMaxCheckinError:
		cmd = config.ProduceUserViolationMsgLambda
	default:
		// cmd = "not-implemented"
		return
	}
	payload, _ := json.Marshal(msg)
	_, _, _ = lambdaPkg.CallLambdaFunc(client, cmd, payload)
}

func serverError(c *gin.Context, errorCode int, err error) {
	resp := protocol.CheckInResponse{
		Code: errorCode,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func clientError(c *gin.Context, errorCode int, err error) {
	resp := protocol.CheckInResponse{
		Code: errorCode,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusBadRequest, resp)
}

func success(c *gin.Context) {
	resp := protocol.CheckInResponse{
		Code: config.CodeOK,
		Msg:  "check in successfully",
	}
	c.JSON(http.StatusOK, resp)
}
