package ms2p;

import (
	"fmt"
	
	"github.com/go-xorm/xorm"	
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"	
)


// preMethods
func sync(engine *xorm.Engine, args ...interface{}) error {

	var arr []interface{}
	
	for _, arg := range args {
		arr = append(arr, arg)
	}
		
	return engine.Sync2(args);
}


/// sqliteEngine
func sqliteEngine() (*xorm.Engine, error) {
	f := "the.fox"
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
		panic(err);
	}
}

// Generate database
func Generate() {
	engines := []engineFunc{ sqliteEngine, postgresEngine };
	for _, enginefunc := range engines {
		Orm, err := enginefunc();
		assert(err);
		
		fmt.Println("--------", Orm.DriverName(), "----------")
		
		Orm.ShowSQL(true)
		err = sync(Orm)
		assert(err);
	}
}


// Read from sqlite3
func Read(args ...interface {}) ([]interface {}){

	engine, err := sqliteEngine();
	assert(err);

	for _, arg := range args {
		err = engine.Find(&arg);
		assert(err);		
	}

	return args
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

