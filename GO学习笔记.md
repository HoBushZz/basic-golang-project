# GoLang学习笔记

## JSON

### JSON是什么

JSON（JavaScript对象表示法 JavaScript Object Notation），是存储和交换文本信息的语法，轻量级的文本数据交换格式，类似XML，但是比XML更小、更快、更容易解析。
JSON==独立于各种编程语言平台==，并且几乎所有的主流编程语言都内置了对JSON数据格式的支持。

### JSON语法规则

JSON语法是JavaScript对象表示法语法的子集：

- 数据以**名值对**的形式表示
- 数据之间以逗号分隔
- 花括号保存对象
- 方括号保存数组

### JSON值

JSON名值对包括字段名称（包含在双引号中），后面一个冒号，然后是值，如：

```javascript
"username": "Jessica"
```

JSON值可以包含以下多种类型：

- 数字（整数和浮点数）
- 字符串（在双引号中）
- 逻辑值（true或false）
- 数组（在方括号中）
- 对象（在花括号中）
- null

### JSON对象和JSON数组

JSON对象包含在花括号中，可以包含多个*名/值*对，其值可以为数组；JSON数组包含在方括号中，可以包含多个对象。**JSON既可以以花括号开头，也可以以方括号开头**，JSON对象和JSON数组可以互相嵌套，如：

```json
{
    "users": [
        {"username":"Jacky",age:30},
        {"username":"Jessica",age:23},
        {"username":"James",age:42}
    ]
}
```

### JSON文件

- json文件的后缀为：”.json”
- JSON文件的MIME类型为：“application/json”

### JSON序列化和反序列化

早期的JSON解析器基本上就是使用JavaScript的`eval()`函数。由于JSON是JavaScript语法的子集，因此`eval()`函数可以解析并返回JavaScript对象和数组。
ECMAScript 5对解析JSON的行为进行了规范，定义了全局对象JSON。

## Shell脚本和编程

### 变量和环境变量

Shell 变量的作用域可以分为三种：
有的变量只能在函数内部使用，这叫做局部变量（local variable）；
有的变量可以在当前 Shell 进程中使用，这叫做全局变量（global variable）；
而有的变量还可以在子进程中使用，这叫做环境变量（environment variable）。

## GoLang

### 可见性

* 包变量的可见性

规定首字母大写是包public

首字母小写是包Private

* 结构体变量变量的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

### 结构体标签

```go
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}
```

### channel

关闭后的通道有以下特点：

1. 对一个关闭的通道再发送值就会导致 panic。
2. 对一个关闭的通道进行接收会一直获取值直到通道为空。
3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4. 关闭一个已经关闭的通道会导致 panic。

那我们如何判断一个通道是否被关闭了呢？

对一个通道执行接收操作时支持使用如下多返回值模式。

```go
value, ok := <- ch
```

其中：

- value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
- ok：通道ch关闭时返回 false，否则返回 true。

## Go依赖管理

### go mod replace

[参考博客](https://blog.csdn.net/qq_24433609/article/details/127323097)

几个典型使用场景：

* 替换无法下载的包（使用代理），但是在import不能修改，因为代理知识镜像仓库，go.mod中定义的module还是原先的包
* 调试依赖包
* 使用fork仓库
* 禁止被依赖
* 引入本地包

### go.work（工作区）

[参考博客](https://segmentfault.com/a/1190000041681242)

[参考博客2](https://juejin.cn/post/7145855715565895710)

* Go 1.18引入的工作区模式，可以让你不用修改Go Module的go.mod，就能同时对多个有依赖的Go Module进行本地开发。
* 一个module一旦放在工作空间下，那么该module的go.mod寻找必须通过工作空间的go.work去实现，而不能直接使用`go build`建立，所以要在工作空间下先`go work use xxx`把，`./xxx`添加到go.work中的use中，这样你才可以在该module项目中引用module项目中的包。

# 前端学习笔记

## URL和URI

URL（Uniform Resource Locator）和 URI（Uniform Resource Identifier）都是用于标识和定位资源的字符串。

URI 是一个更通用的概念，它用于标识任何可以通过网络访问的资源。而 URL 则是 URI 的一种特定形式，它包含了用于定位资源的详细信息。

具体来说：

* URI：统一资源标识符，是用于唯一标识一个资源的字符串。它可以是一个网址、文件路径、协议标识符等。URI 由两个主要部分组成：标识符和位置标识符。标识符指定了资源的身份，位置标识符指定了资源的位置。URI 的例子包括：

* URL：统一资源定位符，标识了资源的位置和访问方式。例如，https://www.example.com 是一个 URL。
  URN：统一资源名称，仅标识了资源的名称而不指定其位置。例如，urn:isbn:1234567890 是一个 URN。
  URL：统一资源定位符，是一种特定的 URI，用于定位网络上的资源。URL 包含了资源的访问协议（如 HTTP、HTTPS）、主机名（如 www.example.com）和可选的路径、查询参数、片段标识等。URL 的例子包括：

  * http://www.example.com/index.html

  * https://www.example.com/search?q=example

  * ftp://ftp.example.com/files/file.txt

    总结起来，URI 是一个用于唯一标识资源的字符串，而 URL 则是一种特定形式的 URI，用于定位网络上的资源。URL 包含了详细的位置信息，而 URI 可以是其他形式的标识符，不仅限于网络资源的访问。

## \<label>和\<input>

将一个 <label> 和一个 <input> 元素匹配在一起，你需要给 <input> 一个 id 属性。而 <label> 需要一个 for 属性，其值和 <input> 的 id 一样。

另外，你也可以将 <input> 直接放在 <label> 里，此时则不需要 for 和 id 属性，因为关联已隐含存在

[参考资料](https://developer.mozilla.org/zh-CN/docs/Web/HTML/Element/label)

## 