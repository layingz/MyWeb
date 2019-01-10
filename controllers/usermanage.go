package controllers

import (
	"github.com/gin-gonic/gin"
	. "superly.club/web/models"
	"fmt"
)

func Register(c *gin.Context) {
	userName := c.PostForm("username")
	passWord := c.PostForm("password")

	sql := Mysql{}
	sql.Sql = fmt.Sprintf("INSERT INTO user(name, passwd, type，ctime) VALUES ('%s', '%s', '%s')", userName, passWord, "admin",GetTime())
	err := sql.Insert()
	if err != nil {
		SendErr(c)
		return
	}

	SendOk(c)
	return
}

func Login(c * gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	//WebLog.Printf("name:%s,passwd:%s\n", username, password)

	var value string
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("select passwd from user where name='%s'", username)
	err := sql.Query(&value)
	if err != nil || value == ""{
		SendErr(c)
		return
	}

	WebLog.Printf("value:%s\n", value)

	if value == password{
		SendOk(c)
		return
	}

	SendErr(c)
	return
}