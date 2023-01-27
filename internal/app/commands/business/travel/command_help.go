package travel

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BusinessTravelCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__business__travel — help\n"+
			"/list__business__travel — list products\n"+
			"/get__business__travel — get a entity\n"+
			"/delete__business__travel — delete an existing entity\n"+
			"/new__business__travel — create a new entity\n"+
			"/edit__business__travel — edit a entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.Help: error sending reply message to chat - %v", err)
	}
}
