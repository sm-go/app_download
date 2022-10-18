package dto

type CreateDomain struct {
	Name        string `json:"name" form:"name" binding:"required"`
	DownloadUrl string `json:"download_url" form:"download_url" binding:"required"`
	Description string `json:"description" form:"description"`
}

type UpdateDomain struct {
	Id          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name"`
	DownloadUrl string `json:"download_url" form:"download_url"`
	Description string `json:"description" form:"description"`
}

type DeleteDomain struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
}
