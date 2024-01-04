package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"iam-server/api/v1/payload"
	"iam-server/db"
	"iam-server/models"
	"iam-server/res"
	"iam-server/token"
	"iam-server/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Authenticator interface {
	SignIn(email string, password string) (string, *token.Payload, error)
	SignUp(req *payload.LoginReq) (string, *token.Payload, error)
}

func Login(c echo.Context) error {
	var req payload.LoginReq
	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.NotOk(err.Error()))
	}

	var authUser models.AuthUser
	err = db.DB.Where("email = ?", req.Email).First(&authUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, res.NotOk(err.Error()))
	}

	result, err := utils.VldPassword(req.Password, authUser.Password)
	if !result || err != nil {
		return c.JSON(http.StatusInternalServerError, res.NotOk("Wrong Password"))
	}

	accessToken, accessPayload, err := token.NewToken(req.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.NotOk(err.Error()))
	}

	return c.JSON(http.StatusOK, res.Ok(&payload.LoginRes{
		Payload: accessPayload,
		Token:   accessToken,
	}))
}

func Register(c echo.Context) error {
	var req payload.RegisterReq
	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.NotOk(err.Error()))
	}

	var authUser models.AuthUser
	err = db.DB.First(&authUser).Where("email = ?", req.Email).Error
	fmt.Println(err)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, res.NotOk("User Exist"))
	}

	hashedPassword, _, err := utils.HashPassword(req.Password, utils.Default())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.NotOk(err.Error()))
	}

	authUser = models.AuthUser{
		UserId:   genUserId(req.Email),
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	}
	err = db.DB.Create(&authUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.NotOk(err.Error()))
	}

	accessToken, accessPayload, err := token.NewToken(req.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.NotOk(err.Error()))
	}

	return c.JSON(http.StatusOK, res.Ok(&payload.RegisterRes{
		Payload: accessPayload,
		Token:   accessToken,
	}))
}

func genUserId(email string) string {
	return strings.Split(email, "@")[0] + time.Now().Weekday().String()[0:1] + strconv.FormatInt(time.Now().Unix(), 10)
}
