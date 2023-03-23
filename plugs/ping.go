//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package plugs

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Ping(bot *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	start := time.Now().UnixNano() / int64(time.Millisecond)
	msg, _ := bot.SendMessage(
		chat.Id,
		"`Processing....`",
		&gotgbot.SendMessageOpts{ParseMode: "markdown"},
	)
	end := time.Now().UnixNano() / int64(time.Millisecond)
	ping := end - start
	msg.EditText(
		bot,
		fmt.Sprintf("<b>Ping:</b> <code>%v</code>", ping),
		&gotgbot.EditMessageTextOpts{ParseMode: "html"},
	)
	return nil
}
