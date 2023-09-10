package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/goland"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

func query() {

}
func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")

	var u user

	sqlStr := "select id, name, age from user where id=?;"

	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	rowObj := db.QueryRow(sqlStr, 1)
	rowObj.Scan(&u.id, &u.name, &u.age)

	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
