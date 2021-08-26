# mvc
版本v0.3.4及以上，持续维护

## 概要
微服务搭建

## 项目结构

```json
demo-account
  |- api
    |- user
      |- login.go
    |- register.go
      |- func Register() ... API注册
  |- ...其它自行扩展 如: common等
  |- main.go // 启动入口
    |- func main() ... 启动服务
    |- func init() ... 初始化服务中间件
```

## 搭建流程
项目基本项目结构，以下流程以demo-account为例进行说明。

### API创建
demo-account/api/user/login.go
```go
package user

import (
	"fmt"

	"github.com/xm-chentl/go-dbfty"
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/actionresult"
	"github.com/xm-chentl/go-mvc/scope"
)

// Login 登录接口
//【固定规范】属性注入 todo: 标签，后续扩展
type Login struct {
	// 中间件注入
	RouteCtx mvc.IRoute
	UserDb   dbfty.IFactory `key:"userdb"`
	// 参数
	Account  string
	Password string
}

// Code 【固定规范】接口编号
func (a Login) Code() string {
	return "1001"
}

// Scope 【固定规范】作用域枚举
func (a Login) Scope() scope.Value {
	return scope.Server
}

// Execute 执行
func (a *Login) Execute() mvc.IActionResult {
	if a.UserDb != nil {
		db := a.UserDb.Db()
		fmt.Println("组件userDb => ", db)
	}
	if a.RouteCtx != nil {
		fmt.Println("routeCtx => ", a.RouteCtx)
	}
	fmt.Println("account => ", a.Account, "password => ", a.Password)
  // 获取参数值
	return actionresult.JSON(fmt.Sprintf("account: %s, password: %s", a.Account, a.Password))
}


```
### API注册
demo-account/api/register.go
```go
package api

import (
	"github.com/xm-chentl/go-mvc/metadata"

	"demo-account/api/user"
)

// Register 注册
func Register() {
	metadata.Register(
		&user.Login{},
	)
}

```
### 服务启动入口

demo-account/main.go
```go
package main

import (
	"demo-account/api"
	"demo-account/conf"

	dbftymock "github.com/xm-chentl/go-dbfty/mock"
	"github.com/xm-chentl/go-mvc/container"
	"github.com/xm-chentl/go-mvc/ginex"
	"github.com/xm-chentl/go-mvc/handler"
)

func main() {
	// 启动 服务
	ginex.New().AddHandler(
		handler.Default(),
	).Run(conf.Get().Port)
}

func init() {
	// api 注册
	api.Register()
	// 配置初始化
	conf.Init()
	// 内置是间件容器, 此处为注入
	container.Set("userdb", dbftymock.New())
}

```