package conf

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

func NewDefaultConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		MySQL: newDefaultMySQL(),
	}
}

type Config struct {
	App   *app   `toml:"app"`
	MySQL *mysql `toml:"mysql"`
}

func newDefaultApp() *app {
	return &app{
		Name: "vblog",
		HTTP: newDefaultHttp(),
	}
}

type app struct {
	//应用名称
	Name string `toml:"name" env:"APP_NAME"`
	HTTP *http  `toml:"http"`
}

func newDefaultHttp() *http {
	return &http{
		Host: "localhost",
		Port: "8080",
	}
}

type http struct {
	Host string `toml: "host" env:"HTTP_HOST"`
	Port string `toml: "port" env:"HTTP_PORT"`
}

func newDefaultMySQL() *mysql {
	return &mysql{
		HOST:     "",
		PORT:     "3306",
		DATABASE: "vblog",
		USERNAME: "",
		PASSWORD: "",
	}
}

type mysql struct {

	// 表示 MySQL 数据库的 IP 地址或域名
	HOST string `toml:"host"    env:"MYSQL_HOST"`
	// 表示 MySQL 数据库的端口号
	PORT string `toml:"port" 	 env:"MYSQL_PORT"`
	// 表示要连接的 MySQL 数据库名称
	DATABASE string `toml:"database" env:"MYSQL_DATABASE"`
	// 表示连接 MySQL 数据库所使用的用户名
	USERNAME string `toml:"username" env:"MYSQL_USERNAME"`
	// 表示连接 MySQL 数据库所使用的密码
	PASSWORD string `toml:"password" env:"MYSQL_PASSWORD"`

	// 表示连接池中最大的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 表示连接池中最大的空闲连接数
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 表示连接池中连接的最大生命周期
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// 表示连接池中连接的最大空闲时间
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`

	// 用于保证连接池的线程安全性
	lock sync.Mutex
	// 连接池中实际的连接对象，是一个 *sql.DB 类型的指针
	dbconn *sql.DB
}

// 数据连接 需要单列模式
func (m *mysql) GetDB() *sql.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.dbconn == nil {
		conn, err := m.getDB()
		if err != nil {
			panic(err)
		}
		m.dbconn = conn
	}
	return m.dbconn
}

// 通过 MySQL 配置获取一个连接池
func (m *mysql) getDB() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.USERNAME,
		m.PASSWORD,
		m.HOST,
		m.PORT,
		m.DATABASE,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("connect to mysql<%s> error: %v", dsn, err)
	}

	//设置连接池参数
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)

	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleConn != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	//通过 ping 来测试 mysql 是否可达
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error,%s", dsn, err.Error())
	}
	return db, nil

}
