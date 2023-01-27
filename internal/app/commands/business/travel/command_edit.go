package travel

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessTravelCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	err = c.travelService.Update(int64(idx), business.Travel{
		Title:     "new2",
		Where:     "new2",
		StartDate: time.Now(),
		Duration:  10,
	})

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
