package service

import (
	"path"
	"strings"
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/setting"
)

//AddImage 记录图片
func AddImage(str string) error {
	name := path.Base(str)
	id := strings.TrimSuffix(name, path.Ext(str))
	img := &model.Image{
		ImageID:   id,
		ImageName: name,
		ImagePath: str,
	}
	return db.AddImage(img)
}

//GetImages 读取图片
func GetImages(page int) ([]*model.Image, error) {
	offset := (page - 1) * setting.PageSize
	return db.GetImage(offset, setting.PageSize)
}
