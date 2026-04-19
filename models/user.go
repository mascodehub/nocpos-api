package models

type User struct {
	IDUser string `json:"iduser" gorm:"column:iduser;primaryKey"`

	Name string `json:"name" gorm:"column:name"`

	Username string `json:"username" gorm:"column:username"`

	Password string `json:"password" gorm:"column:password"`

	IDOutlet string `json:"idoutlet" gorm:"column:idoutlet"`

	Status string `json:"status" gorm:"column:status"`
}

func (User) TableName() string {
	return "master_users"
}
