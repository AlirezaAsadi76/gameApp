package main

import (
	"gameApp/config"
	"gameApp/delivery/httpserver"
	"gameApp/repository/mysql"
	"gameApp/service/authservice"
	"gameApp/service/userservice"
	"gameApp/validator/uservalidator"
	"time"
)

const (
	jwtSecret            = "secret"
	AccessTokenSubject   = "at"
	RefreshTokenSubject  = "rt"
	AccessTokenDuration  = time.Hour * 24
	RefreshTokenDuration = time.Hour * 24 * 7
)

func main() {

	cfg := config.Config{
		Mysql: mysql.Config{
			Host:     "localhost",
			Port:     3308,
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Database: "gameapp_db",
		},
		HttpServer: config.HttpServer{
			Port: 7660,
		},
		Auth: authservice.Config{
			SignKey:              jwtSecret,
			AccessSubject:        AccessTokenSubject,
			RefreshSubject:       RefreshTokenSubject,
			AccessTokenDuration:  AccessTokenDuration,
			RefreshTokenDuration: RefreshTokenDuration,
		},
	}
	//migrate := migrator.New(mysql.Config{
	//	Host:     "localhost",
	//	Port:     3308,
	//	Username: "gameapp",
	//	Password: "gameappt0lk2o20",
	//	Database: "gameapp_db",
	//})
	//migrate.Up()
	authServ, userServ, userUv := setupService(cfg)

	server := httpserver.New(cfg, authServ, userServ, userUv)

	server.Start()

}

func setupService(cfg config.Config) (authservice.Service, *userservice.Service, uservalidator.Validator) {

	authServ := authservice.New(cfg.Auth)

	repo := mysql.NewDB(cfg.Mysql)
	userServ := userservice.New(repo, authServ)
	uV := uservalidator.New(repo)
	return authServ, userServ, uV
}
