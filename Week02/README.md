# 学习笔记

## Week02 作业题目：

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

假设我们有一个API接口 `/blogs?user_id=123` 可以根据所提供的userID获取相关的博客文章。我们的当API request到达API server的时候，
分别经过这三层: 
1. API service handler 负责接受request，调用storage layer来读取数据库，得到的结果写回response
2. storage layer， 负责拆封数据库请求调用DAO层，拿到DAO层的数据之后负责组装业务数据。
3. DAO layer， 负责数据库读写。

DAO层抛出error之后，storage layer直接将error返回到API service handler层，即storage layer不需要知道DAO层发生了什么，只需要知道有错误
发生，把错误处理委任给service handler。

service handler 调用storage layer 返回错误之后，要打印日志
```go
		log.Printf("%+v\n", errors.Wrapf(err, "failed to get blogs for user id %d", userID))
```
日志中附上输入的用户ID，并保存堆栈信息方便debug。

然后返回404 http错误码方便客户端做fallback处理.

```go
		http.Error(w, "blogs not found", http.StatusNotFound)
```

## 允许示例代码
```go
make run
```
会在 localhost:8090 启动http server

在浏览器中输入
``` go
http://localhost:8090/blogs?user_id=123
```

之后你会在终端看到如下错误日志 并在浏览器端收到404
```go
2020/12/02 00:30:36 sql: no rows in result set
failed to get blogs for user id 123
github.com/qihach/go-week-02/handlers.GetBlogs
        /Users/qihang/github/Go-000/Week02/handlers/handlers.go:30
net/http.HandlerFunc.ServeHTTP
        /usr/local/go/src/net/http/server.go:2041
net/http.(*ServeMux).ServeHTTP
        /usr/local/go/src/net/http/server.go:2416
net/http.serverHandler.ServeHTTP
        /usr/local/go/src/net/http/server.go:2836
net/http.(*conn).serve
        /usr/local/go/src/net/http/server.go:1924
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1373
```
