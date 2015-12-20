package auth

import (
	"errors"

	"github.com/dimorinny/echo-jwt/jwt"
	"github.com/dimorinny/queuer/src/database"
	"github.com/dimorinny/queuer/src/response"
	"github.com/dimorinny/queuer/src/util"
	"github.com/labstack/echo"
)

const (
	secret        = "SuperSecret"
	usernameField = "email"
)

var (
	Config jwt.Config
	Jwt    jwt.Jwt
)

func Init() {
	initConfig()
	initJwt()
}

func initConfig() {
	Config = jwt.NewConfig(secret)
	Config.UsernameField = usernameField

	Config.AuthErrorHandler = response.AuthErrorHandler
	Config.HeaderInvalidHandler = response.HeaderInvalidHandler
	Config.TokenInvalidHandler = response.TokenInvalidHandler
	Config.TokenExpireHandler = response.TokenExpireHandler

	Config.RefreshResponseHandler = response.RefreshResponseHandler
	Config.LoginResponseHandler = response.LoginResponseHandler
}

func initJwt() {
	Jwt = jwt.NewJwt(Config, Authenticate, Identity)
}

// Return identity for login user
func Authenticate(email string, password string) interface{} {
	user := database.User{}

	if err := database.Db.Where(&database.User{Email: email}).First(&user).Error; err != nil {
		return nil
	}

	if !isPasswordValid(password, user.Password) {
		return nil
	}

	return user.ID
}

// Return user by id
func Identity(val interface{}) interface{} {
	user := database.User{}
	database.Db.First(&user, int(val.(float64)))
	return user
}

func AdminRequired(c *echo.Context) error {
	user := c.Get(Config.IdentityKey).(database.User)

	if !user.IsSuperAdmin {
		response.AdminRequiredHandler(c)
		return errors.New("User not admin")
	}

	return nil
}

func Register(c *echo.Context) error {
	// TODO: validations
	email := c.Form("email")
	firstName := c.Form("firstName")
	lastName := c.Form("lastName")
	password := c.Form("password")

	if !util.IsAllStringsNotEmpty(email, firstName, lastName, password) {
		response.RegisterParamsError(c)
		return nil
	}

	user := database.User{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		Password:     encryptPassword(password),
		IsSuperAdmin: false,
	}

	if database.Db.Create(&user).Error != nil {
		response.RegisterAlreadyExistsError(c)
		return nil
	}

	access, _ := Jwt.GenerateAccessToken(user.ID)
	refresh, _ := Jwt.GenerateRefreshToken(user.ID)

	response.LoginResponseHandler(c, access, refresh)
	return nil
}
