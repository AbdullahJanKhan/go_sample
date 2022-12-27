package db_name

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/abdullahjankhan/go_sample/config"
	"github.com/abdullahjankhan/go_sample/repository"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gdb *gorm.DB
var storeOnce sync.Once
var store repository.Store

type Store struct {
	db *gorm.DB
}

// SharedStore return global or single instance of postgres connection (bounded in sync once)
func SharedStore() repository.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			panic(err)
		}
		store = NewStore(gdb)
	})
	return store
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	_ = config.GetConfig()

	postgresConf := &Config{
		Host:      "cfg.PostDataSource.Addr",
		Port:      "cfg.PostDataSource.Port",
		DbName:    "cfg.PostDataSource.Database",
		User:      "cfg.PostDataSource.User",
		Password:  "cfg.PostDataSource.Password",
		SSLEnable: false,
	}
	retries := 3 //"cfg.PostDataSource.Retries"
	log.Println(postgresConf)
	var err error
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	for err != nil {
		log.Println(err, fmt.Sprintf("Failed to connect to database (%d)", retries))

		if retries > 1 {
			retries--
			time.Sleep(5 * time.Second)
			gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

			continue
		}
		panic(err)
	}
	if true { //"cfg.PostDataSource.EnableAutoMigrate" {
		var tables = []interface{}{
			// enter your modles you want to migrate to DB
		}

		for _, table := range tables {
			log.Printf("migrating database, table: %v", reflect.TypeOf(table))
			if err = gdb.AutoMigrate(table); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Store) BeginTx() (repository.Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) CommitTx() error {
	return s.db.Commit().Error
}
