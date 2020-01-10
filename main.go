package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	/*
		var a int = 1
		var b string = "test"
		var c []float32
		var e struct {
			x int
			y string
		}
		//声明一个字符串变量
		f := ""
	*/
	fmt.Println("Hello World")
	x := 1
	y := 2
	fmt.Println(x, y)
	//两个变量 change
	x, y = y, x
	fmt.Println(x, y)

	//输出正弦图像 start
	const size = 300
	pic := image.NewGray(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	//fmt.Println(pic)

	file, err := os.Create("pic.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, pic)

	file.Close()

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2

		pic.SetGray(x, int(y), color.Gray{0})
	}

	file2, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file2, pic)

	file2.Close()
	//输出正弦图像 end

	f := false //默认值false
	fmt.Println(f)

	//使用``方式定义的字符串会保留源格式
	s := `one line
two line`

	fmt.Println(s)

	//声明一个切片
	var arr = make([]int, 3)

	arr[0] = 1
	arr[1] = 2
	//runtime error: index out of range [3] with length 3 数据越界异常
	//arr[3] = 3

	//[1 2 0] 打印出来的不是地址值，而是具体值
	fmt.Println(arr)
	//从指定索引位置切开 [)
	fmt.Println(arr[1:]) //[2 0]
	fmt.Println(arr[:1]) //[1]

	var str = "abc"
	ptr := &str //ptr为str的指针
	//https://www.jianshu.com/p/8be8d36e779c 这个网址里面有具体的格式代表输出的内容
	fmt.Printf("%T\n", ptr) //打印类型
	fmt.Printf("%p\n", ptr) //打印指针地址
	fmt.Printf("%T\n", str)

	value := *ptr //value为指针对应的值
	fmt.Printf("%T\n", value)
	fmt.Printf("%s\n", value) //%s 正常输出字符串
	/*
		& 与 * 互为相反的操作， &是取变量的指针  *是取指针对应的值
	*/

	a := 1
	b := 2

	swap(&a, &b)

	fmt.Println(a, b)

	//基于flag包 创建一个 字符串类型的指针 mode
	var mode = flag.String("mode", "", "this is mode")
	fmt.Printf("%T\n", mode) //*string

	//基于new函数创建指针变量
	ptr2 := new(string)
	*ptr2 = "value by ptr2"
	str2 := *ptr2                  //这步相当于声明了一个value值为 value by ptr2 的字符串，会在内存中开辟新的内存空间 类似于java中 new String("");
	fmt.Println(&str2, ptr2, str2) //0xc000030480 0xc000030470 value by ptr2

	//打印字符串的长度
	str3 := "this is str"
	fmt.Println(len(str3)) //11
	str4 := "测试"
	fmt.Println(len(str4)) //6 这里的6代表字节个数

	//打印字符个数
	fmt.Println(utf8.RuneCountInString(str4 + "abc"))

	//字符串遍历
	for _, s := range str4 {
		fmt.Printf("%c  %d\n", s, s)
	}

	tracer := "测试字符串123"
	//获取字符串指定字符所在位置的索引 这里获取是是按照ASCII编码，一个汉字算3个字符长度
	comma := strings.Index(tracer, "1")
	fmt.Println(comma)

	/*这边使用 rune数组来将字符串转换为字符数组
	  有两种字符 byte 和 rune， rune相当于java中的char
		byte对应的是ASCII编码格式
		rune对应的是Unicode
	*/
	str5 := "测试12345"
	byteArr5 := []rune(str5)
	for i := 0; i < len(byteArr5); i++ {
		fmt.Println(string(byteArr5[i]))
	}

	//字符串拼接
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString("abc")
	stringBuilder.WriteString("123")
	fmt.Println(stringBuilder.String())

	//读取kv格式文件
	fileName := "kv.txt"
	iniFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bufio.NewReader(iniFile)
	var key string = "name"
	for {
		lineStr, err := reader.ReadString('\n') //读取字符串直到\n 表示读取一行 因为换行符为\t\n
		if err != nil {
			fmt.Println(err.Error()) //EOF  golang之文件结尾错误（EOF） EOF就表示文件到了结尾，通过异常来判断读文件是否结束
			break
		}
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}

		var pair []string = strings.Split(lineStr, "=")

		if len(pair) == 2 {
			//键值对
			currentKey := strings.TrimSpace(pair[0])
			if currentKey == key {
				println(strings.TrimSpace(pair[1]))
			}
		}
	}
	defer iniFile.Close() //defer和Java中finally相似 表示程序结束前执行

	//声明常量
	const c1 string = "constant"

	//自定义数据类型
	type newInt int     //定义了一个新的类型叫 newInt 默认按照int的类型看待
	type intAlias = int //定义了int的别名，本质就是int

	var newIntA newInt = 1
	var intB intAlias = 2

	fmt.Printf("%T\n", newIntA) //main.newInt
	fmt.Printf("%T\n", intB)    //int

	//声明数组
	arr2 := [...]string{"a", "b", "c"}
	fmt.Println(arr2)

	for k, v := range arr2 {
		fmt.Println(k, v)
	}

	for i := 0; i < len(arr2); i++ {
		fmt.Println(i, arr2[i])
	}

	//声明切片
	var slice1 []int
	fmt.Println(slice1 == nil) //true

}

/*
*int 表示int类型的指针
 */
func swap(a *int, b *int) {
	c := *a //a指针指向的值赋值给c
	*a = *b //b指针指向的值赋值给a
	*b = c  //临时变量c的值赋值给b
}
