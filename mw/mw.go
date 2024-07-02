package mw

import (
	"fmt"
	"net/http"

	"github.com/eylulkadioglu/Music/models"
	"github.com/eylulkadioglu/Music/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Checks the authorization header and JWT token
func CheckAuthorization(ctx *gin.Context) {
	authorizationHeader := ctx.Request.Header.Get("Authorization")

	if authorizationHeader == "" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  "error",
				"message": "Authorization is required!",
				"code":    "001",
			},
		)
		return
	}

	authorizationHeader = authorizationHeader[7:]
	fmt.Printf("Got token: [%s]\n", authorizationHeader)

	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(
		authorizationHeader,
		claims,
		// Validates the token 
		func(token *jwt.Token) (any, error) {
			return utils.GetJwtKey(), nil
		},
	)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  "error",
				"message": "Can't parse token!",
				"error":   err.Error(),
				"code":    "002",
			},
		)
		return
	}

	if !tkn.Valid {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  "error",
				"message": "Invalid token!",
				"code":    "003",
			},
		)
		return
	}

	fmt.Printf("Request from %v\n", claims)

	ctx.Next()
}
