# 抖音推荐系统

## 目录结构

- movies文件夹下存放了所有的视频mp4文件
- titok-web存放了前端代码
- notes文件下存放了项目文档
- config.yaml是配置文件,里面只需要该数据库dns即可,注意该文件在第一次运行代码时自动生成
- log.txt是默认的日志文件位置,记录了后端的日志,可以在config.yaml中修改日志文件位置
- 其他都是后端的代码

## 部署要求

环境需求:需要安装golang和postgresql和nodejs,并且把视频的mp4解压到movies文件夹下,首次运行自动创建数据表和根据mp4文件创建数据.

1. 运行`go mod tidy`下载go依赖
2. 运行`go run main.go`启动后端服务,此时会生成.`config.yaml`文件,不出意外会爆错退出,如果没退出手动退一下
3. 修改`config.yaml`文件中的数据库配置
4. 运行`go run main.go`再次启动后端服务,此时后端成功启动
5. 进入titoke-web下运行`npm install`安装前端依赖
6. 运行`npm run dev`启动前端服务
