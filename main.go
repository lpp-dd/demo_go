package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
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

}

/*
*int 表示int类型的指针
 */
func swap(a *int, b *int) {
	c := *a //a指针指向的值赋值给c
	*a = *b //b指针指向的值赋值给a
	*b = c  //临时变量c的值赋值给b
}
