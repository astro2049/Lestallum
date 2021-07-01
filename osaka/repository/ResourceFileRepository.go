package repository

import (
	"osaka/global"
	"osaka/model"
)

func FindFileById(fid uint) (*model.ResourceFile, error) {
	file := new(model.ResourceFile)
	err := global.DB.Where("id = ?", fid).First(file).Error
	return file, err
}
