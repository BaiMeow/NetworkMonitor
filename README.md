# OSPF-monitor
监测你的OSPF状态并绘制图表

monitor your OSPF status with graph

## 预览
**preview**

[DN11 OSPF Status](https://monitor.dn11.baimeow.cn/)

## 配置
**config**

`/backend/config/sample.yaml` 为配置文件模板，将其复制为 `config.yaml` 并修改其中的配置即可。

ospf-monitor 具有灵活的探针配置方案，一个OSPF探针由 Fetch 和 Parse 两个部分组成，Fetch 负责获取数据，Parse 负责解析数据，两者可以自由组合，也可以fork之后自定义。

`/backend/config/sample.yaml` is the template of config file, copy it to `config.yaml` and modify the config.

ospf-monitor has a flexible probe configuration scheme, an OSPF probe consists of Fetch and Parse, Fetch is responsible for getting data, Parse is responsible for parsing data, the two can be freely combined, or fork and customize.

