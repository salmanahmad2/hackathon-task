package db

import (
	"log"

	db_errors "hackathon/service/db/error"

	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (db *NonFungibleTokenDBImpl) SignupAPIDB(c echo.Context, user *models.User) error {
	tx := db.conn.MustBegin()
	_, err := tx.NamedQuery(`INSERT INTO "Hackathon".user(user_id,name,password,email) 
		 VALUES(:user_id,:username,:password,:email)`, user)
	if err != nil {
		log.Println(err)
		return db_errors.NewInternalServerError("Unable to insert user data")
	}
	err = tx.Commit()
	if err != nil {
		return db_errors.NewInternalServerError(err.Error())
	}
	return nil
}
