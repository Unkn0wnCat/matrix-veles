package model

import (
	"encoding/base64"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type DBUser struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	Username       string `bson:"username" json:"username"` // Username is the username the user has
	HashedPassword string `bson:"password" json:"-"`        // HashedPassword contains the bcrypt-ed password

	Admin *bool `bson:"admin,omitempty" json:"admin,omitempty"` // If set to true this user will have all privileges

	MatrixLinks        []*string `bson:"matrix_links" json:"matrix_links"`                 // MatrixLinks is the matrix-users this user has verified ownership over
	PendingMatrixLinks []*string `bson:"pending_matrix_links" json:"pending_matrix_links"` // PendingMatrixLinks is the matrix-users pending verification

	Password *string `bson:"-" json:"-"` // Password may never be sent out!
}

func (usr *DBUser) ValidateMXID(mxid string) error {
	for i, pendingMxid := range usr.PendingMatrixLinks {
		if strings.EqualFold(*pendingMxid, mxid) {
			usr.PendingMatrixLinks = append(usr.PendingMatrixLinks[:i], usr.PendingMatrixLinks[i+1:]...)

			usr.MatrixLinks = append(usr.MatrixLinks, &mxid)

			return nil
		}
	}

	return errors.New("not pending")
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
	usr.Password = nil
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
