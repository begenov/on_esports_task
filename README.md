# Discord Bot

## Description
This Discord bot is created to polls and fetching weather information. Simply set up the bot according to the instructions below and start using it on your Discord server.

## Installation and Configuration

1. **Clone the repository:**

 ```bash
   git clone https://github.com/begenov/on_esports_task.git
   cd on_esports_task
```

2. **Install dependencies:**
```bash
go mod tidy
```

3. **Configure settings:**
Create a .env file in the project's root and specify the required parameters:

```bash
OPEN_WEATHER_MAP_API_KEY='OPEN_WEATHER_MAP_API_KEY'
DISCORD_BOT_TOKEN='DISCORD_BOT_TOKEN'
```

4. **Run the bot:**

```bash
go run cmd/main.go
```

## User Guide

### Help

- `!help`: Display information about available commands and how to use them.

### Weather

- `!weather <city>`: Get the current weather in the specified city.

### Polls

- `!poll create "Poll Title" "Option 1" "Option 2" ...`: Create a new poll with the specified answer options.
- `!poll vote <poll_number> <option_number>`: Vote for the selected option in the poll.
- `!poll result <poll_number>`: View the results of a specific poll.

**Example:**
```bash
!poll create "Favorite Color" "Red" "Blue" "Green"
!poll vote 1 2
!poll result 1
```

**Make sure to replace `OPEN_WEATHER_MAP_API_KEY` and `DISCORD_BOT_TOKEN` with your actual OpenWeatherMap API key and Discord bot token.**