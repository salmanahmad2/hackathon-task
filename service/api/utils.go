package api

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"hackathon/pkg/utils"
	api_errors "hackathon/service/api/error"
	db_errors "hackathon/service/db/error"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func handleAPIError(err error) error {
	log.Println("error is :", err)
	switch err {
	case &db_errors.InternalServerError{}:
		return api_errors.NewInternalServerError(err.Error())
	case &db_errors.NotFound{}:
		log.Println("i am in handle api error")
		return api_errors.NewUnProcessableRequest(err.Error())
	default:
		return &api_errors.InternalServerError{}
	}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", api_errors.NewInternalServerError(err.Error())
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return api_errors.NewUnauthorizedError(err.Error())
	}
	return nil
}

func CreateUUID() string {
	id := uuid.New()
	return id.String()
}

func SetExpTime(duration time.Duration) int64 {
	clock := utils.NewClock()
	addTime := clock.NowUTC().Add(duration)
	expTimeUnix := utils.ToUnix(&addTime)
	return *expTimeUnix / 1000
}

func GenerateCode(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(max-min+1) + min
	return code
}

func ValidateBidPrice(currentBid float64, newBid float64) error {
	minIncrease := currentBid * 0.05
	minBid := currentBid + minIncrease
	if currentBid < newBid {
		return nil
	}
	return fmt.Errorf("bid must be atleast %v", minBid)
}
func (api *NonFungibleTokenAPIImpl) IsAuctionListed(auctionCloseTime int64) error {
	currentTime := api.clock.NowUnix()
	if auctionCloseTime > currentTime {
		return nil
	}
	return api_errors.NewUnProcessableRequest("you can not update the auction now")
}
func (api *NonFungibleTokenAPIImpl) ValidateAdditionalTime(auctionCloseTime int64) int64 {

	currentTime := api.clock.NowUnix()
	checkAddTime := api.clock.NowUTC().Add(time.Minute * 15)
	checkAddTimeUnix := utils.ToUnix(&checkAddTime)
	if auctionRemainingTime := auctionCloseTime - currentTime; auctionRemainingTime > *checkAddTimeUnix {
		return auctionCloseTime
	} else {
		auctionCloseTime = auctionCloseTime + *checkAddTimeUnix
		return auctionCloseTime
	}
}
func (api *NonFungibleTokenAPIImpl) CheckAuctionValidation(auctionStartTime int64, auctionCloseTime int64) bool {
	currentTime := api.clock.NowUnix()
	if auctionStartTime < currentTime && auctionCloseTime > currentTime {
		return true
	}
	return false
}
