package payments

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type DAO struct {
	Host       string
	Database   string
	Collection string
}

var db *mgo.Database

func (p *DAO) Connect() {
	session, err := mgo.Dial(p.Host)
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB. Host: %s Error: %v",
			p.Host, err)
	}
	db = session.DB(p.Database)
}

func (p *DAO) FindAll() ([]Payment, error) {
	var payments []Payment
	err := db.C(p.Collection).Find(bson.M{}).All(&payments)
	return payments, err
}

func (p *DAO) FindById(id string) (Payment, error) {
	var payment Payment
	err := db.C(p.Collection).FindId(bson.ObjectIdHex(id)).One(&payment)
	return payment, err
}

func (p *DAO) Insert(payment Payment) error {
	err := db.C(p.Collection).Insert(&payment)
	return err
}

func (p *DAO) Update(payment Payment) error {
	err := db.C(p.Collection).UpdateId(payment.ID, &payment)
	return err
}

func (p *DAO) Delete(payment Payment) error {
	err := db.C(p.Collection).Remove(&payment)
	return err
}
