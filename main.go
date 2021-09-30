package main

import (
	"BMS/handlers"
	"fmt"

	"net/http"

	"github.com/akrylysov/algnhsa"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func preHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := connectToDatabase()
		if err != nil {
			fmt.Println("Cannot connect to database")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/home", preHandler(http.HandlerFunc(handlers.HomeHandler)))
	mux.Handle("/signup", preHandler(http.HandlerFunc(handlers.AuthHandler)))
	mux.Handle("/login", preHandler(http.HandlerFunc(handlers.AuthHandler)))
	mux.Handle("/theaters", preHandler(http.HandlerFunc(handlers.TheatersHandler)))
	mux.Handle("/theaters/", preHandler(http.HandlerFunc(handlers.TheatersHandler)))
	mux.Handle("/booking", preHandler(http.HandlerFunc(handlers.BookingHandler)))
	http.ListenAndServe(":9000", mux)
	algnhsa.ListenAndServe(mux, nil)
}

func connectToDatabase() (err error) {
	alias := "default"
	connString := "root:root@123@tcp(127.0.0.1:3306)/bms?interpolateParams=true&charset=utf8mb4"
	//check if db is already registered, if not register it
	_, err = orm.GetDB(alias)
	if err != nil {
		err = orm.RegisterDataBase(alias, "mysql", connString)
		return
	}
	//check if db is registered by now, if not continue to retry
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	//check if connection is still working, if not continue to retry
	err = MysqlTest(alias)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	fmt.Println("Mysql connection successfull")
	return
}

//MysqlTest ... function to test mysql connection
func MysqlTest(alias string) error {
	o := orm.NewOrm()
	o.Using(alias)
	_, err := o.Raw("SELECT 1").Exec()
	return err
}
