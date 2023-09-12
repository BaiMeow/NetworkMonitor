<script lang="ts" setup>
import "echarts";
import { reactive } from "vue";

import VChart from "vue-echarts";
import { Netmask } from "netmask";

import { getASMetaData, getBGP, ASMetaData } from "../api/graph";

interface Edge {
    source: string;
    target: string;
    lineStyle?: any;
    symbol?: string[];
}

interface Node {
    name: string;
    value: string;
    meta?: any;
    peer_num: number;
    symbolSize?: number;
    network: string[];
}

interface Params<T> {
    dataType: string;
    data: T;
}

function mergeObjects(obj1: any, obj2: any): any {
  for (const key in obj2) {
    if (
        obj2.hasOwnProperty(key)
        && (obj1.hasOwnProperty(key)|| !(key in obj1))
    ) {
      if (typeof obj2[key] === 'object' && obj2[key] !== null && typeof obj1[key] === 'object' && obj1[key] !== null) {
        obj1[key] = mergeObjects([key], obj2[key]);
      } else {
        obj1[key] = obj2[key];
      }
    }
  }
}

const option: any = reactive({
    title: {
        text: "DN11 & Vidar Network Monitor",
    },
    tooltip: {
        trigger: "item",
        triggerOn: "mousemove",
        formatter: (params: Params<any>) => {
            if (params.dataType === "edge") {
                params = params as Params<Edge>
                return `${params.data.source} ↔ ${params.data.target}`
            }

            // dataType === node
            console.log(params)
            if (!params.data.meta) {
                return
            }
    
            params = params as Params<Node>

            const metadata: ASMetaData = params.data.meta

            if (metadata.html) {
                return metadata.html
            }

            let output = `ASN: ${params.data.name}`
            if (metadata.display) {
                output += `<br/>name:${metadata.display}`
            }
            if (metadata.appendix) {
                for (let key in metadata.appendix) {
                    output += `<br/>${key}: ${metadata.appendix[key]}`
                }
            }
            output += `<br/> network:`
            params.data.network.forEach((net: string) => {
                output += `<br/>${net}`
            });
            output += `<br/>Peer Count: <div class="peer_count"> ${params.data.peer_num} </div>`
            return output
        },
    },
    roam: "scale",
    symbolSize: 50,
    animationDurationUpdate: 1500,
    animationEasingUpdate: "quinticInOut" as any,
    series: [
        {
            type: "graph",
            layout: "force",
            animation: true,
            force: {
                repulsion: 500,
                gravity: 0.02,
                friction: 0.15,
                edgeLength: 80,
            },
            roam: true,
            label: {
                show: true,
                position: "right",
                formatter: (params: any) => {
                    if (params.data.meta && params.data.meta.display) {
                        return params.data.meta.display;
                    }
                    return params.data.name;
                },
            },
            draggable: true,
            data: [],
            links: [],
        },
    ],
    lineStyle: {
        opacity: 0.9,
        width: 2,
    },
});

getBGP().then((resp) => {
    if (!resp.data.as) {
        alert("no data")
        return
    }

    const nodes = resp.data.as.reduce((nodes, cur) => {
        nodes.push({
            name: cur.asn.toString(),
            value: cur.asn.toString(),
            peer_num: 0,
            network: cur.network.sort((a, b) =>
                parseInt(a.split("/")[1]) - parseInt(b.split("/")[1])
            ).reduce((network, cur) =>
                network.findIndex((net) => {
                    let nmask = new Netmask(net);
                    return nmask.contains(cur) || nmask.toString() === cur;
                }) === -1 ?
                    [...network, cur] : network
                , [] as string[]
            ).sort((a, b) => {
                let an = a.split(/[./]/).map((x) => parseInt(x))
                let bn = b.split(/[./]/).map((x) => parseInt(x))
                for (let i = 0; i < an.length; i++) {
                    if (an[i] > bn[i]) {
                        return 1
                    } else if (an[i] < bn[i]) {
                        return -1
                    }
                }
                return -1
            })
        })
        return nodes;
    }, [] as Node[]);

    nodes.forEach((node) => {
        node.peer_num = resp.data.link.filter((lk) => {
            return lk.src === parseInt(node.name) || lk.dst === parseInt(node.name);
        }).length;
        node.value = '' + node.peer_num;
        node.symbolSize = Math.pow(node.peer_num, 1 / 2) * 7;
        node = reactive(node);
        getASMetaData(parseInt(node.name)).catch((e) => {
            if (e.response.status !== 404) {
                console.log(e)
            }
        }).then((resp) => {
            if (resp === undefined){
                return
            }
            if (resp.customNode){
                mergeObjects(node, resp.customNode)
            }
            node.meta = resp;
        });
    });

    const edges = resp.data.link.reduce((edges, cur) => {
        edges.push({
            source: cur.src.toString(),
            target: cur.dst.toString(),
        });
        return edges;
    }, [] as Edge[]);

    option.series[0].data = [] as Node[];
    option.series[0].data.push(...nodes);
    option.series[0].links = edges;
});
</script>

<template>
    <v-chart :option="option" class="graph" />
</template>
<style scoped>
.graph {
    width: 100vw;
    height: 100%;
    top: 0;
    left: 0;
    position: absolute;
}
</style>