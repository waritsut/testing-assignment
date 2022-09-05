package db_driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pkg/errors"
)

var Db Repo

type Repo interface {
	GetDb() *gorm.DB
	Find(dest interface{}, conds ...interface{}) *DbRepo
	Save(value interface{}) *DbRepo
	Raw(sql string, values ...interface{}) *DbRepo
	Scan(dest interface{}) *DbRepo
	Exec(sql string, values ...interface{}) *DbRepo
	Begin(opts ...*sql.TxOptions) *DbRepo
}

type DbRepo struct {
	*gorm.DB
}

func NewMyStore(db *gorm.DB) *DbRepo {
	return &DbRepo{DB: db}
}

var counts int

func (d *DbRepo) GetDb() *gorm.DB {
	return d.DB
}

func (d *DbRepo) Find(dest interface{}, conds ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Find(dest, conds...))
}

func (d *DbRepo) Save(value interface{}) *DbRepo {
	return NewMyStore(d.DB.Save(value))
}

func (d *DbRepo) Raw(sql string, values ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Raw(sql, values...))
}

func (d *DbRepo) Scan(dest interface{}) *DbRepo {
	return NewMyStore(d.DB.Scan(dest))
}

func (d *DbRepo) Exec(sql string, values ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Exec(sql, values...))
}

func (d *DbRepo) Begin(opts ...*sql.TxOptions) *DbRepo {
	return NewMyStore(d.DB.Begin(opts...))
}

func OpenDb() (Repo, error) {
	dsn := os.Getenv("DSN")
	// dsn := "host=localhost user=postgres password=password dbname=cashiers port=5432 sslmode=disable TimeZone=UTC"

	for {
		db, err := dbConnect(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return &DbRepo{DB: db}, nil
		}

		if counts > 10 {
			log.Println(err)
			return nil, err
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}

func dbConnect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] DSN: %s", dsn))
	}

	return db, nil
}
