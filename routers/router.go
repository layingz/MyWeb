package routers

import (
	"github.com/gin-gonic/gin"
	. "superly.club/web/models"
	. "superly.club/web/controllers"
	"net/http"
)

func InitRouter(router *gin.Engine) {


	router.LoadHTMLGlob(GetAP("views/*"))

	//起始页
	router.GET("/", func(c *gin.Context){c.HTML(http.StatusOK, "index.html", gin.H{})})

	//用户
	router.GET("/login", func(c *gin.Context){c.HTML(http.StatusOK, "login.html", gin.H{})})
	router.POST("/login", Login)
	//router.POST("/register", Register)

	//管理后台
	router.GET("/manage", func(c *gin.Context){c.HTML(http.StatusOK, "manage.html", gin.H{})})

	//文章管理
	router.POST("/articlemanage", PostArticle)
	router.DELETE("/articlemanage", DelArticle)
	router.GET("/article", ListArticle)
	router.GET("/study", ListStudy)
	router.GET("/explore",ListExplore)
	router.GET("/article/:name", GetArticle)

	//文件管理
	router.GET("/download/:name", DFile)
	router.POST("/fileup", UpFile)
	router.GET("/dir", ListFile)
	router.DELETE("/filedel", DelFile)

	//网页管理
	router.POST("/webmanage", PostWeb)
	router.GET("/webmanage", GetWeb)
	router.DELETE("/webmanage", DelWeb)
}
