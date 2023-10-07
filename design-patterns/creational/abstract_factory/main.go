package main

import "fmt"

/*
Abstract Factory Design Pattern is a creational design pattern that lets you create a family of related objects.
It is an abstraction over the factory pattern
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

	file struct {
		name    string
		content string
	}
	ntfs struct {
		files map[string]file
	}
	ext4 struct {
		files map[string]file
	}
	FileSystem interface {
		CreateFile(path string)
		FindFile(path string) file
	}

	Factory func(string) interface{}
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

func DatabaseFactory(env string) interface{} {
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

func (n ntfs) CreateFile(path string) {
	f := file{name: path, content: "NTFS file"}
	n.files[path] = f
	fmt.Println("NTFS")
}

func (n ntfs) FindFile(path string) file {
	return n.files[path]
}

func (e ext4) CreateFile(path string) {
	f := file{name: path, content: "EXT4 file"}
	e.files[path] = f
	fmt.Println("EXT4")
}

func (e ext4) FindFile(path string) file {
	return e.files[path]
}

func FileSystemFactory(env string) interface{} {
	switch env {
	case "production":
		return &ntfs{
			files: make(map[string]file),
		}
	case "development":
		return &ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FileSystemFactory
	default:
		return nil
	}
}

func SetUpConstructors(env string) (Database, FileSystem) {
	fs := AbstractFactory("filesystem")
	db := AbstractFactory("database")

	return db(env).(Database), fs(env).(FileSystem)
}
