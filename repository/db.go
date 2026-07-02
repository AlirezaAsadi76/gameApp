package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"gameApp/entity"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	db *sql.DB
}

func NewDB() *MysqlDB {
	db, err := sql.Open("mysql", "gameapp:gameappt0lk2o20@(localhost:3308)/gameapp_db")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &MysqlDB{
		db: db,
	}
}

func (s *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	row := s.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	user := entity.User{}
	var createdAt []uint8
	rErr := row.Scan(&user.ID, &user.PhoneNumber, &user.Name, &createdAt)
	if rErr != nil {

		if errors.Is(rErr, sql.ErrNoRows) {

			return true, nil
		}
		return false, fmt.Errorf("query row scan error: %w", rErr)
	}

	return false, nil

}

func (s *MysqlDB) Register(user entity.User) (entity.User, error) {

	result, eErr := s.db.Exec(`insert into users(name, phone_number) values(?, ?)`, user.Name, user.PhoneNumber)
	if eErr != nil {
		return entity.User{}, fmt.Errorf("cannot insert user: %w", eErr)
	}

	id, rErr := result.LastInsertId()
	if rErr != nil {
		return entity.User{}, fmt.Errorf("cannot get last insert id: %w", rErr)
	}
	return entity.User{ID: uint(id)}, nil
}
