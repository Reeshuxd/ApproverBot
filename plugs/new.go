//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package plugs

import (
	"GoBot/db"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var OWNER_ID = os.Getenv("OWNER_ID")
var OID, _ = strconv.ParseInt(OWNER_ID, 10, 64)

func NewAcceptor(bot *gotgbot.Bot, ctx *ext.Context) error {
	app := db.GetApproval(ctx.EffectiveUser.Id)
	msg := db.GetMsg(int64(OID))
	if !app {
		db.AddUser(ctx.EffectiveUser.Id)
		bot.SendMessage(
			ctx.EffectiveUser.Id,
			msg,
			&gotgbot.SendMessageOpts{ParseMode: "html"},
		)
	} else {
		bot.ApproveChatJoinRequest(ctx.EffectiveChat.Id, ctx.EffectiveUser.Id, &gotgbot.ApproveChatJoinRequestOpts{})
		bot.SendMessage(
			ctx.EffectiveUser.Id,
			msg,
			&gotgbot.SendMessageOpts{ParseMode: "html"},
		)
	}
	return nil
}

func SetApproval(bot *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	user := ctx.EffectiveUser
	if user.Id != OID {
		return nil
	}
	text := strings.Split(ctx.EffectiveMessage.Text, " ")
	if text == nil {
		bot.SendMessage(chat.Id, "Give an arguement to set\nEnable or Disable", &gotgbot.SendMessageOpts{})
	}
	var mode bool
	curr := db.GetApproval(user.Id)
	if text[1] == "enable" {
		mode = true
	} else if text[1] == "disable" {
		mode = false
	} else {
		bot.SendMessage(
			chat.Id,
			fmt.Sprintf("**Your current approval mode:** `%v`", curr),
			&gotgbot.SendMessageOpts{ParseMode: "markdown"},
		)
	}
	db.Approval(user.Id, mode)
	bot.SendMessage(
		chat.Id,
		fmt.Sprintf("Successfully %s approval!", text),
		&gotgbot.SendMessageOpts{ParseMode: "markdown"},
	)
	return nil
}

func GetApp(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser.Id != OID {
		return nil
	}
	get := db.GetApproval(ctx.EffectiveUser.Id)
	var mode string
	if !get {
		mode = "disabled"
	} else {
		mode = "enabled"
	}
	bot.SendMessage(
		ctx.EffectiveChat.Id,
		fmt.Sprintf("<b>Your current approval mode:</b> %v", mode),
		&gotgbot.SendMessageOpts{ParseMode: "html"},
	)
	return nil
}
