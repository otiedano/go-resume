package db

import (
	"database/sql"
	"fmt"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/zlog"
	"time"
	"unsafe"

	"github.com/jmoiron/sqlx"
)

//AddArticleCategory 新增文章分类
func AddArticleCategory(c *model.ArticleCategory) (intNum int, err error) {
	sqlStr := "INSERT INTO article_category (category_name,category_no) VALUES (?,?)"

	rs, err := db.Exec(sqlStr, c.CategoryName, c.CategoryNo)
	if err != nil {
		return
	}

	id64, err := rs.LastInsertId()
	if err != nil {
		return
	}
	intNum = *(*int)(unsafe.Pointer(&id64))
	return

}

//GetArticleCategories 读取文章分类
func GetArticleCategories() (categories []*model.ArticleCategory, err error) {
	sqlStr := `
	SELECT category_id,category_name,category_no 
	From article_category
	WHERE category_id > 1
	ORDER BY category_no asc,update_time desc
	
	`

	err = db.Select(&categories, sqlStr)
	return
}

//GetAllArticleCategories 读取所有文章分类
func GetAllArticleCategories() (categories []*model.ArticleCategory, err error) {
	sqlStr := `
	SELECT category_id,category_name,category_no 
	From article_category
	ORDER BY category_no desc,update_time desc
	 
	`

	err = db.Select(&categories, sqlStr)
	return
}

//EditArticleCategory 修改文章分类
func EditArticleCategory(category *model.ArticleCategory) (err error) {
	sqlStr := `UPDATE article_category SET category_name=?,category_no=?,update_time=? WHERE category_id=?`

	result, err := db.Exec(sqlStr, category.CategoryName, category.CategoryNo, time.Now(), category.CategoryID)
	if err != nil {
		return
	}

	rs, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rs != 1 {
		err = fmt.Errorf("rs is not equal 1")
		return
	}

	return
}

//EditArticleCategories 批量修改文章分类
func EditArticleCategories(categories *model.ArticleCategories) (err error) {
	l := len(*categories)
	fmt.Printf("categories,%+v\n", categories)
	fmt.Print("l:", l)
	if l < 1 {
		return
	}

	sqlStr := fmt.Sprintf(`
	INSERT INTO article_category (category_name,category_no,category_id) VALUES %s 
	ON DUPLICATE KEY 
	UPDATE category_name=VALUES(category_name),category_no=VALUES(category_no)`, batchStringParam(l))
	fmt.Printf("sql:%v\n", sqlStr)
	query, args, err := sqlx.In(sqlStr, categories.ConvInterfaceArray()...)
	query += ",update_time=?"
	args = append(args, time.Now())

	fmt.Printf("sql:%v\n", query)
	fmt.Printf("args:%+v\n", args)
	if err != nil {
		return
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return
	}
	rs, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rs < 1 {
		err = fmt.Errorf("rs not equal 1")
		return
	}
	return
}

//DelArticleCategories 删除文章分类
func DelArticleCategories(ids []int) (err error) {
	sqlStr := "delete from article_category where category_id in (?) "
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return
	}
	rs, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rs < 1 {
		zlog.Warn("delete rowsAffected not larger than 1")
	}
	return nil
}

//ExistArticleCategory 判断分类是否存在
func ExistArticleCategory(id int) (bool, error) {
	var rac int
	sqlStr := "select category_id from article_category where category_id=? "
	err := db.Get(&rac, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}
