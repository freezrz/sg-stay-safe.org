package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sg-stay-safe.org/checkin-service/protocol"
)

func main() {
	log.Println("checkin-service service...")
	router := gin.Default()
	router.GET("/checkin-service", checkIn)

	router.Run(":5000")
}

// getRestaurantById locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func checkIn(c *gin.Context) {
	var event protocol.CheckInEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Sanitise(event)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "okay"})
}

func Sanitise(event protocol.CheckInEvent) error {
	// Create Lambda service client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("ap-southeast-1")})

	// Get the 10 most recent items
	request := protocol.CheckInEvent{AnonymousId: event.AnonymousId, PlaceId: event.PlaceId}

	payload, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling MyGetItemsFunction request")
		return errors.New("invalid checkin-service")
	}

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("sanitise_checkin"), Payload: payload})
	if err != nil {
		fmt.Println("Error calling MyGetItemsFunction")
		os.Exit(0)
	}
	log.Println(result.GoString())

	var resp protocol.SanitiserResponse

	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		fmt.Println("Error unmarshalling SanitiserResponse")
		return errors.New("invalid checkin-service")
	}
	log.Println(resp.Msg)
	return nil
}
