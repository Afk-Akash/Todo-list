package models

type Todo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Priority string `json:"priority"`
	UserID   int    `json:"userid"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
}
