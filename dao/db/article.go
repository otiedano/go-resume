package db

import (
	"database/sql"
	"fmt"
	"sz_resume_202005/model"
	"time"
	"unsafe"

	"github.com/jmoiron/sqlx"
)

//TotalArticle 计算文章总数
func TotalArticle() (num int, err error) {
	var sqlStr string

	sqlStr = "SELECT COUNT(article_id) from article where status=1"

	err = db.Get(&num, sqlStr)
	return
}

//TotalArticleByCategory 计算文章总数
func TotalArticleByCategory(categoryID int) (num int, err error) {
	var sqlStr string
	if categoryID == 0 {
		sqlStr = "SELECT COUNT(article_id) from article where status=1"
		err = db.Get(&num, sqlStr)
	} else {
		sqlStr = "SELECT COUNT(article_id) from article where category_id=? and status=1"
		err = db.Get(&num, sqlStr, categoryID)
	}

	return
}

//TotalArticleByAuthor 计算文章总数
func TotalArticleByAuthor(userID int) (num int, err error) {

	sqlStr := "SELECT COUNT(article_id) from article where author_id=? and status!=2"
	err = db.Get(&num, sqlStr, userID)
	return

}

//TotalArticleByStatus 计算文章总数
func TotalArticleByStatus(args ...interface{}) (num int, err error) {
	var sqlStr string

	if len(args) == 0 {
		sqlStr = "SELECT COUNT(article_id) from article"
		err = db.Get(&num, sqlStr)
		return
	}
	if n, ok := args[0].(int); ok && n >= 0 && n <= 2 {
		sqlStr = "SELECT COUNT(article_id) from article where status=? "
		err = db.Get(&num, sqlStr, n)
		return
	}

	sqlStr = "SELECT COUNT(article_id) from article"
	err = db.Get(&num, sqlStr)
	return
}

//GetArticle 读取审核过的文章详情
func GetArticle(articleID int) (article *model.ArticleDetail, err error) {
	article = &model.ArticleDetail{}
	sqlStr := `
	SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no,a.content 
	FROM article a 
	LEFT JOIN article_category c ON a.category_id=c.category_id
	LEFT JOIN user u ON a.author_id=u.user_id 
	WHERE a.article_id=? and a.status=1
	`
	err = db.Get(article, sqlStr, articleID)
	return
}

//GetPArticle 读取文章详情
func GetPArticle(userID, articleID int) (article *model.ArticleDetail, err error) {
	article = &model.ArticleDetail{}
	sqlStr := `
	SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no,a.content 
	FROM article a 
	LEFT JOIN article_category c ON a.category_id=c.category_id 
	LEFT JOIN user u ON a.author_id=u.user_id 
	WHERE a.article_id=? And a.author_id=?
	`
	err = db.Get(article, sqlStr, articleID, userID)
	return
}

//GetRPArticle 读取文章详情
func GetRPArticle(articleID int) (article *model.ArticleDetail, err error) {
	article = &model.ArticleDetail{}
	sqlStr := `
	SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no,a.content 
	FROM article a 
	LEFT JOIN article_category c ON a.category_id=c.category_id 
	LEFT JOIN user u ON a.author_id=u.user_id 
	WHERE a.article_id=? 
	`
	err = db.Get(article, sqlStr, articleID)
	return
}

//GetArticlesByAuthor 根据作者返回文章列表,用于后端管理
func GetArticlesByAuthor(authorID int, offset, limit int) (articles []*model.Article, err error) {
	sqlStr := `
	SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,c.category_name,c.category_no,a.category_id
	
	 FROM article AS a 
	LEFT JOIN article_category AS c ON a.category_id=c.category_id 
	LEFT JOIN user AS u ON a.author_id=u.user_id 
	WHERE a.author_id=? AND a.status!=2
	ORDER BY a.update_time desc,a.article_id desc
	LIMIT ? OFFSET ?
	`

	err = db.Select(&articles, sqlStr, authorID, limit, offset)
	return
}

//GetAllArtilesByStatus 根据status返回所有文章列表。
func GetAllArtilesByStatus(offset, limit int, args ...interface{}) (articles []*model.Article, err error) {
	if len(args) == 0 {
		sqlStr := `
SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no
FROM article a 
LEFT JOIN article_category c ON a.category_id=c.category_id
LEFT JOIN user u ON a.author_id=u.user_id 
ORDER BY a.update_time desc,a.article_id desc
LIMIT ? OFFSET ?
`

		err = db.Select(&articles, sqlStr, limit, offset)
		if err != nil {
			return nil, err
		}
		return
	}
	if n, ok := args[0].(int); ok && n >= 0 && n <= 2 {
		sqlStr := `
		SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no
		FROM article a 
		LEFT JOIN article_category c ON a.category_id=c.category_id
		LEFT JOIN user u ON a.author_id=u.user_id 
		WHERE a.status=? 
		ORDER BY a.update_time desc,a.article_id desc
		LIMIT ? OFFSET ?
		`

		err = db.Select(&articles, sqlStr, n, limit, offset)
		if err != nil {
			return nil, err
		}
		return
	}
	sqlStr := `
SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no
FROM article a 
LEFT JOIN article_category c ON a.category_id=c.category_id
LEFT JOIN user u ON a.author_id=u.user_id 
ORDER BY a.update_time desc,a.article_id desc
LIMIT ? OFFSET ?
`

	err = db.Select(&articles, sqlStr, limit, offset)
	if err != nil {
		return nil, err
	}
	return
}

//GetArticlesByCategory 根据分页和分类返回文章列表，用于前端显示
func GetArticlesByCategory(categoryID int, offset, limit int) (articles []*model.Article, err error) {
	if categoryID == 0 {
		//ORDER BY 首先需要a.update_time desc,
		sqlStr := `
SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no
FROM article a 
LEFT JOIN article_category c ON a.category_id=c.category_id
LEFT JOIN user u ON a.author_id=u.user_id 
WHERE a.status=1
ORDER BY a.update_time desc,a.article_id desc
LIMIT ? OFFSET ?
`

		err = db.Select(&articles, sqlStr, limit, offset)
	} else {
		//ORDER BY 首先需要a.update_time desc,
		sqlStr := `
SELECT a.article_id,a.title,a.img,a.view_count,a.status,a.summary,a.create_time,a.update_time,a.author_id,u.user_name,u.avatar,a.category_id,c.category_name,c.category_no
FROM article a 
LEFT JOIN article_category c ON a.category_id=c.category_id 
LEFT JOIN user u ON a.author_id=u.user_id 
WHERE a.category_id=? and a.status=1
ORDER BY a.update_time desc,a.article_id desc
LIMIT ? OFFSET ?
`
		err = db.Select(&articles, sqlStr, categoryID, limit, offset)
	}

	return
}

//AddArticle 新增文章
func AddArticle(userID int, article *model.ArticleDetail) (intNum int, err error) {
	sqlStr := `INSERT INTO article (title,content,img,summary,category_id,author_id) VALUES (?,?,?,?,?,?)`
	rs, err := db.Exec(sqlStr, article.Title, article.Content, article.Img, article.Summary, article.CategoryID, userID)
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

//EditArticle 修改文章
func EditArticle(userID int, article *model.ArticleDetail) (err error) {
	sqlStr := `UPDATE article SET title=?,content=?,img=?,summary=?,category_id=?,update_time=?  WHERE article_id=? AND author_id=?`

	result, err := db.Exec(sqlStr, article.Title, article.Content, article.Img, article.Summary, article.CategoryID, time.Now(), article.ArticleID, userID)
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

//DelArticlesFE 删除文章
func DelArticlesFE(ids []int) (err error) {
	sqlStr := "delete from article a where a.article_id in (?) "
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
		return fmt.Errorf("delete rowsAffected not larger than 1")
	}
	return nil
}

//DelArticles 删除文章
func DelArticles(userID int, ids []int) (err error) {
	sqlStr := "UPDATE article SET status=2 where article_id in (?) and author_id=?"
	query, args, err := sqlx.In(sqlStr, ids, userID)
	if err != nil {
		return
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return
	}

	return nil
}

//ExistArticleByAuth 检查是否有操作文章的权限
func ExistArticleByAuth(userID, id int) (bool, error) {
	sqlStr := "select article_id from article where article_id=? and author_id=? and status!=2"
	var rarticle model.Article
	err := db.Get(&rarticle, sqlStr, id, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

//ExistArticleByID 前端判断文章是否存在
func ExistArticleByID(id int) (bool, error) {
	sqlStr := "select article_id from article where article_id=? and status=1 "
	var rarticle model.Article
	err := db.Get(&rarticle, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

//ExistArticle 管理员查看是否存在
func ExistArticle(id int) (bool, error) {
	sqlStr := "select a.article_id from article a where a.article_id=? "
	var rarticle model.Article
	err := db.Get(&rarticle, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

//CheckArticles 文章审核
func CheckArticles(ids []int, status int) (err error) {
	sqlStr := "UPDATE article a SET a.status=? where a.article_id in (?) "
	query, args, err := sqlx.In(sqlStr, status, ids)
	if err != nil {
		return
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return
	}

	return nil

}

//CountArticle 访问数量
func CountArticle(id int) (err error) {
	sqlStr := "update article set view_count=view_count+1 where article_id=?"
	_, err = db.Exec(sqlStr, id)
	return
}
