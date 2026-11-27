package entity

type Message struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ReceiverID int    `json:"receiver_id"`
	Message    string `json:"message"`
}

func (Message) TableName() string {
	return "message"
}