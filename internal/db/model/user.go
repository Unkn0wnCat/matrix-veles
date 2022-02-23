package model

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

type DBUser struct {
	Username       string `bson:"username" json:"username"`
	HashedPassword string `bson:"password" json:"password"`

	Password *string `bson:"-" json:"-"`
}

func (usr *DBUser) HashPassword() error {
	if usr.Password == nil {
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(*usr.Password), 14)
	if err != nil {
		return err
	}

	usr.HashedPassword = base64.StdEncoding.EncodeToString(hash)
	return nil
}

func (usr *DBUser) CheckPassword(password string) error {
	hash, err := base64.StdEncoding.DecodeString(usr.HashedPassword)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))

	return err
}
