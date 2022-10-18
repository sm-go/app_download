package handler

import (
	"app-download/dto"
	"app-download/model"
	"app-download/repository"
	"app-download/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type installLogController struct {
	R           *gin.Engine
	installRepo *repository.InstallLogRepository
}

func NewInstallLogController(h *Handler) *installLogController {
	return &installLogController{
		R:           h.R,
		installRepo: h.installlogRepo,
	}
}

func (ctr *installLogController) Register() {
	group := ctr.R.Group("/api/install/log")

	group.GET("/all", ctr.all)
	group.POST("/create", ctr.create)
}

func (ctr *installLogController) all(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var install_logs model.InstallLog
	logs, err := ctr.installRepo.FindAll(&install_logs, &pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseObject{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &dto.ResponseObject{
		Code: 200,
		Msg:  "Success",
		Data: logs,
	})
}

func (ctr *installLogController) create(c *gin.Context) {
	req := &dto.CreateInstall{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	install := &model.InstallLog{}
	if err := smapping.FillStruct(install, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}

	ctx := c.Request.Context()
	data, err := ctr.installRepo.Create(ctx, install)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseObject{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseObject{
		Code: 200,
		Msg:  "Success",
		Data: data,
	})

}
