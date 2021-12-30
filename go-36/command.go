package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")

	// 自定义命令源码文件的参数使用说明
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

//https://golang.google.cn/pkg/flag/
func main() {
	// 我们在调用flag包中的一些函数（比如StringVar、Parse等等）的时候，实际上是在调用flag.CommandLine变量的对应方法。
	// 自定义命令源码文件的参数使用说明
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s: \n", "this script")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()

	// 这样做的好处依然是更灵活地定制命令参数容器。但更重要的是，你的定制完全不会影响到那个全局变量flag.CommandLine。
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
