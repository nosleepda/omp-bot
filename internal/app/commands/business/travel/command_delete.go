package travel

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *BusinessTravelCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.travelService.Remove(int64(idx))

	if err != nil {
		log.Print(err)
		return
	}

	var outputMsgText string
	if ok {
		outputMsgText = fmt.Sprintf("Travel %d deleted", idx)
	} else {
		outputMsgText = fmt.Sprintf("Travel %d not deleted", idx)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.Delete: error sending reply message to chat - %v", err)
	}
}
