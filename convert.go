package ms2p;

import (
	"fmt"
	"reflect"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Config struct {
	DBPath  string
	PgPara  string
}

func assert(err error) {
	if err != nil { panic(err); }
}

func Convert(config Config, args ...interface{}) {

	// sqlite
	sqlite, sqlite_err := gorm.Open("sqlite3", config.DBPath);

	assert(sqlite_err);
	defer sqlite.Close();
	
	// progres
	postgres, postgre_err := gorm.Open("postgres", config.PgPara);

	assert(postgre_err);
	defer postgres.Close();
	
	for _, arg := range(args) {
		sqlite.Find(arg);
		
		postgres.AutoMigrate(arg)
		postgres.Create(arg);

		arg_obj := reflect.ValueOf(arg).Elem();
		fmt.Println(arg_obj);
		
		println("Succeed!")
	}
}
