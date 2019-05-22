package main

import "github.com/tecbot/gorocksdb"

func main() {
	concurrency, _ := strconv.Atoi(os.Args[1])
	tasks, _ := strconv.Atoi(os.Args[2])
	disks, _ := strconv.Atoi(os.Args[3])
	partitions, _ := strconv.Atoi(os.Args[4])

	dbPaths := listDBPaths('b', disks, partitions)
	dbs := openDBs(dbPaths)
	ropt := gorocksdb.NewDefaultReadOptions()
	wopt := gorocksdb.NewDefaultWriteOptions()

	wch := make(chan time.Duration, concurrency)
	wsuccess := make(chan int, concurrency)

	rch := make(chan time.Duration, concurrency)
	rsuccess := make(chan int, concurrency)
	srcs := prepareData(1000)

	keys := populateData(1000, srcs, dbs)


}

