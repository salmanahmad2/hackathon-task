package db

import (
	db_errors "hackathon/service/db/error"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (db *NonFungibleTokenDBImpl) UpdatePasswordDB(c echo.Context, user *models.User) error {
	tx := db.conn.MustBegin()
	_, err := tx.NamedQuery(`UPDATE "NFT".user set password=:password WHERE user_id=:user_id`, user)
	if err != nil {
		return db_errors.NewInternalServerError(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return db_errors.NewInternalServerError("unable to update password")
	}
	return nil
}

func (db *NonFungibleTokenDBImpl) UpdateTodoDB(c echo.Context, todo *models.Todo) error {
	tx := db.conn.MustBegin()
	_, err := tx.NamedQuery(`UPDATE "Hackathon".todo set title=:title,discription=:description,status=:status WHERE user_id=:user_id`, todo)
	if err != nil {
		return db_errors.NewInternalServerError(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return db_errors.NewInternalServerError("unable to update todo")
	}
	return nil
}
