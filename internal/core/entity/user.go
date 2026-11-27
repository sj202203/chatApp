package entity

type User struct {
	UserId     int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name"`
	Messages   []Message `json:"messages" gorm:"foreignKey:ReceiverID;references:UserId"`
	UserGroups []UserGroup `json:"user_groups"`
}

func (User) TableName() string {
	return "user"
}

type UserGroupDTO struct {
	Id    int   `json:"id"`
	Users []int
}

type UserGroup struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Users    []User    `json:"users" gorm:"many2many:user_group_mappings"`
	Messages []Message `json:"messages" gorm:"foreignKey:ReceiverID;references:Id"`
}

func (UserGroup) TableName() string {
	return "user_group"
}