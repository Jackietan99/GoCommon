package common

import (
	"database/sql" //这包一定要引用，是底层的sql驱动
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Mydb struct {
	*sql.DB
}

var SqlDebug bool

/**
* 打开数据库连接池
* @param		conStr	连接字符串  $user:$pwd@tcp($ip:$port)/$dbname?charset=$charset
					用户名:密码@tcp(IP：端口)/数据库名?charset=字符集
* @param		debug	是否开启调试
*/
func ConnectMysql(conStr string, debug bool) *Mydb {
	var err error
	SqlDebug = debug

	//db并不是一个连接，而是管理数据库的句柄，这里开启进程池配置
	db := new(Mydb)
	db.DB, err = sql.Open("mysql", conStr)

	if err != nil {
		//连接池失败了
		i := 0
		for {
			db.DB, err = sql.Open("mysql", conStr)
			if err == nil {
				if SqlDebug {
					fmt.Println("连接数据库" + conStr + "，连接成功")
				}
				break
			} else {
				//如果连接数据库失败了，这里不停的重新连接直到成功返回
				if SqlDebug {
					fmt.Printf("连接数据库"+conStr+"，连接失败，错误信息:"+err.Error()+"，第%d次重新连接\n", i+1)
				}
				i = i + 1
			}
			if i >= 30 {
				time.Sleep(2 * time.Minute)
				i = 0
			}
		}
	} else {
		db.SetMaxOpenConns(20000)
		db.SetMaxIdleConns(10000)
		db.SetConnMaxLifetime(5 * time.Second)
		if SqlDebug {
			fmt.Println("连接数据库" + conStr + "，连接成功")
		}
	}
	return db
}

/**
* 检测数据库连接是否正常
* return true=正常
 */
func (db *Mydb) checkConn() bool {
	res := false
	var err error
	for i := 0; i <= 5; i++ {
		err = db.Ping()
		if err == nil {
			res = true
			break
		} else {
			if SqlDebug {
				fmt.Println("conn->", err.Error())
			}
			LogsWithFileName("", "SQL_", "conn:"+err.Error())
		}
	}
	return res
}
