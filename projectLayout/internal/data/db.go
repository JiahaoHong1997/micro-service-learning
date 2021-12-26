package data

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego-component/eredis"
)

func NewDB() *egorm.Component {
	return egorm.Load("mysql.example").Build()
}

func NewCache() *eredis.Component {
	return eredis.Load("redis.example").Build()
}
