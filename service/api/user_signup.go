package api

import (
	api_errors "hackathon/service/api/error"
	"hackathon/service/cache"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) SignUpUserAPI(c echo.Context, user *models.User, signup *models.SignUpRequest) error {
	if *signup.OTP == "" {
		userDB, err := api.db.GetUserDB(c, user)
		if userDB == nil {
			err := api.SendEmailForSignUpAPI(c, user)
			if err != nil {
				return api_errors.NewInternalServerError(err.Error())
			}
			return nil
		} else if len(*userDB.Email) > 0 && *userDB.Email == *user.Email {
			return api_errors.NewError("email already exists")
		}
		if err != nil {
			return api_errors.NewInternalServerError(err.Error())
		}
	} else {
		err := api.ConfirmOTPAPI(c, user, signup)
		if err != nil {
			return err
		}
		hashPassword, err := HashPassword(*user.Password)
		*user.Password = hashPassword
		if err != nil {
			return err
		}
		*user.UserId = CreateUUID()
		err = api.db.SignupAPIDB(c, user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *NonFungibleTokenAPIImpl) ConfirmOTPAPI(c echo.Context, user *models.User, signup *models.SignUpRequest) error {
	cacheData, cacheErr := api.cache.CheckCache(c, cache.UserSignupKey+*user.Email)
	if cacheErr != nil {
		return api_errors.NewInternalServerError(cacheErr.Error())
	}
	cacheOTP := string(cacheData)
	if *signup.OTP == cacheOTP {
		api.cache.DeleteCache(c, cache.UserSignupKey+*user.Email)
		return nil
	}
	return api_errors.NewUnauthorizedError("Invalid OTP")
}
