package config

import (
	"gameApp/repository"
	"gameApp/service/authservice"
)

type HttpServer struct {
	Port int
}
type Config struct {
	Mysql      repository.Config
	HttpServer HttpServer
	Auth       authservice.Config
}
