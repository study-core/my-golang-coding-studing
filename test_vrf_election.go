package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	mrand "math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/secp256k1"
	"github.com/PlatONnetwork/PlatON-Go/crypto/vrf"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/plugin"
	"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
)

var (
	totalWeight     = flag.Uint("totalWeight", 5000000000, "待分配的总权重，为每个候选人从该值中分配随机数量的权重，直至分配完所有权重")
	dividend        = flag.Uint("dividend", 200, "用于计算二项分布的P值，dividend/总权重=P，该值直接影响选举的概率")
	number          = flag.Uint("number", 100, "测试选举的次数，该值表示一共进行多少轮的选举测试")
	candidateNumber = flag.Uint("candidateNumber", 101, "待选举的人数，该值表示候选人数量")
	electionNumber  = flag.Uint("electionNumber", 8, "需要从候选人中选举出来的人数")
	stakeThreshold  = flag.Uint("stakeThreshold", 1000000, "每个候选人的最低质押金，模拟的每个候选人的质押金>=该值")
	randRange       = flag.Uint("randRange", 130000000, "为每个候选人分配的随机质押金的范围，1~N，N为该值")
	file            = flag.String("file", "", "文件地址，该文件内容为每个候选人的权重，如指定了该地址则使用该文件内配置的权重，不自动随机分配权重（格式为：地址\\t权重）")
)

var curve = secp256k1.S256()

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", "根据指定参数利用VRF和二项分布模拟真实选举，对多次选举的结果进行统计（执行结果会输出到该文件所在目录）；注意以下规则：\n"+
			"\t1.候选人权重的分配规则为，首先为每个候选人分配100W的LAT，然后再为每个候选人分配随机的质押金；\n"+
			"\t2.如果指定了file地址，则使用配置的权重否则自动随机分配权重（默认为随机分配权重）；\n"+
			"\t3.二项分布的P值计算是由dividend/totalWeight得来的；\n"+
			"\t4.如果调整了totalWeight的值，则需要视情况调整randRange，否则将导致分配给候选人的质押金偏低或偏高，分配不均匀；\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	TestStakingPlugin_Probability()
}

type temp struct {
	count int
	v     *staking.Validator
	sk    *ecdsa.PrivateKey
}

func TestStakingPlugin_Probability() {
	startTime := time.Now()
	xcom.GetEc(xcom.DefaultMainNet)
	vqList := make(staking.ValidatorQueue, 0)
	statistics := make(map[discover.NodeID]*temp)
	preNonces := make([][]byte, 0)
	r := mrand.New(mrand.NewSource(time.Now().Unix()))
	currentNonce := crypto.Keccak256([]byte(RandString(r, 20)))

	totalShares := new(big.Int)

	if *file == "" {
		sumWeight := int(*totalWeight - (*candidateNumber * *stakeThreshold))
		for i := 0; i < int(*candidateNumber); i++ {
			privKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
			nodeId := discover.PubkeyID(&privKey.PublicKey)
			addr := crypto.PubkeyToAddress(privKey.PublicKey)
			mrand.Seed(time.Now().UnixNano())
			shares := new(big.Int).SetUint64(uint64(*stakeThreshold))
			if sumWeight > 0 {
				randNum := mrand.Intn(int(*randRange))
				if sumWeight >= randNum {
					sumWeight -= randNum
				} else {
					randNum = sumWeight
					sumWeight = 0
				}
				shares.Add(shares, new(big.Int).SetUint64(uint64(randNum)))
			}
			totalShares.Add(totalShares, shares)
			shares.Mul(shares, new(big.Int).SetInt64(1e18))
			v := &staking.Validator{
				NodeAddress:     addr,
				NodeId:          nodeId,
				ProgramVersion:  1,
				StakingTxIndex:  1,
				ValidatorTerm:   1,
				StakingBlockNum: 1,
				Shares:          shares,
			}
			vqList = append(vqList, v)
			t := &temp{
				v:  v,
				sk: privKey,
			}
			statistics[v.NodeId] = t
			time.Sleep(time.Nanosecond)
		}
	} else {
		vqList, statistics, totalShares = ReadCandidateWeight(*file)
		if len(vqList) == 0 || len(statistics) == 0 {
			panic("读取候选人权重信息失败!")
		}
	}

	for i := 0; i < len(vqList); i++ {
		for j := i + 1; j <= len(vqList)-1; j++ {
			if vqList[i].Shares.Cmp(vqList[j].Shares) < 0 {
				vqList[i], vqList[j] = vqList[j], vqList[i]
			}
		}
	}
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if nil != err {
		panic(err)
	}
	var vrfProveTime time.Duration
	var vrfProveCount uint64
	var electionTime time.Duration
	var electionCount uint64

	n := int(*number)
	for i := 0; i < n; i++ {
		preNonces = make([][]byte, 0)
		for _, value := range statistics {
			sTime := time.Now()
			pi, err := vrf.Prove(value.sk, currentNonce)
			if nil != err {
				panic(err)
			}
			vrfProveTime += time.Since(sTime)
			vrfProveCount++
			currentNonce = vrf.ProofToHash(pi)
			preNonces = append(preNonces, currentNonce)
		}
		sTime := time.Now()
		pi, err := vrf.Prove(privKey, currentNonce)
		vrfProveTime += time.Since(sTime)
		vrfProveCount++
		if nil != err {
			panic(err)
		}
		currentNonce = vrf.ProofToHash(pi)
		sTime = time.Now()
		result, err := plugin.ProbabilityElection(vqList, int(*electionNumber), currentNonce, preNonces, *dividend)
		electionTime += time.Since(sTime)
		electionCount++
		if nil != err {
			log.Panic("Failed to ProbabilityElection, err:", err)
			return
		}
		for _, v := range result {
			statistics[v.NodeId].count++
		}
	}

	file, err := os.OpenFile(filepath.Join(GetCurrentPath(), fmt.Sprintf("test_vrf_election_result_%v.txt", time.Now().Nanosecond())), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if nil != err {
		fmt.Println(err)
		return
	}
	title := fmt.Sprintf("节点地址\t权重(LAT)\t被选中总次数")
	log.Println(title)
	_, err = file.WriteString(title + "\n")
	if nil != err {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for _, value := range vqList {
		shares := value.Shares
		electionResult := fmt.Sprintf("%v\t%v\t%v", hex.EncodeToString(value.NodeAddress.Bytes()), shares.Div(shares, new(big.Int).SetInt64(1e18)).Text(10), statistics[value.NodeId].count)
		_, err = file.WriteString(electionResult + "\n")
		if nil != err {
			fmt.Println(err)
			return
		}
		log.Println(electionResult)
	}
	outResult := fmt.Sprintf("测试选举完成!!! 结果已输出到：%v\n\t总耗时：%v\t生成VRF证明平均耗时：%v\t选举平均耗时：%v\n"+
		"\ttotalShares：%v\tdividend：%v\t总选举次数：%v",
		file.Name(), time.Since(startTime).String(), time.Duration(uint64(vrfProveTime)/vrfProveCount).String(), time.Duration(uint64(electionTime)/electionCount).String(),
		totalShares, *dividend, *number)
	_, err = file.WriteString(outResult)
	if nil != err {
		fmt.Println(err)
		return
	}
	log.Println(outResult)
}

func RandString(r *mrand.Rand, len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func GetCurrentPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return string(path[0 : i+1])
}

func ReadCandidateWeight(fileName string) (staking.ValidatorQueue, map[discover.NodeID]*temp, *big.Int) {
	file, err := os.Open(fileName)
	if nil != err {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	statistics := make(map[discover.NodeID]*temp)
	vqList := make(staking.ValidatorQueue, 0)
	totalShares := new(big.Int)
	for {
		input, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		strs := strings.Split(string(input), "\t")
		privKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
		nodeId := discover.PubkeyID(&privKey.PublicKey)
		addr := crypto.PubkeyToAddress(privKey.PublicKey)

		weight, bool := new(big.Int).SetString(strs[1], 10)
		if !bool {
			panic("转换权重失败!")
		}
		totalShares.Add(totalShares, weight)
		weight.Mul(weight, new(big.Int).SetInt64(1e18))
		v := &staking.Validator{
			NodeAddress:     addr,
			NodeId:          nodeId,
			ProgramVersion:  1,
			StakingTxIndex:  1,
			ValidatorTerm:   1,
			StakingBlockNum: 1,
			Shares:          weight,
		}
		vqList = append(vqList, v)
		t := &temp{
			v:  v,
			sk: privKey,
		}
		statistics[v.NodeId] = t
	}
	return vqList, statistics, totalShares
}
