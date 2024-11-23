package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvent(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))  // initialise the bot 

	go printCommandEvent(bot.CommandEvents())  // spawn a goroutine for command analytics

	bot.Command("this is the joke: <joke>", &slacker.CommandDefinition){
		Description:
	}

	ctx, cancel := context.WithCancel(context.Background())  // context for graceful shutdown
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
