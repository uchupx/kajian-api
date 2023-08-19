package mysql

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBPayload struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

// Create new connection to mysql database
// return *sqlx.DB and error
func NewConnection(p DBPayload) (*sqlx.DB, error) {
	newConfig := mysql.Config{
		User:   p.Username,
		Passwd: p.Password,
		DBName: p.Database,
		Net:    "tcp",
		Addr:   p.Host + ":" + p.Port,
	}

	db, err := sqlx.Connect("mysql", newConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed connectin database host:%s, err: %+v", fmt.Sprintf("%s:%s", p.Host, p.Port), err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed ping database host:%s, err: %+v", fmt.Sprintf("%s:%s", p.Host, p.Port), err)
	}

	return db, nil
}
