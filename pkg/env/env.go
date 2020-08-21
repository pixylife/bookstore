package env

import (
	controller "bookstore/pkg/controller"
	gorm "github.com/jinzhu/gorm"
	"os"
)

var RDB *gorm.DB
var CMDCtrl *controller.StorageCtrl
var QueryCtrl *controller.StorageCtrl
var ESCtrl *controller.MessageCtrl

const DB0_DIALET = "DB0_DIALET"

var Db0DIALET = "mongo"

const DB0_DB = "DB0_DB"

var Db0Db = "db"

const DB0_HOST = "DB0_HOST"

var Db0Host = "localhost"

const DB0_PORT = "DB0_PORT"

var Db0Port = "27017"

const DB0_USER = "DB0_USER"

var Db0User = "root"

const DB0_PWD = "DB0_PWD"

var Db0Pwd = "root"

const DB0_URL = "DB0_URL"

var Db0Url = "localhost"

const DB1_DIALET = "DB1_DIALET"

var Db1DIALET = "sqlite3"

const DB1_DB = "DB1_DB"

var Db1Db = "db"

const DB1_HOST = "DB1_HOST"

var Db1Host = ""

const DB1_PORT = "DB1_PORT"

var Db1Port = ""

const DB1_USER = "DB1_USER"

var Db1User = ""

const DB1_PWD = "DB1_PWD"

var Db1Pwd = ""

const DB1_URL = "DB1_URL"

var Db1Url = ""

func LoadEnvs() {

	if val := os.Getenv(DB0_DIALET); val != "" {
		Db0DIALET = val
	}
	if val := os.Getenv(DB0_DB); val != "" {
		Db0Db = val
	}
	if val := os.Getenv(DB0_HOST); val != "" {
		Db0Host = val
	}
	if val := os.Getenv(DB0_PORT); val != "" {
		Db0Port = val
	}
	if val := os.Getenv(DB0_USER); val != "" {
		Db0User = val
	}
	if val := os.Getenv(DB0_PWD); val != "" {
		Db0Pwd = val
	}
	if val := os.Getenv(DB0_URL); val != "" {
		Db0Url = val
	}
	if val := os.Getenv(DB1_DIALET); val != "" {
		Db1DIALET = val
	}
	if val := os.Getenv(DB1_DB); val != "" {
		Db1Db = val
	}
	if val := os.Getenv(DB1_HOST); val != "" {
		Db1Host = val
	}
	if val := os.Getenv(DB1_PORT); val != "" {
		Db1Port = val
	}
	if val := os.Getenv(DB1_USER); val != "" {
		Db1User = val
	}
	if val := os.Getenv(DB1_PWD); val != "" {
		Db1Pwd = val
	}
	if val := os.Getenv(DB1_URL); val != "" {
		Db1Url = val
	}
}
