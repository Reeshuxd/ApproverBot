//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package plugs

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	text := `
<b>Heya,</b> I am developed for accepting new coming join requests in a channel.
I am made with <a href="go.dev">golang</a> langauge for a better performance.

<b>(c) Made by <a href="https://github.com/reeshuxd/ApproverBot">Reeshu</a></b>
	`
	ctx.EffectiveMessage.Reply(bot, text, &gotgbot.SendMessageOpts{
		ParseMode:             "html",
		DisableWebPagePreview: true,
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{
					{Text: "Main Menu", CallbackData: "menu"},
					{Text: "Source code", Url: "https://github.com/reeshuxd/ApproverBot"},
				},
			},
		},
	})
	return nil
}

func MenuCB(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser.Id != OID {
		return nil
	}
	text := `
<b>Help-Menu</b>

/start - To start the bot.
/ping - To ping the bot.
/approval <code>(enable/disable)</code> - To enable or disable approval mode.
/get_approval - To get current approval settings.
/get_message - To get your current welcome message.
/add_message <code>(reply_to_message)</code> - To change your welcome message.
/stats - To get current number of users.
/bcast (reply_to_message) - To send broadcast to the users.

<i>Thanks for using the bot.💖</i>
<b>(c) Made by <a href="github.com/reeshuxd/ApproverBot">Reeshu</a></b>
	`
	cb := ctx.Update.CallbackQuery
	cb.Message.EditText(
		bot,
		text,
		&gotgbot.EditMessageTextOpts{
			ParseMode:             "html",
			DisableWebPagePreview: true,
		},
	)
	return nil
}
