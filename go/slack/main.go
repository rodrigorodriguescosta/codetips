package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
)

func main() {
	api := slack.New("yourtoken")
	headerText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("Hello :\n*Rodrigo*"), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	approvalText := slack.NewTextBlockObject("mrkdwn", "Subject: hey, how are you?", false, false)

	fieldsSection := slack.NewSectionBlock(approvalText, nil, nil)

	_, _, err := api.PostMessage("#dev", slack.MsgOptionBlocks(
		headerSection,
		fieldsSection,
	))
	if err != nil {
		log.Fatal(err)
	}
}
