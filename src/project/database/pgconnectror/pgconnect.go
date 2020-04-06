package pgconnect

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "mimik"
	password = "qwerty143"
	dbname   = "chapter2"
)

type pgconnector struct {
	Connected bool `false`
	psqlInfo  string
	Db        *sql.DB
}

var instance *pgconnector
var once sync.Once

//хочется понять как работает синглтон, для каждого запроса делать свои коннекты к базе не совсем верно, ладно если бы это было через unix сокеты
func GetInstance() *pgconnector {
	once.Do(func() {
		instance = &pgconnector{}
	})
	return instance
}

func (self *pgconnector) OpenConnect() {
	self.psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", self.psqlInfo)

	if err != nil {
		panic(err)
		return
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
		return
	}
	self.Db = db
	self.Connected = true
	fmt.Println("Successfully connected!")

	/*
		по хорошему здесь надо сделать вывод определений табличек в описанные типы, чтобы автоматом создавать все относительно типов,
		хотелось бы понять как это сделать, чтобы можно было при создании нового типа автоматом создавать его в бд если такого нет.

	*/
	result, errs := db.Exec("CREATE TABLE IF NOT EXISTS public.users ( id uuid,name varchar(255),lastname varchar(255), age int, birthdate bigint);")
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(result)

}
