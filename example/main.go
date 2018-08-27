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
