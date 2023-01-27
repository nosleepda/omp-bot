package business

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/travel"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessCommander struct {
	bot             *tgbotapi.BotAPI
	travelCommander travel.TravelCommander
}

func NewBusinessCommander(bot *tgbotapi.BotAPI) *BusinessCommander {
	commander := travel.NewTravelCommander(bot)
	return &BusinessCommander{
		bot:             bot,
		travelCommander: commander,
	}
}

func (c *BusinessCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "travel":
		c.travelCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BusinessCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BusinessCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "travel":
		c.travelCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BusinessCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
