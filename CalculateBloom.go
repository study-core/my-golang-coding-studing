package main

import (
	"fmt"
	"math"
)

func main() {

	//pos := float64(0.1)
	//ln2 := math.Log(2)
	//
	//num := float64(2048) * math.Pow(ln2, 2)/ -1 /math.Log(pos)
	//fmt.Println(num)

	bitm := optimalNumOfBits(427, 0.9)
	fmt.Println(bitm)

	k := optimalNumOfHashFunctions(2048, 427)
	fmt.Println(k)
}



/**
 * 计算 Bloom Filter的bit位数m
 *
 * <p>See http://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives for the
 * formula.
 *
 * @param n 预期数据量
 * @param p 误判率 (must be 0 < p < 1)
 */
func optimalNumOfBits(n uint64, p float64) int64 {
	if 0 == p {
		p = math.MaxFloat64
	}
	//m = -1 * (n * lnP)/(ln2)^2
	//return int64(-float64(n)*math.Log(p)/(math.Log(2)*math.Log(2)))
	ln2 := math.Log(2)
	return int64(-1 * (float64(n) * math.Log(p)) / math.Pow(ln2, 2))
}




/**
 * 计算最佳k值，即在Bloom过滤器中插入的每个元素的哈希数
 *
 * <p>See http://en.wikipedia.org/wiki/File:Bloom_filter_fp_probability.svg for the formula.
 *
 * @param n 预期数据量
 * @param m bloom filter中总的bit位数 (must be positive)
 */
func optimalNumOfHashFunctions (n, m uint64) int64 {
	//k = m/n * ln2    int(math.Ceil(float64(b.m) / nFloat * ln2))
	//return int64(math.Max(float64(1),  math.Round( float64(m) / float64(n) * math.Log(float64(2)))))
	ln2 := math.Log(2)
	return int64(math.Ceil(float64(m) / float64(n) * ln2))
}