# FCM Mock Server GO (docker)

Simple FCM mock for testing server applications.
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
* **POST** */send* - send message to the mock (200 response)
* **GET** */messages* - returns (as `JSON`) history of `FCM` API calls to this mock
* **DELETE** */messages* - clear messages history

## Message model
```js
{
	"to": "123",    // device token or topic
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

## test.http

Simple test requests - work with [REST client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)