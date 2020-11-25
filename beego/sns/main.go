package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must supply a topic name")
		fmt.Println("Usage: go run SnsCreateTopic.go TOPIC")
		os.Exit(1)
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	//create in topic
	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(os.Args[1]),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.TopicArn)

	// print topic list
	list, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, t := range list.Topics {
		fmt.Println(*t.TopicArn)
	}

	//sns print subscriber
	emailPtr := flag.String("e", "shiv.suthar9001@gmail.com", "The email address of the user subscribing to the topic")
	topicPtr := flag.String("t", "arn:aws:sns:ap-south-1:693461261720:shiva", "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *emailPtr == "" || *topicPtr == "" {
		fmt.Println("You must supply an email address and topic ARN")
		fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
		os.Exit(1)
	}
	subreu, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              emailPtr,
		Protocol:              aws.String("email"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*subreu.SubscriptionArn)

	//messageing to all subscriber of given topic
	input := &sns.PublishInput{
		Message:  aws.String("Hello world!"),
		TopicArn: aws.String("arn:aws:sns:ap-south-1:693461261720:shiva"),
	}

	mesreu, err := svc.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println(mesreu)
}
