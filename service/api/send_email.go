package api

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"

	"hackathon/config"
	api_errors "hackathon/service/api/error"
	"hackathon/service/cache"
	"hackathon/service/models"

	"github.com/badoux/checkmail"
	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) SendEmailForSignUpAPI(c echo.Context, user *models.User) error {
	emailErr := api.CheckEmailExist(*user.Email)
	if emailErr != nil {
		return api_errors.NewUnProcessableRequest("Email does not exist")
	}
	code := GenerateCode(1000, 9999)
	signupMessage := "Welcome to myBidify\nYou're all set. Now you will be able to buy & sell NFT's in the World leading NFT marketplace.\nPlease confirm your registration\nHere is your OTP:%d\nNever share this OTP with anyone.\nOur Best,\nTeam MyBidify"
	err := api.SendMailToUser(user, signupMessage, code, false)
	if err != nil {
		return err
	}
	err = api.cache.SaveToCache(c, cache.UserSignupKey+*user.Email, []byte(fmt.Sprint(code)), cache.UserSignupKey_Exp)
	if err != nil {
		return api_errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (api *NonFungibleTokenAPIImpl) SendEmailForForgotPasswordAPI(c echo.Context, user *models.User) error {
	userDB, err := api.db.GetUserDB(c, user)
	if err != nil {
		return err
	}
	if userDB == nil {
		return api_errors.NewError("user not found")
	}
	token := GenerateCode(100000, 999999)
	forgotPasswordMessage := "Welcome to myBidify\nWe received a request to reset your myBidify Password.\nEnter the following password reset token:%d\nAlternatively, you can directly change your password.\nhttp://localhost:5002/swagger/index.html/default/post_forgotpassword?token=%d\nTeam MyBidify"
	err = api.SendMailToUser(user, forgotPasswordMessage, token, true)
	if err != nil {
		return err
	}
	err = api.cache.SaveToCache(c, cache.ForgetPasswordKey+fmt.Sprint(token), []byte(*user.Email), cache.ForgetPasswordKey_Exp)
	if err != nil {
		return api_errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (api *NonFungibleTokenAPIImpl) SendMailToUser(user *models.User, emailMessage string, code int, messageFlag bool) error {
	smtpConfig := config.Cfg.SMTP
	to := *user.Email
	var subject string
	var body []byte
	if messageFlag {
		subject = "[mybidify] Request For Forgot Password"
		body = []byte(fmt.Sprintf(emailMessage, code, code))
	} else {
		subject = "[mybidify] Confirm Your Registration"
		body = []byte(fmt.Sprintf(emailMessage, code))
	}
	message := ComposeMail(to, smtpConfig.MailAddress, subject, body)
	auth := smtp.PlainAuth("", smtpConfig.MailAddress, smtpConfig.Password, smtpConfig.SmtpHost)
	err := smtp.SendMail(smtpConfig.SmtpHost+":"+smtpConfig.SmtpPort, auth, smtpConfig.MailAddress, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}

func ComposeMail(to string, from string, subject string, body []byte) []byte {
	header := make(map[string]string)
	header["From"] = FormatEmailAddress(from)
	header["To"] = FormatEmailAddress(to)
	header["Subject"] = EncodeRFC2047(subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return []byte(message)
}

func FormatEmailAddress(addr string) string {
	e, err := mail.ParseAddress(addr)
	if err != nil {
		return addr
	}
	return e.String()
}

func EncodeRFC2047(str string) string {
	addr := mail.Address{Address: str}
	return strings.Trim(addr.String(), " <>")
}

func (api *NonFungibleTokenAPIImpl) CheckEmailExist(userMail string) error {
	smtpConfig := config.Cfg.SMTP
	err := checkmail.ValidateHostAndUser(smtpConfig.HostName, smtpConfig.MailAddress, userMail)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		return smtpErr
	}
	return nil
}
