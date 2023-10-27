# Tour

tour（工具集）是《Go 语言编程之旅：一起用 Go 做项目》中的项目，是第一章 [命令行应用：打造属于自己的工具集] 的代码。

说明：本仓库是自己学习 Go 时，不是官方仓库，官方仓库请移步：[go-programming-tour-book/tour](https://github.com/go-programming-tour-book/tour)


## 安装对应的库（和书上版本可能不一致）

```shell
export GO111MODULE=on 
go get -u github.com/spf13/cobra@v1.1.1
```
## 功能清单

- 单词格式转换
- 便利的时间工具
- SQL语句到结构体的转换 
- JSON 到结构体的转换

## 单词格式转换 word

### 安装 cobra 

```
go get -u github.com/spf13/cobra@v1.1.1
```

### 结构
- cmd
  - root.go
  - word.go
- internal
- pkg

### 测试命令

```
go run main.go word -s=MiracleWong -m=1
go run main.go word -s=MiracleWong -m=2
go run main.go word -s=miracle_wong -m=3
go run main.go word -s=miracle_wong -m=4
go run main.go word -s=MiracleWong -m=5
```

## 时间计算 time

### 测试命令

```
go run main.go time now
go run main.go time calc -c="2029-09-04 15:04:02" -d=5m
go run main.go time calc -c="2029-09-04 12:02:33" -d=-12h
```


## SQL转换和处理 sql
```
# 启动 mysql（前提是 Mac 上有 MySQL）

go run main.go sql struct --username root --password 123456 --db nba --table player
go run main.go sql struct --username root --password 123456 --db nba --table player_score
```

