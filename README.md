# mvc
版本v0.2.0以上

## 概要
方便搭建微服务

## 项目结构

```json
demo-account
  |- api
    |- wechat
      |- wechat.go
      |- ...以下为请球参数结构的规范（仅作为参考）
      |- request-xxx.go
    |- ...其它接口
    register.go //注册api
  |- ...其它自行扩展 如: common等
  |- mvc
    |- mvc.go // loveyd-lib-go/mvc 流程的封装
  |- main.go // 启动入口
    |- 注册api
    |- demo-account/mvc.Run(端口)
```

## 搭建流程 v0.1.0、v0.2.0 有效
项目基本的创建流程，以下流程以demo-account为例进行说明。

### API创建
demo-account/wechat/request-login.go
```go
package wechat

type requestLogin struct{
    Code string
}

```
demo-account/wechat/wechat.go
```go
package wechat

import (
	"loveyd-lib-go/mvc"
	actionresult "loveyd-lib-go/mvc/action-result"
)

type wechat struct{}

func (w wechat) Login(req *requestLogin) mvc.IActionResult {
    actionresult.JSON(nil)
}

```
### API注册
demo-account/api/register.go
```go
package api

import (
	"demo-account/api/wechat"
	"loveyd-lib-go/mvc/metadata"
)

func Register() {
    // 默认方式
    metadata.Register(wechat.wechat{})
    // 自定义方式
    metadata.RegisterByCustom(map[string]interface{}{
        "自定义服务名": wechat.wechat{}
    })
}

```
### 配置启动处理流程
demo-account/mvc/mvc.go
```go
package mvc

import (
	"demo-account/api"

	"loveyd-lib-go/mvc"
	"loveyd-lib-go/mvc/gin"
	"loveyd-lib-go/mvc/handler"
)

// Run 启动
// @param port 端口
func Run(port int) {
    // api请求方式 => /service/action
    handle := new(handler.Handle)   // 默认处理入口
	handle.Next(
		new(handler.Service),   // service处理节点
	).Next(
		new(handler.Action),    // action处理节点
	).Next(
		new(handler.Verify),    // verify 参数处理节点
	).Next(
		new(handler.Invoke),    // invoke 执行处理节点
    )
    
	mvc.Default().SetHandle(handle).Run(port)
}

func init() {
    // 注册api
    api.Register()
    // 初始化mvc的初始化框架
    mvc.SetDefault(gin.New())
}
```
### 启动服务

demo-account/main.go
```go
package main

import(
    "demo-account/mvc"
)

func main() {
    mvc.Run(8080)
}
```


## v0.3.0以上