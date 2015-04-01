/*
* @Author: xeodou
* @Date:   2015
* @Last Modified by:   xeodou
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := "PRAGMA key = '123456';"
	_, err = db.Exec(p)
	if err != nil {
		fmt.Println(err)
	}
	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, UNIQUE(`name`));"
	_, err = db.Exec(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := "INSERT INTO `users` (name, password) values('xeodou', 123456);"
	_, err = db.Exec(d)
	if err != nil {
		fmt.Println(err)
	}

	e := "select name, password from users where name='xeodou';"
	rows, err := db.Query(e)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var password string
		rows.Scan(&name, &password)
		fmt.Print("{\"name\":\"" + name + "\", \"password\": \"" + password + "\"}")
	}
	rows.Close()
}
