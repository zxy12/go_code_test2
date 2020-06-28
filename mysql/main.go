package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	user := "root"
	pwd := "root"
	host := "127.0.0.1:3307"
	dbname := "test"

	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=500ms&readTimeout=%dms&writeTimeout=%dms&interpolateParams=true", user, pwd, host, dbname, 5000, 2000)

	db, err := sql.Open("mysql", str)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	db.SetMaxOpenConns(1)

	ch := make(chan int)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover->%+v\n", r)
			}
			r, err := db.Query("select 1")
			fmt.Printf("ret=%v, %v\n", r, err)
			ch <- 1
			time.Sleep(5 * time.Second)
			r.Close()
		}()
	}()

	<-ch
	fmt.Println("1....")
	r, err := db.Query("select 2")
	fmt.Println("x....")

	var id int
	for r.Next() {
		r.Scan(&id)
		fmt.Println(id)
	}

	fmt.Println("end.")
}
