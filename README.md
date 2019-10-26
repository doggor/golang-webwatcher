It's my first golang program which will notify me once the Tokyo-Naeba Bus is available for making booking at 7 Jan 2020.

This is a personal project for fun only.

## Build locally
```
go build cmd/webwatcher/main.go
```

## Build to Docker image
```sh
docker build -t doggor/webwatcher:{version} .
```

## Env variable required for execution
- TELEGRAM_CHATBOT_TOKEN:
  The token of the Telegram chatbot to notify me.
- TELEGRAM_USER_ID:
  My Telegram user ID the chatbot can find me.
