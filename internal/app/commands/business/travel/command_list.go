package travel

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *BusinessTravelCommander) List(inputMessage *tgbotapi.Message, callbackData CallbackListData) {
	outputMsgText := "Here all the travels: \n\n"

	travels, err := c.travelService.List(callbackData.Cursor, callbackData.Limit)
	for idx, p := range travels {
		outputMsgText += fmt.Sprintf("%s\n%s\n", strconv.Itoa(idx), p.String())
	}

	currentCount := callbackData.Cursor + int64(len(travels))

	count, err := c.travelService.Count()
	if err != nil {
		log.Fatal(err)
	}

	outputMsgText += fmt.Sprintf("\n%v - %v\n", currentCount, count)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{Cursor: callbackData.Cursor + callbackData.Limit, Limit: 3})

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
