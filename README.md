# JJBot
专供机姬BOT开发的程序

## 1.目录结构
```
├─service
│  ├─bot 机器人服务
│  │  ├─botapi 机器人核心
│  │  ├─live BiliBili直播服务
│  │  └─ys 原神服务
│  ├─conf 配置文件
│  ├─cron 定时任务服务
│  ├─db 数据库服务
│  │  ├─mongodb
│  │  ├─mysql
│  │  └─redis
│  ├─logger 日志服务
│  └─web web服务
└─task 定时任务
  main.go 主入口
```

## 2.开发基本要求
- golang >= 1.18
- 基础的golang知识，IDE（如[GoLand](https://www.jetbrains.com/go/)）
- Git基础
- 阅读代码能力：
- - 定时任务写法
- - 数据库使用
- - 编写配置文件
- 使用搜索引擎查找问题的能力，机器人核心出现异常及时与开发者取得联系

## 3.使用需注意
- 注意打包的平台，如windows/linux x64/x86
- 打包不同平台可使用goreleaser，项目内已包含配置文件
- 启动前建立conf文件夹、cookie文件和config配置文件

## 4.（可选）goreleaser打包方式
- 终端中使用指令下载 `go install github.com/goreleaser/goreleaser@latest`
- 打包 `goreleaser release --snapshot --rm-dist`
- 在dist目录下使用对应平台的文件即可
