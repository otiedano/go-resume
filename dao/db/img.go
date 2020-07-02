package db

import (
	"sz_resume_202005/model"
	"time"
)

//AddImage 添加图片
func AddImage(img *model.Image) (err error) {
	sqlStr := "insert into image (img_id,img_name,img_path,update_time) values (?,?,?,?) on duplicate key update  update_time=values(update_time)"

	_, err = db.Exec(sqlStr, img.ImageID, img.ImageName, img.ImagePath, time.Now())

	return
}

//GetImage 获取图片
func GetImage(offset, limit int) (images []*model.Image, err error) {
	sqlStr := `
	SELECT img_id,img_name,img_path,update_time,create_time FROM image
	ORDER BY update_time desc
	LIMIT ? OFFSET ?
	`

	err = db.Select(&images, sqlStr, limit, offset)
	return
}
