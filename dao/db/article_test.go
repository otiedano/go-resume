package db

import (
	"fmt"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/setting"
	"testing"
)

func TestTotalArticle(t *testing.T) {
	n, err := TotalArticle()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)

}
func TestTotalArticleByAuthor(t *testing.T) {
	n, err := TotalArticleByAuthor(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)

}
func TestTotalArticleByStatus(t *testing.T) {
	n, err := TotalArticleByStatus()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count:%v\n", n)
	n, err = TotalArticleByStatus(0)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count0:%v\n", n)
	n, err = TotalArticleByStatus(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count1:%v\n", n)
	n, err = TotalArticleByStatus(2)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("count2:%v\n", n)
}
func TestAddArticle(t *testing.T) {
	article := model.ArticleDetail{
		Article: model.Article{
			Title:      "文章标题s",
			Img:        "/runtime/upload/images/avatar.jpg",
			CategoryID: 8,
			Summary:    "这个可以没有吧",
		},

		Content: "tiedan",
	}
	article1 := model.ArticleDetail{
		Article: model.Article{
			Title:      "最近生活的感想",
			Img:        "/runtime/upload/images/avatar.jpg",
			CategoryID: 9,
			Summary:    "想起来我就更新",
		},

		Content: "内容太多，想不起来",
	}

	intnum, err := AddArticle(1, &article)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("intnum:%v\n", intnum)
	intnum, err = AddArticle(1, &article1)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("intnum:%v\n", intnum)
}

func TestCheckArticle(t *testing.T) {
	ids := []int{5}
	err := CheckArticles(ids, 1)
	if err != nil {
		t.Error("err:", err)
	}
}

func TestGetArticlesByAuthor(t *testing.T) {

	a, err := GetArticlesByAuthor(1, 0, 15)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("len a:%#v\n\n", len(a))

	// for _, v := range a {
	// 	fmt.Printf("v:%+v\n", v)
	// }
}
func TestGetArticlesByCategory(t *testing.T) {
	a, err := GetArticlesByCategory(2, 0, setting.PageSize)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("len a:%#v\n\n", len(a))

	for _, v := range a {
		fmt.Printf("v:%+v\n", v)
	}
}
func TestGetAllArtilesByStatus(t *testing.T) {
	a, err := GetAllArtilesByStatus(0, 10, 2)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("len a:%#v\n\n", len(a))

	for _, v := range a {
		fmt.Printf("v:%+v\n", v)
	}
	fmt.Printf("len:%+v\n", len(a))
}
func TestGetArticle(t *testing.T) {
	a, err := GetArticle(8)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf(" a:%+v\n\n", a)

}
func TestGetPArticle(t *testing.T) {
	a, err := GetPArticle(1, 8)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("len a:%+v\n\n", a)

}
func TestGetRPArticle(t *testing.T) {
	a, err := GetRPArticle(31)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf(" a:%+v\n\n", a)
}

func TestEditArticle(t *testing.T) {
	userID := 1
	article := &model.ArticleDetail{
		Article: model.Article{
			ArticleID:  10,
			Title:      "重新修改了一下文章",
			Img:        "/runtime/upload/images/avatar.jpg",
			CategoryID: 9,
			Summary:    "文章的摘要修改",
		},

		Content: "这是修改过的文字。",
	}
	err := EditArticle(userID, article)
	if err != nil {
		t.Error("err:", err)
	}
}
func TestExistArticleByID(t *testing.T) {
	b, err := ExistArticleByID(10)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("exit:%v", b)

}
func TestExistArticleByAuth(t *testing.T) {

	b, err := ExistArticleByAuth(1, 10)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("exit:%v", b)

}
func TestExistArticleByAdmin(t *testing.T) {

	b, err := ExistArticle(10)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("exit:%v", b)

}
func TestVisitArticle(t *testing.T) {
	err := CountArticle(3)
	err = CountArticle(3)
	err = CountArticle(3)
	if err != nil {
		t.Error("err:", err)
	}
	a, err := GetArticle(3)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("article:%+v", a)
}

func TestDelArticleFE(t *testing.T) {
	ids := []int{3, 4, 5}
	err := DelArticlesFE(ids)
	if err != nil {
		t.Error("err:", err)
	}
}

func TestDelArticle(t *testing.T) {
	idss := map[string][]int{
		"first":  {39, 38, 37},
		"second": {39, 38, 37},
		"none":   {},
	}
	for k, v := range idss {
		fmt.Println("key:", k)
		err := DelArticles(1, v)
		if err != nil {
			t.Error("err:", err)
		}
	}

}
func BenchmarkGetArticlesByCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetArticlesByCategory(8, 0, setting.PageSize)
		if err != nil {
			b.Error("err:", err)
		}

	}
}
func BenchmarkGetArticlesByCategory0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetArticlesByCategory(0, 0, setting.PageSize)
		if err != nil {
			b.Error("err:", err)
		}

	}
}
func BenchmarkGetArticlesByAuthor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetArticlesByAuthor(1, 0, 15)
		if err != nil {
			b.Error("err:", err)
		}

	}
}
