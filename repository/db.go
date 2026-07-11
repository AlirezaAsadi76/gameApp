package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"gameApp/entity"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}
type MysqlDB struct {
	config Config
	db     *sql.DB
}

// "gameapp:gameappt0lk2o20@(localhost:3308)/gameapp_db"
func NewDB(config Config) *MysqlDB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &MysqlDB{
		db:     db,
		config: config,
	}
}

func (s *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	row := s.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	_, rErr := scanUser(row)

	if rErr != nil {

		if errors.Is(rErr, sql.ErrNoRows) {

			return true, nil
		}
		return false, fmt.Errorf("query row scan error: %w", rErr)
	}

	return false, nil

}

func (s *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {

	row := s.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	user, rErr := scanUser(row)

	if rErr != nil {

		if errors.Is(rErr, sql.ErrNoRows) {

			return entity.User{}, false, nil
		}
		return entity.User{}, false, fmt.Errorf("query row scan error: %w", rErr)
	}

	return user, true, nil

}

func (s *MysqlDB) GetUserByID(UserID uint) (entity.User, error) {

	row := s.db.QueryRow(`select * from users where id = ?`, UserID)

	// TODO - use function for scan user
	user, rErr := scanUser(row)
	if rErr != nil {

		if errors.Is(rErr, sql.ErrNoRows) {

			return entity.User{}, fmt.Errorf("user not found: %w", rErr)
		}
		return entity.User{}, fmt.Errorf("query row scan error: %w", rErr)
	}

	return user, nil

}
func (s *MysqlDB) Register(user entity.User) (entity.User, error) {

	result, eErr := s.db.Exec(`insert into users(name, phone_number, password) values(?, ?, ?)`, user.Name, user.PhoneNumber, user.Password)
	if eErr != nil {
		return entity.User{}, fmt.Errorf("cannot insert user: %w", eErr)
	}

	id, rErr := result.LastInsertId()
	if rErr != nil {
		return entity.User{}, fmt.Errorf("cannot get last insert id: %w", rErr)
	}
	return entity.User{ID: uint(id)}, nil
}

func scanUser(row *sql.Row) (entity.User, error) {

	user := entity.User{}
	var createdAt []uint8
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	return user, err
}
