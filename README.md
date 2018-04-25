# Send Slack Message
This activity allows you to send a message on slack channel along with various attachment options. 
You need to make sure that the slack API access token have enough permissions to publish the message on channel.

## Installation
### Flogo CLI
```bash
flogo install github.com/pawarvishal123/slacksend
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
    {
      "name": "AccessToken",
      "type": "string",
      "required": true
    },
    {
      "name": "Channel",
      "type": "string",
      "required": true
    },
    {
      "name": "Message",
      "type": "string",
      "required": true
    },
    {
      "name": "Attachment",
      "type": "string"
    }
  ],
  "output": [
    {
      "name": "timestamp",
      "type": "string"
    }
  ]
}
```

## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| AccessToken  | True     | The Access Token from Slack |
| Channel       | True     | Name of the slack channel |
| Message     | True     | The text message to send |
| Attachment     | False     | Message Attachment |

## Example Inputs
```json
{
            "id": "slacksend_2",
            "name": "Send Slack message",
            "description": "Publish a message to a Slack Channel",
            "activity": {
              "ref": "github.com/pawarvishal123/slacksend",
              "input": {
                "AccessToken": "<<YOUR-ACCESS-TOKEN>>",
                "Channel": "general",
                "Message": "msg attachment test",
                "Attachment": {
                  "attachments": [
                    {
                      "fallback": "Required plain-text summary of the attachment.",
                      "color": "#36a64f",
                      "pretext": "Optional text that appears above the attachment block",
                      "author_name": "Mr abc",
                      "author_link": "http://flickr.com/bobby/",
                      "author_icon": "http://flickr.com/icons/bobby.jpg",
                      "title": "Slack API Documentation",
                      "title_link": "https://api.slack.com/",
                      "text": "Optional text that appears within the attachment",
                      "fields": [
                        {
                          "title": "Priority",
                          "value": "High",
                          "short": false
                        }
                      ],
                      "image_url": "http://my-website.com/path/to/image.jpg",
                      "thumb_url": "http://example.com/path/to/thumb.png",
                      "footer": "Slack API",
                      "footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
                      "ts": 123456789
                    }
                  ]
                }
              }
            }
          }
```

## Third Party Library
Slack API in Go - [https://github.com/nlopes/slack](https://github.com/nlopes/slack)
