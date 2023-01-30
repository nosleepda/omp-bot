package travel

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *BusinessTravelCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	fields := strings.Split(args, ";")

	if len(fields) != 5 {
		log.Println("wrong args", args)
		return
	}

	idx, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	mappedTravel, err := mapTravel(fields[1:])

	err = c.travelService.Update(int64(idx), mappedTravel)

	if err != nil {
		log.Print(err)
		return
	}

	outputMsgText := fmt.Sprintf("Travel %d updated", idx)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.List: error sending reply message to chat - %v", err)
	}
}
