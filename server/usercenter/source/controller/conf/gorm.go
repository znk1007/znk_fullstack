package userconf

//GormConf gorm服务器配置
type GormConf struct {
	Host     string
	Port     string
	Username string
	Password string
	Dialect  string
	DB       string
}

var gorms []GormConf

var gormmap map[Env]GormConf

func init() {
	gormmap = map[Env]GormConf{
		Dev: GormConf{
			Host:     "localhost",
			Port:     "3306",
			Username: "znk",
			Password: "man_znk-1007",
			Dialect:  "mysql",
			DB:       "znk",
		},
		Test: GormConf{
			Host:     "47.105.85.107",
			Port:     "3308",
			Username: "znk",
			Password: "man_znk-1007",
			Dialect:  "mysql",
			DB:       "znk",
		},
		Prod: GormConf{
			Host:     "47.105.85.107",
			Port:     "3308",
			Username: "znk",
			Password: "man_znk-1007",
			Dialect:  "mysql",
			DB:       "znk",
		},
	}
}

//GetGormConf 获取当前gorm服务器配置
func getGormConf(env Env) GormConf {
	if g, ok := gormmap[env]; ok {
		return g
	}
	return gormmap[Dev]
}

//SetGormConf 设置gorm服务配置
func setGormConf(env Env, host string, port string, username string, password string, dialect string, db string) {
	gc := GormConf{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Dialect:  dialect,
		DB:       db,
	}
	gormmap[env] = gc
}
