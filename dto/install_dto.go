package dto

type CreateInstall struct {
	DeviceId string `json:"device_id" form:"device_id" binding:"required"`
}
