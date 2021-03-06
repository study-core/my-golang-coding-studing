package main
//import (
//	"sync"
//	"crypto/sha256"
//	"encoding/hex"
//	"strconv"
//	"time"
//	"bufio"
//	"fmt"
//	"log"
//	"net"
//	"io"
//	"encoding/json"
//	"math/rand"
//	"os"
//	"github.com/joho/godotenv"
//	"github.com/davecgh/go-spew/spew"
//)
//
//type Block struct {
//	Index 			int  	// 区块索引
//	Timestemp		string 	// 出块时间
//	BPM 			int 	// 数据
//	Hash 		 	string 	// 当前块 Hash
//	PreHash 		string 	// 前一个块的Hash
//	Validator 		string  // 校验器
//}
//
//
//var (
//	Blockchain 	[]Block					// 链数据
//	tempBlocks 	[]Block 				// 临时区块
//	candidateBlocks = make(chan Block)  // 接收新块的通道
//	announcements 	= make(chan string) // TCP Server 向所有节点广播最新区块
//	validators		= make(map[string]int) // 使用map存储验证的token
//)
//
//var mutex = &sync.Mutex{} 	// 用于并发读写的锁
//
///**
//计算Hash的函数
// */
//func calculateHash(s string) string {
//	h := sha256.New()
//	h.Write([]byte(s))
//	hashed := h.Sum(nil)
//	return hex.EncodeToString(hashed)
//}
//
///**
//计算区块Hash
// */
//func calculateBlockHash(block Block) string {
//	record := strconv.Itoa(block.Index) +
//		block.Timestemp +
//		strconv.Itoa(block.BPM) +
//		block.PreHash
//	return calculateHash(record)
//}
//
//func generateBlock(oldBlock Block, BPM int, address string) (Block, error){
//	var newBlock Block
//	t := time.Now()
//	newBlock.Index = oldBlock.Index + 1
//	newBlock.Timestemp = t.String()
//	newBlock.BPM = BPM
//	newBlock.PreHash = oldBlock.Hash
//	newBlock.Validator = address
//	return  newBlock, nil
//}
//
//func isBlockValid(newBlock Block, oldBlock Block) bool {
//	if oldBlock.Index + 1 != newBlock.Index {return false}
//	if oldBlock.Hash != newBlock.PreHash {return false}
//	if newBlock.Hash != calculateBlockHash(newBlock) {return false}
//	return true
//}
//
//
////处理逻辑
//func handleConn(conn net.Conn) {
//	defer conn.Close() //在方法结束前关闭连接
//
//	go func() {
//		for {
//			//被选中的winner进行传入到 announcements中
//			msg := <-announcements
//			io.WriteString(conn, msg)
//		}
//	}()
//	// 校验 address
//	var address string
//	// 允许验证者输入他想要加入的tokens的数量
//	// 拥有足够多的tokens，就更有机会获得新的块
//	io.WriteString(conn, "Enter token balance:") //使用natcat进行输入值
//	scanBalance := bufio.NewScanner(conn)
//	for scanBalance.Scan() {
//		balance, err := strconv.Atoi(scanBalance.Text())
//		if err != nil {
//			log.Printf("%v not a number: %v", scanBalance.Text(), err)
//			return
//		}
//
//		t := time.Now()
//		address = calculateHash(t.String())
//		validators[address] = balance
//		fmt.Println(validators)
//		break
//	}
//
//	io.WriteString(conn, "\nEnter a new BPM:")
//
//	scanBPM := bufio.NewScanner(conn)
//
//	go func() {
//		for {
//			//进行输入的BPM验证
//			for scanBPM.Scan() {
//				bpm, err := strconv.Atoi(scanBPM.Text())
//				//如果恶意方试图用错误的输入来改变链,则将此map删除
//				//在这使用了一个简单的逻辑，就是判断输入的BMP是否为一个整数格式
//				if err != nil {
//					log.Printf("%v not a number: %v", scanBPM.Text(), err)
//					delete(validators, address)
//					conn.Close()
//				}
//
//				mutex.Lock()
//				oldLastIndex := Blockchain[len(Blockchain)-1]
//				mutex.Unlock()
//
//				//创建新块block，并考虑起是否伪造
//				newBlock, err := generateBlock(oldLastIndex, bpm, address)
//				if err != nil {
//					log.Println(err)
//					continue //输出所有log err
//				}
//				if isBlockValid(newBlock, oldLastIndex) {
//					candidateBlocks <- newBlock
//				}
//				io.WriteString(conn, "\nEnter a new BPM:")
//			}
//		}
//	}()
//
//	//模拟接收广播
//	for {
//		time.Sleep(time.Minute)
//		mutex.Lock()
//		//用一个规整的json格式输出区块
//		output, err := json.MarshalIndent(Blockchain, "", "\t")
//		mutex.Unlock()
//		if err != nil {
//			log.Fatal(err)
//		}
//		io.WriteString(conn, string(output)+"\n")
//	}
//
//}
//
////选择winner，通过随机选择块来选择验证者来伪造一个区块链，并通过标记的数量加权
//func pickWinner() {
//	time.Sleep(time.Second * 30)
//	mutex.Lock()
//	temp := tempBlocks
//	mutex.Unlock()
//
//	lotteryPool := []string{}
//	if len(temp) > 0 {
//
//	OUTER:  //使用这个循环来判断是否已经存在相同的验证在temp当中
//		for _, block := range temp { //索引值不用，设"_"
//			// if already in lottery pool, skip
//			for _, node := range lotteryPool {
//				if block.Validator == node {
//					continue OUTER
//				}
//			}
//
//			mutex.Lock()
//			setValidators := validators
//			mutex.Unlock()
//
//			k, ok := setValidators[block.Validator]
//			if ok {
//				for i := 1; i < k; i++ {
//					lotteryPool = append(lotteryPool, block.Validator)
//				}
//			}
//		}
//
//		//从池(lotteryPool)中随机选取winner
//		s := rand.NewSource(time.Now().Unix())
//		r := rand.New(s)
//		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]
//
//		//添加winner中的块block 并让所有节点知道
//		for _, block := range temp {
//			if block.Validator == lotteryWinner {
//				mutex.Lock()
//				Blockchain = append(Blockchain, block)
//				mutex.Unlock()
//				for range validators {
//					announcements <- "\n winning validator: " + lotteryWinner + "\n"
//				}
//				break
//			}
//		}
//
//		mutex.Lock()
//		tempBlocks = []Block{}
//		mutex.Unlock()
//	}
//}
//
//
//
//func main() {
//	//在同目录下创建prop.env文件("PORT=8088")
//	err := godotenv.Load("prop.env")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//构建创世块genesisBlock
//	t := time.Now()
//	genesisBlock := Block{}
//	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}
//	spew.Dump(genesisBlock)
//
//	Blockchain = append(Blockchain, genesisBlock)
//	//读取.env文件，获取Server端口8088
//	httpPort := os.Getenv("PORT")
//	//监听server
//	server, err := net.Listen("tcp", ":"+httpPort)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("HTTP Server Listening on port : ", httpPort)
//	defer server.Close()
//
//	go func() {
//		for candidate := range candidateBlocks {
//			mutex.Lock()
//			tempBlocks = append(tempBlocks, candidate)
//			mutex.Unlock()
//		}
//	}()
//
//	go func() {
//		for {
//			pickWinner() //选举winner
//		}
//	}()
//
//	for {
//		conn, err := server.Accept() //开启服务
//		if err != nil {
//			log.Fatal(err)
//		}
//		go handleConn(conn) //处理连接
//	}
//
//}
