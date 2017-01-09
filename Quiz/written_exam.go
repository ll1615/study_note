//************************************************************************************************************************************
// 1.有如下代码
//   msg := []int{1,2,3,4,5,6,7,8,9};
//   sli1 := msg[2:3:4]
//   sli2 := msg[2:3]
//   请说明两个切片有什么不同?

// 两者的长度相同,都为1(3-2).
// sli1显示声明了容量,为首位下标只差:2(4-2)
// sli2容量默认为起始下标至原切片末节点:7(9-2)

package main

import "fmt"

func main() {
	msg := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli1 := msg[2:3:4]
	sli2 := msg[2:3]
	fmt.Println("len(sli1)", len(sli1), "cap(sli1)", cap(sli1))
	fmt.Println("len(sli2)", len(sli2), "cap(sli2)", cap(sli2))
}
// 输出:
// len(sli1)=1, cap(sli1)=2
// len(sli2)=1, cap(sli2)=7

//************************************************************************************************************************************
// 2.方法中,方法的接受者类型,带*与不带*有什么区别?

//   带*:传入的是接收者类型的指针,可使用指针直接对原调用对象进行修改
// 不带*:传入的是接收者类型的实例,即原调用对象的拷贝,无法对原调用对象进行操作
// 如果方法的接收者类型是接口,实现该接口的对象不可调用带*的方法,实现该接口的对象的指针则无限制
package main

import "fmt"

func main() {
	smp := Sample{ID: 2}
	fmt.Println("raw ID:", smp.ID)
	smp.DoSomethingElse()
	fmt.Println("ID after calling function DoSomethingElse:", smp.ID)
	smp.DoSomething()
	fmt.Println("ID after calling DoSomething:", smp.ID)
}

type Sample struct {
	ID int
}

func (s *Sample) DoSomething() {
	fmt.Println("ID in function DoSomething:", s.ID)
	s.ID *= 10
}

func (s Sample) DoSomethingElse() {
	fmt.Println("ID in function DoSomethingElse:", s.ID)
	s.ID *= 10
}
// 输出:
// raw ID: 2
// ID in function DoSomethingElse: 2
// ID after calling function DoSomethingElse: 2
// ID in function DoSomething: 2
// ID after calling DoSomething: 20

//************************************************************************************************************************************
// 3.写一个简单的实例方法,实现面向对象的多态

package main

import "fmt"

func main() {
	if _, err := Divide(2, 0); err != nil {
		fmt.Println(err)
	}
}

type DivisorError string

func (e DivisorError) Error() string {
	return string(e)
}

func Divide(divedend, divisor int32) (int32, error) {
	if divisor == 0 {
		return 0, DivisorError("Error:divisor can't be zero!")
	}
	return divedend / divisor, nil
}
// 输出:
// Error:divisor can't be zero!

//************************************************************************************************************************************
// 4.编写一个函数,实现字符串反转,函数输入类型为string

package main

import "fmt"

func main() {
	fmt.Println(strReverse("hello, gopher!"))
}

func strReverse(input string) string {
	rinput := []rune(input)
	ln := len(rinput)
	hln := ln / 2

	for i := 0; i < hln; i++ {
		rinput[i], rinput[ln-i-1] = rinput[ln-i-1], rinput[i]
	}
	return string(rinput)
}

// 输出:
// !rehpog ,olleh

//************************************************************************************************************************************
// 5.编写一个程序,实现两个go程之间互相通讯

package main

import "fmt"

func main() {
	const maxLength = 10
	ch := make(chan Fact)

	fmt.Println("Factorial:")
	for i := 1; i <= maxLength; i++ {
		go factorial(i, ch)
	}
	for i := 1; i <= maxLength; i++ {
		f := <-ch
		fmt.Printf("%2d!=%-d\n", f.Num, f.Factorial)
	}
}

//Fact 存储阶乘结果
type Fact struct {
	Num       int
	Factorial int64
}

func factorial(n int, rst chan Fact) {
	res := int64(1)
	for i := 1; i <= n; i++ {
		res *= int64(i)
	}

	rst <- Fact{n, res}
}

// 输出(随机):
// Factorial:
//  1!=1
//  2!=2
//  3!=6
//  4!=24
//  6!=720
//  5!=120
//  8!=40320
//  7!=5040
//  9!=362880
// 10!=3628800
