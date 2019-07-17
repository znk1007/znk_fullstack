package cockroachz

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Manager 管理对象实例
var Client *gorm.DB

//"postgresql://maxroach@localhost:26257/bank?
// ssl=true&
// sslmode=require&
// sslrootcert=certs/ca.crt&
// sslkey=certs/client.maxroach.key&
// sslcert=certs/client.maxroach.crt")

func init() {
	const addr = "postgresql://maxroach:znk1007@localhost:26257/znk? " +
		"ssl=true&" +
		"sslmode=require&" +
		"sslrootcert=certs/ca.crt&" +
		"sslkey=certs/client.maxroach.key&" +
		"sslcert=certs/client/maxroach.crt"

	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatalf("connect to Cockroachdb failed: %v", err.Error())
		return
	}
	Client = db
}

// Close 关闭数据库
func Close() {
	if Client != nil {
		Client.Close()
		Client = nil
	}
}
