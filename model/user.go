package model

type User struct {
	Id         int    `gorm:"type:int;primary_key"`
	Name       string `gorm:"type:varchar(255);unique;not null "`
	IsVerified bool
}
