package travel

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *BusinessTravelCommander) New(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products, err := c.travelService.List(int64(0), int64(0))
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{Cursor: 0, Limit: 3})

	callbackPath := path.CallbackPath{
		Domain:       "business",
		Subdomain:    "travel",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.List: error sending reply message to chat - %v", err)
	}
}
