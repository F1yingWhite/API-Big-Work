# 抖音推荐系统

## git仓库

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
2. 进入titoke-web目录下运行`npm install`安装前端依赖
3. 运行`go run main.go --dev`同时启动前后端服务,此时会生成.`config.yaml`文件，在里面可以配置postgresql的dsn，日志文件位置等，默认数据库是sqlite，所以不配置也可以运行
4. (可选)修改`config.yaml`文件中的数据库配置后可以使用go run main.go --postgresql --dev指定postgresql当数据库
5. (可选)运行`npm run dev`启动前端服务
