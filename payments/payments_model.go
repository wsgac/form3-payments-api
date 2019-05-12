package payments

import "github.com/globalsign/mgo/bson"

type Payment struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Type        string        `bson:"type" json:"type"`
	Title       string        `bson:"title" json:"title"`
	Source      string        `bson:"source" json:"source"`
	Destination string        `bson:"destination" json:"destination"`
	Amount      string        `bson:"amount" json:"amount"`
}
