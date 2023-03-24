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
	pass := 0
	for _, u := range users {
		user := strconv.Itoa(int(u["user_id"].(int64)))
		use, _ := strconv.ParseInt(user, 10, 64)
		bot.SendMessage(
			use,
			ctx.EffectiveMessage.ReplyToMessage.Text,
			&gotgbot.SendMessageOpts{ParseMode: "html"},
		)
		pass += 1
	}
	msg.EditText(
		bot,
		fmt.Sprintf("Broadcast succesfuull!\nFailed: %v", (len(users)-pass)),
		&gotgbot.EditMessageTextOpts{ParseMode: "markdown"},
	)
	return nil
}
