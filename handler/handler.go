package handler

import (
	"app-download/ds"
	"app-download/repository"
	"app-download/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	R *gin.Engine

	tokenService    *service.TokenService
	userRepo        *repository.UserRepository
	downloadlogRepo *repository.DownloadLogRepo
	installlogRepo  *repository.InstallLogRepository
	appRepo         *repository.AppRepository
	DmRepo          *repository.DmRepository
}

type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {

	// token Repository
	tokenRepo := repository.NewTokenRepository(c.DS)
	tokenService := service.NewTokenService(&service.TSConfig{
		TokenRepo: tokenRepo,
	})

	// User Repository
	userRepo := repository.NewUserRepository(c.DS)

	// Download Log Repository
	downloadRepo := repository.NewDownloadLogRepository(c.DS)

	// Install Log Repository
	installRepo := repository.NewInstallLogRepository(c.DS)

	// App Repository
	appRepo := repository.NewAppRepository(c.DS)

	// Dm Repository
	dmRepo := repository.NewDmRepository(c.DS)

	return &Handler{
		R: c.R,

		tokenService:    tokenService,
		userRepo:        userRepo,
		downloadlogRepo: downloadRepo,
		installlogRepo:  installRepo,
		appRepo:         appRepo,
		DmRepo:          dmRepo,
	}
}

func (h *Handler) Register() {
	authController := NewAuthController(h)
	authController.Register()

	//domain controller
	dmController := NewDmController(h)
	dmController.Register()

	//apps controller
	appController := NewAppController(h)
	appController.Register()

	//install log controller
	installLogController := NewInstallLogController(h)
	installLogController.Register()

	//download log controller
	downloadlogController := NewDownloadLogController(h)
	downloadlogController.Register()
}
