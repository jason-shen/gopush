package apn

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
)

func Apn(cert string, environment string, devicetoken string, topic string, mode, string, priority int, pushType apns2.EPushType, message []byte) (*apns2.Response, error) {
	var client *apns2.Client
	certpath, err := certificate.FromP12File(cert, "")
	if err != nil {
		log.Fatal("Cert Error:", err)
		return nil, err
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = devicetoken
	notification.Topic = topic
	notification.Priority = priority
	notification.PushType = pushType
	notification.Payload = message

	if environment == "production" {
		client = apns2.NewClient(certpath).Production()
	}

	if environment == "development" {
		client = apns2.NewClient(certpath).Production()
	}

	res, err := client.Push(notification)
	if err != nil {
		log.Fatal("Error:", err)
		return nil, err
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)

	return res, nil
}