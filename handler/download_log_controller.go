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

type downloadlogController struct {
	R            *gin.Engine
	downloadRepo *repository.DownloadLogRepo
}

func NewDownloadLogController(h *Handler) *downloadlogController {
	return &downloadlogController{
		R:            h.R,
		downloadRepo: h.downloadlogRepo,
	}
}

func (ctr *downloadlogController) Register() {
	group := ctr.R.Group("/api/download/log")

	group.GET("/all", ctr.all)
	group.POST("/create", ctr.create)
}

func (ctr *downloadlogController) all(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var downloadlogs model.DownloadLog
	logs, err := ctr.downloadRepo.FindAll(&downloadlogs, &pagination)

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

func (ctr *downloadlogController) create(c *gin.Context) {
	req := &dto.CreateDownloadRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	download := &model.DownloadLog{}
	if err := smapping.FillStruct(download, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}
	clientIpAddr := utils.GetClientIP(c)
	download.IpAddress = clientIpAddr

	ctx := c.Request.Context()
	data, err := ctr.downloadRepo.Create(ctx, download)
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
