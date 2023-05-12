<script setup lang="ts">
import "echarts";
import { reactive } from "vue";

import VChart from "vue-echarts";
import axios from "axios";

interface Router {
  router_id: string;
  metadata?: Object;
}

interface Link {
  src: string;
  dst: string;
  cost: number;
}

interface Area {
  area_id: string;
  router: Router[];
  links: Link[];
}

interface Edge {
  source: string
  target: string
  value: number
  lineStyle?: any
  symbol?: string[]
}

interface Node {
  name: string
  value: string
  meta?: any
  area: string[]
}

const option: any = reactive({
  title: {
    text: 'DN11 OSPF Status',
  },
  tooltip: {
    trigger: 'item',
    triggerOn: 'mousemove',
    formatter: (params: any) => {
      if (params.dataType === 'edge') {
        params.data as Link
        return `link: ${params.data.source} <==> ${params.data.target} <br/> cost: ${params.data.value}`;
      } else {
        params.data as Edge
        let output = `RouterId: ${params.data.value}`;
        if (params.data.meta) {
          output += '<br/>';
          for (let key in params.data.meta) {
            output += `${key}: ${params.data.meta[key]} <br/>`;
          }
          output += `area: ${params.data.area.join()}`
        }
        return output
      }
    },
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
        repulsion: 1200,
        gravity: 0.1,
      },
      roam: true,
      label: {
        show: true,
        position: 'right',
        formatter: (params: any) => {
          if (params.data.meta && params.data.meta.name) {
            return params.data.meta.name;
          }
          return params.data.value;
        },
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



axios.get('https://monitor.dn11.baimeow.cn/api/graph').then(response => {
  let area: Array<Area> = response.data;
  let { nodes, edges } = area.reduce(({ nodes, edges, }, cur) => {
    if (cur.links && cur.links.length !== 0) {
      cur.links.forEach((link: Link) => {
        edges.push((() => {
          let eg: Edge = {
            source: '' + link.src,
            target: link.dst,
            value: link.cost,
          }
          if (cur.links.findIndex((l) => {
            return l.src === link.dst && l.dst === link.src && l.cost !== link.cost
          }) !== -1) {
            eg.lineStyle = {
              curveness: 0.07
            }
            eg.symbol = ['', 'arrow'];
          }
          return eg;
        })());
      });
    }
    if (cur.router as Array<Router> && cur.router.length !== 0) {
      cur.router.forEach((router: Router) => {
        let index = nodes.findIndex((r) => r.name == router.router_id)
        if (index !== -1) {
          nodes[index].area.push(cur.area_id)
          return
        }
        nodes.push({
          name: router.router_id,
          value: router.router_id,
          meta: router.metadata ? router.metadata : {},
          area: [cur.area_id]
        });
      });
    }
    return { nodes, edges };
  }, { nodes: [] as Array<Node>, edges: [] as Array<Edge> });

  option.series[0].data = nodes;
  option.series[0].links = edges;
});

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