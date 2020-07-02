package db

import (
	"encoding/json"
	"fmt"
	"sz_resume_202005/dao/redisdb"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/zlog"
	"testing"
)

func init() {
	setting.Init()
	zlog.Init()
	defer zlog.Sync()
	redisdb.Init()
	Init()
	zlog.Debug("test~")
}

func TestAddUser(t *testing.T) {
	zlog.Debug("TestAddUser")
	user := model.UserAuth{
		UserName: "普通人",
		Password: "imtiedan",
	}
	id, err := AddUser(user.UserName, user.Password)
	if err != nil {
		zlog.Errorf("adduser failed,err:%v\n", err)
	}
	zlog.Info("id:", id)
}
func TestCheckUserAuth(t *testing.T) {

	user := model.UserAuth{
		UserName: "tiedan",
		Password: "imtiedan",
	}
	ifExist, err := CheckUserAuth(&user)

	if err != nil {
		t.Error("err:", err)
		return
	}
	if ifExist {
		fmt.Print("tiedan存在")
		return
	}
	fmt.Print("用户名密码错误")
}
func TestGetUser(t *testing.T) {
	user := model.UserAuth{
		UserName: "tiedan",
		Password: "imtiedan",
	}
	u, err := GetUser(&user)
	if err != nil {
		t.Error("err:", err)
		return
	}
	fmt.Printf("u:%v\n", u)
}
func TestGetUserInfo(t *testing.T) {
	var id int
	id = 1

	userInfo, err := GetUserInfo(id)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("userInfo:%#v\n", userInfo)
}

func TestAddExp(t *testing.T) {
	// exp1 := &model.Experience{
	// 	StartTime:   "2010",
	// 	EndTime:     "2011",
	// 	LeaveReason: "不详",
	// 	Company:     "公司1",
	// 	Content:     "坐班",
	// 	Salary:      4000,
	// }
	// exp2 := &model.Experience{
	// 	StartTime:   "2010",
	// 	EndTime:     "2011",
	// 	LeaveReason: "不详",
	// 	Company:     "公司2",
	// 	Content:     "坐班",
	// 	Salary:      4000,
	// }
	// exp3 := &model.Experience{
	// 	StartTime:   "2010",
	// 	EndTime:     "2011",
	// 	LeaveReason: "不详",
	// 	Company:     "公司3",
	// 	Content:     "坐班",
	// 	Salary:      4000,
	// }
	// exps := model.Experiences{exp1, exp2, exp3}

}

// func TestDelExperience(t *testing.T) {
// 	// err := DelExperiences([]int{1})
// 	// if err != nil {
// 	// 	t.Error("DelExperiences function err:", err)
// 	// }
// }

func TestBatchStringParam(t *testing.T) {
	batchStringParam(0)
	batchStringParam(1)
	batchStringParam(2)
	batchStringParam(3)
}

// func TestAddExp1(t *testing.T) {
// 	exp1 := &model.Experience{
// 		StartTime:   "2010",
// 		EndTime:     "2011",
// 		LeaveReason: "不详",
// 		Company:     "公司1",
// 		Content:     "坐班",
// 		Salary:      4000,
// 		UserID:      1,
// 	}
// 	users := model.Experiences{exp1, exp1}
// 	b, _ := json.Marshal(users)
// 	fmt.Printf("users:%v\n", string(b))
// 	fmt.Println(users)
// 	var u model.Experiences
// 	json.Unmarshal(b, &u)
// 	fmt.Printf("u:%#v\n", u)

// 	for _, v := range u {
// 		fmt.Printf("v:%#v\n", v)
// 	}
// }

type Ecp struct {
	Exxx []model.Experience `json:"experience"`
}

func TestJson(t *testing.T) {
	str := `{"experience":[{"company":"公司1","content":"坐班","start_time":"2010","end_time":"2011","salary":4000,"leave_reason":"不详","user_id":1},{"company":"公司1","content":"坐班","start_time":"2010","end_time":"2011","salary":4000,"leave_reason":"不详","user_id":1}]}`
	var ecp Ecp
	json.Unmarshal([]byte(str), &ecp)

	fmt.Printf("ecp:%#v\n", ecp)

	for _, v := range ecp.Exxx {
		fmt.Printf("Exxx:%#v\n", v)
	}
}
func TestEditUserInfo(t *testing.T) {

	userInfo := &model.UserInfo{
		User: model.User{
			Avatar: "wenzi",
			UserID: 1,
		},
		Career:        "全栈工程师",
		Workingage:    11,
		Mobile:        "18686481411",
		Location:      "cc",
		Major:         "计算机",
		Edubackground: "本科",
		Introduce:     "一段介绍",
	}
	err := EditUserInfo(userInfo)
	if err != nil {
		t.Error(err)
	}
}
func TestIsAdmin(t *testing.T) {
	b, err := IsAdmin(2)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("用户权限是管理员？：%v", b)
}
