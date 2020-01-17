package main

import (
	"bufio"
	"bytes"
	list2 "container/list"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strings"
	"sync"
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

	//声明切片 切片的声明不需要声明长度，凡是带有长度声明的都是数组
	var slice1 []int
	fmt.Println(slice1 == nil) //true
	//基于make函数声明切片
	slice2 := make([]int, 2, 10) //声明一个大小为2 初始化分配内存空间10的切片
	slice3 := make([]int, 2)     //单纯的声明一个大小为2的切片
	fmt.Println(slice2)          //[0 0] len() = 2
	fmt.Println(slice3)          //[0 0] len() = 2
	slice2 = append(slice2, 1, 2, 3)
	fmt.Println(slice2) //[0 0 1 2 3]

	//切片复制 基于copy复制的是值，复制完是两个具有相同元素的不同切片
	const sliceLength int = 10
	var srcSlice = make([]int, sliceLength)
	for i := 0; i < sliceLength; i++ {
		srcSlice[i] = i
	}
	var destSlice = make([]int, sliceLength) //目标切片必须要先定义长度，否则会报错
	copy(destSlice, srcSlice)
	var refSlice = srcSlice
	srcSlice[0] = 10
	fmt.Println(destSlice[0], refSlice[0]) //0 10

	//切片删除指定元素 没有api只能手动切然后再拼 如果需要做大量的增删不建议使用切片的数据结构
	slice4 := []int{1, 2, 3, 4, 5}
	deleteIndex := 1
	slice4 = append(slice4[:1], slice4[deleteIndex+1:]...) //在新的切片后面添加上...表示添加整个切片
	fmt.Println(slice4)

	//map
	//声明
	map1 := make(map[string]int) //声明一个key为string， value为int的map
	//添加元素
	map1["xiaoming"] = 90
	map1["xiaohong"] = 80
	map1["xiaohei"] = 85
	fmt.Println("after add")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//删除元素
	fmt.Println("after delete")
	delete(map1, "xiaoming")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//修改元素
	map1["xiaohei"] = 88
	fmt.Println("after change")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//遍历查询
	fmt.Println("final println")
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	//获取指定key ok表示是否含有当前元素
	xiaohei_score, ok1 := map1["xiaohei"]
	xiaoming_score, ok2 := map1["xiaoming"]
	fmt.Println(xiaohei_score, ok1)  //88 true
	fmt.Println(xiaoming_score, ok2) //0 false

	//concurrent map
	var ccmap sync.Map
	ccmap.Store("name", "xiaoming")
	ccmap.Store("age", 11)
	map_value, ok3 := ccmap.Load("name")
	fmt.Println(map_value, ok3)
	ccmap.Delete("name")

	//list
	my_list := list2.New()
	//var my_list2 list2.List

	//add
	my_list.PushBack(1)  //列表后插入
	my_list.PushFront(2) //列表前插入

	element := my_list.PushBack(3) //返回值为当前元素的地址值
	my_list.InsertBefore(4, element)
	my_list.PushBack(5)

	//遍历并删除指定元素 不存在并发修改异常
	for i := my_list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
		if i.Value == 5 {
			my_list.Remove(i)
		}
	}

	/*流程控住语句
	if
	for
	for-range
	switch-case
	break
	continue 和java中用法一样，扩展了和break一样的loop标识
	*/

	i := 0
	//定义一个标签
	if i == 2 {
	} else {
		goto goTag
	}
goTag:
	fmt.Println("this is goto content and i is ", i)
	i++

	//方法是值传递
	var st1 = struct1{
		a: 0,
		b: "",
		c: &struct1{
			a: 0,
			b: "",
			c: nil,
		},
	}
	fmt.Printf("before st1: %p\n", &st1)
	fmt.Printf("before st1 value: %+v\n", st1)
	fmt.Printf("before st1.c.a value: %+v\n", st1.c.a)
	fmt.Printf("before st1.a value: %+v\n", st1.a)
	st2 := passByValue(st1)
	fmt.Printf("after st1: %p\n", &st1)
	fmt.Printf("after st1 value: %+v\n", st1)
	fmt.Printf("after st1.c.a value: %+v\n", st1.c.a)
	fmt.Printf("after st1.a value: %+v\n", st1.a)
	fmt.Printf("st2: %p\n", &st2)
	/*
		before st1: 0xc0000046e0
		before st1 value: {a:0 b: c:0xc000004700}
		before st1.c.a value: 0
		before st1.a value: 0
		func inner s: 0xc000004760
		func inner s value: {a:0 b: c:0xc000004700}
		after st1: 0xc0000046e0
		after st1 value: {a:0 b: c:0xc000004700}
		after st1.c.a value: 1
		after st1.a value: 0
		st2: 0xc000004740

		结论：
		golang中结构体作为方法入参时，会重新复制一块内存地址，结构体的属性，按照值传递的方式进行传递
			  基本数据类型的属性，在方法中修改值，是修改的值副本，不会对原结构体造成影响，而结构体类型的属性在方法中修改结构体的属性，会生效，因为结构体属性传递的是引用副本
			  所以为了实现类似于java中的传递引用副本的相似功能，避免对属性进行完全copy，方法参数数据类型为*struct,直接传递结构体的指针
		java中类对象作为方法入参时，会直接按照值传递的方式进行，直接传递引用副本
	*/

}

/*
*int 表示int类型的指针
 */
func swap(a *int, b *int) {
	c := *a //a指针指向的值赋值给c
	*a = *b //b指针指向的值赋值给a
	*b = c  //临时变量c的值赋值给b
}

type struct1 struct {
	a int
	b string
	c *struct1
}

func passByValue(s struct1) struct1 {
	fmt.Printf("func inner s: %p\n", &s)
	fmt.Printf("func inner s value: %+v\n", s)
	s.c.a = 1
	s.a = 1
	return s
}
