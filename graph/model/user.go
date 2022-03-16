package model

import model "github.com/Unkn0wnCat/matrix-veles/internal/db/model"

type User struct {
	ID                 string    `json:"id"`
	Username           string    `json:"username"`
	Admin              *bool     `json:"admin"`
	MatrixLinks        []*string `json:"matrixLinks"`
	PendingMatrixLinks []*string `json:"pendingMatrixLinks"`
}

func MakeUser(dbUser *model.DBUser) *User {
	return &User{
		ID:                 dbUser.ID.Hex(),
		Username:           dbUser.Username,
		Admin:              dbUser.Admin,
		MatrixLinks:        dbUser.MatrixLinks,
		PendingMatrixLinks: dbUser.PendingMatrixLinks,
	}
}
