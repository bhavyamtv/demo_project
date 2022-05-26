package config

import (
	"bytes"

	"database/sql"

	"fmt"

	"log"

	"strconv"
	"time"

	//package for mysql

	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

// DB : Master DB handle

var DB *sql.DB

var MGO_SESSION *mgo.Session

//DB_READ : Slave DB Handle for read

var DB_READ *sql.DB

type dbConfig struct {
	hostName string

	username string

	password string

	database string

	port int
}

type mongo_config struct {
	hostName string

	userName string

	password string
}

//InitDB : initialises db handles

func InitDB() (*sql.DB, *mgo.Session, error) {

	var err error

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(2)
	fmt.Println(num)
	if num == 1 {
		slaveDB := dbConfig{
			"localhost",
			"root",
			"mtv",
			"demoproject",
			3306,
		}
		var slaveURL bytes.Buffer
		slaveURL.WriteString(slaveDB.username)
		slaveURL.WriteString(":")
		slaveURL.WriteString(slaveDB.password)
		slaveURL.WriteString("@tcp(")
		slaveURL.WriteString(slaveDB.hostName)
		slaveURL.WriteString(":")
		slaveURL.WriteString(strconv.Itoa(slaveDB.port))
		slaveURL.WriteString(")/")
		slaveURL.WriteString(slaveDB.database)
		slaveURLString := slaveURL.String()
		DB_READ, err = sql.Open("mysql", slaveURLString)
		if err != nil {
			log.Panic(err)
		}
		//DB_READ.SetMaxIdleConns(250)
		DB_READ.SetConnMaxLifetime(time.Second)
		DB_READ.SetMaxOpenConns(0)
		DB_READ.SetMaxIdleConns(0)
		if err = DB_READ.Ping(); err != nil {
			log.Panic(err)
		}
		fmt.Println("DB Connected Successfully to Slave DB 1.")
	} else {
		slaveDB := dbConfig{
			"localhost",
			"root",
			"mtv",
			"demoproject",
			3306,
		}
		var slaveURL bytes.Buffer
		slaveURL.WriteString(slaveDB.username)
		slaveURL.WriteString(":")
		slaveURL.WriteString(slaveDB.password)
		slaveURL.WriteString("@tcp(")
		slaveURL.WriteString(slaveDB.hostName)
		slaveURL.WriteString(":")
		slaveURL.WriteString(strconv.Itoa(slaveDB.port))
		slaveURL.WriteString(")/")
		slaveURL.WriteString(slaveDB.database)
		slaveURLString := slaveURL.String()
		DB_READ, err = sql.Open("mysql", slaveURLString)
		if err != nil {
			log.Panic(err)
		}
		DB_READ.SetConnMaxLifetime(time.Second)
		DB_READ.SetMaxOpenConns(0)
		DB_READ.SetMaxIdleConns(0)
		if err = DB_READ.Ping(); err != nil {
			log.Panic(err)
		}
		fmt.Println("DB Connected Successfully to Slave DB 2." + slaveDB.hostName)
	}

	if err = DB_READ.Ping(); err != nil {

		log.Panic(err)

	}
	// info := &mgo.DialInfo{

	// 	Addrs: []string{"3.109.203.97"},

	// 	Database: "admin",

	// 	Username: "event_builder",

	// 	Password: "jaiHOmultitV",
	// }

	// MGO_SESSION, err = mgo.DialWithInfo(info)

	if err != nil {

		log.Println(err)

		log.Panic(err)

	}

	fmt.Println("Mongo DB. Connected successfuly")

	return DB_READ, MGO_SESSION, nil

}
