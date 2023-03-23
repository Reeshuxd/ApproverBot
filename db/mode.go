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
	cdb = database.Collection("Chats")
)

func GetApproval(user_id int64) bool {
	find := cdb.FindOne(context.TODO(), bson.M{"user_id": user_id})
	return find.Err() == nil
}

func Approval(user_id int64, mode bool) (bool, error) {
	xmp := new(bool)
	*xmp = true
	opt := bson.M{"$set": bson.M{"approval": mode}}
	_, err := cdb.UpdateOne(context.TODO(), bson.M{"user_id": user_id}, opt, &options.UpdateOptions{Upsert: xmp})
	if err != nil {
		return false, err
	}
	return true, nil
}
