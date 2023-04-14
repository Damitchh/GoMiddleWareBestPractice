package controllers

import (
	"Hacktiv10JWT/database"
	"Hacktiv10JWT/helpers"
	"Hacktiv10JWT/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	appJSON = "application/JSON"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	if err := database.PingDB(db); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err":     "Internal Server Error",
			"message": "Failed to ping database",
		})
		return
	}
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	var NewUser models.User

	if contentType == appJSON {
		ctx.ShouldBindJSON(&NewUser)
	} else {
		ctx.ShouldBind(&NewUser)
	}

	err := db.Debug().Create(&NewUser).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        NewUser.ID,
		"email":     NewUser.Email,
		"full_name": NewUser.Fullname,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	User := models.User{}
	password := ""

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}
	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Role)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
