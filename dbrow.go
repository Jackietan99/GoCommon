package common

import (
	"database/sql" //这包一定要引用，是底层的sql驱动
	"fmt"
	"strconv"
	"time"
)

/**
* 获得插入的sql语句
* @tableName string 要插入的数据表名
* @d map[string]string 要插入的数据map，格式：map["字段名"]=值
* return string sql查询语句
 */
func (db *Mydb) GetInsertSql(tableName string, d map[string]string) string {
	field := ""
	values := ""
	for k, v := range d {
		//拼接sql语句
		if field == "" {
			field = k
			values = "'" + v + "'"
		} else {
			field = field + "," + k
			values = values + "," + "'" + v + "'"
		}
	}
	sql := "insert IGNORE into " + tableName + "(" + field + ") values(" + values + ")"
	return sql
}

/**
* 获得插入的sql语句
* @tableName string 要插入的数据表名
* @d map[string]string 要插入的数据map，格式：map["字段名"]=值
* return string sql查询语句
 */
func (db *Mydb) GetInsertTrueSql(tableName string, d map[string]string) string {
	field := ""
	values := ""
	for k, v := range d {
		//拼接sql语句
		if field == "" {
			field = k
			values = "'" + v + "'"
		} else {
			field = field + "," + k
			values = values + "," + "'" + v + "'"
		}
	}
	sql := "insert into " + tableName + "(" + field + ") values(" + values + ")"
	return sql
}

/**
* 插入一条数据到数据表当中
* @tableName string 要插入的数据表名
* @d map[string]string 要插入的数据map，格式：map["字段名"]=值
* @return res int,id string 返回影响记录的条数和最后影响的id，插入成功res=1，id有可能为空不能做为判断依据
 */
func (db *Mydb) InsertRow(tableName string, d map[string]string) (int, string) {
	//记录开始时间
	start := time.Now()
	res := -1
	id := ""

	if db.checkConn() == false {
		return res, id
	}

	lens := len(d)
	if lens < 1 {
		return res, id
	}

	sqlStr := db.GetInsertSql(tableName, d)

	if SqlDebug {
		fmt.Println("sqlStr->", sqlStr)
	}

	sdb, err := db.Exec(sqlStr)
	if err != nil {
		LogsWithFileName("", "SQL_", "InsertRow Exec Error->"+err.Error()+"\n InsertRow SQL->"+sqlStr)
	} else {
		r, err := sdb.RowsAffected()
		if err == nil {
			res = int(r)
		} else {
			LogsWithFileName("", "SQL_", "InsertRow RowsAffected Error->"+err.Error()+"\n InsertRow SQL->"+sqlStr)
		}
		r2, err2 := sdb.LastInsertId()
		if err2 == nil {
			id = strconv.FormatInt(r2, 10)
		} else {
			LogsWithFileName("", "SQL_", "InsertRow LastInsertId Error->"+err2.Error()+"\n InsertRow SQL->"+sqlStr)
		}
	}

	if SqlDebug {
		end := time.Now()
		fmt.Println("Run time ->", end.Sub(start).Seconds(), "\r\n")
	}
	return res, id
}

/**
* 获得更新的sql语句
* @tableName string 要插入的数据表名
* @d map[string]string 要插入的数据map，格式：map["字段名"]=值
* @w  更新条件
* return string sql查询语句
 */
func (db *Mydb) GetUpdateSql(tableName string, d map[string]string, w string) string {
	fieldStr := ""
	for k, v := range d {
		//拼接sql语句
		if fieldStr == "" {
			fieldStr = k + "='" + v + "'"
		} else {
			fieldStr = fieldStr + "," + k + "='" + v + "'"
		}
	}

	sql := "update " + tableName + " set " + fieldStr + " where " + w
	return sql
}

/**
* 更新一条数据到数据表当中
* @tableName string 要更新的数据表名
* @d map[string]string 要更新的数据map，格式：map["字段名"]=值
* @w 更新条件
* return res int 返回影响记录的条数
 */
func (db *Mydb) UpdateRow(tableName string, d map[string]string, w string) int {
	//记录开始时间
	start := time.Now()
	res := -1
	if db.checkConn() == false {
		return res
	}

	lens := len(d)

	if lens < 1 {
		return res
	}

	sqlStr := db.GetUpdateSql(tableName, d, w)

	if SqlDebug {
		fmt.Println("sqlStr->", sqlStr, "\r\n")
	}

	sdb, err := db.Exec(sqlStr)
	if err != nil {
		LogsWithFileName("", "SQL_", "UpdateRow Exec Error->"+err.Error()+"\n UpdateRow SQL->"+sqlStr)
	} else {
		r, err := sdb.RowsAffected()
		if err == nil {
			res = int(r)
		} else {
			LogsWithFileName("", "SQL_", "UpdateRow RowsAffected Error->"+err.Error()+"\n UpdateRow SQL->"+sqlStr)
		}
	}
	if SqlDebug {
		end := time.Now()
		fmt.Println("Run time ->", end.Sub(start).Seconds(), "\r\n")
	}
	return res
}

/**
* 传入一条查询语句，返回一行数据
* @sqlStr string sql语句
* return map类型
 */
func (db *Mydb) GetRow(sqlStr string) map[string]string {
	//记录开始时间
	start := time.Now()
	if SqlDebug {
		fmt.Println("sqlStr->", sqlStr, "\r\n")
	}
	record := make(map[string]string)
	//需要去检查数据库连接是否异常
	if db.checkConn() == false {
		return record
	}

	rows, err := db.Query(sqlStr)
	if err != nil {
		//如果异常，返回空
		LogsWithFileName("", "SQL_", "UpdateRow RowsAffected Error->"+err.Error()+"\n UpdateRow SQL->"+sqlStr)
		return record
	}
	defer rows.Close()

	//读取查询的字段
	columns, err2 := rows.Columns()
	if err2 != nil {
		//如果异常返回空
		LogsWithFileName("", "SQL_", "UpdateRow RowsAffected Error->"+err2.Error()+"\n UpdateRow SQL->"+sqlStr)
		return record
	}
	//创建有效切片
	values := make([]interface{}, len(columns))
	//行扫描，必须复制到这样切片的内存地址中去
	scanArgs := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				_, Isint := col.(int32)
				if Isint {
					record[columns[i]] = strconv.Itoa(int(col.(int32)))
				} else {
					record[columns[i]] = string(col.([]byte))
				}
			}
		}
	}
	if SqlDebug {
		end := time.Now()
		fmt.Println("Run time ->", end.Sub(start).Seconds(), "\r\n")
	}
	return record
}

/**
* 删除数据
* @tablename string 表名
* @where string 删除的where条件,不能为空
* @return 影响行数和错误
 */
func (db *Mydb) DeleteRowsByWhere(tablename, where string) int {
	res := -1
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s", tablename, where)
	if db.checkConn() == false {
		return res
	}
	sdb, err := db.Exec(sql)
	if err == nil {
		num, err2 := sdb.RowsAffected()
		if err2 == nil {
			res = int(num)
		} else {
			return res
		}
	}
	return res
}

/*
* 传入一条查询语句，返回数据列表多行的数组切片
* @sqlStr string sql语句
* return slice类型
 */

func (db *Mydb) GetRows(sqlStr string) [](map[string]string) {
	//记录开始时间
	start := time.Now()
	if SqlDebug {
		fmt.Println(sqlStr, "\r\n")
	}
	//定义一个map类型切片，用来返回数据列表
	//需要去检查数据库连接是否异常
	records := make([](map[string]string), 0)
	if db.checkConn() == false {
		return records
	}
	// 执行查询语句
	rows, err := db.Query(sqlStr)
	if err != nil {
		//如果异常，返回空
		LogsWithFileName("", "SQL_", "UpdateRow RowsAffected Error->"+err.Error()+"\n UpdateRow SQL->"+sqlStr)
		return records
	}
	defer rows.Close()

	// 返回列名称
	columns, err2 := rows.Columns()
	if err2 != nil {
		LogsWithFileName("", "SQL_", "UpdateRow RowsAffected Error->"+err2.Error()+"\n UpdateRow SQL->"+sqlStr)
		return records
	}

	//创建一个切片内存空间
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// 将从数据库接口中读取到的数据copy到切片
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//数据列游标滚动
	for rows.Next() {
		//得到一行数据
		err3 := rows.Scan(scanArgs...)
		if err3 != nil {
			LogsWithFileName("", "SQL_", "UpdateRow RowsAffected Error->"+err3.Error()+"\n UpdateRow SQL->"+sqlStr)
			return records
		}

		//创建map型变量
		record := make(map[string]string)
		//将数据从内存里面读取出来，转载到map中(取数据列字段)
		for i, col := range values {
			//将byte类型数据装入map
			if col != nil {
				record[columns[i]] = string(col)
			} else {
				record[columns[i]] = ""
			}
		}
		//将map加入到数据切片中去(取数据行)
		records = append(records, record)
	}
	if SqlDebug {
		end := time.Now()
		fmt.Println("Run time ->", end.Sub(start).Seconds(), "\r\n")
	}
	return records
}

/**
* 事务处理(独占行锁定)
* 因为mysql的特殊性，这里锁定的行必须走索引，不然将形成表锁
* @sqlArr	[]string	。。传递进来的sql语句

mysql事务调试,进程1:
BEGIN;
set autocommit=0;
#COMMIT;(如果进程不提交的话，其他update查询将被行锁挂起)
end;
select sleep(100);  //需要hold住事务请求，才方便调试第二个请求操作

mysql事务调试，进程2：
explain select availablebalance from userfund where userid=5286

*@param	sqlArr   传入sql数组
*
* return	error	返回错误的内容
*/
func (db *Mydb) CommitSql(sqlArr []string) error {
	//记录开始时间
	start := time.Now()
	//赋值和判断输入是否合法)
	sql := "begin;"
	sql = sql + "set autocommit=0;"

	//开启事务
	cmt, err := db.Begin()
	if err != nil {
		LogsWithFileName("", "SQL_Transaction_", "创建事务失败 ->"+err.Error())
		return err
	}
	if SqlDebug {
		fmt.Println("开始创建事务成功", sqlArr)
	}

	for _, v := range sqlArr {
		_, err2 := cmt.Exec(v)
		if err2 != nil {
			//如果发现有错误，回滚这个事务
			cmt.Rollback()
			LogsWithFileName("", "SQL_Transaction_", "执行事务中的sql失败 ->"+err2.Error()+"sql语句："+v)
			return err2
		}
		sql = sql + v
	}

	err3 := cmt.Commit()
	if err3 != nil {
		LogsWithFileName("", "SQL_Transaction_", "事务提交失败 ->"+err3.Error())
		return err3
	}
	sql = sql + "COMMIT;END;"

	if SqlDebug {
		end := time.Now()
		fmt.Println("Run time ->", end.Sub(start).Seconds(), "\r\n")
	}

	return err3
}

func (db *Mydb) Do(sql string) error {
	var err error
	//需要去检查数据库连接是否异常
	for i := 0; i <= 5; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}
	_, err = db.Exec(sql)
	if err != nil {
		LogsWithFileName("", "MySQL_", "Query Error->"+err.Error()+"\n Query SQL->"+sql)
	}
	return err
}
