package slacksend

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/nlopes/slack"
	"fmt"
)

// log is the default package logger
var flogoLogger = logger.GetLogger("activity-tibco-slacksend")

// SlackSendActivity is a stub for your Activity implementation
type SlackSendActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	flogoLogger.Debugf("SlackSend NewActivity")
	return &SlackSendActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *SlackSendActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *SlackSendActivity) Eval(context activity.Context) (done bool, err error) {
	flogoLogger.Debugf("SlackSend Eval")

	accesstoken := context.GetInput("AccessToken").(string)
	message := context.GetInput("Message").(string)
	channel := context.GetInput("Channel").(string)

	api := slack.New(accesstoken)
	params := slack.PostMessageParameters{}
	
	//attachment := context.GetInput("Attachment").(map[string]string)
	fmt.Printf("\n Attachment: %+v", context.GetInput("Attachment"))
	fmt.Printf("\n Attachment.attachments: %+v", context.GetInput("Attachment").attachments)
	
	if attachments, ok := context.GetInput("Attachment").(map[string][]slack.Attachment); ok {
		fmt.Printf("\n\n attachments json object:::: %+v", attachments)
		params.Attachments = []slack.Attachment{}
		for _, attachElem := range attachments["attachments"] {
			fmt.Printf("\n attachElem json object:::: %+v", attachElem)
			params.Attachments = append(params.Attachments, attachElem)
		}
	} else {
		fmt.Printf("Can not parse message attachment content...")
		return
	}
	
	channelID, timestamp, err := api.PostMessage(channel, message, params)
	if err != nil {
		flogoLogger.Debugf("%s\n", err)
		return
	}
	context.SetOutput("timestamp", timestamp)
	flogoLogger.Debugf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return true, nil
}
