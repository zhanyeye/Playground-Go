package ch19

import (
	"fmt"
	"os"
)

const name = "zhanyeye"

func main() {
	os.Mkdir("tmp", 0666)
	// 「飞雪无情」这个字符串逃逸到了堆上，这是因为「飞雪无情」这个字符串被已经逃逸的指针变量引用，所以它也跟着逃逸了，引用代码如下
	/**
	 ** func (p *pp) printArg(arg interface{}, verb rune) {
	 ** p.arg = arg
	 ** //省略其他无关代码
	 ** }
	 ** 被已经逃逸的指针引用的变量也会发生逃逸。
	 */
	fmt.Println("飞雪无情")

	// 从这一结果可以看到，变量 m 没有逃逸，反而被变量 m 引用的变量 s 逃逸到了堆上。所以被map、slice 和 chan 这三种类型引用的指针一定会发生逃逸的。
	m := map[int]*string{}
	s := "飞雪无情"
	m[0] = &s
}

// 通过 new 函数申请了一块内存；
// 然后把它赋值给了指针变量 s；
// 最后通过 return 关键字返回。
// 在这一命令中，-m 表示打印出逃逸分析信息，-l 表示禁止内联，可以更好地观察逃逸。从以上输出结果可以看到，发生了逃逸，也就是说指针作为函数返回值的时候，一定会发生逃逸
func newString() *string {
	s := new(string)
	*s = "zhanyeye"
	return s
}
