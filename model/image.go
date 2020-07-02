package model

import "time"

//Image 上传图片
type Image struct {
	ImageID    string    `json:"img_id" db:"img_id" redis:"img_id"`                //图片id
	ImageName  string    `json:"img_name" db:"img_name" redis:"img_name"`          //图片全名
	ImagePath  string    `json:"img_path" db:"img_path" redis:"img_path"`          //图片路径
	CreateTime time.Time `json:"create_time" db:"create_time" redis:"create_time"` //图片id
	UpdateTime time.Time `json:"update_time" db:"update_time" redis:"update_time"` //更新时间
}
