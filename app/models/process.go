package models

import (
	"github.com/goravel/framework/database/orm"
)

type Process struct {
	orm.Model
	Folder string
	File   string
}

func (r *User) TableName() string {
	return "process"
}
