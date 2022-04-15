## day3 前缀树路由Router

### 内容：

- 使用前缀树实现动态路由匹配（dynamic route）
- 实现两种匹配模式`:name`和`*filepath`

### 步骤：

- `Trie`树，主要是两个功能：插入`insert`、查找`search`，这两个功能完成后，前缀树可以作为一个模块抽离出来。
- 将前缀树作为一个模块用于路由`router`之中。
- 动态匹配过程中，通配符动态匹配到的值需要进行存储，给`Context`结构体中加一个属性`Params`用于保存动态匹配到的属性值。

---

## day4-group(2022.4.11)

### 内容

- 实现路由分组控制功能，支持路由分组嵌套
- 路由分组方便后续加入中间件（`day5-middlewares`完成）

### 步骤

- 增加`RouterGroup`类，记录每一个路由分组
- `RouterGroup`同时嵌入`Engine`中，`Engine`作为整个框架顶层的分组，记录所有的路由分组信息。
- 在`RouterGroup`中加入`*Engine`属性，可以由任意一个路由分组到达作为顶层分组的`Engine`中。

### 新增加结构体

```go
type Engine struct {
	*RouterGroup //这里我不确定，是表示Engine本身也作为一个顶层的路由分组吗？
	router       *router
	groups       []*RouterGroup //存储所有的路由分组
}
```
```go
type RouterGroup struct {
	prefix      string        //路由组匹配前缀
	middlewares []HandlerFunc //中间件
	parent      *RouterGroup  //支持嵌套的路由分组
	engine      *Engine       //所有的路由分组都共享一个Engine实例
}
```
---

## day5-middlewares（2022.4.11）

### 内容

- 设计并实现Web框架的中间件机制
- 实现一个超简单的通用`Logger`中间件(也就三两行代码)

### 步骤

- 每一个路由分组`RouterGroup`对象都有一个属性`middlewares`（类型为`[]HandlerFunc`）存储该分组的所有的中间件
- 进行路由匹配的时候，请求路径每匹配到一个路由分组就会把该路由分组的所有中间件加载到一个`slice`中，之后将该`slice`作为一个`Context`类型对象的`handlers`属性值，再把该`Context`
  对象交给全局的`Engine`实例（也就是`main`函数中创建的`r := gee.New()`）执行

### bug

在一个地方写错了，找了几个小时的bug

```go
v1 := r.Group("/v1")
 ```

这里我写成`"v1"`了，furious

---

## day7-panic-recovery(2022.4.11)

### 内容

-实现Gee框架的错误处理机制，避免因为一个panic而整个服务崩溃

### 步骤

-简单处理，就是添加一个全局中间件`recovery`

### 疑问

我在`gee`包中，想实现一个`gee.Default()`
函数返回一个默认的`Engine`实例，就是想其中预先加入部分中间件，那就会引用到`middlewares`包中的函数。 

在`gee`目录下的`go.mod`文件里面引入`middlewares`，同时我也已经在`middlewares`
包里面引入了`gee`。 此时都没报错，但是我在`day7-panic-recover`目录下的`go.mod`引入`gee`包和`middlewares`包就会报错（因为在`main`函数里面会用到`gee.Default()`）

