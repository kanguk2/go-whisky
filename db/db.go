package db

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	"log"
	"os"

	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq" //import postgres
)

var db *sql.DB
var db2 *gorp.DbMap

// Init ...
func Init() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB ...
func ConnectDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	db2 = dbmap
	return db, nil
}

// GetDB ...
func GetDB() *sql.DB {
	return db
}

// GetDB ...
func GetDB2() *gorp.DbMap {
	return db2
}

// RedisClient ...
var RedisClient *_redis.Client

// InitRedis ...
func InitRedis(selectDB ...int) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

}

// GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}
