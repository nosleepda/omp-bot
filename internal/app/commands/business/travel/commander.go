package travel

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/business/travel"
)

type TravelCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message, callbackData CallbackListData)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessTravelCommander struct {
	bot           *tgbotapi.BotAPI
	travelService travel.TravelService
}

func NewTravelCommander(bot *tgbotapi.BotAPI) *BusinessTravelCommander {
	service := travel.NewDummyTravelService()
	return &BusinessTravelCommander{
		bot:           bot,
		travelService: service,
	}
}

func (c *BusinessTravelCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BusinessTravelCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BusinessTravelCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg, CallbackListData{Cursor: 0, Limit: 3})
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
