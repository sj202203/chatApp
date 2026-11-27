package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandlerImpl struct {
	status int
	db     *gorm.DB
}

func NewHandlerImpl(status int, db *gorm.DB) HealthHandlerImpl {
	return HealthHandlerImpl{status: status, db: db}
}

func (h HealthHandlerImpl) GetStatus(ctx *gin.Context) {
	if h.status == 200 {
		ctx.JSON(h.status, gin.H{"message": "ok"})
	} else {
		ctx.JSON(h.status, gin.H{"Some issue": "not ok"})
	}
}