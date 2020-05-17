package lib

import (

	//"database/sql"
	//"fmt"
	"fmt"
	"lib"
	"os"

	//"os"
	//"strings"

	l4g "code.google.com/p/log4go"
)

type Db struct {
	l4gLogger *l4g.Logger
	Mysql     *lib.Mysql
}

func NewDb(logger *l4g.Logger, mysql *lib.Mysql) *Db {
	return &Db{l4gLogger: logger, Mysql: mysql}
}

//写入监控数据
/*func (d *Db) InsertMonitor(sql string) bool {
	return d.Mysql.Exec(sql)
}

//更新监控数据
func (d *Db) UpdateMonitor(sql string) bool {
	return d.Mysql.Exec(sql)
}*/

//获取mysql proxy表数据
func (d *Db) GetProxyData(sql string) []Proxy {
	var id int
	var url string
	var template string
	var paramList string
	var list string
	var checkStr string

	rows, retStatus := d.Mysql.Query(sql)
	defer rows.Close()

	ret := make([]Proxy, 0)
	if !retStatus {
		d.l4gLogger.Error(fmt.Sprintf("find %s data fail.", lib.GetCurrentFuncName()))
		fmt.Println(fmt.Sprintf("find %s data fail.", lib.GetCurrentFuncName()))
		return ret
	}

	var line Proxy
	for rows.Next() {
		line = Proxy{}
		if err := rows.Scan(&id, &url, &template, &paramList, &list, &checkStr); err != nil {
			fmt.Println("get project data fail", err)
			os.Exit(-1)

		}
		line.ID = id
		line.Url = url
		line.Template = template
		line.ParamList = paramList
		line.List = list
		line.CheckStr = checkStr
		ret = append(ret, line)
	}
	return ret
}

//获取订单偏差数据
/*func (d *Db) GetOrderDeviation(where map[string]interface{}) []OrderDeviationInfo {
	var supplier_id int
	var book_date string
	var create_time string

	sql := "SELECT supplier_id,book_date,create_time FROM stat_mini_order_table"
	if len(where) > 0 {
		sql_where := ""
		sql_where_arr := make([]string, 0)
		for k, v := range where {
			if _, ok := v.(int); ok {
				sql_where_arr = append(sql_where_arr, fmt.Sprintf("%s %d", k, v))
			} else {
				sql_where_arr = append(sql_where_arr, fmt.Sprintf("%s '%s'", k, v))
			}
		}
		sql_where = strings.Join(sql_where_arr, " AND ")
		sql = fmt.Sprintf("%s WHERE %s", sql, sql_where)
	}
	sql = fmt.Sprintf("%s order by concat(book_date, create_time)", sql)

	rows, retStatus := d.Mysql.Query(sql)
	defer rows.Close()

	ret := make([]OrderDeviationInfo, 0)
	if !retStatus {
		d.l4gLogger.Error(fmt.Sprintf("find %s data fail.", lib.GetCurrentFuncName()))
		fmt.Println(fmt.Sprintf("find %s data fail.", lib.GetCurrentFuncName()))
		return ret
	}

	var line OrderDeviationInfo
	for rows.Next() {
		line = OrderDeviationInfo{}
		if err := rows.Scan(&supplier_id, &book_date, &create_time); err != nil {
			fmt.Println("get project data fail", err)
			os.Exit(-1)

		}
		line.SupplierID = supplier_id
		line.BookDate = book_date
		line.CreateTime = create_time
		ret = append(ret, line)
	}
	return ret
}*/
