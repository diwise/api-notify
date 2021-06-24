# api-notify
Api for supporting subscriptions and notifications
## API
GET /subscriptions
POST /subscriptions
GET /subscriptions/{subscriptionId}
PUT /subscriptions/{subscriptionId}
DELETE /subscriptions/{subscriptionId}

## Examples
### Add subscription
POST to /subscriptions with body
```
{
  "entities": [
    {
      "id": "string",
      "type": "string",
      "idPattern": "string"
    },
     {
      "id": "str2",
      "type": "Device",
      "idPattern": "*.??"
    }   
  ],
  "name": "sub2",
  "description": "string",
  "watchedAttributes": [
    "string"
  ],
  "timeInterval": 0,
  "expires": "2021-06-17T12:19:13.877Z",
  "isActive": true,
  "throttling": 0,
  "q": "string",

  "csf": "string",
  "id": "sub2",
  "type": "Subscription",
  "notification": {
    "attributes": [
      "string"
    ],
    "format": "string",
    "endpoint": {
      "uri": "http://localhost:8668/v2/notify",
      "accept": "application/json"
    },
    "status": "ok",
    "timesSent": 0,
    "lastNotification": "2021-06-17T12:19:13.877Z",
    "lastFailure": "2021-06-17T12:19:13.877Z",
    "lastSuccess": "2021-06-17T12:19:13.877Z"
  },
  "status": "active",
  "createdAt": "2021-06-17T12:19:13.877Z",
  "modifiedAt": "2021-06-17T12:19:13.877Z"
}
```
```Id```: Id for Subscription. If exists the subscription will be updated
```entities```: What entities should the subscription create notifications for. Equal ID or equal type is implementetd.
