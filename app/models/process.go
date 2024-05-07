package models

import (
	"github.com/goravel/framework/database/orm"
)

type Process struct {
	orm.Model
	File      string
	Extension string
	Folder    string
	Route     string
}

func (r *User) TableName() string {
	return "process"
}
