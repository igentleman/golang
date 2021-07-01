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

func f(str string){}


func IndexFunc(){

}

func main(){
	IndexAnyTest()
}