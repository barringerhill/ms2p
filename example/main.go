// Struct

package main

import "github.com/udtrokia/ms2p"

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

func main() {
	ms2p.Generate();
	ms2p.Write(ms2p.Read(Block, Tx))
}
