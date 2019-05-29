
# Gin-api

一个基于Go-gin框架的API


## 安装
```
$ go get github.com/vsiryxm/gin-api
```

## 运行

### 前置条件

- Mysql
- Redis

### 准备工作

Create a **blog database** and import [SQL](https://github.com/vsiryxm/gin-api/blob/master/docs/sql/blog.sql)

### 项目配置

修改文件 `config/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = db_blog
TablePrefix = tb_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### 运行
```
$ cd $GOPATH/src/gin-api

$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /export/*filepath         --> github.com/vsiryxm/gin-api/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStati
cHandler.func1 (3 handlers)
[GIN-debug] HEAD   /export/*filepath         --> github.com/vsiryxm/gin-api/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStati
cHandler.func1 (3 handlers)
[GIN-debug] GET    /upload/images/*filepath  --> github.com/vsiryxm/gin-api/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStati
cHandler.func1 (3 handlers)
[GIN-debug] HEAD   /upload/images/*filepath  --> github.com/vsiryxm/gin-api/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStati
cHandler.func1 (3 handlers)
[GIN-debug] GET    /qrcode/*filepath         --> github.com/vsiryxm/gin-api/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStati
cHandler.func1 (3 handlers)
[GIN-debug] HEAD   /qrcode/*filepath         --> github.com/vsiryxm/gin-api/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStati
cHandler.func1 (3 handlers)
[GIN-debug] GET    /auth                     --> github.com/vsiryxm/gin-api/controller/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/vsiryxm/gin-api/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (3
handlers)
[GIN-debug] POST   /upload                   --> github.com/vsiryxm/gin-api/controller/api.UploadImage (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/vsiryxm/gin-api/controller/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/vsiryxm/gin-api/controller/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/vsiryxm/gin-api/controller/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/vsiryxm/gin-api/controller/api/v1.DeleteTag (4 handlers)
[GIN-debug] POST   /tags/export              --> github.com/vsiryxm/gin-api/controller/api/v1.ExportTag (3 handlers)
[GIN-debug] POST   /tags/import              --> github.com/vsiryxm/gin-api/controller/api/v1.ImportTag (3 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/vsiryxm/gin-api/controller/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/vsiryxm/gin-api/controller/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/vsiryxm/gin-api/controller/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/vsiryxm/gin-api/controller/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/vsiryxm/gin-api/controller/api/v1.DeleteArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles/poster/generate --> github.com/vsiryxm/gin-api/controller/api/v1.GenerateArticlePoster (4 handlers
)
[info] start http server listening :8000

```

## 特征

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis

