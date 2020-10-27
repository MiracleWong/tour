# Tour

## 单词格式转换 word

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