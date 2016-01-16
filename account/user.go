package account

import (
	"errors"

	"github.com/gernest/cifer/core"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

const (
	accountCollection = "accounts"
)

func Create(c *core.Config, usr *core.User) (*core.User, error) {
	collection, err := c.Collection(accountCollection)
	if err != nil {
		return nil, err
	}
	err = collection.Find(bson.M{"email": usr.Email}).One(&core.User{})
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// hash password
	p, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	usr.Password = string(p)
	usr.ID = bson.NewObjectId()
	err = collection.Insert(usr)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
