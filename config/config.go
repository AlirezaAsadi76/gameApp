package config

import (
	"gameApp/repository/mysql"
	"gameApp/service/authservice"
)

type HttpServer struct {
	Port int
}
type Config struct {
	Mysql      mysql.Config
	HttpServer HttpServer
	Auth       authservice.Config
}
