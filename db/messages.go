//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mdb = database.Collection("Messages")
)

type User struct {
	Message string `bson:"message"`
}

func GetMsg(user_id int64) string {
	check := mdb.FindOne(context.TODO(), bson.M{"user_id": user_id})
	if check.Err() == nil {
		var result User
		check.Decode(&result)
		return result.Message
	}
	return "NoMsg()"
}

func AddMsg(user_id int64, message string) (bool, error) {
	xmp := new(bool)
	*xmp = true
	opt := bson.M{"$set": bson.M{"message": message}}
	_, err := mdb.UpdateOne(context.TODO(), bson.M{"user_id": user_id}, opt, &options.UpdateOptions{Upsert: xmp})
	if err != nil {
		return false, err
	}
	return true, nil
}
