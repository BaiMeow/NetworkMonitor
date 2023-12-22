<script lang="ts" setup>
import "echarts";
import { reactive, watchEffect } from "vue";

import VChart from "vue-echarts";
import { getOSPF } from "../api/ospf";

import { inject ,ref } from "vue";
import { ASDataKey } from "../inject/key";
import { ASData } from "../api/meta";

const props = defineProps<{
  asn: number;
  loaded: () => void;
}>();

interface Edge {
  source: string;
  target: string;
  value: number;
  cost: number;
  lineStyle?: any;
  symbol?: string[];
}

interface Node {
  name: string;
  value: string;
  meta?: any;
  peer_num?: number;
  symbolSize?: number;
  area: string[];
}

interface Params<T> {
  dataType: string;
  data: T;
}

const loading = ref(true);

const asdata = inject(ASDataKey)?.value as ASData;

const option: any = reactive({
  title: {
    text: '',
  },
  tooltip: {
    trigger: "item",
    triggerOn: "mousemove",
    formatter: (params: Params<any>) => {
      if (params.dataType === "edge") {
        params = params as Params<Edge>;
        return `Link: ${params.data.source} ↔ ${params.data.target} <br/> Cost: <div class="cost">${params.data.cost}</div>`;
      } else {
        params = params as Params<Node>;
        let output = `Router ID: ${params.data.value}`;
        if ("meta" in params.data) {
          output += "<br/>";
          for (let key in params.data.meta) {
            output += `${key}: ${params.data.meta[key]} <br/>`;
          }
          output += `Area: ${params.data.area.join(", ")}`;
          output += ` <br/> Peer Count: <div class="peer_count"> ${params.data.peer_num} </div>`;
        }
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
      circular: {
        rotateLabel: true,
      },
      animation: true,
      force: {
        initLayout: "circular",
        repulsion: 120,
        gravity: 0.02,
        edgeLength: [40, 300],
        friction: 0.15
      },
      zoom: 1,
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
      edgeLabel: {
        show: true,
        formatter: (params: any) => params.data.cost
      },
      data: [],
      links: [],
    },
  ],
  lineStyle: {
    opacity: 0.9,
    width: 2,
  },
});

watchEffect(async () => {
  loading.value = true
  getOSPF(props.asn).then(async areas => {
    const nodes = areas.reduce((nodes, cur) => {
      if ((cur.router) && cur.router.length !== 0) {
        cur.router.forEach((router) => {
          let index = nodes.findIndex((r) => r.name == router.router_id);
          if (index !== -1) {
            nodes[index].area.push(cur.area_id);
            return;
          }
          nodes.push({
            name: router.router_id,
            value: router.router_id,
            meta: router.metadata ? router.metadata : {},
            area: [cur.area_id],
          });
        });
      }
      return nodes;
    }, [] as Node[]);

    const all_links = areas.flatMap(
      area => area.links
    );

    const all_routers = areas.reduce(
      (routers, cur) => {
        if (cur.router === undefined || cur.router.length === 0) {
          return routers;
        }
        cur.router.forEach((r) => {
          if (routers.findIndex((router) => router.router_id === r.router_id) === -1) {
            routers.push(r);
          }
        })
        return routers
      },
      [] as (typeof areas)[number]['router']
    );

    // calculate node peers and size
    let { min, max } = all_links.reduce(
      ({ min, max }, cur) => {
        return {
          min: Math.min(min, cur.cost),
          max: Math.max(max, cur.cost),
        };
      },
      { min: all_links[0].cost, max: all_links[0].cost }
    );

    nodes.forEach((node) => {
      node.peer_num = all_links.filter((lk) => {
        return lk.src === node.name;
      }).length;
      node.symbolSize = Math.pow(node.peer_num, 1 / 2) * 7;
    });

    let edges: Edge[] = [];
    // prepare edges fro render
    all_routers.forEach((a) => {
      all_routers.forEach((b) => {
        if (a.router_id >= b.router_id) return;

        let links = all_links.filter((lk) => {
          return (lk.src === a.router_id && lk.dst === b.router_id) || (lk.src === b.router_id && lk.dst === a.router_id);
        });

        let lines: typeof all_links = [];
        let arrows: typeof all_links = [];
        while (links.length !== 0) {
          let cur = links.pop() as NonNullable<(typeof links)[number]>;
          let pair_idx = links.findIndex((lk) =>
            lk.src === cur.dst && lk.dst === cur.src && lk.cost === cur.cost
          );
          if (pair_idx !== -1) {
            links.splice(pair_idx, 1)[0];
            lines.push(cur);
          } else {
            arrows.push(cur);
          }
        }

        let curveness = 0.07;
        if (lines.length % 2 === 1) {
          const line = lines.pop() as NonNullable<(typeof lines)[number]>;
          edges.push({
            source: line.src,
            target: line.dst,
            value: 100 / line.cost,
            cost: line.cost
          })
        }
        let next_curveness = false
        while (lines.length !== 0) {
          const l1 = lines.pop() as NonNullable<(typeof lines)[number]>;
          edges.push({
            source: l1.src,
            target: l1.dst,
            value: 100 / l1.cost,
            cost: l1.cost,
            lineStyle: {
              curveness: next_curveness ? -curveness : curveness
            }
          });
          if (next_curveness) {
            curveness += 0.07;
            next_curveness = false;
          } else
            next_curveness = true;
        }

        let pre_source = arrows[arrows.length - 1]?.src;
        while (arrows.length !== 0) {
          const arrow = arrows.pop() as NonNullable<(typeof lines)[number]>;
          edges.push({
            source: arrow.src,
            target: arrow.dst,
            value: 100 / arrow.cost,
            cost: arrow.cost,
            lineStyle: {
              curveness: pre_source === arrow.src && next_curveness ? -curveness : curveness
            },
            symbol: ['', 'arrow']
          })
          if (next_curveness) {
            curveness += 0.07;
            next_curveness = false;
          } else
            next_curveness = true;
        }
      });
    });

    option.series[0].force.edgeLength = [(150 * min) / max, 150];
    option.series[0].force.repulsion = [(100 * max) / min, 100];
    option.series[0].data = nodes;
    option.series[0].links = edges;
    option.title.text = `${asdata.metadata[props.asn].display} Network`;
    loading.value = false;
  })
})

</script>

<template>
  <div v-if="loading" class="graph loading">
    Loading...
  </div>
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

.loading {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 2rem;
  font-weight: bold;
  color: #2242a3;
}
</style>
