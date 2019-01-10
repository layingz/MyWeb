package controllers

import (
	"github.com/gin-gonic/gin"
	. "superly.club/web/models"
	"fmt"
)

type WebInfo struct {
	Type	string
	Name	string
	Address string
	Ctime 	string
}

func PostWeb(c *gin.Context){

	webType := c.PostForm("type")
	webName := c.PostForm("name")
	webAddress := c.PostForm("address")

	WebLog.Printf("type:%s,name:%s,ad:%s\n", webType, webName, webAddress)

	sql := Mysql{}
	sql.Sql = fmt.Sprintf("INSERT INTO web(type, name, address, ctime) VALUES ('%s', '%s', '%s', '%s')", webType, webName, webAddress, GetTime())
	err := sql.Insert()
	if err != nil {
		SendErr(c)
		return
	}

	SendOk(c)
	return
}

func GetWeb(c *gin.Context){

	webType := c.Query("type")

	sql := Mysql{}
	if webType == "" {
		sql.Sql = fmt.Sprintf("select type,name,address,ctime from web")
		r, _, err := sql.Search(4)
		if err != nil {
			SendErr(c)
			return
		}

		for _, i := range *r {
			println(i.First, i.Second, i.Third, i.Fourth)
		}

	}else {
		sql.Sql = fmt.Sprintf("select name,address,ctime from web where type = '%s'", webType)
		r, _, err := sql.Search(3)
		if err != nil {
			SendErr(c)
			return
		}

		for _, i := range *r {
			println(i.First, i.Second, i.Third)
		}
	}
}

func DelWeb(c *gin.Context){
	WebName := c.Query("name")
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("DELETE FROM web WHERE name='%s'", WebName)
	err := sql.ModDel()
	if err != nil {
		SendErr(c)
		return
	}

	SendOk(c)
	return
}