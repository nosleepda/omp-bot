package travel

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessTravelCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var idx int
	var err error

	if len(args) == 0 {
		idx = 0
	} else {
		idx, err = strconv.Atoi(args)
		if err != nil {
			log.Println("wrong args", args)
			return
		}
	}

	product, err := c.travelService.Describe(int64(idx))
	if err != nil {
		log.Printf("fail to get travel with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.Get: error sending reply message to chat - %v", err)
	}
}
