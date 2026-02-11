<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts" setup>
import { watch, computed, ref } from 'vue'

import { use, ECElementEvent, ElementEvent } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { GraphChart } from 'echarts/charts'
import { TooltipComponent, TitleComponent } from 'echarts/components'

import {
  Link,
  getBetweenness,
  getCloseness,
  getOSPF,
  getPathBetweenness,
  Router,
} from '../api/ospf'
import { ElementOf, useDark } from '@vueuse/core'
import { onBeforeRouteLeave, useRoute } from 'vue-router'
import { useGraph, useGraphEvent } from '@/state/graph'

import { dispatchEchartAction } from '@/state/graph'
import { useASMeta } from '@/state/meta'
import { fontColor } from '@/state/font'
import { setUpdatedTime } from '@/state/updated_time'

const route = useRoute()
use([CanvasRenderer, GraphChart, TooltipComponent, TitleComponent])
const isDark = useDark()
const focusedRouterId = ref('')

const { option, selectList, loading } = useGraph()

let firstLoading = true

loading.value = true

interface Edge {
  source: string
  target: string
  value: number
  cost: number
  betweenness?: number
  lineStyle?: any
  symbol?: string[]
  type: 'bidirectional' | 'unidirectional'
}

interface Node {
  name: string
  value: string
  betweenness?: number
  closeness?: number
  meta?: any
  peer_num?: number
  subnet?: string[]
  symbolSize?: number
  area: string[]
}

interface Params<T> {
  dataType: string
  data: T
}

const ASMeta = useASMeta()

const asn = computed<string>(() => route.params.asn as string)

import { updatedData } from '@/state/event'
import { calcEdgeWidthFromPathBetweenness } from '@/utils/edge_width'
watch(updatedData, (data) => {
  if (data?.type === 'ospf' && data.key === asn.value + '') loadData()
})

option.title = {
  text: computed(() =>
    ASMeta.value?.metadata?.[asn.value]?.display
      ? `${ASMeta.value.metadata[asn.value].display} Network`
      : `AS ${asn.value}`,
  ),
  left: '10',
  top: '10',
  textStyle: {
    color: fontColor,
  },
  subtext: computed(
    () => `Nodes: ${nodes.value?.length || 0}  Peers: ${peers.value || 0}`,
  ),
}
option.tooltip = {
  trigger: 'item',
  triggerOn: 'mousemove',
  backgroundColor: computed(() => (isDark.value ? '#333' : 'white')),
  textStyle: {
    color: computed(() => (isDark.value ? 'white' : 'black')),
  },
  confine: true,
  enterable: true,
  formatter: (params: Params<any>) => {
    if (params.dataType === 'edge') {
      params = params as Params<Edge>
      let output = `Link: ${params.data.source} ↔ ${params.data.target}`
      if (params.data.betweenness != undefined) {
        output += `<br/>Betweenness: ${params.data.betweenness.toFixed(4)}`
      }
      output += `<br/>Cost: <div class="cost">${params.data.cost}</div>`
      return output
    } else {
      const node = params as Params<Node>
      let output = `Router ID: ${node.data.name}`
      output += `<br/>Area: ${params.data.area.join(', ')}`
      if (Object.keys(node.data.meta).length !== 0) {
        output += '<br/>'
        for (let key in node.data.meta) {
          output += `${key}: ${node.data.meta[key]}`
        }
      }
      if (params.data.betweenness != undefined) {
        output += `<br/>Betweenness: ${params.data.betweenness.toFixed(3)}`
      }
      if (params.data.closeness != undefined) {
        output += `<br/>Closeness: ${params.data.closeness.toFixed(3)}`
      }
      if (node.data.subnet?.length) {
        output += '<br/>Subnet:'
        node.data.subnet.forEach((net) => {
          output += `<br/>${net}`
        })
      }
      output += ` <br/> Peer Count: <div class="peer_count"> ${params.data.peer_num} </div>`
      return output
    }
  },
  position: function () {
    return [20, 50]
  },
}
option.symbolSize = 50
option.animationDurationUpdate = 1500
option.animationEasingUpdate = 'quinticInOut'
option.series = [
  {
    type: 'graph',
    layout: 'force',
    circular: {
      rotateLabel: true,
    },
    force: {
      initLayout: 'circular',
      repulsion: 120,
      gravity: 0.02,
      edgeLength: [40, 300],
      friction: 1,
      layoutAnimation: false,
    },
    roam: true,
    label: {
      show: true,
      position: 'right',
      color: 'inherit',
      fontWeight: 1000,
      fontFamily: 'Microsoft YaHei',
      formatter: (params: any) => {
        if (params.data.meta && params.data.meta.name) {
          return params.data.meta.name
        }
        return params.data.name
      },
    },
    draggable: true,
    edgeLabel: {
      show: true,
      formatter: (params: any) => params.data.cost,
      padding: 0,
      color: fontColor,
    },
    data: [],
    links: [],
    emphasis: {
      focus: 'adjacency',
      lineStyle: {
        width: 10,
      },
    },
  },
]
option.lineStyle = {
  opacity: 0.9,
  width: 2,
}

// load data
const ospfData = ref<Awaited<ReturnType<typeof getOSPF>>>()
const betweenness = ref<Awaited<ReturnType<typeof getBetweenness>>>()
const closeness = ref<Awaited<ReturnType<typeof getCloseness>>>()
const pathBetweenness = ref<Awaited<ReturnType<typeof getPathBetweenness>>>()
async function loadData() {
  let asnInt = parseInt(asn.value)
  closeness.value = await getCloseness(asnInt)
  betweenness.value = await getBetweenness(asnInt)
  pathBetweenness.value = await getPathBetweenness(asnInt)
  ospfData.value = await getOSPF(asnInt)
}

// auto refresh
let autoRefreshInterval: ReturnType<typeof setTimeout>
watch(
  [asn],
  () => {
    focusedRouterId.value = ''
    if (autoRefreshInterval) clearInterval(autoRefreshInterval)
    firstLoading = true
    loadData()
    autoRefreshInterval = setInterval(() => {
      loadData()
    }, 60 * 1000)
  },
  {
    immediate: true,
  },
)
onBeforeRouteLeave(() => {
  if (autoRefreshInterval) clearInterval(autoRefreshInterval)
})

const all_links = computed(() =>
  ospfData.value?.graph
    .flatMap((area) => area.links)
    .filter((link) => link !== undefined),
)

const all_routers = computed(() =>
  ospfData.value?.graph.reduce<Array<Router>>(
    (routers, cur) => {
      cur.router.forEach((r) => {
        if (
          routers.findIndex((router) => router.router_id === r.router_id) === -1
        )
          routers.push(r)
      })
      return routers
    },
    [] as ElementOf<typeof ospfData.value>['router'],
  ),
)

const nodes = computed(() =>
  ospfData.value?.graph.reduce((nodes, cur) => {
    cur.router.forEach((router) => {
      let index = nodes.findIndex((r) => r.name == router.router_id)
      if (index !== -1) {
        nodes[index].area.push(cur.area_id)
        return
      }
      let markedPeer = new Set<string>()
      const peer_num =
        all_links.value?.filter((lk) => {
          if (lk.src === router.router_id && !markedPeer.has(lk.dst)) {
            markedPeer.add(lk.dst)
            return true
          }
          return false
        }).length || 0
      nodes.push({
        name: router.router_id,
        value: '' + peer_num,
        betweenness: betweenness.value?.[router.router_id],
        closeness: closeness.value?.[router.router_id],
        meta: router.metadata ? router.metadata : {},
        subnet: router.subnet,
        area: [cur.area_id],
        peer_num: peer_num,
        symbolSize: Math.pow(peer_num + 3, 1 / 2) * 7,
      })
    })
    return nodes
  }, [] as Node[]),
)

const edges = computed(() =>
  all_routers.value
    ?.flatMap((value, index, array) => {
      return array.slice(index + 1).map((r) => [value, r])
    })
    .flatMap(([a, b]) => {
      let edges: Edge[] = []

      // all links between two nodes
      let links =
        all_links.value?.filter((lk) => {
          return (
            (lk.src === a.router_id && lk.dst === b.router_id) ||
            (lk.src === b.router_id && lk.dst === a.router_id)
          )
        }) || []

      let lines: NonNullable<typeof all_links.value> = []
      let arrows: typeof all_links.value = []

      // convert to line or arrow
      while (links.length !== 0) {
        let cur = links.pop() as NonNullable<ElementOf<typeof links>>
        let pair_idx = links.findIndex(
          (lk) =>
            lk.src === cur.dst && lk.dst === cur.src && lk.cost === cur.cost,
        )
        if (pair_idx !== -1) {
          links.splice(pair_idx, 1)
          lines.push(cur)
        } else {
          arrows.push(cur)
        }
      }

      lines = lines.map((line) =>
        line.src < line.dst
          ? line
          : {
              src: line.dst,
              dst: line.src,
              cost: line.cost,
            },
      )

      // middle line
      if (lines.length % 2 === 1) {
        const line = lines.pop() as NonNullable<(typeof lines)[number]>
        const betweenness =
          (pathBetweenness.value?.find(
            (p) =>
              p.src == line.src && p.dst == line.dst && p.cost == line.cost,
          )?.betweenness || 0) +
          (pathBetweenness.value?.find(
            (p) =>
              p.src == line.dst && p.dst == line.src && p.cost == line.cost,
          )?.betweenness || 0)

        edges.push({
          source: line.src,
          target: line.dst,
          value: 100 / line.cost,
          cost: line.cost,
          betweenness: betweenness,
          lineStyle: {
            width: calcEdgeWidthFromPathBetweenness(
              betweenness,
              nodes.value?.length,
            ),
          },
          type: 'bidirectional',
        })
      }

      let curveness = 0.07
      let next_curveness = false
      while (lines.length !== 0) {
        const l1 = lines.pop() as NonNullable<(typeof lines)[number]>
        const betweenness =
          (pathBetweenness.value?.find(
            (p) => p.src == l1.src && p.dst == l1.dst && p.cost == l1.cost,
          )?.betweenness || 0) +
          (pathBetweenness.value?.find(
            (p) => p.src == l1.dst && p.dst == l1.src && p.cost == l1.cost,
          )?.betweenness || 0)
        edges.push({
          source: l1.src,
          target: l1.dst,
          value: 100 / l1.cost,
          cost: l1.cost,
          betweenness: betweenness,
          lineStyle: {
            curveness: next_curveness ? -curveness : curveness,
            width: calcEdgeWidthFromPathBetweenness(
              betweenness,
              nodes.value?.length,
            ),
          },
          type: 'bidirectional',
        })
        if (next_curveness) {
          curveness += 0.07
          next_curveness = false
        } else next_curveness = true
      }

      let pre_source = arrows[arrows.length - 1]?.src

      while (arrows.length !== 0) {
        const arrow = arrows.pop() as NonNullable<(typeof lines)[number]>
        const betweenness =
          pathBetweenness.value?.find(
            (p) =>
              p.src == arrow.src && p.dst == arrow.dst && p.cost == arrow.cost,
          )?.betweenness || 0
        edges.push({
          source: arrow.src,
          target: arrow.dst,
          value: 100 / arrow.cost,
          cost: arrow.cost,
          betweenness: betweenness,
          lineStyle: {
            curveness:
              pre_source === arrow.src && next_curveness
                ? -curveness
                : curveness,
            width: calcEdgeWidthFromPathBetweenness(
              betweenness,
              nodes.value?.length,
            ),
          },
          symbol: ['', 'arrow'],
          type: 'unidirectional',
        })
        if (next_curveness) {
          curveness += 0.07
          next_curveness = false
        } else next_curveness = true
      }
      return edges
    }),
)

const peers = computed(() => {
  let markedPeer = new Set<string>()
  for (const link of all_links.value || []) {
    if (!markedPeer.has(link.src + link.dst)) {
      markedPeer.add(link.src + link.dst)
      markedPeer.add(link.dst + link.src)
    }
  }
  return markedPeer.size / 2
})

const renderData = async () => {
  if (!nodes.value || !edges.value) return
  option.series[0].force.edgeLength = [30, 150]
  option.series[0].force.repulsion = 200
  option.series[0].data = nodes
  option.series[0].links = edges
  if (ospfData.value) setUpdatedTime(ospfData.value?.updated_at)
}

watch([nodes], () => {
  if (!nodes.value) return
  selectList.value = nodes.value.map((n) => {
    return {
      label: n.name,
      value: n.name,
      onselected: () => {
        dispatchEchartAction({
          type: 'highlight',
          seriesIndex: 0,
          name: n.name,
        })
      },
    }
  })
})

let timer: ReturnType<typeof setTimeout> | null = null
watch(ospfData, async () => {
  if (!firstLoading) {
    option.series[0].force.layoutAnimation = true
  }
  await renderData()
  loading.value = false
  if (!firstLoading) {
    if (timer) {
      clearTimeout(timer)
    }
    timer = setTimeout(() => {
      option.series[0].force.layoutAnimation = false
    }, 3000)
  }
  firstLoading = false
})

const { handleClick, handleZrClick, handleMouseUp, handleMouseDown } =
  useGraphEvent()
handleMouseDown.value = () => {
  if (timer) {
    clearTimeout(timer)
  }
  option.series[0].force.friction = 0.15
  option.series[0].force.layoutAnimation = true
}

handleMouseUp.value = () => {
  timer = setTimeout(() => {
    option.series[0].force.layoutAnimation = false
  }, 6000)
}

onBeforeRouteLeave(() => {
  if (timer) {
    clearTimeout(timer)
  }
})

interface LinkWithIndex {
  index: number
  link: Link
}

function getEdgeIndexesInShortestPathOf(src: string) {
  let node_distance:Map<string, number> = new Map();
  node_distance.set(src, 0);
  let unvisited_links: Array<LinkWithIndex> =
    all_links.value?.map((e, i) => {
      return {
        index: i,
        link: e,
      }
    }) || []
  let visited_links: Array<LinkWithIndex> = []
  let found_new_node_pos = 1;

  while (node_distance.size <= (nodes.value?.length || 0) && unvisited_links.length > 0) {
    let allow_start = Array.from(node_distance.keys());
    let min_cost_edge = unvisited_links.filter((e) => allow_start.includes(e.link.src)).reduce<LinkWithIndex | null>((min_cost_edge, e) => {
      if (min_cost_edge === null) return e;
      if (e.link.cost+node_distance.get(e.link.src)! < min_cost_edge.link.cost+node_distance.get(min_cost_edge.link.src)!) return e;
      return min_cost_edge;
    }, null)
    if (min_cost_edge === null) break;
    unvisited_links = unvisited_links.filter((e) => e.index !== min_cost_edge.index);
    if (node_distance.get(min_cost_edge.link.dst) === undefined || node_distance.get(min_cost_edge.link.dst)! === node_distance.get(min_cost_edge.link.src)! + min_cost_edge.link.cost) {
      node_distance.set(min_cost_edge.link.dst, node_distance.get(min_cost_edge.link.src)! + min_cost_edge.link.cost);
      visited_links.push(min_cost_edge);
      found_new_node_pos = visited_links.length;
    }
  }
  let tree_links = visited_links.slice(0, found_new_node_pos)
  return edges.value?.map((e, i) => {
    return {
      index: i,
      ...e,
    }
  }).flatMap((e) => {
    if (e.type === 'bidirectional') {
      return tree_links.some((link) => link.link.src === e.source && link.link.dst === e.target && link.link.cost === e.cost || link.link.src === e.target && link.link.dst === e.source && link.link.cost === e.cost) ? e.index : [];
    } else if (e.type === 'unidirectional') {
      return tree_links.some((link) => link.link.src === e.source && link.link.dst === e.target && link.link.cost === e.cost) ? e.index : [];
    }
    return null;
  }).flat().filter((i) => i !== null) || [];
}

handleClick.value = (e: ECElementEvent) => {
  if (e.dataType === 'node') {
    const data = e.data as Node
    let edge_idxes = getEdgeIndexesInShortestPathOf(data.name)
    let node_idxes = nodes.value?.map((n,i) => {
      return {
        index: i,
        ...n
      }
    }).filter((n) => edge_idxes.some((i) => edges.value?.[i]?.source === n.name || edges.value?.[i]?.target === n.name)).map((n) => n.index)
    dispatchEchartAction({
      type: 'highlight',
      batch: [{
        dataType: 'edge',
        dataIndex: edge_idxes,
        }, {
        dataType: 'node',
        dataIndex: node_idxes,
        notBlur: true
        }],
    })
    focusedRouterId.value = data.name
  }
}

handleZrClick.value = (e: ElementEvent) => {
  if (e.target === undefined) {
    focusedRouterId.value = ''
  }
}
</script>

<template>
  <Transition name="fade" appear> </Transition>
</template>
<style scoped>
.uptime {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: 80vw;
  height: 80vh;
  margin: auto;
}

.fade-enter-active {
  transition: all 0.2s ease-in;
}

.fade-leave-active {
  transition: all 0.16s ease-out;
}

.fade-enter-from {
  opacity: 0;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
