package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "io"
    "log"
    "net/http"
    "os"
    "time"
    "fmt"
    "github.com/davecgh/go-spew/spew"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)


/*
*它代表组成区块链的每一个块的数据模型
*/
type Block struct {
    Index     int       //是这个块在整个链中的位置
    Timestamp string    //显而易见就是块生成时的时间戳
    BPM       int       //每分钟心跳数，也就是心率   采用人类平静时的心跳数据 (存储在区块链上的数据)
    Hash      string    //是这个块通过 SHA256 算法生成的散列值
    PrevHash  string    //代表前一个块的 SHA256 散列值
}



/*
*定义一个结构表示整个链，最简单的表示形式就是一个 Block 的 slice
*/
var Blockchain []Block


// 我们使用散列算法（SHA256）来确定和维护链中块和块正确的顺序，确保每一个块的 PrevHash 值等于前一个块中的 Hash 值




/**
用来计算给定的数据的 SHA256 散列值
*/
func calculateHash(block Block) string {
    record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}



/**
生成块的函数

接受一个“前一个块”参数，和一个 BPM 值
*/
func generateBlock(oldBlock Block, BPM int) (Block, error) {
    var newBlock Block

    t := time.Now()
    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = t.String()
    newBlock.BPM = BPM
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Hash = calculateHash(newBlock)

    return newBlock, nil
}



/**
校验块

1、检查 Index 来看这个块是否正确得递增
2、检查 PrevHash 与前一个块的 Hash 是否一致
3、再来通过 calculateHash 检查当前块的 Hash 值是否正确
*/
func isBlockValid(newBlock, oldBlock Block) bool {
    if oldBlock.Index+1 != newBlock.Index {
        return false
    }
    if oldBlock.Hash != newBlock.PrevHash {
        return false
    }
    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }
    return true
}




/**
将本地的过期的链切换成最新的链
*/
func replaceChain(newBlocks []Block) {
    if len(newBlocks) > len(Blockchain) {
        Blockchain = newBlocks
    }
}


//下面是搭建一个web入口来查询 区块的信息


/**
初始化web 服务
*/
func run() error {
    mux := makeMuxRouter()
    httpAddr := os.Getenv("ADDR")
    fmt.Println("ADDR:=" + httpAddr)
    // log.Println("Listening on ", os.Getenv("ADDR"))
    s := &http.Server{
        // Addr:           ":" + httpAddr,
         Addr:           ":" + "8080",
        Handler:        mux,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    if err := s.ListenAndServe(); err != nil {
        return err
    }

    return nil
}






func makeMuxRouter() http.Handler {
    muxRouter := mux.NewRouter()
    muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
    muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
    return muxRouter
}







/**
GET 请求的 handler
*/
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
    bytes, err := json.MarshalIndent(Blockchain, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    io.WriteString(w, string(bytes))
}



//POST请求的东东

type Message struct {
    BPM int
}


/**
POST 请求的 handler
*/
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
    var m Message

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&m); err != nil {
        respondWithJSON(w, r, http.StatusBadRequest, r.Body)
        return
    }
    defer r.Body.Close()

    newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m.BPM)
    if err != nil {
        respondWithJSON(w, r, http.StatusInternalServerError, m)
        return
    }
    if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
        newBlockchain := append(Blockchain, newBlock)
        replaceChain(newBlockchain)
        spew.Dump(Blockchain)
    }

    respondWithJSON(w, r, http.StatusCreated, newBlock)

}



/**
处理post的响应
*/
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
    response, err := json.MarshalIndent(payload, "", "  ")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("HTTP 500: Internal Server Error"))
        return
    }
    w.WriteHeader(code)
    w.Write(response)
}




func main() {
    // fileNameArr := []string{"example.json"}
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        t := time.Now()

        // genesisBlock （创世块）是 main 函数中最重要的部分，通过它来初始化区块链，毕竟第一个块的 PrevHash 是空的
        genesisBlock := Block{0, t.String(), 0, "", ""}   
        //使用spew.Dump 这个函数可以以非常美观和方便阅读的方式将 struct、slice 等数据打印在控制台里，方便我们调试
        spew.Dump(genesisBlock)
        Blockchain = append(Blockchain, genesisBlock)
    }()
    log.Fatal(run())

}