package routers

import (
	"sz_resume_202005/routers/api"
	"sz_resume_202005/routers/api/apipublic"
	v1 "sz_resume_202005/routers/api/v1"
	"sz_resume_202005/utils/middleware"

	"github.com/gin-gonic/gin"
)

//LoadRouters 载入路由
func LoadRouters(r *gin.Engine) {

	API := r.Group("/api")
	API.POST("/login", api.Login)

	AUTH := API.Group("/auth")
	AUTH.Use(middleware.CheckToken())

	loadBack(AUTH)
	loadFront(API)
}

func loadBack(g *gin.RouterGroup) {
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.POST("/logout", api.Logout)
	g.GET("/userinfo", api.GetUserInfo)
	g.PUT("/userinfo", api.EditUserInfo)

	g.POST("/experiences", api.AddExperiences)
	g.GET("/experiences", api.GetExperiences)
	g.PUT("/experiences", api.EditExperience)
	g.DELETE("/experiences", api.DelExperiences)

	//技能
	//-获取全部技能
	g.GET("/skills", api.GetSkills)
	//-添加技能
	g.POST("/skills", api.AddSkill)
	//-编辑技能
	g.PUT("/skills/id/:id", api.EditSkill)
	//-批量编辑技能
	g.PUT("/skills", api.EditSkill)
	//-删除技能
	g.DELETE("/skills", api.DelSkills)

	//软件
	//-获取全部软件
	g.GET("/softwares", api.GetSkills)
	//-添加软件
	g.POST("/softwares", api.AddSoftware)
	//-编辑软件
	g.PUT("/softwares/id/:id", api.EditSoftware)
	//-批量编辑软件
	g.PUT("/softwares", api.EditSoftware)
	//-删除软件
	g.DELETE("/softwares", api.DelSoftwares)

	//文章  ？加入审核，根据状态获取全部文章，管理员读取文章详情
	//-根据用户获取文章列表
	g.GET("/articles", v1.GetArticle)
	//-获取文章详情
	g.GET("/articles/id/:id", v1.GetArticleDetail)
	//-添加文章
	g.POST("/articles", v1.AddArticle)
	//-编辑文章
	g.PUT("/articles/id/:id", v1.EditArticle)
	//-删除文章
	g.DELETE("/articles", v1.DelArticle)
	//-
	//-管理员审核文章
	g.POST("/articles/a/check", v1.RACheckArticle)
	//-管理员根据状态读取文章列表
	g.GET("/articles/a", v1.RAGetArticle)
	//-管理员读取文章详情
	g.GET("/articles/a/id/:id", v1.RAGetArticleDetail)
	//-管理员删除文章
	g.DELETE("/articles/a", v1.RADelArticle)

	//文章分类
	//-获取全部分类
	g.GET("/articles/categories", v1.GetArticleCategory)
	//-添加分类
	g.POST("/articles/categories", v1.AddArticleCategory)
	//-编辑分类
	g.PUT("/articles/categories/id/:id", v1.EditArticleCategory)
	//-批量编辑分类
	g.PUT("/articles/categories", v1.BatchEditArticleCategory)
	//-删除分类
	g.DELETE("/articles/categories", v1.DelArticleCategory)

	//作品   ？加入审核，根据状态获取全部作品，管理员读取作品详情
	//-根据用户获取作品列表
	g.GET("/works", v1.GetWork)
	//-获取作品详情
	g.GET("/works/id/:id", v1.GetWorkDetail)
	//-添加作品
	g.POST("/works", v1.AddWork)
	//-编辑编辑
	g.PUT("/works/id/:id", v1.EditWork)
	//-删除作品
	g.DELETE("/works", v1.DelWork)
	//-
	//-管理员审核作品
	g.POST("/works/a/check", v1.RACheckWork)
	//-管理员根据状态读取作品列表
	g.GET("/works/a", v1.RAGetWork)
	//-管理员读取作品详情
	g.GET("/works/a/id/:id", v1.RAGetWorkDetail)
	//-管理员删除作品
	g.DELETE("/works/a", v1.RADelWork)

	//作品标签
	//-获取全部标签
	g.GET("/works/tags", v1.GetWorkTag)
	//-添加标签
	g.POST("/works/tags", v1.AddWorkTag)
	//-编辑标签
	g.PUT("/works/tags/id/:id", v1.EditWorkTag)
	//-批量编辑标签
	g.PUT("/works/tags", v1.BatchEditWorkTag)
	//-删除标签
	g.DELETE("/works/tags", v1.DelWorkTag)

	//权限管理
	//-给予权限
	//-删除权限
	g.POST("/upload", api.UploadImage)

}

func loadFront(g *gin.RouterGroup) {

	//前端不需要验证token，即可返回
	//首页
	g.GET("/userinfocpt", apipublic.GetUserInfoCpt)

	//技能和软件
	g.GET("/tec", apipublic.GetTec)
	//文章列表
	g.GET("/articles", apipublic.GetArticles)
	//文章详情
	g.GET("/articles/:id", apipublic.GetArticleByID)
	//作品列表
	g.GET("/works", apipublic.GetWork)
	//作品详情
	g.GET("/works/:id", apipublic.GetWorkByID)

}
