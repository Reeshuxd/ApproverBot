//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package plugs

import (
	"ApproverBot/db"
	"fmt"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Stats(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser.Id != OID {
		return nil
	}
	users := db.GetUsers()
	msg := fmt.Sprintf("**Users:** %v", len(users))
	bot.SendMessage(
		ctx.EffectiveChat.Id,
		msg,
		&gotgbot.SendMessageOpts{ParseMode: "markdown"},
	)
	return nil
}

func Broadcast(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser.Id != OID {
		return nil
	}
	chat := ctx.EffectiveChat
	msg, _ := bot.SendMessage(
		chat.Id,
		"`Processing....`",
		&gotgbot.SendMessageOpts{ParseMode: "markdown"},
	)
	users := db.GetUsers()
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		msg.EditText(
			bot,
			"Reply to a message to broadcast.",
			&gotgbot.EditMessageTextOpts{},
		)
	}
	fail := 0
	for _, user := range users {
		ui := fmt.Sprintf("%v", user["user_id"])
		uid, _ := strconv.ParseInt(ui, 10, 64)
		_, err := bot.SendMessage(
			uid,
			ctx.EffectiveMessage.ReplyToMessage.Text,
			&gotgbot.SendMessageOpts{ParseMode: "html"},
		)
		if err != nil {
			fmt.Println("Failed Bcast:", err.Error())
			fail++
			continue
		}
	}
	bot.SendMessage(
		chat.Id,
		fmt.Sprintf("Broadcast succesfuull!\nFailed: %v", fail),
		&gotgbot.SendMessageOpts{ParseMode: "markdown"},
	)
	return nil
}
