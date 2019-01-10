package models

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type MysqlHandle interface {
	Insert() error
	ModDel() error
	Query(QueryResult ...interface{}) error
	Search(count int) (* []MsqlMap, error)
}

type Mysql struct {
	Sql string
}

type MsqlMap struct {
	First	string
	Second	string
	Third	string
	Fourth	string
	Fifth	string
	count	int
}

func (m *Mysql)Insert() error {
	rs, err := SqlDB.Exec(m.Sql)
	if err != nil {
		fmt.Printf("sql exec failed!")
		return err
	}
	id, err := rs.LastInsertId()
	fmt.Println(id)
	if err!=nil{
		fmt.Printf("sql insert failed!")
		return err
	}
	return nil
}

func(m *Mysql)ModDel() error {
	rs, err := SqlDB.Exec(m.Sql)
	if err != nil {
		return err
	}
	id, err := rs.RowsAffected()
	fmt.Println(id)
	if err!=nil{
		return err
	}
	return nil
}


func(m *Mysql)Query(QueryResult ...interface{}) error {
	err := SqlDB.QueryRow(m.Sql).Scan(QueryResult ...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func(m *Mysql)Search(count int) (* []MsqlMap, int, error) {
	rows, err := SqlDB.Query(m.Sql)
	if err != nil {
		fmt.Printf("query err...")
		return nil,  0, nil
	}
	defer rows.Close()

	var r []MsqlMap
	reslutCount := 0
	switch count {
	case 1:
		for rows.Next() {
			var info MsqlMap
			rows.Scan(&info.First)
			r = append(r, info)
			reslutCount++
			}
	case 2:
		for rows.Next() {
			var info MsqlMap
			rows.Scan(&info.First, &info.Second)
			r = append(r, info)
			reslutCount++
			}
	case 3:
		for rows.Next() {
			var info MsqlMap
			rows.Scan(&info.First, &info.Second, &info.Third)
			r = append(r, info)
			reslutCount++
			}
	case 4:
		for rows.Next() {
			var info MsqlMap
			rows.Scan(&info.First, &info.Second, &info.Third, &info.Fourth)
			r = append(r, info)
			reslutCount++
			}
	case 5:
		for rows.Next() {
			var info MsqlMap
			rows.Scan(&info.First, &info.Second, &info.Third, &info.Fourth, &info.Fifth)
			r = append(r, info)
			reslutCount++
			}
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("parse err...")
		return nil, 0 , err
	}
	return &r, reslutCount, nil
}

var SqlDB *sql.DB

func MysqlInit() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:zly802311@tcp(127.0.0.1:3306)/web?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

/*
创建目录：
create DATABASE web;
use web;

用户：
CREATE TABLE user(Id INT UNSIGNED AUTO_INCREMENT, name VARCHAR(256) NOT NULL, passwd VARCHAR(100) NOT  NULL,type VARCHAR(64) NOT NULL,ctime VARCHAR(100) NOT NULL, PRIMARY KEY (Id))ENGINE=InnoDB DEFAULT CHARSET=utf8;
文章表：
CREATE TABLE article(Id INT UNSIGNED AUTO_INCREMENT, type VARCHAR(100) NOT NULL, title VARCHAR(100) NOT NULL,content TEXT NOT NULL, ctime VARCHAR(100) NOT NULL, PRIMARY KEY (Id))ENGINE=InnoDB DEFAULT CHARSET=utf8;
网站表：
CREATE TABLE web(Id INT UNSIGNED AUTO_INCREMENT, type VARCHAR(100) NOT NULL, name VARCHAR(100) NOT NULL,address VARCHAR(256) NOT NULL,ctime VARCHAR(100) NOT NULL, PRIMARY KEY (Id))ENGINE=InnoDB DEFAULT CHARSET=utf8;
文件索引：
CREATE TABLE file(Id INT UNSIGNED AUTO_INCREMENT, type VARCHAR(64) NOT NULL, name VARCHAR(128) NOT  NULL,size VARCHAR(64) NOT NULL,ctime VARCHAR(100) NOT NULL, PRIMARY KEY (Id))ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/