package models

import (
	"database/sql/driver"
	"errors"
	"strings"

	"gorm.io/gorm"
)
type Lists []string

type User struct{
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName string `json:"lastName" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Lists Lists `json:"lists" gorm:"type:VARCHAR(255)"`	
}

type Song struct {
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey"`
	ListName string 
	Title  string
	Artist string
	Remark string
	Owner string `json:"owner"`
	// SongRefer uint
}

type Movie struct {
	gorm.Model
	ListName string 
	ID     uint `json:"id" gorm:"primaryKey"`
	Title  string	`json:"title"`
	Remark string	`json:"remark"`
	Owner string `json:"owner"`
	// MovieRefer uint
}

type Book struct {
	gorm.Model
	ListName string `json:"listName"`
	ID     uint `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Writer string	`json:"writer"`
	Remark string	`json:"remark"`
	Owner string `json:"owner"`
	// BookRefer uint
}

type Other struct {
	gorm.Model
	ListName string `json:"listName"`
	ID     uint `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Remark string	`json:"remark"`
	Owner string `json:"owner"`
	// OtherRefer uint
}

func (l *Lists) Scan(src any)error{
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot be cast to []byte")
	}
	*l = strings.Split(string(bytes), ",")
	return nil
}

func (l Lists) Value() (driver.Value,error){
	if len(l) == 0 {
		return nil,nil
	}
	return strings.Join(l, ","),nil
}