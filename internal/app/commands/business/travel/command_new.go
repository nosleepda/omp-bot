package travel

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *BusinessTravelCommander) New(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	fields := strings.Split(args, ";")

	if len(fields) != 4 {
		log.Println("wrong args", args)
		return
	}

	mappedTravel, err := mapTravel(fields)

	createdTravel, err := c.travelService.Create(mappedTravel)
	if err != nil {
		log.Print(err)
		return
	}

	outputMsgText := fmt.Sprintf("Travel created:\n\n%v", createdTravel)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.New: error creating new travel - %v", err)
	}
}
