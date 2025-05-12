<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts" setup>
import { watch, computed, ref } from 'vue'

import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { GraphChart } from 'echarts/charts'
import { ECElementEvent, ElementEvent } from 'echarts'
import { TooltipComponent, TitleComponent } from 'echarts/components'

import { getOSPF, Router } from '../api/ospf'
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
const uptimeRouterId = ref("")

const { option, selectList, loading } = useGraph()

loading.value = true

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
watch(updatedData, (data) => {
  if (data?.type === 'ospf' && data.key === asn.value + '') loadData()
})

option.title = {
  text: computed(() =>
    ASMeta.value?.metadata?.[asn.value]?.display
      ? `${ASMeta.value.metadata[asn.value].display} Network`
      : `AS ${asn.value}`,
  ),
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
      return `Link: ${params.data.source} ↔ ${params.data.target} <br/> Cost: <div class="cost">${params.data.cost}</div>`
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
async function loadData() {
  ospfData.value = await getOSPF(parseInt(asn.value))
}

// auto refresh
let autoRefreshInterval: ReturnType<typeof setTimeout>
watch(
  [asn],
  () => {
    uptimeRouterId.value = ""
    if (autoRefreshInterval) clearInterval(autoRefreshInterval)
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
        edges.push({
          source: line.src,
          target: line.dst,
          value: 100 / line.cost,
          cost: line.cost,
        })
      }

      let curveness = 0.07
      let next_curveness = false
      while (lines.length !== 0) {
        const l1 = lines.pop() as NonNullable<(typeof lines)[number]>
        edges.push({
          source: l1.src,
          target: l1.dst,
          value: 100 / l1.cost,
          cost: l1.cost,
          lineStyle: {
            curveness: next_curveness ? -curveness : curveness,
          },
        })
        if (next_curveness) {
          curveness += 0.07
          next_curveness = false
        } else next_curveness = true
      }

      let pre_source = arrows[arrows.length - 1]?.src
      while (arrows.length !== 0) {
        const arrow = arrows.pop() as NonNullable<(typeof lines)[number]>
        edges.push({
          source: arrow.src,
          target: arrow.dst,
          value: 100 / arrow.cost,
          cost: arrow.cost,
          lineStyle: {
            curveness:
              pre_source === arrow.src && next_curveness
                ? -curveness
                : curveness,
          },
          symbol: ['', 'arrow'],
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

watch(
  ospfData,
  async () => {
    await renderData()
    option.series[0].force.layoutAnimation = false
    loading.value = false
  },
  { immediate: true },
)

let timer: ReturnType<typeof setTimeout> | null = null

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

handleClick.value = (e: ECElementEvent) => {
  if (e.dataType === 'node') {
    const data = e.data as Node
    uptimeRouterId.value = data.name
  }
}

handleZrClick.value = (e: ElementEvent) => {
  if (e.target === undefined) {
    uptimeRouterId.value = ""
  }
}
</script>

<template>
  <Transition name="fade" appear>
    <OSPFUptime
      class="uptime"
      v-if="uptimeRouterId"
      :routerId="uptimeRouterId"
      :asn="asn"
    />
  </Transition>
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
