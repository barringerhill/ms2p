# ms2p

Migrate data from sqlite3 to postgreSQL with xorm


### ASSERT

JUST FOR ALLBLUE

***

BUILD FAILED

***

### Install

```
go get -u https://github.com/udtrokia/ms2p

```

#### Content

+ Struct

```go
type Config struct {
    DBPath string
    PgPara string
}

func Convert (config Config, args ...interface) {}

```

#### Example

```Golang

package main;
import "github.com/udtrokia/ms2p"

type Tx struct {
	Block_number int
	Hash string
	Input string
}

func (Tx) TableName() string {return "tx"}

func main() {
	config := ms2p.Config{
		DBPath: "the.fox",
		PgPara: "host=127.0.0.1 port=5432 dbname=edata sslmode=disable",
	}
	
	ms2p.Convert(config, &Tx{});
}

```

#### TODO

More intro.


#### LICENSE

GPLv3.0
