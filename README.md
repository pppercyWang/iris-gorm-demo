# iris-gorm-demo
与lib-ui对应的服务端代码，使用iris+gorm+mysql搭建的一个restful项目模板

#### 12.26更新
1. 优化代码结构
2. jwt实现
3. json传参
4. md5处理
```
conf  配置文件
controllers  控制器 入参处理 api的入口
datasource 数据库配置 
models  结构体
db  sql数据文件 postman接口文件
repo 数据库的操作
middleware 中间件 jwt实现
route  注册路由
service 业务逻辑代码
utils  工具类
config.json 配置文件的映射
main.go 主程序入口
```
### 启动项目
```
1.安装依赖 go get
2.go run main.go
```
1. 使用go get直接下载依赖，或在github手动下载包放到gopath/src/github.com/
2. 导包时使用相对路径需要将项目放在你配置的GOPATH目录下
3. 这里我没有使用go module来管理依赖，是因为下载包很容易被墙。你可以使用go mod init用go mod来管理 

#### 前端
https://github.com/pppercyWang/lib-ui
