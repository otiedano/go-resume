package routers

import (
	"sz_resume_202005/routers/api"
	"sz_resume_202005/routers/api/apipublic"
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
	//-添加技能
	//-编辑技能
	//-批量编辑技能
	//-删除技能

	//软件
	//-获取全部软件
	//-添加软件
	//-编辑软件
	//-批量编辑软件
	//-删除软件

	//文章  ？加入审核，根据状态获取全部文章，管理员读取文章详情
	//-
	//文章分类
	//-获取全部分类
	//-添加分类
	//-编辑分类
	//-批量编辑分类
	//-删除分类

	//作品   ？加入审核，根据状态获取全部作品，管理员读取作品详情
	//作品标签

	//权限管理
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
