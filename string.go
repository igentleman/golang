package main

import (
	"fmt"
	"strings"
)

// EqualFoldTest 不区分大小写对比
func EqualFoldTest(){ //true
	var str = "zfck"
	fmt.Println(strings.EqualFold("Zfck",str))
}

/**
     ********     	HasPrefix 判断前后缀字符串 		***********
 */

// HasPrefixTest 判断是否以固定前缀字符串
func HasPrefixTest(){ //true
	fmt.Println(strings.HasPrefix("www.zfck.net","www.z"))
}

// HasSuffixTest 判断是否以固定后缀字符串
func HasSuffixTest(){ //true
	fmt.Println(strings.HasSuffix("www.zfck.net","k.net"))
}


/**
		*******      Contains 包含函数       *********
 */

// ContainsTest 判断所给字符串是否包含特定子字符串
func ContainsTest(){ //true
	fmt.Println(strings.Contains("www.zfck.net","ck."))
}

// ContainsRuneTest 判断所给字符串是否包含特定子字符(指的是字符！！！！！  包括ascii码值 ！！！)
func ContainsRuneTest(){ //true
	fmt.Println(strings.ContainsRune("www.zfck.net",'c'))  //true
	fmt.Println(strings.ContainsRune("www.zfck.net",117))  //false
	fmt.Println(strings.ContainsRune("www.zfck.net",46))   //true
}

// ContainsAnyTest 判断所给字符串是否包含某一个子字符(只要有一个字符串在目标字符串中，那么都返回true)
func ContainsAnyTest(){ //true
	fmt.Println(strings.ContainsAny("www.zfck.net","abcde"))  //true
	fmt.Println(strings.ContainsAny("www.zfck.net","sdcda"))  //true
	fmt.Println(strings.ContainsAny("www.zfck.net","llddd"))   //false
}


// CountTest 判断所给字符串在子字符串的个数
func CountTest(){
	fmt.Println(strings.Count("www.zfck.net","ck."))  //1
	fmt.Println(strings.Count("www.zfck.net","e"))	//1
	fmt.Println(strings.Count("www.zfck.net","w"))	//3
	fmt.Println(strings.Count("www.zfck.net","wwww"))	//0
}

// IndexTest 子串(第二个参数)在字符串(第一个参数)中第一次出现的位置，不存在则返回-1。
func IndexTest(){
	fmt.Println(strings.Index("www.zfck.net","ck."))  //6
	fmt.Println(strings.Index("www.zfck.net","e"))	//10
	fmt.Println(strings.Index("www.zfck.net","w"))	//0
	fmt.Println(strings.Index("www.zfck.net","wwww"))	//-1
}

// IndexByteTest 子字符(第二个参数)在字符串(第一个参数)中第一次出现的位置，不存在则返回-1。
func IndexByteTest(){
	fmt.Println(strings.IndexByte("www.zfck.net",'z'))     //4
	fmt.Println(strings.IndexByte("www.zfck.net",'s'))	//-1
	fmt.Println(strings.IndexByte("www.zfck.net",'a'))	//-1
	fmt.Println(strings.IndexByte("www.zfck.net",'v'))	//-1
}

// IndexRuneTest 子字符(第二个参数)在字符串(第一个参数)中第一次出现的位置，不存在则返回-1。
func IndexRuneTest(){
	fmt.Println(strings.IndexRune("www.zfck.net",'z'))     //4
	fmt.Println(strings.IndexRune("www.zfck.net",'s'))	//-1
	fmt.Println(strings.IndexRune("www.zfck.net",'a'))	//-1
	fmt.Println(strings.IndexRune("www.@zfck.net",64))	//4
}

// IndexAnyTest 子字符(第二个参数)在字符串(第一个参数)中第一次出现的位置，不存在则返回-1。   类似 ContainAny
func IndexAnyTest(){
	fmt.Println(strings.IndexAny("www.zfck.net","wwwz"))   //0
	fmt.Println(strings.IndexAny("www.zfck.net","seee"))	//10
	fmt.Println(strings.IndexAny("www.zfck.net","asd"))	//-1
	fmt.Println(strings.IndexAny("www.@zfck.net","asdf"))	//6
}

/***********  indexFunc的用法 **************/

// f p为字符类型
func f(p rune) bool{
	if p == 'w'{
		return true
	}else {
		return false
	}
}

// IndexFuncTest 该函数可以调用一个函数（包括匿名函数）,不存在则返回-1。
func IndexFuncTest(){
	fmt.Println(strings.IndexFunc("www.zfck.net",f)) // 0
	fmt.Println(strings.LastIndexFunc("www.zfck.net",f))  // 2
}

// LastIndexTest 子串(第二个参数)在字符串(第一个参数)中最后一次出现的位置，不存在则返回-1。
func LastIndexTest(){
	fmt.Println(strings.LastIndex("www.zfck.net","w"))    //2
}

// LastIndexAnyTest 子串任意一个字符(第二个参数)在字符串(第一个参数)中最后一次出现的位置，不存在则返回-1。
func LastIndexAnyTest(){
	fmt.Println(strings.LastIndexAny("www.zfck.net","wasdfwte"))    //11
}

// TitleTest 将每句话的第一个数变成大写
func TitleTest(){
	fmt.Println(strings.Title("hellow world")) //	Hellow World
	fmt.Println(strings.Title("www.zfck.net world")) //	Www.Zfck.Net World
}

// ToLowerTest 将所有字母都转为对应的小写(只进行值拷贝)
func ToLowerTest(){
	var upStr = "WWW.ZFCK.NET"
	fmt.Println(strings.ToLower(upStr)) // www.zfck.net
	fmt.Println(upStr) 	// WWW.ZFCK.NET
}

// ToUpperTest 将所有字母都转为对应的小写(只进行值拷贝)
func ToUpperTest(){
	var upStr = "www.zfck.net"
	fmt.Println(strings.ToUpper(upStr)) // www.zfck.net
	fmt.Println(upStr) 	// www.zfck.net
}

// ToTitleTest 将每句话都变成大写（类似ToUpper）
func ToTitleTest(){
	fmt.Println(strings.ToTitle("hellow world")) //	HELLOW WORLD
	fmt.Println(strings.ToTitle("www.zfck.net world")) //	WWW.ZFCK.NET WORLD
}

// RepeatTest 拼接并重复提供给的字符串
func RepeatTest(){
	fmt.Println("刘" + strings.Repeat("粤新",2)) // 刘粤新粤新
}

// ReplaceTest 替换老字符 数字代表替换几个字符
func ReplaceTest(){
	fmt.Println(strings.Replace("www.zfck.net","w","T",3))
	fmt.Println(strings.Replace("www.zfck.net","w","T",2))
	fmt.Println(strings.Replace("www.zfck.net","w","T",1))
	fmt.Println(strings.Replace("www.zfck.net","w","T",-1))
}

func r(p rune) rune{
	fmt.Println(string(p)) // www.zfck.net
	if p == 't'{
		return '正'
	}
	if p == 'w'{
		return '夫'
	}
	return p
}

// MapTest 该方法会循环将提供的字符串传到 r 函数进行匹配与兑换
func MapTest(){
	fmt.Println(strings.Map(r,"www.zfck.net")) // 夫夫夫.zfck.ne正
}

// TrimTest 去除提供字符串两边给定的字符
func TrimTest(){
	fmt.Printf("[%v]",strings.Trim("www.zfck.netwwww","w"))  //[.zfck.net]
}

// TrimSpaceTest 除去两边的空格以及制表符
func TrimSpaceTest(){
	fmt.Println(strings.TrimSpace("     \n\t\bwww.zf\tck.net    \b\n")) // www.zf  ck.net
}

// TrimLeftRightTest 匹配前/后端
func TrimLeftRightTest(){
	//前
	fmt.Println(strings.TrimLeft("www.zfck.net","wwww")) // .zfck.net

	//后
	fmt.Println(strings.TrimRight("www.zfck.netw","wwww")) // www.zfck.net
}

// TrimPrefixSuffixTest 返回去除（第一个参数）可能的前缀prefix的字符串。
func TrimPrefixSuffixTest(){
	//前
	fmt.Println(strings.TrimPrefix("www.zfck.net","wwww")) // www.zfck.net
	fmt.Println(strings.TrimPrefix("www.zfck.net","www.")) // zfck.net

	//后
	fmt.Println(strings.TrimSuffix("www.zfck.net","wwww")) // www.zfck.net
	fmt.Println(strings.TrimSuffix("www.zfck.net",".net")) // www.zfck
}

func main(){
	TrimPrefixSuffixTest()
}