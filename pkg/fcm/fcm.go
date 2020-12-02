package fcm

import (
	"github.com/appleboy/go-fcm"
	"log"
)

func Fcm(apikey string, devicetoken string, data map[string]interface{}) (*fcm.Response, error) {
	msg := &fcm.Message{
		To: devicetoken,
		Data: data,
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient(apikey)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Printf("%#v\n", response)
	return response, nil
}