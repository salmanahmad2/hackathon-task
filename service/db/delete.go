package db

import (
	db_errors "hackathon/service/db/error"

	"github.com/labstack/echo/v4"
)

func (db *NonFungibleTokenDBImpl) DeleteUserDB(c echo.Context, userId string) error {
	tx := db.conn.MustBegin()
	_, err := tx.Exec(`delete FROM "NFT".user where user_id=$1;`, userId)
	if err != nil {
		return db_errors.NewInternalServerError(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return db_errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (db *NonFungibleTokenDBImpl) DeleteTodoDB(c echo.Context, userId string, taskID string) error {
	tx := db.conn.MustBegin()
	_, err := tx.Exec(`delete FROM "Hackathon".todo where task_id=$1 and userId=$2;`, taskID, userId)
	if err != nil {
		return db_errors.NewInternalServerError(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return db_errors.NewInternalServerError(err.Error())
	}
	return nil
}
