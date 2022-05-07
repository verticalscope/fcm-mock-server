# FCM Mock Server GO (docker)

Simple FCM mock for both old fcm HTTP and HTTP v1
Image for this repo on [DockerHub](https://hub.docker.com/repository/docker/andruxxxa/fcm-mock-server-go)

## Quick start

To run in docker simply run this:
```bash
docker run -p 4004:4004 -d andruxxxa/fcm-mock-server-go
```
The docker exposes one port - 4004.

To run localy just:
```bash
go run main.go
```

## API

Endpoints:
* **POST** */send* - send message to the mock (200 response) - old HTTP
* **POST** */v1/projects/[project_id]/messages:send* - send message to the mock (200 response) - HTTP V1
* **GET** */api/messages* - returns (as `JSON`) history of `FCM` API calls to this mock
* **DELETE** */api/messages* - clear messages history

## Message model
### Old HTTP
```js
{
    "to": "string",    // device token or topic
    "data": {       // [string]string dictionary
        "field1": "value1",
	    "field2": "value2",
        ...
    },
    "notification": {
        "title": "title",
        "body": "body"
    }
}
```
### HTTP V1
```js
{
    "Message": {
        "token": "string",    
        "topic" : "string",
        "condition" : "string",
        "data": {       // [string]string dictionary
            "field1": "value1",
            "field2": "value2",
            ...
        },
        "notification": {
            "title": "string",
            "body": "string"
        }
    }
}
```
## test.http

Simple test requests - work with [REST client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)