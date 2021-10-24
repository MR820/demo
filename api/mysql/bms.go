/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/8/28
 * Time 2:12 下午
 */

package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "bms"
	PASSWORD = "GH5tzde73r2KWpfZ"
	NETWORK  = "tcp"
	SERVER   = "81.68.69.84"
	PORT     = 3306
	DATABASE = "test"
)

var DbBms = &sql.DB{}

func init() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DbBms, _ = sql.Open("mysql", conn)
	//defer db.Close()
	DbBms.SetMaxOpenConns(50)
	DbBms.SetMaxIdleConns(10)
}
