package models

import (
	"os"
	"log"
	"io"
	"mime/multipart"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"strconv"
)

type FileManage interface{
	Post() error
	Put() error
	Del() error
	Get() error
	List() (*[]File, error)
}

type File struct {
	Path string `json:"path"`
	//Context  *gin.Context
	File multipart.File	`json:"file"`
	Info string  `json:"info"`
	Type string  `json:"type"`
	Size string `json:"size"`
}

func (f *File)Post() error {
	out, err := os.Create(UpDownFile + f.Path)
	if err != nil {
		fmt.Printf("create file Error!, err:%s", err)
		log.Fatal(err)
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, f.File)
	if err != nil {
		fmt.Printf("copy file Error!, err:%s", err)
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *File)Put() error {
	d1 := []byte(f.Info)
	err := ioutil.WriteFile(UpDownFile + f.Path, d1, 0644)
	if err!=nil{
		fmt.Printf("write file Error!, err:%s", err)
		log.Fatal(err)
		return err
	}

	return nil
}

func (f *File)Del() error {
	err := os.Remove(UpDownFile + f.Path)
	if err != nil {
		fmt.Printf("file remove Error!, err:%s", err)
		return err
	}

	return nil
}

func (f *File)Get() error {
	context, err := ioutil.ReadFile(UpDownFile + f.Path)
	if err != nil {
		fmt.Printf("read file Error!, err:%s", err)
		return err
	}
	f.Info = string(context)
	return nil
}

func  (f *File)List() (*[]File, error) {
	lm := make([]File, 0)
	//遍历目录，读出文件名、大小
	err := filepath.Walk(UpDownFile + f.Path, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return err
		}
		//	fmt.Println(fi.Name(), fi.Size()/1024)
		var m File
		m.Path = fi.Name()
		m.Size = strconv.FormatInt(fi.Size()/1024, 10)
		lm = append(lm, m)
		return nil
	})
	if err != nil{
		return &lm, err
	}

	return &lm, nil
}