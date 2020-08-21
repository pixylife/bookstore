package controller

import (
	env "bookstore/pkg/env"
	model "bookstore/pkg/model"
)

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

type Payload struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberme"`
}
type JwtCustomClaims struct {
	Sub  string
	Auth string

	jwt.StandardClaims
}

func LoginController(g *echo.Echo, contextRoot string) {

	g.POST(contextRoot+"/authenticate", login)
}
func login(ctx context.Context) error {

	loginRequest := new(Payload)
	er1 := c.Bind(loginRequest)

	if er1 != nil {
		return er1
	}
	db := env.RDB
	users := []*model.User{}
	db.Find(&users)
	var targetUser *model.User
	if len(users) == 0 {
		targetUser = &model.User{Username: "admin", Password: "admin", Firstname: "", Lastname: ""}
		db.Create(&targetUser)
	}
	for _, user := range users {
		if strings.EqualFold(user.Username, loginRequest.Username) && strings.EqualFold(user.Password, loginRequest.Password) {
			targetUser = user
		}
	}
	if targetUser == nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	claims := JwtCustomClaims{"admin", "ROLE_ADMIN,ROLE_USER", jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	c.Response().Header().Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", t))
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Origin, X-Requested-With, Content-Type, Accept")

	if err != nil {
		return err
	}
	targetUser.Token = t
	return c.JSON(http.StatusOK, targetUser)
}
