package model

import (
	"database/sql/driver"
	"time"
)

//UserToken 用户的token信息
type UserToken struct {
	Token string `form:"token" json:"token" binding:"required"`
}

//UserAuth 用户
type UserAuth struct {
	Password string `json:"password" db:"password" redis:"password" valid:"Required"`    //用户密码
	UserName string `json:"user_name" db:"user_name" redis:"user_name" valid:"Required"` //用户昵称
}

//User 用户
type User struct {
	UserID int `json:"user_id" db:"user_id" redis:"user_id"` //用户ID
	//Password  string    `json:"password" db:"password" redis:"password"`          //用户密码
	CreateTime time.Time `json:"create_time" db:"create_time" redis:"create_time"` //创建时间
	Avatar     string    `json:"avatar" db:"avatar" redis:"avatar"`                //头像
	//Token     string    `json:"token" db:"token" redis:"token"`                   //token
	UserName string `json:"user_name" db:"user_name" redis:"user_name" ` //用户昵称
	Role     string `json:"role" db:"role" redis:"role" `                //权限
}

//UserInfo 用户信息
type UserInfo struct {
	User                 //--用户类型
	Career        string `json:"career" db:"career" redis:"career" valid:"Max(20)"`                             //职位
	Mobile        string `json:"mobile" db:"mobile" redis:"mobile" valid:"Mobile" `                             //手机
	Mail          string `json:"mail" db:"mail" redis:"mail" valid:"Email"`                                     //邮件
	Location      string `json:"location" db:"location" redis:"location" valid:"Required"`                      //地点
	Major         string `json:"major" db:"major" redis:"major" valid:"Required"`                               //专业
	Edubackground string `json:"edu_background" db:"edu_background" redis:"edu_background" valid:"MaxSize(20)"` //教育背景
	Workingage    int    `json:"working_age" db:"working_age" redis:"working_age" valid:"Range(1, 40)"`         //工龄
	Introduce     string `json:"introduce" db:"introduce" redis:"introduce" valid:"Required"`                   //自我介绍
	HasExp        bool   `json:"has_experiences" db:"has_experiences" redis:"has_experiences"`                  //是否有工作经验
	HasEdit       bool   `json:"has_edit" db:"has_edit" redis:"has_edit"`                                       //自我介绍

}

//UserInfoCpt 用户完整信息
type UserInfoCpt struct {
	UserInfo
	Experiences []*Experience `json:"experience" db:"experience" redis:"experience"` //--工作经历类型数组
}

//Experience 工作经历
type Experience struct {
	Company     string    `json:"company" db:"company" redis:"company" valid:"Required"`                //工作公司
	Content     string    `json:"content" db:"content" redis:"content" valid:"Required"`                //工作内容
	StartTime   time.Time `json:"start_time" db:"start_time" redis:"start_time" valid:"Required"`       //开始时间
	EndTime     time.Time `json:"end_time" db:"end_time" redis:"end_time" valid:"Required"`             //结束时间
	Salary      int64     `json:"salary" db:"salary" redis:"salary" valid:"Required;Max(1000000)"`      //薪资
	LeaveReason string    `json:"leave_reason" db:"leave_reason" redis:"leave_reason" valid:"Required"` //离职原因
	UserID      int       `json:"user_id" db:"user_id" redis:"user_id"`
	ExpID       int       `json:"exp_id" db:"exp_id" redis:"exp_id"`
}

//Experiences Experience数组
type Experiences []*Experience

//ConvInterfaceArray 转换成接口数组
func (a Experiences) ConvInterfaceArray() (n []interface{}) {
	n = make([]interface{}, len(a))
	for i, v := range a {
		n[i] = v
	}
	return
}

//Value 实现driver.value接口，使得sqlx.In可以使用
func (e *Experience) Value() (driver.Value, error) {
	return []interface{}{e.Company, e.Content, e.StartTime, e.EndTime, e.Salary, e.LeaveReason, e.UserID}, nil
}

//UserSkill 技术
type UserSkill struct {
	UserInfo
	Skills
	Softwares
}

//Skill 技能
type Skill struct {
	SkillID   int    `json:"skill_id" db:"skill_id" redis:"skill_id"`
	SkillName string `json:"skill_name" db:"skill_name" redis:"skill_name" valid:"Required"`
	Img       string `json:"img" db:"img" redis:"img" valid:"Required"`
	Per       uint   `json:"per" db:"per" redis:"per" valid:"Required"`
	UserID    int    `json:"user_id" db:"user_id" redis:"user_id"`
	SkillNO   int    `json:"skill_no" db:"skill_no" redis:"skill_no" valid:"Required;Min(0)"`
}

//Skills Skill数组
type Skills []*Skill

//ConvInterfaceArray 转换成接口数组
func (a Skills) ConvInterfaceArray() (n []interface{}) {
	n = make([]interface{}, len(a))
	for i, v := range a {
		n[i] = v
	}
	return
}

//Value 实现driver.value接口，使得sqlx.In可以使用
func (e Skill) Value() (driver.Value, error) {
	return []interface{}{e.SkillID, e.SkillName, e.Img, e.Per, e.UserID, e.SkillNO}, nil
}

//Software 技能
type Software struct {
	SoftwareID   int    `json:"software_id" db:"software_id" redis:"software_id"`
	SoftwareName string `json:"software_name" db:"software_name" redis:"software_name" valid:"Required"`
	Img          string `json:"img" db:"img" redis:"img" valid:"Required"`
	UserID       int    `json:"user_id" db:"user_id" redis:"user_id"`
	SoftwareNO   int    `json:"software_no" db:"software_no" redis:"software_no" valid:"Required;Min(0)"`
}

//Softwares Software数组
type Softwares []*Software

//ConvInterfaceArray 转换成接口数组
func (a Softwares) ConvInterfaceArray() (n []interface{}) {
	n = make([]interface{}, len(a))
	for i, v := range a {
		n[i] = v
	}
	return
}

//Value 实现driver.value接口，使得sqlx.In可以使用
func (e Software) Value() (driver.Value, error) {
	return []interface{}{e.SoftwareID, e.SoftwareName, e.Img, e.UserID, e.SoftwareNO}, nil
}
