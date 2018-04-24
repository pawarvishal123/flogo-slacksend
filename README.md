# Send Slack Message
This activity allows you to send a message on slack channel

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
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
