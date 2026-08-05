# HasiiChatBot

A fast and smart Telegram chatbot built with Golang and MongoDB. The bot automatically learns from user conversations and stickers in your group to generate natural replies.

## Features
- Written in Golang for extreme speed and low memory usage
- Learns dynamically from user replies (both text and stickers)
- Group admins can turn the chatbot on or off using `/chatbot [on/off]`
- Fast MongoDB storage for learned data
- Clean architecture and simple to deploy

## Requirements
- Go 1.20+
- MongoDB Database
- Telegram Bot Token

## Installation
1. Clone the repository
```bash
git clone https://github.com/hasindu-nagolla/HasiiChatBot.git
cd HasiiChatBot
```

2. Setup environment variables in a `.env` file
```env
BOT_TOKEN=your_telegram_bot_token_here
MONGO_URL=your_mongodb_connection_string_here
```

3. Build and Run
```bash
go mod tidy
go build -o bot ./cmd/bot/main.go
./bot
```
## Legacy Version

Looking for the previous Python (Pyrogram) version? 

It is no longer actively maintained, but you can find the legacy source code here:
[HasiiChatBot Legacy Version (Python)](https://github.com/hasindu-nagolla/HasiiChatBot/tree/0c98f265ac4ac7c3c5c5d7fdcf098aaad3fd160d)

## Support
- Owner: [@Hasindu_Lakshan](https://t.me/Hasindu_Lakshan)
- Source: [GitHub](https://github.com/hasindu-nagolla/HasiiChatBot)
