<script lang="ts" setup>
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
  cost: number
  lineStyle?: any
  symbol?: string[]
}

interface Node {
  name: string
  value: string
  meta?: any
  peer_num?: number
  symbolSize?: number
  area: string[]
}

interface Params<T> {
  dataType: string
  data: T
}

const option: any = reactive({
  title: {
    text: 'DN11 OSPF Status',
  },
  tooltip: {
    trigger: 'item',
    triggerOn: 'mousemove',
    formatter: (params: Params<any>) => {
      if (params.dataType === 'edge') {
        params = params as Params<Edge>
        return `Link: ${params.data.source} ↔ ${params.data.target} <br/> Cost: <div class="cost">${params.data.cost}</div>`;
      } else {
        params = params as Params<Node>
        let output = `Router ID: ${params.data.value}`;
        if ("meta" in params.data) {
          output += '<br/>';
          for (let key in params.data.meta) {
            output += `${key}: ${params.data.meta[key]} <br/>`;
          }
          output += `Area: ${params.data.area.join(', ')}`
          output += ` <br/> Peer Count: <div class="peer_count"> ${params.data.peer_num} </div>`
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
      circular: {
        rotateLabel: true
      },
      animation: true,
      force: {
        initLayout: 'circular',
        repulsion: 180,
        gravity: 0.02,
        edgeLength: [40, 200]
      },
      zoom: 1,
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
        formatter: (params: any) => {
          return params.data.cost;
        },
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
  let { nodes, edges } = area.reduce(({ nodes, edges }, cur) => {
    if (cur.links && cur.links.length !== 0) {
      cur.links.forEach((link: Link) => {
        edges.push((() => {
          let eg: Edge = {
            source: '' + link.src,
            target: link.dst,
            value: 100 / link.cost,
            cost: link.cost
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
          area: [cur.area_id],
        });
      });
    }
    return { nodes, edges };
  }, { nodes: [] as Array<Node>, edges: [] as Array<Edge> });
  let { min, max } = edges.reduce(
    ({ min, max }, cur) => {
      return {
        min: Math.min(min, cur.cost),
        max: Math.max(max, cur.cost)
      }
    }, { min: edges[0].cost, max: edges[0].cost });
  nodes.forEach((node) => {
    node.peer_num = edges.filter((edge) => {
      return edge.source === node.name
    }).length
    node.symbolSize = Math.pow(node.peer_num, 1 / 2) * 7;
  })
  option.series[0].force.edgeLength = [300 * min / max, 300];
  option.series[0].force.repulsion = [180 * max / min, 180];
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