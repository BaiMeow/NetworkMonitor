<script lang="ts" setup>
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { GraphChart } from "echarts/charts";
import {
    TooltipComponent,
    TitleComponent,
} from "echarts/components";

import { reactive, inject, ref } from "vue";

import VChart from "vue-echarts";
import { ECElementEvent, ECharts } from "echarts";
import { Netmask } from "netmask";

import { getBGP } from "../api/bgp";
import { prettierNet } from "../utils/colornet";
import { ASData } from "../api/meta";

import { ASDataKey } from "../inject/key"

const chart = ref<ECharts | null>(null)

const loading = ref(true);

const asdata = inject(ASDataKey)?.value;

interface Edge {
    source: string;
    target: string;
    value: number
    lineStyle?: any;
    symbol?: string[];
    emphasis?: any;
}

interface Node {
    name: string;
    value: string;
    meta?: any;
    peer_num: number;
    symbolSize?: number;
    network: string[];
    itemStyle?: any;
}

interface Params<T> {
    dataType: string;
    data: T;
}

use([
    CanvasRenderer,
    GraphChart,
    TooltipComponent,
    TitleComponent
]);

function mergeObjects(obj1: any, obj2: any): any {
    for (const key in obj2) {
        if (
            obj2.hasOwnProperty(key)
            && (obj1.hasOwnProperty(key) || !(key in obj1))
        ) {
            if (typeof obj2[key] === 'object' && obj2[key] !== null && typeof obj1[key] === 'object' && obj1[key] !== null) {
                mergeObjects(obj1[key], obj2[key]);
            } else {
                obj1[key] = obj2[key];
            }
        }
    }
}

const option: any = reactive({
    title: {
        text: "DN11 & Vidar Network",
        subtext: "",
    },
    tooltip: {
        trigger: "item",
        triggerOn: "mousemove",
        confine: true,
        enterable: true,
        formatter: (params: Params<any>) => {
            if (params.dataType === "edge") {
                params = params as Params<Edge>
                return `${params.data.source} ↔ ${params.data.target}`
            }

            // dataType === node
            params = params as Params<Node>
            let output = `ASN: ${params.data.name}`

            if (params.data.meta) {
                const metadata: ASData['metadata'][''] = params.data.meta
                if (metadata.display) {
                    output += `<br/>name: ${metadata.display}`
                }
                if (metadata?.monitor?.appendix) {
                    const { monitor: { appendix } } = metadata
                    for (let key in appendix) {
                        const value = appendix[key] as string | string[];
                        if (typeof value === 'string') {
                            output += `<br/>${key}: ${value}`;
                        } else if (Array.isArray(value)) {
                            output += `<br/>${key}:`;
                            for (let i in value) {
                                output += `<br/> - ${value[i]}`;
                            }
                        }
                    }
                }
            }
            output += `<br/> network:<br/>`
            if (asdata) {
                output += prettierNet(params.data.network, params.data.name, asdata.announcements)
            } else {
                output += params.data.network.join('<br/>')
                output += `<br/>`
            }
            output += `Peer Count: <div class="peer_count"> ${params.data.peer_num} </div>`
            return output
        },
        position: function () {
            return [20, 50];
        }
    },
    series: [{
        type: "graph",
        symbolSize: 50,
        layout: "force",
        force: {
            repulsion: 500,
            gravity: 0.02,
            friction: 1,
            edgeLength: [10, 140],
            layoutAnimation: false
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
        itemStyle: {
            borderColor: "#000000",
            borderWidth: 0.4,
            shadowColor: "#2242a3",
        },
        draggable: true,
        data: [],
        links: [],
        emphasis: {
            focus: 'adjacency',
            lineStyle: {
                width: 10
            }
        }
    }],
});

getBGP().then(async (resp) => {
    if (!resp.as) {
        alert("no data")
        return
    }

    const nodes = resp.as.reduce((nodes, cur) => {
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

    nodes.map((node) => {
        node = reactive(node);
        node.peer_num = resp.link.filter((lk) => {
            return lk.src === parseInt(node.name) || lk.dst === parseInt(node.name);
        }).length;
        node.value = '' + node.peer_num;
        node.symbolSize = Math.pow(node.peer_num, 1 / 2) * 7;
        node.itemStyle = {
            shadowBlur: Math.pow(node.peer_num, 1 / 2) * 2,
        }
        if (asdata && asdata.metadata && node.name in asdata?.metadata) {
            const customNode = asdata.metadata[node.name].monitor?.customNode
            if (customNode) {
                mergeObjects(node, asdata.metadata[node.name].monitor?.customNode)
            }
            node.meta = asdata.metadata[node.name];
        }
    })

    const edges = resp.link.reduce((edges, cur) => {
        const src = nodes.find((node) => node.name === cur.src.toString());
        const dst = nodes.find((node) => node.name === cur.dst.toString());
        if (src == null || dst == null) {
            return edges;
        }
        edges.push({
            source: cur.src.toString(),
            target: cur.dst.toString(),
            value: 1 / Math.pow(Math.min(src.peer_num, dst.peer_num), 1 / 2) * 100,
        });
        return edges;
    }, [] as Edge[]);

    option.series[0].data = [] as Node[];
    option.series[0].data.push(...nodes);
    option.series[0].links = edges;
    option.title.subtext = `Nodes: ${nodes.length} Peers: ${edges.length}`;
    loading.value = false;
});

let timer: NodeJS.Timeout | null = null

const handle_mouse_down = (_: ECElementEvent) => {
    if (timer) {
        clearTimeout(timer)
    }
    option.series[0].force.friction = 0.15
    option.series[0].force.layoutAnimation = true
}

const handle_mouse_up = (_: ECElementEvent) => {
    timer = setTimeout(() => {
        option.series[0].force.layoutAnimation = false
    }, 6000);
}

</script>

<template>
    <div v-if="loading" class="graph loading">
        Loading...
    </div>
    <v-chart ref="chart" :option="option" class="graph" autoresize @mousedown="handle_mouse_down"
        @mouseup="handle_mouse_up" />
</template>

<style scoped>
.graph {
    height: 100dvh;
    width: 100vw;
    top: 0;
    left: 0;
    position: absolute;
}

.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 2rem;
    font-weight: bold;
    color: #2242a3;
}
</style>