<script setup lang="ts">
import "echarts";
import { reactive } from "vue";

import VChart from "vue-echarts";
import axios from "axios";

const option = reactive({
    title: {
        text: 'DN11 OSPF Status',
    },
    tooltip: {
        trigger: 'item',
        triggerOn: 'mousemove',
        formatter: '{c}',
    },
    roam: 'scale',
    symbolSize: 50,
    animationDurationUpdate: 1500,
    animationEasingUpdate: 'quinticInOut' as any,
    series: [
        {
            type: 'graph',
            layout: 'force',
            animation: true,
            force: {
                initLayout: 'circular',
                repulsion: 1000,
                gravity: 0,
            },
            roam: true,
            label: {
                show: true,
                position: 'right',
            },
            draggable: true,
            edgeLabel: {
                show: true,
                formatter: '{c}',
            },
            data: [],
            links: [],
        }
    ],
    lineStyle: {
        opacity: 0.9,
        width: 2,
    }
})

axios.get('/api/graph').then(response => {
    let data: Array<any> = response.data;
    let { nodes, edges } = data.reduce(({ nodes, edges }, cur) => {
        if (cur['links'] as Array<any> && cur['links'].length !== 0) {
            cur['links'].forEach((link: any) => {
                edges.push({
                    source: link['src'],
                    target: link['dst'],
                    value: link['cost'],
                });
            });
        }
        if (cur['router'] as Array<any> && cur['router'].length !== 0) {
            cur['router'].forEach((router: any) => {
                if (nodes.findIndex((e: any) => e.name === router['router_id']) != -1) return;
                nodes.push({
                    name: router['router_id'],
                    value: router['router_id'],
                });
            });
        }
        return { nodes, edges };
    }, { nodes: [], edges: [] });

    option.series[0].data = nodes;
    option.series[0].links = edges;
});

//let backend = "[{\"area_id\":\"0.0.0.0\",\"router\":[{\"router_id\":\"172.16.255.2\"},{\"router_id\":\"172.16.255.4\"},{\"router_id\":\"172.16.255.3\"},{\"router_id\":\"172.16.255.7\"},{\"router_id\":\"172.16.255.5\"}],\"links\":[{\"src\":\"172.16.255.2\",\"dst\":\"172.16.255.4\",\"cost\":5},{\"src\":\"172.16.255.2\",\"dst\":\"172.16.255.3\",\"cost\":5},{\"src\":\"172.16.255.3\",\"dst\":\"172.16.255.4\",\"cost\":5},{\"src\":\"172.16.255.3\",\"dst\":\"172.16.255.7\",\"cost\":5},{\"src\":\"172.16.255.4\",\"dst\":\"172.16.255.5\",\"cost\":5},{\"src\":\"172.16.255.4\",\"dst\":\"172.16.255.7\",\"cost\":5}]}]"

</script>

<template>
    <v-chart class="graph" :option="option" />
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