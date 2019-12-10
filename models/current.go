package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"path/filepath"
)

//var mypath = "/home/myproject/src/superly.club/web/"
const (
	//UpDownFile = "/home/data/"
	UpDownFile = "D:/go/src/superly.club/web/file/"
)


func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetAP(path string) string{
	//后台程序获取不到gopath
	//return filepath.Join(os.Getenv("GOPATH"), "src/superly.club/web/" + path)
	return filepath.Join("E:/code/MyWeb/" + path)
	//return  mypath + path
}

func SendOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
	})
}

func SendErr(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
	})
}
