package travel

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor int64 `json:"cursor"`
	Limit  int64 `json:"limit"`
}

func (c *BusinessTravelCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	callbackData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &callbackData)
	if err != nil {
		log.Fatalf("BusinessTravelCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)

		return
	}

	travels, err := c.travelService.List(callbackData.Cursor, callbackData.Limit)
	if err != nil {
		//log.Fatal(err)

		return
	}

	outputMsgText := "Here all the products: \n\n"
	for idx, p := range travels {
		outputMsgText += fmt.Sprintf("%s\n%s\n", strconv.Itoa(idx), p.String())
	}
	currentCount := callbackData.Cursor + int64(len(travels))

	count, err := c.travelService.Count()
	if err != nil {
		log.Fatal(err)
	}

	outputMsgText += fmt.Sprintf("\n%v - %v\n", currentCount, count)

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		outputMsgText,
	)

	inlineKeyboardRow := c.newInlineKeyboardRow(count, currentCount, callbackData, callbackPath)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineKeyboardRow)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BusinessTravelCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}

func (c *BusinessTravelCommander) newInlineKeyboardRow(totalCount int64, currentCount int64, callbackData CallbackListData, callbackPath path.CallbackPath) []tgbotapi.InlineKeyboardButton {

	var buttons = make([]tgbotapi.InlineKeyboardButton, 0, 2)

	if callbackData.Cursor == 0 {
		buttons = append(buttons, newInlineKeyboardButton("Next page", callbackData.Cursor+callbackData.Limit, callbackPath))
	} else if currentCount > 0 {
		buttons = append(buttons, newInlineKeyboardButton("Previous page", callbackData.Cursor-3, callbackPath))
		if currentCount != totalCount {
			buttons = append(buttons, newInlineKeyboardButton("Next page", callbackData.Cursor+callbackData.Limit, callbackPath))
		}
	} else {
		buttons = append(buttons, newInlineKeyboardButton("Next page", callbackData.Cursor+callbackData.Limit, callbackPath))
	}

	return tgbotapi.NewInlineKeyboardRow(buttons...)
}

func newInlineKeyboardButton(text string, cursor int64, callbackPath path.CallbackPath) tgbotapi.InlineKeyboardButton {
	serializedData, _ := json.Marshal(CallbackListData{Cursor: cursor, Limit: 3})
	callbackPath.CallbackData = string(serializedData)

	return tgbotapi.NewInlineKeyboardButtonData(text, callbackPath.String())
}
