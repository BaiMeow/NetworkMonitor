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
- [x] RouterOS API
- [x] HTTP
- [ ] Telnet
- [ ] OSPF Broadcast

### 数据解析 Parse

- [x] Bird
- [x] MRT
- [x] RouterOS
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

`monitor-metadata.json` 需要放在运行目录下，可以为图表提供一些别名和样式，以及 BGP 宣告信息，这个信息大部分都用于 BGP 绘图，可以参照 [演示站 metadata](https://metadata.dn11.baimeow.cn/monitor-metadata.json) 编写。

如果你只有一个 OSPF 网络，没有 BGP 和对应的 ASN，可以随便取一个 uint16 作为 ASN，保持 config 和 metadata 文件中的 ASN 一致即可。

OSPF 图暂不支持自定义节点样式，metadata 中只有 display 字段会生效，所以 OSPF only 的可以这样写。

```json
{
  "metadata": {
    "4211110000": {
      "display": "AS NAME"
    }
  }
}
```

`/backend/config/sample.yaml` is the template of config file, copy it to `config.yaml` and modify the config.

NetworkMonitor has a flexible probe configuration scheme, a probe consists of Fetch and Parse, Fetch is responsible for getting data, Parse is responsible for parsing data, the two can be freely combined, or fork and customize.

`monitor-metadata.json` should be put at workdir, and providing some alias and styles, also BGP announces. Most of them are used in BGP graph. You can write metadata file with reference of [demo metadata](https://metadata.dn11.baimeow.cn/monitor-metadata.json)。

if you only have single OSPF network, no BGP and its ASN, please choose a rand uint16 as ASN, keepping ASN in config and metadata file same.

OSPF graph don't support custom node style, so only display field  applied in metadata file, you can write metadata file as following.

```json
{
  "metadata": {
    "4211110000": {
      "display": "AS NAME"
    }
  }
}
```

## 构建 build

1. 前端打包（如果需要） pack frontend (if needed)

    cd 到 fronted目录下，执行 `pnpm build`，构建好的文件会出现在 `backend/static` 目录下
   
    cd to fronted directory, run `pnpm build`，the built files will appear in `backend/static` directory

3. 构建后端 build backend

    cd 到 backend 目录下，执行 `go build`

    cd to backend directory, run `go build`

## 运行 run

将`config.yaml` 和构建好的二进制文件放在同一目录下，执行程序即可

put `config.yaml` and the built binary in the same dir and run it

## 自有化部署 self-host

配置中没有提供 title icon 一类的定制方式，需要自行修改前端代码，重新编译。

influxdb uptime 功能还在试验中，不建议开启。

custom title and icon is not supported in config, you should edit fronted code by yourself and compile them.

influxdb uptime is under experiment, not recommanded.
