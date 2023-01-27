package travel

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessTravelCommander) New(inputMessage *tgbotapi.Message) {

	product, err := c.travelService.Create(business.Travel{
		Title:     "new",
		Where:     "new",
		StartDate: time.Now(),
		Duration:  10,
	})
	if err != nil {
		log.Print(err)
		return
	}

	outputMsgText := fmt.Sprintf("Travel created:\n\n%v", product)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.New: error creating new travel - %v", err)
	}
}
