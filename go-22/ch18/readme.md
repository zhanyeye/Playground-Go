## 质量保证: Go 语言如何通过测试保证质量

#### 单元测试
单元测试
```shell
go test -v
```
单元测试覆盖率
```shell
go test -v --coverprofile=ch18.cover
```
生成HTML格式的单元测试覆盖率报告
```
go tool cover -html='ch18.cover' -o='ch18.html'
```
#### 基准测试
基准测试，测试函数性能go test -bench=. -benchtime=3s
```shell
go test -bench='.'
```
基准测试指定执行时间
```shell
go test -bench='.' -benchtime=3s
```
基准测试进行内存统计
```shell
go test -bench='.' -benchmem
```
