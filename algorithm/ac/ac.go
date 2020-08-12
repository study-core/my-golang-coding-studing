package main


// https://zhuanlan.zhihu.com/p/80325757
// https://www.cnblogs.com/nullzx/p/7499397.html   todo 正在看的
func main() {

}


// todo  ac 自动机  (多模式 字符串匹配算法)

// 我们经常用的字符串方法indexOf，都是判定两个字符串的包含关系，底层使用类似KMP，BM， Sunday这样的算法。
// todo 如果我们要判断一个长字符串是否包含多个短字符串呢？
// 比如在一篇文章找几个敏感词，在DNA串中找几个指定的基因对pattern进行预处理，
// 如果我们的模式串存在多个，则不适合了，我们就需要用到一种多模式匹配算法。


// AC自动机的核心算法仍然是寻找模式串内部规律，达到在每次失配时的高效跳转。这一点与单模式匹配KMP算法是一致的。
// todo 不同的是，AC算法寻找的是模式串之间的相同前缀关系。

// todo AC自动机  其实就是  前缀Trie + KMP