package genorm

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	sh "github.com/codeskyblue/go-sh"
	"github.com/domego/gokits"
	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
	"github.com/domego/gopt/gens/funcs"
	_ "github.com/go-sql-driver/mysql"

	yaml "gopkg.in/yaml.v1"
)

var typeMap = [][]string{
	{"int", "int32", "*int32"},
	{"tinyint", "byte", "*byte"},
	{"bigint", "int64", "*int64"},
	{"varchar", "string", "*string"},
	{"char", "string", "*string"},
	{"text", "string", "*string"},
	{"tinytext", "string", "*string"},
	{"datetime", "time.Time", "*time.Time"},
	{"date", "time.Time", "*time.Time"},
	{"timestamp", "time.Time", "*time.Time"},
	{"decimal", "float64", "*float64"},
}

// ColumnInfo table column info
type ColumnInfo struct {
	Field      string
	Type       string
	Null       string
	Key        string
	Default    *string
	Extra      string
	GoType     string
	ThriftType string
	Required   string
}

// Gen gen
func Gen(dbFile string) {
	databaseDir := "model/" + genutils.Values["AppName"].(string)
	dirs := []string{
		"model/all",
		databaseDir,
	}
	model := map[string]interface{}{
		"Name":     genutils.Values["AppName"],
		"RootPath": genutils.Values["RootPath"],
	}
	funcs.FuncMap["setDefault"] = SetDefault
	funcs.FuncMap["getTableFieldNames"] = GetTableFieldNames

	// mkdirs
	genutils.InitDirs(dirs)

	dbs := parseDBFile(dbFile)
	for _, v := range dbs {
		db := v.(map[interface{}]interface{})
		genSource(db)
		loadDBMetaInfo(databaseDir, db)
		// log.Debugf("%v", db)
		genutils.GenFileWithTargetPath("model/database/db.go.tmpl", databaseDir+"/gen_db.go", db)
		sh.Command("gofmt", "-w", ".", sh.Dir(databaseDir)).Run()
	}
	model["DB"] = dbs
	// gen model/all/all.go
	//	log.Debugf("%v", model)
	genutils.GenFile("model/all/all.go.tmpl", model)

	sh.Command("gofmt", "-w", "model/all/all.go").Run()

}

// parseDBFile 解析出db配置文件信息
func parseDBFile(dbFile string) []interface{} {
	var bs []byte
	var err error
	var dbs interface{}
	if bs, err = ioutil.ReadFile(dbFile); err != nil {
		log.Fatalf("%s", err)
	}
	if err = yaml.Unmarshal(bs, &dbs); err != nil {
		log.Fatalf("%s", err)
	}
	return dbs.([]interface{})
}

// genSource 生成db连接字符串
func genSource(db map[interface{}]interface{}) {
	pwd := db["Password"].(string)
	if pwd != "" {
		pwd = ":" + pwd
	}
	port := 3306
	portObj, ok := db["Port"]
	if ok {
		port = portObj.(int)
	}
	source := fmt.Sprintf("%s%s@tcp(%s:%d)/%s?charset=utf8", db["Username"], pwd, db["Host"], port, db["Database"])
	db["Source"] = source
}

// loadDBMetaInfo 查询db元信息
func loadDBMetaInfo(databaseDir string, dbInfo map[interface{}]interface{}) {
	var (
		db         *sql.DB
		rows       *sql.Rows
		err        error
		tableMeta  interface{}
		tableMetas []interface{}
	)
	source := dbInfo["Source"].(string)
	if db, err = sql.Open("mysql", source); err != nil {
		log.Fatalf("%s", err)
	}
	defer db.Close()
	if rows, err = db.Query("SHOW TABLES"); err != nil {
		log.Fatalf("%s", err)
	}
	dbName := dbInfo["Name"].(string)
	tables := dbInfo["Table"].(string)
	genTablesArray := strings.Split(tables, ",")
	for rows.Next() {
		var rowName string
		if err = rows.Scan(&rowName); err != nil {
			log.Fatalf("%s", err)
		}
		// 只加载配置的表
		if tables == "*" || utils.IsInArray(genTablesArray, "rowName") {
			tableMeta = loadTableMetaInfo(db, rowName, dbName)
			model := tableMeta.(map[interface{}]interface{})
			genutils.GenFileWithTargetPath("model/database/table.go.tmpl", databaseDir+"/gen_"+model["TableName"].(string)+".go", tableMeta)
			tableMetas = append(tableMetas, tableMeta)
		}
	}
	dbInfo["TableMetas"] = tableMetas
}

// laodTableMetaInfo 查询表结构元信息
func loadTableMetaInfo(db *sql.DB, tableName, dbName string) interface{} {
	var (
		rows            *sql.Rows
		err             error
		primaryKey      = "id"
		primaryKeyType  = "int"
		primaryKeyExtra = ""
		autoIncrement   = false
		columnInfoList  []*ColumnInfo
	)
	if rows, err = db.Query("SHOW COLUMNS FROM `" + tableName + "`"); err != nil {
		log.Fatalf("%s", err)
	}

	for rows.Next() {
		c := new(ColumnInfo)
		if err = rows.Scan(&c.Field, &c.Type, &c.Null, &c.Key, &c.Default, &c.Extra); err != nil {
			log.Fatalf("%s", err)
		}
		c.GoType = toGoType(c.Type, c.Null)
		c.Required = "required"
		if c.Null == "YES" {
			c.Required = ""
		}
		// log.Debugf("%v", c)
		columnInfoList = append(columnInfoList, c)
		if c.Key == "PRI" {
			primaryKey = c.Field
			primaryKeyType = c.GoType
			primaryKeyExtra = c.Extra
			if c.Extra == "auto_increment" {
				autoIncrement = true
			}
		}
	}
	return map[interface{}]interface{}{
		"DBName":          dbName,
		"TableName":       tableName,
		"PrimaryKey":      primaryKey,
		"PrimaryKeyType":  primaryKeyType,
		"PrimaryKeyExtra": primaryKeyExtra,
		"AutoIncrement":   autoIncrement,
		"Columns":         columnInfoList,
	}
}

func toGoType(s, null string) string {
	for _, v := range typeMap {
		if strings.HasPrefix(s, v[0]) {
			if null == "YES" {
				return v[2]
			}
			return v[1]
		}
	}
	log.Fatalf("unsupport type %s", s)
	return ""
}

func SetDefault(ci *ColumnInfo) string {
	if ci.Default == nil {
		return ""
	}
	if *ci.Default == "CURRENT_TIMESTAMP" {
		return ""
	}
	if strings.Contains(ci.GoType, "string") && *ci.Default == "" {
		return ""
	}
	if (strings.Contains(ci.GoType, "int") || strings.Contains(ci.GoType, "byte") || strings.Contains(ci.GoType, "float")) && !strings.Contains(ci.GoType, "*") && *ci.Default == "0" {
		return ""
	}
	r := funcs.UpperCamel(ci.Field) + ": "
	if ci.Null == "YES" {
		if strings.Contains(ci.GoType, "Time") {
			r += "utils.ToTime(\"2006-01-02 15:04:05\", \""
		} else {
			r += "utils." + funcs.CapitalizeFirst(ci.GoType[1:]) + "("
		}

	}
	if strings.Contains(ci.GoType, "string") {
		r += "\""
	}
	r += *ci.Default
	if strings.Contains(ci.GoType, "string") {
		r += "\""
	}
	if ci.Null == "YES" {
		if strings.Contains(ci.GoType, "Time") {
			r += "\""
		}
		r += ")"
	}
	r += ","
	return r
}

func GetTableFieldNames(args []*ColumnInfo) string {
	names := []string{}
	for _, a := range args {
		names = append(names, fmt.Sprintf("`%s`", a.Field))
	}
	return strings.Join(names, ", ")
}
