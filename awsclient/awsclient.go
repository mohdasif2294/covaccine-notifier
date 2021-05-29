package awsclient

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// awsClientCfg ...
type awsClientCfg struct {
	Region string `toml:"region"`
}

var (
	awsCfg    awsClientCfg
	svcClient *sns.SNS
)

//Initialize ...
func Initialize() {
	awsCfg = awsClientCfg{Region: "ca-central-1"}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsCfg.Region)},
	)
	if err != nil {
		log.Fatal("Error in Initializing the aws client ", err.Error())
	}

	// Create SNS service
	svcClient = sns.New(sess)
}

//SendSMS ...
func SendSMS(phoneNumber string, message string) error {

	// Pass the phone number and message.
	params := &sns.PublishInput{
		PhoneNumber: aws.String(phoneNumber),
		Message:     aws.String(message),
	}

	// sends a text message (SMS message) directly to a phone number.
	resp, err := svcClient.Publish(params)

	if err != nil {
		log.Printf("Error in Sending SMS: %s", err.Error())
		return err
	}

	log.Println("SMS Sent to address: ", phoneNumber)
	log.Println("Sms Message ID: ", resp.GoString())

	return nil
}
