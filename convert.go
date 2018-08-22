package main;

import (
	"fmt"
	
	"github.com/go-xorm/xorm"	
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"	
)


// Struct
type Block struct {
	Id int64
	Height int64  `xorm:"unique"`
	Time string   `xorm:"unique"`
	Txs_n int64
	Inner_txs_n int64
	Txs string
}

type Tx struct {
	Id int64
	Tx_id string   `xorm: "unique"`
	Height int64   `xorm: "unique"`
	Content string 
}


// preMethods
func sync(engine *xorm.Engine) error {
	return engine.Sync2(&Block{}, &Tx{});
}


/// sqliteEngine
func sqliteEngine() (*xorm.Engine, error) {
	f := "the.db"
	// os.Remove("the.db")

	return xorm.NewEngine("sqlite3", f);
}


/// postgresEngine
func postgresEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("postgres", "dbname=xorm_test sslmode=disable");
}

/// engineFunc
type engineFunc func() (*xorm.Engine, error)


// assert-tool
func assert(err error) {
	if err != nil {
		fmt.Printf("{}",err);
	}
}

// Generate database
func generate() {
	engines := []engineFunc{ sqliteEngine, postgresEngine };
	for _, enginefunc := range engines {
		Orm, err := enginefunc();
		assert(err);
		
		fmt.Println("--------", Orm.DriverName(), "----------")
		
		Orm.ShowSQL(true)
		err = sync(Orm)
	}
}


// Read from sqlite3
func read() {
	
	var blocks []Block;
	var txs []Tx;

	engine, err := sqliteEngine();
	assert(err);
	
	err = engine.Find(&blocks);
	assert(err);

	err = engine.Find(&txs);
	assert(err);

	fmt.Println("\n---Blocks---: {}\n", blocks);
	fmt.Println("\n---Txs---: {}\n", txs);
}


func main() {
	generate();
	read();
}
