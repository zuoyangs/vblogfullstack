package impl

import (
	"database/sql"

	"github.com/zuoyangs/vblog_backend/conf"
)

// 构造函数，返回一个 Impl 实例
func NewImpl() *Impl {
	return &Impl{}
}

// 负责实现 Blog Service
type Impl struct {
	db *sql.DB
}

//当这个对象初始化时，会获取该对象需要的依赖
//需要db这个依赖，从配置文件中获取

func (i *Impl) Init() error {
	i.db = conf.C().MySQL.GetDB()
	return nil
}
