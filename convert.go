package main;

import (
	"fmt"
	// "reflect"
	// "strings"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Config struct {
	DBPath  string
	PgPara  string
}

type Tx struct {
	Block_number int
	Hash string
	Input string
}

func (Tx) TableName() string {return "tx"}


func assert(err error) {
	if err != nil { panic(err); }
}

func main() {
	config := Config {
		DBPath: "the.fox",
		PgPara: "host=127.0.0.1 port=5432 dbname=edata sslmode=disable",
	}
	
	// sqlite
	sqlite, sqlite_err := gorm.Open("sqlite3", config.DBPath);
	assert(sqlite_err);
	defer sqlite.Close();
	
	// progres
	postgres, postgre_err := gorm.Open("postgres", config.PgPara);
	assert(postgre_err);
	defer postgres.Close();

	postgres.AutoMigrate(&Tx{})
	
	var txs []Tx;
	sqlite.Find(&txs);

	for id, tx := range(txs) {
		fmt.Printf("\r>>>  Process %d / %d ", id + 1, len(txs))
		postgres.Create(tx);
	}
	println(" <<<")
}
