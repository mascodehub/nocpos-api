package models

type User struct {
	IDUser string `json:"iduser" gorm:"column:iduser;primaryKey"`

	Name string `json:"name" gorm:"column:name"`

	Username string `json:"username" gorm:"column:username"`

	Password string `json:"password" gorm:"column:password"`

	AccessLevel string `json:"access_level" gorm:"column:access_level"`

	Status string `json:"status" gorm:"column:status"`
}

func (User) TableName() string {
	return "master_users"
}
