package dto

type CreateApp struct {
	Title       string `json:"title" form:"title"`
	Version     string `json:"version" form:"version" binding:"required"`
	Description string `json:"description" form:"description"`
	Size        string `json:"size" form:"size" binding:"required"`
	OsType      string `json:"os_type" form:"os_type" binding:"required"`
	DomainId    uint64 `json:"domain_id" form:"domain_id" binding:"required"`
}

type UpdateApp struct {
	Id          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" `
	Version     string `json:"version" form:"version"`
	Description string `json:"description" form:"description"`
	Size        string `json:"size" form:"size"`
	OsType      string `json:"os_type" form:"os_type"`
	DomainId    uint64 `json:"domain_id" form:"domain_id" `
}

type DeleteApp struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
}
