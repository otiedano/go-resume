package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
)

//AddRequestLog 记录请求信息
func TestAddRequestLog(t *testing.T) {

	rlog := &model.RequestLog{
		URL:       "http://sfsf.com",
		UserAgent: "字符串",
		TokenInfo: "userid:1,",
		Cookie:    "token:234sf234243fsdf24234,timestamp:2323424",
		Method:    "get",
	}
	err := AddRequestLog(rlog)
	if err != nil {
		t.Error("err:", err)
	}
}

//GetRequestLog 读取请求信息
func TestGetRequestLog(t *testing.T) {
	rlogs, err := GetRequestLog(0, 2)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("rlogs:%v\n", rlogs)

}
