package db

import (
	"log"

    "database/sql"
	_ "github.com/microsoft/go-mssqldb" // sql server driver

)

type Knowledgedb interface {
	GetConnection() *sql.DB
}

type knowledgeDbImpl struct { }

var db *sql.DB
var dataSource Knowledgedb

func (k knowledgeDbImpl) GetConnection() *sql.DB {
	return db
}

func GetKnowledgeDb() Knowledgedb {
	if dataSource == nil {
		dataSource = knowledgeDbImpl{}
	}

	return dataSource
}

func init() {
	connString := "server=localhost;port=1433;user id=sa;password=abcdeF1+"

    conn, err := sql.Open("sqlserver", connString) // load sql server driver

    if err != nil {    
		log.Fatalf("Error connecting to the database: %v", err)
    }

    // Verify the connection
    err = conn.Ping()
    if err != nil {
        
		log.Fatalf("Error connecting to the database: %v", err)
    }

	db = conn
}

// func main() {
// 	GetKnowledgeDb().GetConnection()
// }