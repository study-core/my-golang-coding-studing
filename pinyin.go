


package main

import (
    "fmt"
    // "regexp"
    "unicode"
    // "reflect"
    "bytes"
    pinyin "github.com/linkedin-inc/go-pinyin"
)







func main() {
    // letterArr := []string{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1", "i1", "z1", "k1", "l1", "m1", "n1", "o1", "p1", "q1", "r1", "s1", "t1", "u1", "v1", "w1", "x1", "y1", "z1",
    //                         "a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2", "i2", "z2", "k2", "l2", "m2", "n2", "o2", "p2", "q2", "r2", "s2", "t2", "u2", "v2", "w2", "x2", "y2", "z2",
    //                         "a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3", "i3", "z3", "k3", "l3", "m3", "n3", "o3", "p3", "q3", "r3", "s3", "t3", "u3", "v3", "w3", "x3", "y3", "z3",
    //                         "a4", "b4", "c4", "d4", "e4",  "f4", "g4", "h4", "i4", "z4", "k4", "l4", "m4", "n4", "o4", "p4", "q4", "r4", "s4", "t4", "u4", "v4", "w4", "x4", "y4", "z4",

    //                     }
   
   fmt.Println(GetNextAttrName(31))
}


func IsChineseChar(str string) bool {
    for _, r := range str {
        if unicode.Is(unicode.Scripts["Han"], r) {
            return true
        }
    }
    return false
}


func PickChineseChar(str string) string {
    /**
    strings.Join 最慢
    fmt.Sprintf 和 string + 差不多
    bytes.Buffer又比上者快约500倍
     */
    var buf bytes.Buffer  //用这种拼接字符串 效率最快
    fmt.Print("原始字符串为:" + str + "\n最终取出的字符串为:")
    for _, r := range str {
        // 判断字符是否为汉字
        if unicode.Is(unicode.Scripts["Han"], r) {
            fmt.Printf("%c", r)
            buf.WriteString(string(r))
        }
    }
    fmt.Println("\n")
    return buf.String()
}



func GetFirstEChar(str string) string {
    var firstChar string  //英文字符
    for i, r := range str {
        if i == 0 {
            firstChar = string(r)
        }
    }
    return ConvPinyinSpellString(firstChar, false, true, ",")
}




func GetHeadEChar(str string) string {
    var fHasChinese bool  //用这种拼接字符串 效率最快
    var fChar string
    for i, r := range str {
        // 判断字符是否为汉字
        if i == 0 && unicode.Is(unicode.Scripts["Han"], r) {
            fHasChinese = true
        }
        if i == 0 {
            fChar = string(r)
        }
    }
    if fHasChinese {
        return ConvPinyinSpellString(str, false, true, ",")
    }
    return fChar
}


func GetF3EChar(str string) string {
    var buf bytes.Buffer
    var i int
    for _, r := range str {
        
        if i == 0 || i == 1 || i == 2 {
            buf.WriteString(string(r))
            fmt.Println("i:=" + fmt.Sprint(i) + ",r:=" + string(r))
        }
        i++
    }
    return ConvPinyinSpellString(buf.String(), false, true, ",")
}





func GetNextLetterCharBy26Index(index  int) string {
    letterArr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
    if index < 0 || index > len(letterArr) - 1 {
        return letterArr[0]
    }
    return letterArr[index]
}


func GetNextAttrName(count int) string {
    //求除数
    divNum := count/26
    //求余数
    remNum := count%26
    fmt.Println("总个数为:" + fmt.Sprint(count) + ",除于26=" + fmt.Sprint(divNum) + ",余数=" + fmt.Sprint(remNum))
    return GetNextLetterCharBy26Index(remNum) + fmt.Sprint(divNum + 1)
}



/*
    str 字符串
    isPolyphone 是否开启多音字模式
    isFirstLetter 是否开启首字母模式
 */
func ConvPinyinSpell(str string, isPolyphone, isFirstLetter bool) []string {

    args := pinyin.NewArgs()
    args.Heteronym = isPolyphone
    pys := pinyin.Pinyin(str, args)

    var ss = make([]string, 0)
    var cp = make([]string, 0)

    for _, py := range pys {
        var set = make(map[string]int)
        for _, pyl := range py {
            if isFirstLetter {
                k := string(pyl[0])
                set[k] = 0
            } else {
                set[pyl] = 0
            }
        }
        if len(ss) != 0 {
            cp = ss
            ss = make([]string, 0)
            for k := range set {
                for _, str := range cp {
                    ss = append(ss, str+k)
                }

            }
        } else {
            for k := range set {
                ss = append(ss, k)
            }
        }
    }

    return ss
}

/*
    str 字符串
    isPolyphone 是否开启多音字模式
    isFirstLetter 是否开启首字母模式
    sep 分割符
 */
func ConvPinyinSpellString(str string, isPolyphone, isFirstLetter bool, sep string) string {
    if len(str) == 0 {
        return ""
    }
    var buffer bytes.Buffer
    ss := ConvPinyinSpell(str, isPolyphone, isFirstLetter)

    for _, str := range ss {
        buffer.WriteString(sep + str)
    }
    spell := buffer.String()[len(sep):]
    return string(spell)
}

func ConvFullPinyinSpell(str string) string {
    if len(str) == 0 {
        return ""
    }
    spell := ConvPinyinSpellString(str, true, true, ",") + "," +
            ConvPinyinSpellString(str, true, false, ",")
    return spell
}