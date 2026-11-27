package handler

import (
	"fmt"
	"prob13/internal/core/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type chanRelay struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type ChanRelayHandlerImpl struct {
	log   func(c *gin.Context, code int, message string)
	db    *gorm.DB
	chann chan entity.Message
}

func NewChannelHandlerImpl(db *gorm.DB, chann chan entity.Message) ChanRelayHandlerImpl {
	return ChanRelayHandlerImpl{log: respondError, db: db, chann: chann}
}

func (ch ChanRelayHandlerImpl) RelayMessage(c *gin.Context) {
	var messageRequest chanRelay
	err := c.ShouldBindJSON(&messageRequest)
	if err != nil {
		ch.log(c, 400, "malformed payload")
		return
	}

	var userGroup entity.UserGroup
	err = ch.db.Where("id = ?", messageRequest.Id).Find(&userGroup).Error

	if err != nil {
		ch.log(c, 404, "user group not found")
		return
	}

	message := entity.Message{
		ReceiverID: messageRequest.Id,
		Message:    messageRequest.Message,
	}
	ch.chann <- message

	go ch.ConsumeMessage(c)
	c.JSON(202, "message sent")
}

func (ch ChanRelayHandlerImpl) ConsumeMessage(c *gin.Context) {
	for {
		select {
		case msg, ok := <-(ch.chann):
			{
				if !ok {
					ch.log(c, 500, "some issue with channel")
					return
				}

				var userGroup entity.UserGroup
				err := ch.db.Preload("Users").Where("id=?", msg.ReceiverID).First(&userGroup).Error
				if err != nil {
					fmt.Println("Couldn't find user group")
					continue
				}

				for _, u := range userGroup.Users {
					userMsg := entity.Message{ReceiverID: u.UserId, Message: msg.Message}

					err = ch.db.Create(&userMsg).Error

					if err != nil {
						fmt.Printf("Unable to send message to user %v", u.UserId)
						continue
					}
				}
			}
		default:
			{
				c.JSON(202, "All messages consumed")
				return
			}
		}
	}
}