## 性能优化: Go 语言如何进行代码检查和优化

#### 代码规范检查

golangci-lint 是一个集成工具，它集成了很多静态代码分析工具。通过配置这一工具，可以很灵活地启用需要的代码规范检查。
```shell
go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
```
检测是否安装成功
```shell
golangci-lint version
```
演示 golangci-lint 的使用
```shell
golangci-lint run ch19/
```
golangci-lint 配置
golangci-lint 的配置比较灵活，比如你可以自定义要启用哪些 linter。golangci-lint 默认启用的 linter，包括这些:
```
deadcode - 死代码检查
errcheck - 返回错误是否使用检查
gosimple - 检查代码是否可以简化
govet - 代码可疑检查, 比如格式化字符串和类型不一致
ineffassign - 检查是否有未使用的代码
staticcheck - 静态分析检查
structcheck - 查找未使用的结构体字段
typecheck - 类型检查
unused - 未使用代码检查
varcheck - 未使用的全局变量和常量检查
```
golangci-lint linters 命令查看每个 linter 的说明

如果要修改默认 linter，就需要对 golangci-lint 进行配置。即在项目根目录下新建一个名字为 .golangci.yml 的文件，这就是 golangci-lint 的配置文件。假设我只启用 unused 检查，可以这样配置：
.golangci.yml
```shell
linters:
  disable-all: true
  enable:
    - unused
```

#### 性能优化
Go 语言有两部分内存空间：栈内存和堆内存。
+ 栈内存由编译器自动分配和释放，开发者无法控制。栈内存一般存储函数中的局部变量、参数等，函数创建的时候，这些内存会被自动创建；函数返回的时候，这些内存会被自动释放。
+ 堆内存的生命周期比栈内存要长，如果函数返回的值还会在其他地方使用，那么这个值就会被编译器自动分配到堆上。堆内存相比栈内存来说，不能自动被编译器释放，只能通过垃圾回收器才能释放，所以栈内存效率会很高。

逃逸分析  
+ 既然栈内存的效率更高，肯定是优先使用栈内存。那么 Go 语言是如何判断一个变量应该分配到堆上还是栈上的呢？这就需要逃逸分析了
```shell
go build -gcflags="-m -l" ./ch19/main.go 
```
+ `-m` 表示打印出逃逸分析信息
+ `-l` 表示禁止内联，可以更好地观察逃逸

逃逸到堆内存的变量不能马上被回收，只能通过垃圾回收标记清除，增加了垃圾回收的压力，所以要尽可能地避免逃逸，让变量分配在栈内存上，这样函数返回时就可以回收资源，提升效率。

优化技巧  
通过前面小节的介绍，相信你已经了解了栈内存和堆内存，以及变量什么时候会逃逸，那么在优化的时候思路就比较清晰了，因为都是基于以上原理进行的。下面我总结几个优化的小技巧：
+ 第 1 个需要介绍的技巧是尽可能避免逃逸，因为栈内存效率更高，还不用 GC。比如小对象的传参，array 要比 slice 效果好。
+ 第 2 个技巧,如果避免不了逃逸，还是在堆上分配了内存，那么对于频繁的内存申请操作，我们要学会重用内存，比如使用 sync.Pool。
+ 第 3 个技巧就是选用合适的算法，达到高性能的目的，比如空间换时间。