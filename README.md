# iris-gorm-demo
与lib-ui对应的服务端代码，使用iris+gorm+mysql搭建的一个restful项目模板


```
conf  配置文件
controllers  控制器 接受参数 api的入口
datasource 数据库配置
models  结构体模型
repo 数据库的操作
route  注册路由
service 业务逻辑代码
utils  工具类
config.json 配置文件的映射
main.go 主程序入口
```
### 启动项目
```
// 安装依赖
go run main.go
```
导包时使用相对路径需要将项目放在你配置的GOPATH目录下 import ( "iris-gorm-demo/controllers" )
或者使用go mod

#### 前端
https://github.com/pppercyWang/lib-ui
