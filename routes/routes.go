package routes

import (
	"fmt"
	"net/http"

	"github.com/eylulkadioglu/Music/db"
	"github.com/eylulkadioglu/Music/mailer"
	"github.com/eylulkadioglu/Music/models"
	"github.com/eylulkadioglu/Music/utils"
	"github.com/gin-gonic/gin"
)

// Change the email field below for an accurate test 
func Landing(ctx *gin.Context) {
	go mailer.SendMail("your_email@test.com", "Test Title", "Hello World!")

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "welcome",
		},
	)
}

func GetArtists(ctx *gin.Context) {
	artists, err := db.GetArtists()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"artists": artists,
		},
	)

}

func CreateArtist(ctx *gin.Context) {
	var artist models.Artist

	err := ctx.BindJSON(&artist)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}
	err = db.CreateArtist(artist)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "artist created successfully",
		},
	)
}

func Login(ctx *gin.Context) {
	var loginRequest models.User

	err := ctx.BindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": "Username and password are mandatory fields!",
			},
		)
		return
	}

	ok, user := db.CheckLogin(loginRequest)
	if !ok {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": "Invalid password!",
			},
		)
		return
	}

	tokenString, err := utils.GetJwtToken(user)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "error",
				"message": "Can't login at this time",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Login succesfull!",
			"token":   tokenString,
		},
	)
}

func DeleteArtist(ctx *gin.Context) {
	var artist models.Artist

	err := ctx.BindJSON(&artist)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}

	err = db.DeleteArtist(artist)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "artist deleted successfully",
		},
	)
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}

	err = db.CreateUser(user)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "user created successfully",
		},
	)
}

func LostPassword(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}

	if user.Email == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": "Email field is empty!",
			},
		)
	}

	err = db.CheckUser(user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": "User not exists!",
			},
		)
	}

	// Generates a password reset code and assign it to the user
	code := utils.GetCode()
	db.CreatePasswordCode(user, code)

	// Sends the password reset code via email
	body := fmt.Sprintf("Your code for to change your password is: %s\n", code)
	err = mailer.SendMail(user.Email, "Lost Password Code", body)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "error",
				"message": "Email couldn't send!",
			},
		)
		db.DeletePasswordCode(user)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Code send successfully!",
		},
	)

}

func ChangePassword(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)
		return
	}

	if user.Email == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": "Email field is empty!",
			},
		)
	}

	if user.Code == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": "Password reset code field is empty!",
			},
		)
	}

	if user.Password == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "error",
				"message": "Password reset code field is empty!",
			},
		)
	}

	//Checks the give code and then changes the password
	db.CheckCode(user, user.Code)
	db.ChangePasswordWithCode(user)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Password changed successfully!",
		},
	)

}
