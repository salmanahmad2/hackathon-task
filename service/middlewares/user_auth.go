package middlewares

import (
	"io/ioutil"
	"log"

	"hackathon/pkg/utils"
	server_errors "hackathon/service/server/errors"

	"hackathon/service/models"

	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	PblicKeyPath = "config/mypubkey.pem"
	PvtKeyPath   = "config/mykey.pem"
)

type JWT struct {
	privateKey []byte
	publicKey  []byte
}

func UserAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		TokenString := c.Request().Header.Get("Authorization")
		if TokenString == "" {
			return server_errors.NewUnauthorizedError("You need to login first")
		}
		pubkey, err := ioutil.ReadFile(PblicKeyPath)
		if err != nil {
			log.Fatalf("Public path not found %v", err)
		}
		claims, err := ValidateToken(TokenString, pubkey)
		if err != nil {
			return server_errors.NewUnauthorizedError(err.Error())
		}
		user := models.NewUser()

		*user.UserId = claims["user_id"].(string)
		c.Set("user-id", *user.UserId)
		return next(c)

	}
}

func ValidateToken(token string, j []byte) (jwt.MapClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j)
	if err != nil {
		return nil, err
	}
	tok, validErr := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return key, nil
	})
	v, _ := validErr.(*jwt.ValidationError)
	if err != nil && v.Errors != jwt.ValidationErrorExpired {
		return nil, err
	}
	claims, ok := tok.Claims.(jwt.MapClaims)

	if (!ok) || (!tok.Valid && v.Errors != jwt.ValidationErrorExpired) {
		return nil, err
	}
	return claims, validErr
}

func NewJWT(private, public string) *JWT {
	privateKey, err := ioutil.ReadFile(private)
	if err != nil {
		log.Fatalf("Private key path not found %v", err)
	}
	publicKey, err := ioutil.ReadFile(public)
	if err != nil {
		log.Fatalf("Public path not found %v", err)
	}
	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}
func (j JWT) GenerateToken(user *models.User, expTime int64) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", err
	}

	now := utils.NewClock().NowUnix()
	currentTime := now / 1000
	claims := make(jwt.MapClaims)
	claims["user_id"] = user.UserId
	claims["exp"] = expTime                    // The expiration time after which the token must be disregarded.
	claims["iat"] = currentTime                // The time at which the token was issued.
	claims["duration"] = expTime - currentTime // The time duration of token.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j JWT) GetPublicKey() []byte {
	return j.publicKey
}
