package main

import (
	"fmt"
	"os"

	"github.com/icrowley/fake"
	log "github.com/sirupsen/logrus"

	"github.com/bayugyug/slack"
)

func main() {
	// init
	var (
		servicePath = os.Getenv("SLACK_TOKEN")
		alert       = slack.NewNotification(servicePath)
	)

	log.Println("ALL-URL:", servicePath)

	if len(servicePath) <= 0 {
		fmt.Println("Token is required.")
		os.Exit(1)
	}

	// push
	err := alert.Notify([]*slack.Payload{
		{
			Title:       "Event Push",
			Message:     fake.Sentences(),
			WithDivider: true,
			Icon:        slack.IconTypeSpeechBalloon,
		},
		{
			Title:       "Notify",
			Message:     fake.Sentences(),
			WithDivider: true,
			Icon:        slack.IconTypeHeart,
		},
		{
			Title:       "Star",
			Message:     fake.Sentences(),
			WithDivider: true,
			Icon:        slack.IconTypeStar,
		},
		{
			Title:       "Warning",
			Message:     fake.Sentences(),
			WithDivider: true,
			Icon:        slack.IconTypeWarning,
		},
		{
			Title:       "Critical",
			Message:     fake.Sentences(),
			WithDivider: true,
			Priority:    slack.PriorityCritical,
			Icon:        slack.IconTypeCritical,
			Here:        "Please check ...",
		},
		{
			Title:       "Success",
			Message:     fake.Sentences(),
			WithDivider: true,
			Icon:        slack.IconTypeThumbsUp,
		},
	})
	if err != nil {
		log.Println("fail", err)
		return
	}

	log.Println("sent success")
}
