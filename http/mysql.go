package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	Db *sql.DB
)

func init() {
	Db, _= sql.Open("mysql", "root:root@/foo?charset=utf8")
}

func main() {
	insert(UserInfo{1, "jack", "研发部门", "2018-05-27"})
	//update(*Db, UserInfo{id: 1, Name: "rose"})
	//remove(*Db, 1)
	findAll()
	defer Db.Close()

}

type UserInfo struct {
	id         int64
	Name       string
	Department string
	Created    string
}

func insert(info UserInfo) int64 {
	//插入数据
	stmt, err := Db.Prepare("insert userinfo set username=?,depart_name=?,created=?")
	checkError(err)

	res, err := stmt.Exec(info.Name, info.Department, info.Created)
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println("id:", id)
	return id
}

func update(info UserInfo) {
	stmt, err := Db.Prepare("update userinfo set username=? where uid =?")
	checkError(err)
	result, err := stmt.Exec(info.Name, info.id)
	affect, err := result.RowsAffected()
	checkError(err)
	fmt.Println(affect)
}

func remove(id int64) {
	stmt, err := Db.Prepare("delete from userinfo where uid=?")
	checkError(err)
	res, err := stmt.Exec(id)
	checkError(err)
	affect, err := res.RowsAffected()
	checkError(err)
	if affect == 1 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}

}

func findAll() UserInfo {
	rows, err := Db.Query("select * from userinfo")
	checkError(err)
	for rows.Next() {
		u := UserInfo{}
		err = rows.Scan(&u.id, &u.Name, &u.Department, &u.Created)
		checkError(err)
		fmt.Println(u)
	}

	return UserInfo{}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
