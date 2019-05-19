package main

import (
	"github.com/ethereum/go-ethereum/common/hexutil"

	"flag"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"strings"
	"math/big"
	"github.com/PlatONnetwork/PlatON-Go/core/ppos_storage"
	"database/sql"
)

var (
	lpath = flag.String("in", "", "input source leveldb path")
	s3path = flag.String("out", "", "output source sqlite3 path")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "[-in <path>] [-out <path>]")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, `Give leveldb transform to sqlite3`)
	}
}

func main() {
	flag.Parse()
	fmt.Println("in: ", *lpath, " out: ", *s3path)

	//check and ger table name
	var paths []string
	paths = strings.Split(*lpath, "/")
	if len(paths) == 1 {
		paths = strings.Split(*lpath, "\\")
		if len(paths) == 1 {
			log.Fatal("input leveldb path format error path: ", *lpath)
		}
	}

	db, err := sql.Open("sqlite3", *s3path)
	if err != nil {
		log.Fatal("open database err: ", err)
	}
	defer db.Close()

	table := paths[len(paths)-1]
	sqlStmt := fmt.Sprintf(`create table if not exists %s (key text, value text);`, table)
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("create table %q: %s\n", err, sqlStmt)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	txsql := fmt.Sprintf("insert into %s(key, value) values(?, ?)", table)
	stmt, err := tx.Prepare(txsql)
	if err != nil {
		log.Fatal("tx.Prepare err: ", err)
	}
	defer stmt.Close()

	//open leveldb
	ledb, err := leveldb.OpenFile(*lpath, nil)
	if err!=nil {
		log.Fatal("open err: ", err.Error())
	}
	defer db.Close()
	iter := ledb.NewIterator(nil, nil)
	for iter.Next() {
		_, err = stmt.Exec(hexutil.Encode(iter.Key()), hexutil.Encode(iter.Value()))
		if err != nil {
			log.Fatal(err)
		}
	}
	iter.Release()
	tx.Commit()

	fmt.Println("success!")
}



func test () {
	//ldb, err := ethdb.NewLDBDatabase("E:/platon-data/platon/ppos_storage", 0, 0)
	//ldb, err := ethdb.NewPPosDatabase("E:/platon-data/platon/ppos_storage")
	if err!=nil {
		fmt.Println("NewLDBDatabase faile")
	}
	//pposTemp := ppos_storage.NewPPosTemp(ldb)
	//t.Logf("pposTemp info, pposTemp=%+v", pposTemp)

	pposStorage := ppos_storage.NewPPOS_storage()

	pposStorage.t_storage.Sq = 51200

	nodeId := discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345")

	voteOwner := common.HexToAddress("0x20")



	for i := 0; i < 10; i++ {


		deposit := new(big.Int).SetUint64(uint64(rand.Int63()))

		for i := 0; i < 51200; i++ {

			txHash := common.Hash{}
			txHash.SetBytes(crypto.Keccak256([]byte(strconv.Itoa(time.Now().Nanosecond() + i))))
			blockNumber := new(big.Int).SetUint64(uint64(i))
			ticket := &types.Ticket{
				voteOwner,
				deposit,
				nodeId,
				blockNumber,
				2,
			}

			pposStorage.SetExpireTicket(blockNumber, txHash)
			pposStorage.AppendTicket(nodeId, txHash, ticket)
		}




		blockHash := common.Hash{}
		blockHash.SetBytes(crypto.Keccak256([]byte(strconv.Itoa(time.Now().Nanosecond() + i))))
		startTempTime := time.Now().UnixNano()
		pposTemp.SubmitPposCache2Temp(new(big.Int).SetUint64(uint64(i)), new(big.Int).SetUint64(1), blockHash, pposStorage)
		endTempTime := time.Now().UnixNano()
		t.Log("测试Cache2Temp效率", "startTime", startTempTime, "endTime", endTempTime, "time", endTempTime/1e6-startTempTime/1e6)
		startTime := time.Now().UnixNano()
		pposTemp.Commit2DB(ldb, new(big.Int).SetUint64(uint64(i)), blockHash)
		endTime := time.Now().UnixNano()
		t.Log("测试Commit2DB效率", "startTime", startTime, "endTime", endTime, "time", endTime/1e6-startTime/1e6)
	}
}