package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
)

func TestAddACategory(t *testing.T) {
	a := model.ArticleCategory{
		CategoryName: "新分类",
		CategoryNo:   10,
	}
	num, err := AddArticleCategory(&a)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("idnum:%v", num)

}
func TestGetACategory(t *testing.T) {

	num, err := GetArticleCategories()
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("idnum:%+v\n", num[0])

}

func TestDeleteACategory(t *testing.T) {
	u, err := GetArticleCategories()
	if err != nil {
		t.Error("get articleCategories failed,err:", err)
	}
	var a = make([]int, 0)
	for _, v := range u {
		a = append(a, v.CategoryID)
	}

	err = DelArticleCategories(a)
	if err != nil {
		t.Error("delete category failed:", err)
	}
}
func TestEditACategory(t *testing.T) {
	c := model.ArticleCategory{
		CategoryNo:   1,
		CategoryName: "随笔",
		CategoryID:   4,
	}
	err := EditArticleCategory(&c)
	if err != nil {
		t.Error("edit articleCategory failed,err:", err)
	}
}
func TestEditACategories(t *testing.T) {
	a := make(model.ArticleCategories, 0)
	a = append(a,
		&model.ArticleCategory{
			CategoryID:   8,
			CategoryName: "艺术",
			CategoryNo:   8,
		}, &model.ArticleCategory{
			CategoryID:   9,
			CategoryName: "拼图",
			CategoryNo:   9,
		})
	err := EditArticleCategories(&a)
	if err != nil {
		t.Error("EditArticleCategories failed,err:", err)
	}
}
func TestExistArticleCategory(t *testing.T) {
	b, err := ExistArticleCategory(8)
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("not expect result")
	}
}
