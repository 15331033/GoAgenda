package entities

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	//"database/sql"
)

var myEngine *xorm.Engine

func init() {
	fmt.Println("22222222")
	var err error
	myEngine, err = xorm.NewEngine("sqlite3","./foo.db")
	//myEngine, err = xorm.NewEngine("mysql","root:root@tcp(127.0.0.1:3308)/test?charset=utf8&parseTime=true")
	if err != nil {
        panic(err)
	}
	myEngine.SetMapper(core.SameMapper{})
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}