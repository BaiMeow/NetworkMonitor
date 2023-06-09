<script lang="ts" setup>
import "echarts";
import { reactive } from "vue";

import VChart from "vue-echarts";
import axios from "axios";

interface Resp<T>{
   status_code:number
   status_msg:string
   data:T
}

interface BGP {
    as: AS[]
    link: Link[]
}

interface AS{
    asn:number
    network: string[]
    metadata?:object
}

interface Link {
    src: number;
    dst: number;
}

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

const option: any = reactive({
    title: {
        text: "DN11 & Vidar Network Monitor",
    },
    tooltip: {
        trigger: "item",
        triggerOn: "mousemove",
        formatter: (params: Params<any>) => {
            if (params.dataType === "edge") {
                params = params as Params<Edge>;
                return `${params.data.source} ↔ ${params.data.target}`;
            } else {
                params = params as Params<Node>;
                let output = `ASN: ${params.data.name}`;
                if ("meta" in params.data) {
                    output += "<br/>";
                    for (let key in params.data.meta) {
                        output += `${key}: ${params.data.meta[key]} <br/>`;
                    }
                }
                output += `network: <br/>`
                params.data.network.forEach((net:string) => {
                    output += `${net} <br/>`;
                });
                output += `Peer Count: <div class="peer_count"> ${params.data.peer_num} </div>`;
                return output;
            }
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
                    if (params.data.meta && params.data.meta.name) {
                        return params.data.meta.name;
                    }
                    return params.data.value;
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

axios.get("/api/bgp").then((response) => {
    let resp: Resp<BGP> = response.data;
    const nodes = resp.data.as.reduce((nodes, cur) => {
        nodes.push({
            name: cur.asn.toString(),
            value: cur.asn.toString(),
            meta: cur.metadata ? cur.metadata : {},
            peer_num:0,
            network: cur.network.sort((a, b) => {
                let an = a.split(/[./]/).map((x) => parseInt(x))
                let bn = b.split(/[./]/).map((x) => parseInt(x))
                for (let i = 0; i < an.length; i++) {
                    if (an[i] > bn[i]) {
                        return 1
                    }else if (an[i] < bn[i]) {
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
        node.value = ''+node.peer_num;
        node.symbolSize = Math.pow(node.peer_num, 1 / 2) * 7;
    });

    const edges = resp.data.link.reduce((edges, cur) => {
        edges.push({
            source: cur.src.toString(),
            target: cur.dst.toString(),
        });
        return edges;
    }, [] as Edge[]);

    option.series[0].data = nodes;
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
