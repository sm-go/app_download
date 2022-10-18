package handler

import (
	"app-download/config"
	"app-download/dto"
	"app-download/middleware"
	"app-download/model"
	"app-download/repository"
	"app-download/service"
	"app-download/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authController struct {
	R            *gin.Engine
	tokenService *service.TokenService
	userRepo     *repository.UserRepository
}

func NewAuthController(h *Handler) *authController {
	return &authController{
		R:            h.R,
		tokenService: h.tokenService,
		userRepo:     h.userRepo,
	}
}

func (ctr *authController) Register() {
	group := ctr.R.Group("/api/admin")
	group.POST("/login", ctr.login)

	group.Use(middleware.AuthMiddleware(ctr.userRepo))
	group.POST("/refresh", ctr.refSecret)
	group.POST("/logout", ctr.logout)
}

func (ctr *authController) login(c *gin.Context) {
	req := &dto.RequestUserLogin{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
			Code: http.StatusUnauthorized,
			Msg:  err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	user, err := ctr.userRepo.FindByUsername(ctx, req.Username)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, &dto.ResponseObject{
			Code: 505,
			Msg:  "Not Found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
			Code: 401,
			Msg:  err.Error(),
		})
		return
	}

	ok, _ := utils.ComparePasswords(user.Password, req.Password)
	if !ok {
		c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
			Code: 401,
			Msg:  "Wrong Password",
		})
		return
	}

	pair, err := ctr.tokenService.GenerateTokenPairs(ctx, user, "")
	if err != nil {
		c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
			Code: http.StatusUnauthorized,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseObject{
		Code: http.StatusOK,
		Msg:  "Success",
		Data: pair,
	})
}

func (ctr *authController) refSecret(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	req := &dto.RequestUserRefreshToken{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
			Code: http.StatusUnauthorized,
			Msg:  err.Error(),
		})
		return
	}

	_, err := utils.ValidateUserRefreshToken(req.RefreshToken, config.RefreshSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
			Code: http.StatusUnauthorized,
			Msg:  err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	pair, err := ctr.tokenService.GenerateTokenPairs(ctx, user, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &dto.ResponseObject{
			Code: http.StatusUnauthorized,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseObject{
		Code: http.StatusOK,
		Msg:  "Success",
		Data: pair,
	})
}

func (ctr *authController) logout(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	c.JSON(http.StatusOK, &dto.ResponseObject{
		Code: http.StatusOK,
		Msg:  "Success",
		Data: &dto.ResponseUser{
			Id:        user.Id,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		},
	})
}
