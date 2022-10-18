package dto

type CreateDownloadRequest struct {
	AppId uint64 `json:"app_id" form:"app_id" binding:"required"`
}
