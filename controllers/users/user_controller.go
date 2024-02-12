package users

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"anggi.blog/domain/users"
	"anggi.blog/services"
	"anggi.blog/utils/errors"
	"anggi.blog/utils/logs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	SecretKey = os.Getenv("SECRET_KEY")
)

func Register(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Error when trying to bind JSON", "Controller")
		logs.Error.Println(err)
		c.IndentedJSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		logs.Error.Println(saveErr)
		c.IndentedJSON(saveErr.Status, saveErr)
		return
	}

	resultWp := &users.User{ID: result.ID, Username: result.Username, Email: result.Email}
	c.IndentedJSON(http.StatusCreated, resultWp)
}

func Login(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("Invalid JSON body", "Controller")
		logs.Error.Println(err)
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		logs.Error.Println(getErr)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		err := errors.NewInternalServerError("Error when trying to generate token", "Controller")
		logs.Error.Println(err)
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerError("Error when trying to get cookie", "Controller")
		logs.Error.Println(getErr)
		c.JSON(getErr.Status, getErr)
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		restErr := errors.NewInternalServerError("Error when trying to parse token", "Controller")
		c.JSON(restErr.Status, restErr)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid token", "Controller")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
