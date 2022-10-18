package handler

import (
	"app-download/dto"
	"app-download/model"
	"app-download/repository"
	"app-download/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type appController struct {
	R       *gin.Engine
	appRepo *repository.AppRepository
}

func NewAppController(h *Handler) *appController {
	return &appController{
		R:       h.R,
		appRepo: h.appRepo,
	}
}

func (ctr *appController) Register() {
	group := ctr.R.Group("/api/app")

	group.GET("/all", ctr.all)
	group.POST("/create", ctr.create)
	group.POST("/update", ctr.update)
	group.POST("/delete", ctr.delete)
}

func (ctr *appController) all(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var applists model.App
	lists, err := ctr.appRepo.FindAll(&applists, &pagination)
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
		Data: lists,
	})
}

func (ctr *appController) create(c *gin.Context) {
	req := &dto.CreateApp{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	err = c.SaveUploadedFile(file, "upload/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	app := &model.App{}
	if err := smapping.FillStruct(app, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}

	ctx := c.Request.Context()
	data, err := ctr.appRepo.Create(ctx, app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseObject{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseObject{
		Code: 200,
		Msg:  "Upload App Success",
		Data: data,
	})
}

func (ctr *appController) update(c *gin.Context) {
	req := &dto.UpdateApp{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	app := &model.App{}
	if err := smapping.FillStruct(app, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}

	ctx := c.Request.Context()
	data, err := ctr.appRepo.UpdateApp(ctx, app)
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

func (ctr *appController) delete(c *gin.Context) {
	req := &dto.DeleteApp{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	app := &model.App{}
	if err := smapping.FillStruct(app, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}

	ctx := c.Request.Context()
	err := ctr.appRepo.DeleteApp(ctx, app)
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
	})

}
