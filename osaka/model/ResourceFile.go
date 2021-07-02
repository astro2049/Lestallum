package model

import "gorm.io/gorm"

type ResourceFile struct {
	gorm.Model
	Url  string
	Name string
	Type string
}
