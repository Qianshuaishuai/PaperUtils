package models

import (
	"dreamEbagPapers/helper"
	"os"

	"github.com/astaxie/beego/config"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mconfig struct {
	dbHost     string
	dbName     string
	dbUsername string
	dbPassword string
	dbMaxIdle  int
	dbMaxConn  int

	SnowFlakDomain           string
	SnowFlakAuthUser         string
	SnowFlakAuthUserSecurity string

	ConfigMyResponse map[int]string

	manageDbName     string
	manageDbHost     string
	manageDbUserName string
	manageDbPassword string
}

var (
	Config       Mconfig
	dbOrmDefault *gorm.DB
	dbOrmManager *gorm.DB
)

const (
	//公共响应码
	RESP_OK         = 10000
	RESP_ERR        = 10001
	RESP_PARAM_ERR  = 10002
	RESP_TOKEN_ERR  = 10003
	RESP_NO_ACCESS  = 10004
	RESP_APP_NOT_ON = 10005

	//应用响应码
	RESP_RESOURCE_NOT_FOUND = 13300
	RESP_SEARCH_ERR         = 13301
)

var flakCurl MSnowflakCurl

//获取对应的db对象
func GetDb() *gorm.DB {
	return dbOrmDefault
}

//获取管理后台数据库对象
func GetManageDb() *gorm.DB {
	return dbOrmManager
}

func init() {
	DREAMENV := os.Getenv("DREAMENV")
	if len(DREAMENV) <= 0 {
		DREAMENV = "DEV"
	}

	appConf, _ := config.NewConfig("ini", "conf/app.conf")

	Config = Mconfig{}
	Config.dbHost = appConf.String(DREAMENV + "::dbHost")
	Config.dbName = appConf.String(DREAMENV + "::dbName")
	Config.dbUsername = appConf.String(DREAMENV + "::dbUsername")
	Config.dbPassword = appConf.String(DREAMENV + "::dbPassword")
	Config.dbMaxIdle, _ = appConf.Int(DREAMENV + "::dbMaxIdle")
	Config.dbMaxConn, _ = appConf.Int(DREAMENV + "::dbMaxConn")
	Config.manageDbName = appConf.String(DREAMENV + "::manageDbName")
	Config.manageDbHost = appConf.String(DREAMENV + "::manageDbHost")
	Config.manageDbUserName = appConf.String(DREAMENV + "::manageDbUserName")
	Config.manageDbPassword = appConf.String(DREAMENV + "::manageDbPassword")
	Config.SnowFlakDomain = appConf.String(DREAMENV + "::snowFlakDomain")
	Config.SnowFlakAuthUser = appConf.String(DREAMENV + "::snowFlakAuthUser")
	Config.SnowFlakAuthUserSecurity = appConf.String(DREAMENV + "::snowFlakAuthUserSecurity")

	getResponseConfig()
	//db
	db, _ := gorm.Open("mysql", Config.dbUsername+":"+Config.dbPassword+"@tcp("+Config.dbHost+")/"+Config.dbName+"?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	managerDb, _ := gorm.Open("mysql", Config.manageDbUserName+":"+Config.manageDbPassword+"@tcp("+Config.manageDbHost+")/"+Config.manageDbName+"?charset=utf8&parseTime=True&loc=Asia%2FShanghai")

	db.DB().SetMaxIdleConns(Config.dbMaxIdle)
	db.DB().SetMaxOpenConns(Config.dbMaxConn)

	managerDb.DB().SetMaxIdleConns(Config.dbMaxIdle)
	managerDb.DB().SetMaxOpenConns(Config.dbMaxConn)

	InitGorm(db)
	dbOrmDefault = db
	dbOrmManager = managerDb

	//建立表格
	// CleanTable(db)
	// CreateTable(db)
	// AddData(db)
	// FixAnswers()
	// getPaperQuestion(db)
	// getPaperQuestion(db)
	// writeSql()
	// GetZsPrimarySchoolPaper()
	// writeKnowledge()
	writeOrderSql()
}

//获取config
func getResponseConfig() {
	Config.ConfigMyResponse = make(map[int]string)
	Config.ConfigMyResponse[RESP_OK] = "成功"
	Config.ConfigMyResponse[RESP_ERR] = "失败,未知错误"
	Config.ConfigMyResponse[RESP_PARAM_ERR] = "参数错误"
	Config.ConfigMyResponse[RESP_TOKEN_ERR] = "token错误"
	Config.ConfigMyResponse[RESP_NO_ACCESS] = "没有访问权限"
}

//分库,分表算法
func splitDbAndTable(flagId string) (dbSuffix string, tableSuffix string) {
	if len(flagId) > 0 {
		md5Str := helper.Md516(flagId)
		md5CharList := []byte(md5Str)
		index0 := int(md5CharList[1])
		index := (int(md5CharList[0])) % 100
		dbSuffix = string((index0 % 10) + int('0'))
		tableSuffix = string((index/10)+int('0')) + string((index%10)+int('0'))
	}
	return
}

//获取对应的table名称
func GetTable(baseTableName string, flagId string) (tableName string) {
	if len(flagId) == 0 { //返回默认的库
		tableName = baseTableName
	} else { //获取对应的库
		_, tableSuffix := splitDbAndTable(flagId)
		tableName = baseTableName + "_" + tableSuffix
	}
	return
}

//建表（有表，先删除表格）
// func CreateTable(db *gorm.DB) {
// 	db.CreateTable(&Grade{})
// 	db.CreateTable(&Literary{})
// 	db.CreateTable(&WordNumber{})
// 	db.CreateTable(&Material{})
// 	db.CreateTable(&Accessorie{})
// 	db.CreateTable(&Question{})
// 	db.CreateTable(&SolutionAccessorie{})
// 	db.CreateTable(&Option{})
// 	db.CreateTable(&KeyPoint{})
// 	db.CreateTable(&Answer{})
// }

// func CleanTable(db *gorm.DB) {
// 	db.DropTableIfExists(&Grade{})
// 	db.DropTableIfExists(&Literary{})
// 	db.DropTableIfExists(&WordNumber{})
// 	db.DropTableIfExists(&Material{})
// 	db.DropTableIfExists(&Accessorie{})
// 	db.DropTableIfExists(&Question{})
// 	db.DropTableIfExists(&SolutionAccessorie{})
// 	db.DropTableIfExists(&Option{})
// 	db.DropTableIfExists(&KeyPoint{})
// 	db.DropTableIfExists(&Answer{})
// }

// func AddData(db *gorm.DB) {
// 	AddGradeList(db)
// 	AddLiteraryList(db)
// 	AddWordNumberList(db)
// 	// GenerateMaterialsForzuoye(db, "1")
// 	// readAllFile(db)
// 	GenerateMaterialsForManager(dbOrmManager)
// }
