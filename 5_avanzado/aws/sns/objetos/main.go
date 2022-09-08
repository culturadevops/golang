package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var SNS *SNSClient

func init() {
	SNS = new(SNSClient)
}

type SNSClient struct {
	Region string
	Sess   *session.Session
	Svc    *sns.SNS
	Topic  string
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
func (t *SNSClient) NewSession(region string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		exitErrorf("PROBLEMA DE SESSION CON SNS, %v", err)
	}
	t.Sess = sess

	t.Svc = sns.New(t.Sess)
}
func (t *SNSClient) ChangeArnTopic(Arntopic string) {
	fmt.Println("se cambia el topic " + t.Topic + " por " + Arntopic)
	t.Topic = Arntopic
}
func (t *SNSClient) Publicar(msg string) (error, sns.PublishOutput) {
	result, err := t.Svc.Publish(&sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: aws.String(t.Topic),
	})
	if err != nil {
		exitErrorf("PROBLEMA Publicar CON SNS, %v", err)
		return err, *result
	}
	return err, *result
}
func (t *SNSClient) CreateSubscribeEMail(Endpoint string) {
	result, err := t.Svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              aws.String(Endpoint),
		Protocol:              aws.String("email"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              aws.String(t.Topic),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(*result.SubscriptionArn)
}

func (t *SNSClient) CreateTopic(newtopic string, autochance bool) string {
	result, err := t.Svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(newtopic),
	})
	if err != nil {
		fmt.Println(err.Error())
		t.Topic = ""
		return ""
	}
	if autochance == true {
		t.ChangeArnTopic(*result.TopicArn)
	}

	fmt.Println(*result.TopicArn)
	return *result.TopicArn
}

func (t *SNSClient) ListSubscriptions() {
	result, err := t.Svc.ListSubscriptions(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, s := range result.Subscriptions {
		fmt.Println(*s.SubscriptionArn)
		fmt.Println("  " + *s.TopicArn)
		fmt.Println("")
	}
}
func (t *SNSClient) ListTopics() {

	result, err := t.Svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, t := range result.Topics {
		fmt.Println(*t.TopicArn)
	}
}
func main() {
	SNS.Region = "us-east-1"
	SNS.NewSession(SNS.Region)
	SNS.ChangeArnTopic("arn:aws:sns:us-east-1:561607169148:standartopic")
	SNS.CreateSubscribeEMail("tucorreo@gmail.com")
	//SNS.Publicar("mi nuevo codigo mejorado")
	//SNS.ListTopics()
	//SNS.ListSubscriptions()

}
