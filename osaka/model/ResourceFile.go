package model

import "gorm.io/gorm"

type ResourceFile struct {
	gorm.Model
	Url  string
	Name string
	Type uint // 1: video, 2:avatar, 3:picture, 4: document
}
