package main

import (
	. "superly.club/web/routers"
	"github.com/gin-gonic/gin"
	. "superly.club/web/models"
	"net/http"
)

func addststic(router *gin.Engine){
	router.StaticFS("/showDir", http.Dir(GetAP(".")))
	//router.Static("/static", "./static")
	router.StaticFile("/1.jpg", GetAP("static/img/1.jpg"))
	router.StaticFile("/2.jpg", GetAP("static/img/2.jpg"))
	router.StaticFile("/3.jpg", GetAP("static/img/3.jpg"))
	router.StaticFile("/4.jpg", GetAP("static/img/4.jpg"))

	router.StaticFile("/filemanage.html", GetAP("views/filemanage.html"))
	router.StaticFile("/articlemanage.html", GetAP("views/articlemanage.html"))
	router.StaticFile("/webmanage.html", GetAP("views/webmanage.html"))
}

func init(){
	MysqlInit()
	WebLogInit()
	gin.SetMode(gin.DebugMode)
}

func main() {
	defer SqlDB.Close()
	defer WebLog.Close()

	router := gin.Default()

	//路由信息
	InitRouter(router)

	//静态资源
	addststic(router)

	//运行的端口
	router.Run("127.0.0.1:8000")
}

