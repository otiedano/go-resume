package model

import (
	"database/sql/driver"
	"time"
)

//Work 作品
type Work struct {
	WorkID     int        `json:"work_id" db:"work_id" redis:"work_id"`             //作品ID
	Title      string     `json:"title" db:"title" redis:"title"`                   //作品标题
	CoverImg   string     `json:"cover_img" db:"cover_img" redis:"cover_img"`       //作品封面图
	StartTime  time.Time  `json:"start_time" db:"start_time" redis:"start_time"`    //开始时间
	EndTime    time.Time  `json:"end_time" db:"end_time" redis:"end_time"`          //结束时间
	UserID     int        `json:"user_id" db:"user_id" redis:"user_id"`             //用户ID
	WorkNO     int        `json:"work_no" db:"work_no" redis:"work_no"`             //作品序号
	Tags       WorkTagRls `json:"tag" db:"tag" redis:"tag"`                         //作品标签
	Status     int        `json:"status" db:"status" redis:"status"`                //作品审核状态，0未审核，1审核通过，2软删除
	CreateTime time.Time  `json:"create_time" db:"create_time" redis:"create_time"` //创建时间
	UpdateTime time.Time  `json:"update_time" db:"update_time" redis:"update_time"` //更新时间
	UserName   string     `json:"user_name" db:"user_name" redis:"user_name"`
	Avatar     string     `json:"avatar" db:"avatar" redis:"avatar"`
	ViewCount  int        `json:"view_count" db:"view_count" redis:"view_count"`
}

//WorkDetail 工作详情
type WorkDetail struct {
	Work                  //--作品类型
	Img        string     `json:"img" db:"img" redis:"img"`                         //作品头图
	Content    string     `json:"content" db:"content" redis:"content"`             //作品内容
	WorkImgs   WorkAlbums `json:"albums" db:"albums" redis:"albums"`                //作品轮播图
	IfHorizons bool       `json:"if_horizons" db:"if_horizons" redis:"if_horizons"` //轮播图是否水平
}

//WorkAlbum 工作图片
type WorkAlbum struct {
	AlbumID    int       `json:"album_id" db:"album_id" redis:"album_id"`          //作品id
	WorkID     int       `json:"work_id" db:"work_id" redis:"work_id"`             //作品id
	ImgPath    string    `json:"img_path" db:"img_path" redis:"img_path"`          //图片id
	ImgNo      int       `json:"img_no" db:"img_no" redis:"img_no"`                //图片地址
	UpdateTime time.Time `json:"update_time" db:"update_time" redis:"update_time"` //作品id
	CreateTime time.Time `json:"create_time" db:"create_time" redis:"create_time"` //作品id
}

//WorkAlbums WorkAlbum数组
type WorkAlbums []*WorkAlbum

//ConvInterfaceArray 转换成接口数组
func (a WorkAlbums) ConvInterfaceArray() (n []interface{}) {
	n = make([]interface{}, len(a))
	for i, v := range a {
		n[i] = v
	}
	return
}

//Value 实现driver.value接口，使得sqlx.In可以使用
func (e WorkAlbum) Value() (driver.Value, error) {
	return []interface{}{e.AlbumID, e.WorkID, e.ImgPath, e.ImgNo}, nil
}

//WorkTagRl 绑定的作品标签
type WorkTagRl struct {
	RlID       int       `json:"rl_id" db:"rl_id" redis:"ri_id"`
	TagID      int       `json:"tag_id" db:"tag_id" redis:"tag_id"`
	TagName    string    `json:"tag_name" db:"tag_name" redis:"tag_name"`
	TagNo      int       `json:"tag_no" db:"tag_no" redis:"tag_no"`
	WorkID     int       `json:"work_id" db:"work_id" redis:"work_id"`
	CreateTime time.Time `json:"create_time" db:"create_time" redis:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time" redis:"update_time"`
}

//WorkTagRls 绑定的WorkTagRl数组
type WorkTagRls []*WorkTagRl

//ConvInterfaceArray 转换成接口数组
func (a WorkTagRls) ConvInterfaceArray() (n []interface{}) {
	n = make([]interface{}, len(a))
	for i, v := range a {
		n[i] = v
	}
	return
}

//Value 实现driver.value接口，使得sqlx.In可以使用
func (e WorkTagRl) Value() (driver.Value, error) {
	return []interface{}{e.RlID, e.TagID, e.WorkID}, nil
}

//WorkTag 作品标签
type WorkTag struct {
	TagID      int       `json:"tag_id" db:"tag_id" redis:"tag_id"`
	TagName    string    `json:"tag_name" db:"tag_name" redis:"tag_name"`
	CreateTime time.Time `json:"create_time" db:"create_time" redis:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time" redis:"update_time"`
	TagNO      int       `json:"tag_no" db:"tag_no" redis:"tag_no"`
}

//WorkTags 作品标签数组
type WorkTags []*WorkTag

//ConvInterfaceArray 转换成接口数组
func (a WorkTags) ConvInterfaceArray() (n []interface{}) {
	n = make([]interface{}, len(a))
	for i, v := range a {
		n[i] = v
	}
	return
}

//Value 实现driver.value接口，使得sqlx.In可以使用
func (e WorkTag) Value() (driver.Value, error) {
	return []interface{}{e.TagID, e.TagName, e.TagNO}, nil
}
