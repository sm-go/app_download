package handler

import (
	"app-download/dto"
	"app-download/model"
	"app-download/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type DmController struct {
	R      *gin.Engine
	DmRepo *repository.DmRepository
}

func NewDmController(h *Handler) *DmController {
	return &DmController{
		R:      h.R,
		DmRepo: h.DmRepo,
	}
}

func (ctr *DmController) Register() {
	group := ctr.R.Group("/api/domain")

	group.GET("/all", ctr.all)
	group.POST("/create", ctr.create)
	group.POST("/update", ctr.update)
	group.POST("/delete", ctr.delete)
}

func (ctr *DmController) all(c *gin.Context) {
	ctx := c.Request.Context()
	data := ctr.DmRepo.All(ctx)
	c.JSON(http.StatusOK, &dto.ResponseObject{
		Code: 200,
		Msg:  "Success",
		Data: data,
	})
}

func (ctr *DmController) create(c *gin.Context) {
	req := &dto.CreateDomain{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	dm := &model.Domain{}
	if err := smapping.FillStruct(dm, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}

	ctx := c.Request.Context()
	data, err := ctr.DmRepo.CreateDm(ctx, dm)
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

func (ctr *DmController) update(c *gin.Context) {
	req := &dto.UpdateDomain{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}
	dm := &model.Domain{}
	if err := smapping.FillStruct(dm, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}
	ctx := c.Request.Context()
	data, err := ctr.DmRepo.UpdateDomain(ctx, dm)
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

func (ctr *DmController) delete(c *gin.Context) {
	req := &dto.DeleteDomain{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseObject{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}
	dm := &model.Domain{}
	if err := smapping.FillStruct(dm, smapping.MapFields(req)); err != nil {
		panic(err.Error())
	}

	ctx := c.Request.Context()
	err := ctr.DmRepo.DeleteDomain(ctx, dm)
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
