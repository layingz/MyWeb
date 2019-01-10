package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	. "superly.club/web/models"
	"net/http"
	"os"
	"net/url"
)

func UpFile(c *gin.Context){
	file, header, err := c.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		SendErr(c)
		return
	}

	fileType := c.PostForm("type")
	fileName := header.Filename
	fileSize := header.Size
	fileTime := GetTime()

	fmt.Printf("t:%s,n:%s,s:%s,ct:%s\n", fileType, fileName, fileSize, fileTime)


	if _, err = os.Stat(UpDownFile+ fileType); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(UpDownFile+ fileType, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
				SendErr(c)
				return
			} else {
				fmt.Printf("mkdir success!\n")
			}
		}
	}

	if fileSize > 1024*1024*1024 {
		fmt.Printf("filr to big!\n")
		c.String(http.StatusForbidden, "file size too big!")
		return
	}

	WebLog.Printf("type:%s,name:%s,size:%d\n", fileType, fileName, fileSize)
	fmt.Printf("type:%s,name:%s,size:%d\n", fileType, fileName, fileSize)

	//先写数据库
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("INSERT INTO file(type, name, size, ctime) VALUES ('%s', '%s', '%d', '%s')", fileType, fileName, fileSize, fileTime)
	err = sql.Insert()
	if err != nil {
		SendErr(c)
		return
	}

	//再写文件
	filePath := fmt.Sprintf("%s/%s", fileType, fileName)
	filePost := File{Path: filePath, File: file}
	err = filePost.Post()
	if err != nil {
		sql.Sql = fmt.Sprintf("DELETE FROM file WHERE name='%s'", fileName)
		err = sql.ModDel()
		if err != nil {
			WebLog.Printf("delete file sql failed!,info[type;%s,name:%s,size:%d,ctime:%s]\n", fileType, fileName, fileSize, fileTime)
			SendErr(c)
			return
		}

		SendErr(c)
		return
	}

	SendOk(c)
}

func ListFile(c *gin.Context){
	filePath := c.DefaultQuery("path", "")

	file := File{Path: filePath}
	l, err := file.List()
	if err != nil {
		SendErr(c)
		return
	}

	for _, f := range *l{
		fmt.Println(f.Path, f.Size)
	}

	c.JSON(http.StatusOK, *l)
}

func DFile(c *gin.Context){
	fileName := c.Param("name")

	fmt.Printf("get filename:%s\n", fileName)

	//数据库查询类型
	var fileType string
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("select type from file where name='%s'", fileName)
	err := sql.Query(&fileType)
	if err != nil || fileName == ""{
		fmt.Printf("query failed!")
		SendErr(c)
		return
	}

	path := fmt.Sprintf("%s/%s", fileType, fileName)
	fmt.Printf("df path:%s\n", path)


	c.Header("Content-Disposition", "attachment; filename="+ url.PathEscape(fileName))
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	c.File(UpDownFile + path)

/*
	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= " + fileName}

	file, err := os.Open(UpDownFile + path)
	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}
	defer file.Close()

	io.Copy(c.Writer, file)
	*/
}

func DelFile(c *gin.Context){
	fileName := c.Query("name")
	fmt.Printf("name:%s\n", fileName)

	//数据库查询类型
	var fileType string
	sql := Mysql{}
	sql.Sql = fmt.Sprintf("select type from file where name='%s'", fileName)
	err := sql.Query(&fileType)
	if err != nil || fileType == "" {
		fmt.Printf("query failed!")
		SendErr(c)
		return
	}

	//删除索引
	sql.Sql = fmt.Sprintf("DELETE FROM file WHERE name='%s'", fileName)
	err = sql.ModDel()
	if err != nil {
		WebLog.Printf("delete file sql failed!,info[type;%s,name:%s,ctime:%s]\n", fileType, fileName)
		fmt.Printf("delete file sql failed!,info[type;%s,name:%s,ctime:%s]\n", fileType, fileName)
		SendErr(c)
		return
	}

	filePath := fmt.Sprintf("%s/%s", fileType, fileName)
	fileOp := File{Path: filePath}
	err = fileOp.Del()
	if err != nil {
		SendErr(c)
		return
	}

	SendOk(c)
}