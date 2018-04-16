package slacksend

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var flogoLogger = logger.GetLogger("activity-tibco-slacksend")

// MyActivity is a stub for your Activity implementation
type SlackSendActivity struct {
	metadata        *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	flogoLogger.Debugf("SlackSend NewActivity")
	slackSendActivity := &SlackSendActivity{metadata: metadata}
	return slackSendActivity
}

// Metadata implements activity.Activity.Metadata
func (a *SlackSendActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *SlackSendActivity) Eval(context activity.Context) (done bool, err error) {
		flogoLogger.Debugf("SlackSend Eval")
	
		accesstoken := context.GetInput("AccessToken")
		message := context.GetInput("Message")
		channel := context.GetInput("Channel")
		
		api := slack.New(accesstoken)
		params := slack.PostMessageParameters{}
		/*attachment := slack.Attachment{
		Pretext: "Flogo",
		Text:    message,
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
		//}
		//params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channel, message, params)
		if err != nil {
			fmt.Errorf("%s\n", err)
			return
		}
		context.SetOutput("timestamp", timestamp)
		flogoLogger.Debugf("Message successfully sent to channel %s at %s", channelID, timestamp)
		return true, nil
	}
	return false, fmt.Errorf("SlackSend called without a message to publish")
}
