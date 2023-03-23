//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx, _ = context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	client, _ = mongo.Connect(
		ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URL")),
	)
	database = client.Database("AcceptorBot")
)
