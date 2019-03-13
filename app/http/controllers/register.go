package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/config"
	"github.com/totoval/framework/helpers"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/model"
	"github.com/totoval/framework/utils/crypt"
	"github.com/totoval/framework/utils/jwt"
	"net/http"
	"totoval/app/http/requests"
	"totoval/app/models"
)

type Register struct {
	controller.BaseController
}

func (r *Register) Register(c *gin.Context) {
	// validate and assign requestData
	var requestData requests.UserRegister
	if !r.Validate(c, &requestData) {
		return
	}

	// determine if exist
	user := models.User{
		Email:    &requestData.Email,
	}
	if model.H.Exist(&user, true) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": helpers.L(c, "auth.register.failed_existed")})
		return
	}

	// create user
	// encrypt password //@todo move to model setter later
	encryptedPassword := crypt.Bcrypt(requestData.Password)
	user.Password = &encryptedPassword
	if err := model.H.Create(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error":  helpers.L(c, "auth.register.failed_system_error")})
		return
	}

	// create jwt
	newJwt := jwt.NewJWT(config.GetString("auth.sign_key"))
	username := ""
	if user.Name != nil{
		username = *user.Name
	}
	if token, err := newJwt.CreateToken(string(*user.ID), username); err == nil{
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{"error":  helpers.L(c, "auth.register.failed_token_generate_error")})
	return
}
