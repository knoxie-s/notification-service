# Notification service

[![Actions Status](https://github.com/nightlord189/tcp-pow-go/workflows/main/badge.svg)](https://github.com/nightlord189/tcp-pow-go/actions)

## 1. Description
### This project is a solution for the next problem:

Write notification service to notify user for card activities on his card. You need to write HTTP server which accepts events in JSON format via POST method and stores them in some sort of storage
(could store either in db or some basic memory storage). You also need to write worker/job which will notify those events to client. Notification can be mocked just by printing them to terminal.

- Try to use only standard library
- Provide instructions on how to build and run the code
- Keep your solution simple
- Add your solution to github and share the link with us
- We would like to see some tests too

#### Sample Events
``` JSON
{
"orderType": "Purchase",
"sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
"card": "4433**1409",
"eventDate": "2023-01-04 13:44:52.835626 +00:00",
"websiteUrl": "https://amazon.com"
},

{
"orderType": "CardVerify",
"sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
"card": "4433**1409",
"eventDate": "2023-04-07 05:29:54.362216 +00:00",
"websiteUrl": "https://adidas.com"
},

{
"orderType": "SendOtp",
"sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
"card": "4433**1409",
"eventDate": "2023-04-06 22:52:34.930150 +00:00",
"websiteUrl": "https://somon.tj"
}
```

## 2. Getting started
### 2.1 Requirements
+ [Docker](https://docs.docker.com/engine/install/) installed (to run docker-compose)
+ [Minimock](https://github.com/gojuno/minimock) installed (to generate mocks)

### 2.2 Start server (Go)
```
go run cmd/main.go
```

### 2.3 Build notification-service image (Docker)
```
make docker-build
```

### 2.4 Run server (Docker)
```
make docker-rm
make docker-run
```

### 2.5 Launch tests
```
make test
```

## 3. Demo
https://github.com/user-attachments/assets/2e4fbe66-85e8-4e12-a0b6-caa2b761e422
