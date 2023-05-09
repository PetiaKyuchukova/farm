package handlers

import (
	"farm/backend/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const SecretKey = "secret"

type UserHandler interface {
	Register(ctx *gin.Context)
}
type defaultUserHandler struct {
}

func NewUserHandler() UserHandler {
	return &defaultUserHandler{}
}

func (uh *defaultUserHandler) Register(ctx *gin.Context) {
	var data map[string]string
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(data)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	id := uuid.New()

	user := domain.Farmer{
		ID:       id,
		Name:     data["name"],
		Username: data["username"],
		Email:    data["email"],
		Password: string(password),
	}

	fmt.Println(user)
	//database.DB.Create(&user)

	ctx.JSON(http.StatusOK, data)
}

//func Login(c *gin.Ctx) error {
//	var data map[string]string
//
//	if err := c.BodyParser(&data); err != nil {
//		return err
//	}
//
//	var user models.User
//
//	database.DB.Where("email = ?", data["email"]).First(&user)
//
//	if user.Id == 0 {
//		c.Status(fiber.StatusNotFound)
//		return c.JSON(fiber.Map{
//			"message": "user not found",
//		})
//	}
//
//	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
//		c.Status(fiber.StatusBadRequest)
//		return c.JSON(fiber.Map{
//			"message": "incorrect password",
//		})
//	}
//
//	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
//		Issuer:    strconv.Itoa(int(user.Id)),
//		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
//	})
//
//	token, err := claims.SignedString([]byte(SecretKey))
//
//	if err != nil {
//		c.Status(fiber.StatusInternalServerError)
//		return c.JSON(fiber.Map{
//			"message": "could not login",
//		})
//	}
//
//	cookie := fiber.Cookie{
//		Name:     "jwt",
//		Value:    token,
//		Expires:  time.Now().Add(time.Hour * 24),
//		HTTPOnly: true,
//	}
//
//	c.Cookie(&cookie)
//
//	return c.JSON(fiber.Map{
//		"message": "success",
//	})
//}
//
//func User(c *gin.Ctx) error {
//	cookie := c.Cookies("jwt")
//
//	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(SecretKey), nil
//	})
//
//	if err != nil {
//		c.Status(fiber.StatusUnauthorized)
//		return c.JSON(fiber.Map{
//			"message": "unauthenticated",
//		})
//	}
//
//	claims := token.Claims.(*jwt.StandardClaims)
//
//	var user models.User
//
//	database.DB.Where("id = ?", claims.Issuer).First(&user)
//
//	return c.JSON(user)
//}
//
//func Logout(c *fiber.Ctx) error {
//	cookie := fiber.Cookie{
//		Name:     "jwt",
//		Value:    "",
//		Expires:  time.Now().Add(-time.Hour),
//		HTTPOnly: true,
//	}
//
//	c.Cookie(&cookie)
//
//	return c.JSON(fiber.Map{
//		"message": "success",
//	})
//}
