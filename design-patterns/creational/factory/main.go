package main

import (
	"fmt"
	"reflect"
)

/*
Factory method pattern is a creational pattern that uses factory methods to deal with the problem
of creating objects without having to specify the exact class of the object that will be created.
*/

type (
	mongoDB struct {
		database map[string]string
	}
	mysqlDB struct {
		database map[string]string
	}
	Database interface {
		GetData(k string) string
		PutData(k, v string)
	}
)

func (mdb mongoDB) GetData(k string) string {
	v, ok := mdb.database[k]
	if !ok {
		return ""
	}
	fmt.Println("MongoDB")
	return v
}

func (mdb mongoDB) PutData(k, v string) {
	mdb.database[k] = v
}

func (sql mysqlDB) GetData(k string) string {
	v, ok := sql.database[k]
	if !ok {
		return ""
	}
	fmt.Println("MysqlDB")
	return v
}

func (sql mysqlDB) PutData(k, v string) {
	sql.database[k] = v
}

func databaseFactory(env string) Database {
	switch env {
	case "production":
		return &mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return &mysqlDB{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

func main() {
	mongo := databaseFactory("production")
	mysql := databaseFactory("development")

	mongo.PutData("test", "this is mongodb")
	fmt.Println(mongo.GetData("test"))
	mysql.PutData("test", "this is mysqldb")
	fmt.Println(mysql.GetData("test"))

	fmt.Printf("%T\n", mongo)
	fmt.Printf("%T\n", mysql)
	fmt.Println(reflect.TypeOf(&mongo).Elem())
	fmt.Println(reflect.TypeOf(&mysql).Elem())
}
