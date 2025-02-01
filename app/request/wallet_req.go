package request

import "time"

type WalletRequest struct {
	Name   string  `json:"name" binding:"required" example:"admin"`
	Token  string  `json:"token" binding:"required" example:"token"`
	Uuid   string  `json:"uuid" binding:"required" example:"uuid"`
	Value  float64 `json:"value" binding:"required" example:"1000000"`
	UserID uint    `json:"user_id" binding:"required" example:"1"`
}

type UpdateWallet struct {
	Name      string    `json:"name" binding:"required" example:"admin"`
	Token     string    `json:"token" binding:"required" example:"token"`
	Uuid      string    `json:"uuid" binding:"required" example:"uuid"`
	UpdatedAt time.Time `json:"-" swagger:"-"`
}
