package db

import (
	db_errors "hackathon/service/db/error"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (db *NonFungibleTokenDBImpl) GetUserDB(c echo.Context, user *models.User) (*models.User, error) {
	userData := models.NewUser()
	rows, err := db.conn.NamedQuery(`SELECT * FROM "NFT".user WHERE username=:username or email=:email;`, user)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&userData)
		return userData, err
	}
	return nil, err
}

func (db *NonFungibleTokenDBImpl) GetAllUsersDB(c echo.Context) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := db.conn.Select(&users, `SELECT * FROM "NFT".user;`)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return users, nil
}

func (db *NonFungibleTokenDBImpl) GetProfileDB(c echo.Context, userId string) (*models.User, error) {
	userData := models.NewUser()
	err := db.conn.Get(userData, `SELECT * FROM "NFT".user WHERE user_id=$1;`, userId)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return userData, nil
}

func (db *NonFungibleTokenDBImpl) GetTodoListDB(c echo.Context, userId string) ([]*models.Todo, error) {
	listData := make([]*models.Todo, 0)
	err := db.conn.Select(&listData, `SELECT taskId,title,description,status FROM "Hackathon".todo WHERE user_id=$1;`, userId)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return listData, nil
}

func (db *NonFungibleTokenDBImpl) GetTodoTaskFilteredDB(c echo.Context, userId string, status string) ([]*models.Todo, error) {
	listData := make([]*models.Todo, 0)
	err := db.conn.Select(&listData, `SELECT taskId,title,description,status FROM "Hackathon".todo WHERE user_id=$1 and status=$2;`, userId, status)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return listData, nil
}
