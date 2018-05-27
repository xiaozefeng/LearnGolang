package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)
func init() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/foo?charset=utf8", 30)
	// 注册自定义model
	orm.RegisterModel(new(Userinfo))
	//创建table
	orm.RunSyncdb("default",false,true)
	orm.Debug=true
}

type Userinfo struct {
	Id     int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username    string
	Departname  string
	Created     time.Time
}

func main() {
	o := orm.NewOrm()
	u := Userinfo{Username:"rose",Departname:"研发",Created:time.Now()}
	id, err:=o.Insert(&u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

}

