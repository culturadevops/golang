package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type MyEvent struct {
	InstanceRegion string   `json:"InstanceRegion"`
	InstanceIDList []string `json:"InstanceIdList"`
}

func HandleLambdaEvent(event MyEvent) (string, error) {
	// instances id list
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{},
	}
	for _, id := range event.InstanceIDList {
		input.InstanceIds = append(input.InstanceIds, aws.String(id))
	}

	// stop instances
	svc := ec2.New(session.New())
	result, err := svc.StopInstances(input)
	if err == nil {
		fmt.Println(result)

	} else {
		fmt.Println(err.Error())
	}

	return fmt.Sprintf("Stop instances %d!", len(input.InstanceIds)), nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}