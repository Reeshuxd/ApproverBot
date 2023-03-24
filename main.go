//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package main

import (
	"ApproverBot/plugs"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

func main() {
	bot, err := gotgbot.NewBot(os.Getenv("TOKEN"), &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})
	if err != nil {
		fmt.Println("Error while logging as a bot:" + err.Error())
	}
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})
	dp := updater.Dispatcher
	dp.AddHandler((handlers.NewCommand("start", plugs.Start)))
	dp.AddHandler((handlers.NewCommand("ping", plugs.Ping)))
	dp.AddHandler(handlers.NewCommand("approval", plugs.SetApproval))
	dp.AddHandler(handlers.NewCallback(callbackquery.Equal("menu"), plugs.MenuCB))
	dp.AddHandler(handlers.NewCommand("get_approval", plugs.GetApp))
	dp.AddHandler(handlers.NewCommand("get_message", plugs.GetMessage))
	dp.AddHandler(handlers.NewCommand("message", plugs.AddMessage))
	dp.AddHandler(handlers.NewCommand("stats", plugs.Stats))
	dp.AddHandler(handlers.NewCommand("bcast", plugs.Broadcast))
	dp.AddHandler(handlers.NewChatJoinRequest(nil, plugs.NewAcceptor))

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("@%v has been successfully started!", bot.Username)
	updater.Idle()
}
