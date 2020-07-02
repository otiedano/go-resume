package model

import (
	"database/sql/driver"
	"time"
)

//Article 文章
type Article struct {
	ArticleID    int       `json:"article_id" db:"article_id" redis:"article_id"`    //文章id
	Title        string    `json:"title" db:"title" redis:"title"`                   //文章标题
	ViewCount    int       `json:"view_count" db:"view_count" redis:"view_count"`    //文章浏览次数
	Status       int       `json:"status" db:"status" redis:"status"`                //文章状态：待审核0，已审核1
	Summary      string    `json:"summary" db:"summary" redis:"summary"`             //文章摘要
	CreateTime   time.Time `json:"create_time" db:"create_time" redis:"create_time"` //文章创建时间
	UpdateTime   time.Time `json:"update_time" db:"update_time" redis:"update_time"` //文章修改时间
	UserID       int       `json:"author_id" db:"author_id" redis:"author_id"`
	UserName     string    `json:"user_name" db:"user_name" redis:"user_name"`             //用户名称
	Avatar       string    `json:"avatar" db:"avatar" redis:"avatar"`                      //头像
	CategoryID   int       `json:"category_id" db:"category_id" redis:"category_id"`       //分类id
	CategoryName string    `json:"category_name" db:"category_name" redis:"category_name"` //分类名称
	CategoryNo   int       `json:"category_no" db:"category_no" redis:"category_no"`       //分类序号
	Img          string    `json:"img" db:"img" redis:"img"`                               //图片
	//	CreateTime   time.Time `json:"create_time" db:"create_time" redis:"create_time"`       //分类创
	Tag []*ArticleTag
}

//ArticleCategory 文章分类
type ArticleCategory struct {
	CategoryID   int    `json:"category_id" db:"category_id" redis:"category_id"`       //分类id
	CategoryName string `json:"category_name" db:"category_name" redis:"category_name"` //分类名称
	CategoryNo   int    `json:"category_no" db:"category_no" redis:"category_no"`       //分类序号
	//	CreateTime   time.Time `json:"create_time" db:"create_time" redis:"create_time"`       //分类创建时间
	//	UpdateTime   time.Time `json:"update_time" db:"update_time" redis:"update_time"`       //分类修改时间
}

//ArticleDetail 文章详情
type ArticleDetail struct {
	Article //--文章类型

	Content string `json:"content" db:"content" redis:"content"` //文章内容详情

}

//ArticleTag 文章标签
type ArticleTag struct {
	TagID   int `json:"tag_id" db:"tag_id" redis:"tag_id"`
	TagName int `json:"tag_name" db:"tag_name" redis:"tag_name"`
	TagNo   int `json:"tag_no" db:"tag_no" redis:"tag_no"`
}

//ArticleCategories Experience数组
type ArticleCategories []*ArticleCategory

//ConvInterfaceArray 转换成接口数组
func (a ArticleCategories) ConvInterfaceArray() (n []interface{}) {
	n = make([]interface{}, len(a))
	for i, v := range a {
		n[i] = v
	}
	return
}

//Value 实现driver.value接口，使得sqlx.In可以使用
func (a *ArticleCategory) Value() (driver.Value, error) {
	return []interface{}{a.CategoryName, a.CategoryNo, a.CategoryID}, nil
}
