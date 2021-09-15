package main

import (
	"github.com/cookienyancloud/tgMag/api"

)

const (
	//tokenA = "TOKEN_A"
	tokenB ="TOKEN_B"
)

var Users map[int64]string

func main() {

	bot, updates := api.StartSotaBot(tokenB)

	for update := range updates {
		if update.Message.IsCommand() {
			Users[update.Message.Chat.ID] = update.Message.Command()
			continue
		}
		currentCommand:= Users[update.Message.Chat.ID]

	}


}