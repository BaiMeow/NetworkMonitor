# NetworkMonitor
监测你的Network状态并绘制图表

monitor your network status with graph

## 预览 preview

[dn11 & vidar network status](https://monitor.dn11.baimeow.cn/)

## 特性 Feature

### 通用功能 Common

- [x] 配置文件热更新 Hot config update
- [x] 可配置的数据刷新间隔
- [ ] 可用性监测
### 获取数据 Fetch
- [x] SSH
- [x] Command
- [x] SFTP (模板式语法)
- [ ] RouterOS API
- [ ] HTTP
- [ ] Telnet
- [ ] OSPF Broadcast

### 数据解析 Parse

- [x] Bird
- [x] MRT
- [ ] RouterOS
- [ ] Quagga
- [ ] OSPF Broadcast

### 图像展示 Graph Display

- [x] RouterID
- [x] AreaID
- [X] Cost
- [x] 自定义额外标识 Custom appended data (metadata)
- [x] 子网 Subnet
- [x] ASN

## 配置 config

`/backend/config/sample.yaml` 为配置文件模板，将其复制为 `config.yaml` 并修改其中的配置即可。

NetworkMonitor 具有灵活的探针配置方案，一个探针由 Fetch 和 Parse 两个部分组成，Fetch 负责获取数据，Parse 负责解析数据，两者可以自由组合，也可以fork之后自定义。

`/backend/config/sample.yaml` is the template of config file, copy it to `config.yaml` and modify the config.

NetworkMonitor has a flexible probe configuration scheme, a probe consists of Fetch and Parse, Fetch is responsible for getting data, Parse is responsible for parsing data, the two can be freely combined, or fork and customize.

## 构建 build

1. 前端打包（如果需要） pack frontend (if needed)
    cd 到 fronted目录下，执行 `pnpm build`，构建好的文件会出现在 `backend/static` 目录下
    cd to fronted directory, run `pnpm build`，the built files will appear in `backend/static` directory

2. 构建后端 build backend
    cd 到 backend 目录下，执行 `go build`    
    cd to backend directory, run `go build`

## 运行 run

将`config.yaml` 和构建好的二进制文件放在同一目录下，执行程序即可。


