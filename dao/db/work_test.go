package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
	"time"
)

func TestTotalWork(t *testing.T) {
	n, err := TotalWork()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)
	n, err = TotalWork(0)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)

	n, err = TotalWork(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)
}
func TestTotalWorkByAuth(t *testing.T) {
	n, err := TotalWorkByAuthor(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)
	n, err = TotalWorkByAuthor(1, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)

	n, err = TotalWorkByAuthor(1, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)
	n, err = TotalWorkByAuthor(1, 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)
}
func TestTotalWorkByStatus(t *testing.T) {
	n, err := TotalWorkByStatus()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)
	n, err = TotalWorkByStatus(0)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count0:%v\n", n)
	n, err = TotalWorkByStatus(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count1:%v\n", n)
	n, err = TotalWorkByStatus(2)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count2:%v\n", n)
}
func TestAddWork(t *testing.T) {
	wimg := model.WorkAlbums{}
	stime, _ := time.Parse("2006-01-02 15:04:05", "2018-12-27 18:44:55")
	wimg = append(wimg, &model.WorkAlbum{
		ImgPath: "/runtime/upload/images/first.jpg",
		ImgNo:   1,
	}, &model.WorkAlbum{
		ImgPath: "/runtime/upload/images/second.jpg",
		ImgNo:   2,
	})
	tags := model.WorkTagRls{}
	tags = append(tags, &model.WorkTagRl{

		TagID:   3,
		TagName: "其他",
	}, &model.WorkTagRl{

		TagID:   4,
		TagName: "Flash",
	})
	w := &model.WorkDetail{
		WorkImgs:   wimg,
		IfHorizons: false,
		Img:        "/runtime/upload/images/avatar.jpg",
		Content:    "可能是一段很长的文字",
		Work: model.Work{
			Title:     "第一个作品",
			WorkNO:    1,
			CoverImg:  "/runtime/upload/images/cover_img.jpg",
			StartTime: stime,
			EndTime:   time.Now(),
			UserID:    1,
			Tags:      tags,
		},
	}

	err := AddWork(1, w)
	if err != nil {
		t.Error("err:", err)
	}
}
func genWorkDetail() (w *model.WorkDetail) {
	wimg := model.WorkAlbums{}
	stime, _ := time.Parse("2006-01-02 15:04:05", "2010-01-27 18:44:55")
	wimg = append(wimg, &model.WorkAlbum{

		ImgPath: "/runtime/upload/images/first.jpg",
		ImgNo:   122,
	}, &model.WorkAlbum{

		ImgPath: "/runtime/upload/images/second.jpg",
		ImgNo:   126,
	})
	tags := model.WorkTagRls{}
	tags = append(tags, &model.WorkTagRl{
		RlID:    35,
		TagID:   6,
		TagName: "后端",
	}, &model.WorkTagRl{

		TagID:   7,
		TagName: "前端",
	})
	w = &model.WorkDetail{
		WorkImgs:   wimg,
		IfHorizons: false,
		Img:        "/runtime/upload/images/avatar.jpg",
		Content:    "再次修改过的很长的文字，以后会插入图片",
		Work: model.Work{
			WorkID:    9,
			Title:     "新的作品,测试",
			WorkNO:    20,
			CoverImg:  "/runtime/upload/images/cover_img.jpg",
			StartTime: stime,
			EndTime:   time.Now(),
			UserID:    1,
			Tags:      tags,
		},
	}
	return
}
func TestGetWork(t *testing.T) {
	w, err := GetWork(9)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("w:%+v", w)
	// w, err = GetWorksByTag(3, 1, 10)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Printf("w:%+v", w)
}
func TestGetPWork(t *testing.T) {
	w, err := GetPWork(1, 9)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("w:%+v", w)
	// w, err = GetWorksByTag(3, 1, 10)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Printf("w:%+v", w)
}
func TestGetRPWork(t *testing.T) {
	w, err := GetRPWork(3)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf(" w:%+v\n\n", w)
}

func TestGetWorksByTag(t *testing.T) {
	w, err := GetWorksByTag(0, 0, 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("w:%+v", w)
	w, err = GetWorksByTag(3, 0, 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("w:%+v", w)
	// w, err = GetWorksByTag(3, 1, 10)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Printf("w:%+v", w)
}
func TestGetWorksByAuth(t *testing.T) {
	w, err := GetWorksByAuthor(1, 0, 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("w:%+v", w)
	// w, err = GetWorksByTag(3, 1, 10)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Printf("w:%+v", w)
}
func TestGetAllWorksByStatus(t *testing.T) {
	a, err := GetAllWorksByStatus(0, 10, 2)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("len a:%#v\n\n", len(a))

	for _, v := range a {
		fmt.Printf("v:%+v\n", v)
	}
	fmt.Printf("len:%+v\n", len(a))

}

//EditWork
func TestEditWork(t *testing.T) {
	work := genWorkDetail()
	err := EditWork(1, work)
	if err != nil {
		t.Error("err:", err)
	}
}

//DelWorksFE
func TestDelWorksFE(t *testing.T) {
	ids := []int{4, 5, 6}
	err := DelWorksFE(ids)
	if err != nil {
		t.Error("err:", err)
	}
}

//DelWorks
func TestDelWorks(t *testing.T) {
	ids := []int{8}
	err := DelWorks(1, ids)
	if err != nil {
		t.Error("err:", err)
	}
}

//ExistWorkByAuth
func TestExistWorkByAuth(t *testing.T) {

	b, err := ExistWorkByAuth(1, 9)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("b:%v\n", b)
	if !b {
		t.Error("b is not expect result")
	}
}

//ExistWorkByID
func TestExistWorkByID(t *testing.T) {
	ids := []int{9}
	_ = CheckWorks(ids, 1)
	b, err := ExistWorkByID(9)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("b:%v\n", b)
	if !b {
		t.Error("b is not expect result")
	}

}

//ExistWorkByID
func TestExistWorkByAdmin(t *testing.T) {

	b, err := ExistWork(9)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("b:%v\n", b)
	if !b {
		t.Error("b is not expect result")
	}

}
func TestCheckWorks(t *testing.T) {
	ids := []int{1, 2, 3}
	err := CheckWorks(ids, 1)
	if err != nil {
		t.Error(err)
	}
}

//VisitWork
func TestCountWork(t *testing.T) {
	err := CountWork(9)
	if err != nil {
		t.Error("err:", err)
	}
}
