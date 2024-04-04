package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"messageBot/conf"
	"time"
)

type MysqlDBClient struct {
	db *gorm.DB
}

var MysqlClient MysqlDBClient

type GormLogger struct {
	ctx *gin.Context
}

func (client *MysqlDBClient) SetCtx(ctx *gin.Context) *gorm.DB {
	gormLog := GormLogger{ctx: ctx}
	client.db.SetLogger(&gormLog)
	return client.db
}

// Print handles log events from Gorm for the custom logger.
func (log *GormLogger) Print(v ...interface{}) {
	logger := Logger(log.ctx)
	switch v[0] {
	case "sql":
		logger.WithFields(
			logrus.Fields{
				"module":        "gorm",
				"type":          "sql",
				"rows_returned": v[5],
				"src":           v[1],
				"values":        v[4],
				"duration":      v[2],
			},
		).Debug(v[3])
	case "log":
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

func InitMysql() {
	var err error
	gormClient, err := InitMysqlClient(MysqlConf{
		User:            conf.Conf.Mysql.User,
		Password:        conf.Conf.Mysql.PassWord,
		Addr:            conf.Conf.Mysql.Addr,
		DataBase:        conf.Conf.Mysql.DataBase,
		MaxIdleConns:    conf.Conf.Mysql.MaxIdleConns,
		MaxOpenConns:    conf.Conf.Mysql.MaxOpenConns,
		ConnMaxLifeTime: 3600 * time.Second,
		LogMode:         true,
	})
	if err != nil {
		PanicfLogger(nil, "mysql connect error: %v", err)
	}
	MysqlClient = MysqlDBClient{db: gormClient}
}

type MysqlConf struct {
	User            string
	Password        string
	Addr            string
	DataBase        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime time.Duration
	LogMode         bool
}

// InitMysqlClient
func InitMysqlClient(conf MysqlConf) (client *gorm.DB, err error) {
	client, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
		conf.User,
		conf.Password,
		conf.Addr,
		conf.DataBase))

	if err != nil {
		return client, err
	}
	client.SingularTable(true)
	client.DB().SetMaxIdleConns(conf.MaxIdleConns)
	client.DB().SetMaxOpenConns(conf.MaxOpenConns)
	client.DB().SetConnMaxLifetime(conf.ConnMaxLifeTime)
	client.SetLogger(&GormLogger{})
	client.LogMode(conf.LogMode)

	return client, nil
}
