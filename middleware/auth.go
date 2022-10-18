package middleware

import (
	"app-download/config"
	"app-download/dto"
	"app-download/repository"
	"app-download/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AccessToken string `header:"Authorization"`
}

func AuthMiddleware(r *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
				Code: http.StatusUnauthorized,
				Msg:  "Must provide `Authorization` header in format of `Bearer {token}`",
			})
			c.Abort()
			return
		}

		accessToken := strings.Split(h.AccessToken, "Bearer ")

		if len(accessToken) != 2 {
			c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
				Code: http.StatusUnauthorized,
				Msg:  "Must provide `Authorization` header in format of `Bearer {token}`",
			})
			c.Abort()
			return
		}

		accessTokenClaim, err := utils.ValidateUserAccessToken(accessToken[1], config.PublicKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
				Code: http.StatusUnauthorized,
				Msg:  "Permission denied",
			})
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		user, err := r.FindByUserId(ctx, accessTokenClaim.User.Id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
				Code: http.StatusUnauthorized,
				Msg:  "Permission denied",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()

	}
}
