//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package plugs

import (
	"ApproverBot/db"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GetMessage(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser.Id != OID {
		return nil
	}
	msg := db.GetMsg(ctx.EffectiveUser.Id)
	if msg == "NoMsg()" {
		ctx.EffectiveMessage.Reply(
			bot,
			"there is no message has been added by you!",
			&gotgbot.SendMessageOpts{ParseMode: "markdown"},
		)
	} else {
		ctx.EffectiveMessage.Reply(
			bot,
			msg,
			&gotgbot.SendMessageOpts{ParseMode: "markdown"},
		)

	}
	return nil
}

func AddMessage(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser.Id != OID {
		return nil
	}
	reply := ctx.EffectiveMessage.ReplyToMessage
	chat := ctx.EffectiveChat
	if reply == nil {
		bot.SendMessage(chat.Id, "Please reply to a message to add", &gotgbot.SendMessageOpts{})
		return nil
	}
	db.AddMsg(ctx.EffectiveUser.Id, reply.Text)
	bot.SendMessage(
		chat.Id,
		"Succesfully updated new message.",
		&gotgbot.SendMessageOpts{},
	)
	return nil
}
