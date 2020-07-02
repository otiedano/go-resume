package redisdb

import (
	"fmt"

	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/zlog"
	"testing"
)

func TestMain(t *testing.T) {
	setting.Init()
	zlog.Init()
	defer zlog.Sync()
	Init()

}
func TestSet(t *testing.T) {
	err := Set("name", "tiedan", 5)
	if err != nil {
		t.Errorf("set name err,err:%v\n", err)
	}
}

func TestNil(t *testing.T) {
	a, exist, err := Get("haha")
	if err != nil {

		t.Errorf("testnil err:%v\n", err)
		return
	}
	if !exist {
		t.Errorf("haha字段不存在\n")
		return
	}
	fmt.Printf("haha存在，值:%v\n", a)

	ttm, err := GetTTL("haha")
	if err != nil {
		t.Errorf("liu err:%v\n", err)
	}
	fmt.Printf("haha的ttl:%v\n", ttm)

}

func TestNil1(t *testing.T) {
	name, exist, err := Get("name")

	if err != nil {

		t.Errorf("testnil err:%v\n", err)
	}
	if !exist {
		t.Errorf("name字段不存在\n")
	}
	fmt.Printf("name 字段存在,name:%v\n", name)
	ttm, err := GetTTL("name")
	if err != nil {
		t.Errorf("TTL err:%v\n", err)
	}
	fmt.Printf("name的ttl:%v\n", ttm)
}

func TestDel(t *testing.T) {
	Set("t1", "ttt", 60)
	n, _ := Exist("t1")
	fmt.Printf("判断是否存在的结果n:%v\n", n)
	num, _ := Del("t1")

	fmt.Printf("判断删除的结果num:%v\n", num)
	m, _ := Exist("t1")
	fmt.Printf("判断是否存在的结果n:%v\n", m)
}
func TestExist(t *testing.T) {
	Set("haha", "heihei", 60)
	num, err := Exist("name", "haha")
	if err != nil {
		t.Errorf("Exist keys failed,err:%v\n", err)
	}
	fmt.Printf("num:%v\n", num)
}
