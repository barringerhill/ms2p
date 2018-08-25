package ms2p;

import (
	"fmt"
	
	"github.com/go-xorm/xorm"	
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"	
)

// struct

type Block struct {
        Difficulty int64
        Gas_limit int64
        Gas_used int64
        Hash string   `xorm:"unique"`
        Number int64  `xorm:"unique"`
        Size int64
        Timestamp int64
        Total_difficulty int64
        Txs_n int64
        Finished int64
}

type Tx struct {
        Id int64
        Block_hash string
        Gas int64
        Gas_price int64
        Hash string
        Input string
        Value float64
        Finished int64
}

// preMethods
func sync(engine *xorm.Engine ) error {
	return engine.Sync2(&Block{}, &Tx{});
}


/// sqliteEngine
func sqliteEngine() (*xorm.Engine, error) {
	f := "the.fox"
	// os.Remove("the.fox")

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
		panic(err);
	}
}

// Generate database
func Generate(args ...interface{}) {
	engines := []engineFunc{ sqliteEngine, postgresEngine };
	for _, enginefunc := range engines {
		Orm, err := enginefunc();
		assert(err);
		
		fmt.Println("--------", Orm.DriverName(), "----------")
		
		// Orm.ShowSQL(true)
		err = sync(Orm)
		assert(err);
	}
}


// Read from sqlite3
func Read() ([]interface{}){

	engine, err := sqliteEngine();
	assert(err);

	var retarr []interface{}
	
	target := make([]Block, 0);
	err = engine.Find(&target);
	assert(err);

	retarr = append(retarr, target);
	
	return retarr;
}


// Write into postgres
func Write(args ...[]interface{}) {
	
 	Orm, err := postgresEngine()
	fmt.Println("-------", Orm.DriverName(), "-------");
	assert(err);

	err = sync(Orm)
	assert(err);

	for _, arg := range args {
		for _, data := range arg {
			_, err = Orm.Insert(data);
			assert(err);
		}		
	}
	
	print("\n\nSucceed!\n\n");
}

