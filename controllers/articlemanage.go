package controllers

import (
	"github.com/gin-gonic/gin"
	. "superly.club/web/models"
	"fmt"
	"net/http"
)

func PostArticle(c *gin.Context){

	aType := c.PostForm("type")
	aTitle := c.PostForm("title")
	aContent := c.PostForm("content")

	WebLog.Printf("type:%s,title:%s,context:%s\n", aType, aTitle, aContent)

	sql := Mysql{}
	sql.Sql = fmt.Sprintf("INSERT INTO article(type, title, content, ctime) VALUES ('%s', '%s', '%s', '%s')", aType, aTitle, aContent, GetTime())
	err := sql.Insert()
	if err != nil {
		SendErr(c)
	}

	SendOk(c)
}

func DelArticle(c *gin.Context){
	aName := c.Query("name")
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("DELETE FROM article WHERE name='%s'", aName)
	err := sql.ModDel()
	if err != nil {
		SendErr(c)
	}

	SendOk(c)
}

func GetArticle(c *gin.Context) {
	fileName := c.Param("name")
	if fileName == ""{
		c.HTML(http.StatusOK, "article.html", gin.H{
			"exist": false,
		})
	}

	sql := Mysql{}
	sql.Sql = fmt.Sprintf("select content,ctime from article where title = '%s'", fileName)
	r, _,err := sql.Search(2)
	if err != nil {
		c.HTML(http.StatusOK, "article.html", gin.H{
			"exist": false,
		})
	}
	/*
	for _, i := range *r {
		println(i.First, i.Second, i.Third, i.Fourth)
	}
	*/
	c.HTML(http.StatusOK, "article.html", gin.H{
		"article": *r,
		"exist": true,
	})
}

func ListArticle(c *gin.Context){
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("select id, title, ctime from article")
	r, count, err := sql.Search(3)
	if err != nil {
		c.HTML(http.StatusOK, "articlelist.html", gin.H{
			"exist": false,
		})
	}

	c.HTML(http.StatusOK, "articlelist.html", gin.H{
		"articles": *r,
		"count": count,
		"exist": true,
	})
}

func ListStudy(c *gin.Context){
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("select id, title, ctime from article where type = 'study'")
	r, count,err := sql.Search(3)
	if err != nil {
		c.HTML(http.StatusOK, "articlelist.html", gin.H{
			"exist": false,
		})
	}
	c.HTML(http.StatusOK, "articlelist.html", gin.H{
		"articles": *r,
		"count": count,
		"exist": true,
	})
}

func ListExplore(c *gin.Context){
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("select id, title, ctime from article where type = 'explore'")
	r, count,err := sql.Search(3)
	if err != nil {
		c.HTML(http.StatusOK, "articlelist.html", gin.H{
			"exist": false,
		})
	}
	c.HTML(http.StatusOK, "articlelist.html", gin.H{
		"articles": *r,
		"count": count,
		"exist": true,
	})
}