# Telegram Bot with OpenAI GPT-3 Integration
<p float="left">
<img src="https://img.icons8.com/color/48/000000/golang.png" width="100" alt="Go icon" />
<img src="https://cdn.iconscout.com/icon/premium/png-512-thumb/openai-1523664-1290202.png" width="100" alt="GPT-3 icon" />
<img src="https://img.icons8.com/color/48/000000/cloud-function.png" width="100" alt="Google Cloud Function" />
</p>

This is a Telegram bot that uses OpenAI GPT-3 to generate responses to user input. The bot is written in Go and deployed on Google Cloud Function, which allows it to be highly scalable and cost-efficient.

The bot uses the `X-Telegram-Bot-Api-Secret-Token` header to recognize valid Telegram messages, and it uses the `TG_USER_ID` environment variable as the Telegram user ID to recognize valid Telegram users. This means that only authorized users with the correct Telegram user ID and authentication token will be able to send messages to the bot.

To deploy the bot, you will need to obtain a Telegram bot token and an OpenAI API key, as well as have Go 1.138 or later installed on your local machine. The bot also uses a Telegram webhook to receive messages, which must be set up as part of the deployment process.

## Prerequisites
Before deploying the bot, you will need to have the following:

- A Telegram bot token, which can be obtained by following the instructions in the [Telegram Bot API documentation](https://core.telegram.org/bots#creating-a-new-bot).
- An OpenAI API key, which can be obtained by signing up for the [OpenAI API](https://beta.openai.com/signup/) and creating a new API key.
- Go 1.18 or later installed on your local machine.
- GNU Make installed on your local machine.

## Installation
To deploy the bot, follow these steps:
1. Clone this repository to your local machine.
2. Set up a Google Cloud Function by following the instructions in the [Google Cloud Function documentation](https://cloud.google.com/functions/docs/quickstart-console).
3. Set the following environment variables in a local file `.env.yaml`:
    - `OPENAI_API_URL`: "https://api.openai.com/v1/completions"
    - `TG_API_URL`: "https://api.telegram.org"
    - `TG_BOT_TOKEN`: The Telegram bot token obtained in the Prerequisites section.
    - `OPENAI_API_KEY`: The OpenAI API key obtained in the Prerequisites section.
    - `TG_HEADER_TOKEN`: The secret token to make sure that the webhook was set by the Telegram bot. (The `X-Telegram-Bot-Api-Secret-Token` header)
    - `TG_USER_ID`: The telegram user ID to recognize valid telegram user.
4. Deploy to the Google Cloud Function with `make deploy`.
5. Set up a Telegram webhook by following the instructions in the [Telegram Bot API documentation](https://core.telegram.org/bots/api#setwebhook). The webhook URL should be set to the URL of the Google Cloud Function, and the secret token should be set to the value of `TG_HEADER_TOKEN`.

## Usage
To use the bot, follow these steps:

1. Start a chat with the bot by searching for its name in the Telegram app and clicking on it.
2. Send a message to the bot.
3. The bot will generate a response to your message using OpenAI GPT-3 and send it back to you.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
