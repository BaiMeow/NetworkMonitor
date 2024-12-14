<script lang="ts" setup>
import { inject, watch, computed } from 'vue'

import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { GraphChart } from 'echarts/charts'
import { TooltipComponent, TitleComponent } from 'echarts/components'
import { ECElementEvent } from 'echarts'

import { getOSPF } from '../api/ospf'
import { ASDataKey } from '../inject/key'
import { ASData } from '../api/meta'
import { useDark } from '@vueuse/core'
import { onBeforeRouteLeave, useRoute } from 'vue-router'
import { useGraph, useGraphEvent } from '@/state/graph'

import { dispatchEchartAction } from '@/state/graph'

const route = useRoute()
use([CanvasRenderer, GraphChart, TooltipComponent, TitleComponent])
const isDark = useDark()

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

const asdata = inject(ASDataKey)?.value as ASData

option.title = {
  text: '',
  textStyle: {
    color: computed(() => (isDark.value ? '#E5EAF3' : 'black')),
  },
  subtext: '',
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
      let output = `Router ID: ${node.data.value}`
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
      color: computed(() => (isDark.value ? '#E5EAF3' : 'black')),
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

const load_data = async (asn: string) => {
  const areas = await getOSPF(parseInt(asn as string))
  const nodes = areas.reduce((nodes, cur) => {
    if (cur.router && cur.router.length !== 0) {
      cur.router.forEach((router) => {
        let index = nodes.findIndex((r) => r.name == router.router_id)
        if (index !== -1) {
          nodes[index].area.push(cur.area_id)
          return
        }
        nodes.push({
          name: router.router_id,
          value: router.router_id,
          meta: router.metadata ? router.metadata : {},
          subnet: router.subnet,
          area: [cur.area_id],
        })
      })
    }
    return nodes
  }, [] as Node[])

  const all_links = areas
    .flatMap((area) => area.links)
    .filter((link) => link !== undefined)

  const all_routers = areas.reduce(
    (routers, cur) => {
      if (cur.router === undefined || cur.router.length === 0) {
        return routers
      }
      cur.router.forEach((r) => {
        if (
          routers.findIndex((router) => router.router_id === r.router_id) === -1
        )
          routers.push(r)
      })
      return routers
    },
    [] as (typeof areas)[number]['router'],
  )

  nodes.forEach((node) => {
    let markedPeer = new Set<string>()
    node.peer_num = all_links.filter((lk) => {
      if (lk.src === node.name && !markedPeer.has(lk.dst)) {
        markedPeer.add(lk.dst)
        return true
      }
      return false
    }).length
    node.value = '' + node.peer_num
    node.symbolSize = Math.pow(node.peer_num + 3, 1 / 2) * 7
  })

  let edges: Edge[] = []
  // prepare edges for render
  all_routers.forEach((a) => {
    all_routers.forEach((b) => {
      if (a.router_id >= b.router_id) return

      let links = all_links.filter((lk) => {
        return (
          (lk.src === a.router_id && lk.dst === b.router_id) ||
          (lk.src === b.router_id && lk.dst === a.router_id)
        )
      })

      let lines: typeof all_links = []
      let arrows: typeof all_links = []
      while (links.length !== 0) {
        let cur = links.pop() as NonNullable<(typeof links)[number]>
        let pair_idx = links.findIndex(
          (lk) =>
            lk.src === cur.dst && lk.dst === cur.src && lk.cost === cur.cost,
        )
        if (pair_idx !== -1) {
          links.splice(pair_idx, 1)[0]
          lines.push(cur)
        } else {
          arrows.push(cur)
        }
      }

      let curveness = 0.07
      if (lines.length % 2 === 1) {
        const line = lines.pop() as NonNullable<(typeof lines)[number]>
        edges.push({
          source: line.src,
          target: line.dst,
          value: 100 / line.cost,
          cost: line.cost,
        })
      }
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
    })
  })

  option.series[0].force.edgeLength = [30, 150]
  option.series[0].force.repulsion = [30, 150]
  option.series[0].data = nodes
  option.series[0].links = edges
  option.title.text = asdata?.metadata?.[asn as string]?.display
    ? `${asdata.metadata[asn as string].display} Network`
    : `AS ${asn}`
  let markedPeer = new Set<string>()
  for (const link of all_links) {
    if (!markedPeer.has(link.src + link.dst)) {
      markedPeer.add(link.src + link.dst)
      markedPeer.add(link.dst + link.src)
    }
  }
  option.title.subtext = `Nodes: ${nodes.length}  Peers: ${markedPeer.size / 2}`
  selectList.value = nodes.map((n) => {
    return {
      label: n.name,
      value: n.value,
      onselected: () => {
        dispatchEchartAction({
          type: 'highlight',
          seriesIndex: 0,
          name: n.name,
        })
      },
    }
  })
}

watch(
  () => route.params.asn,
  async (new_asn) => {
    await load_data(new_asn as string)
    option.series[0].force.layoutAnimation = false
    loading.value = false
  },
  { immediate: true },
)

let timer: NodeJS.Timeout | null = null

const { handleMouseUp, handleMouseDown } = useGraphEvent()
handleMouseDown.value = (_: ECElementEvent) => {
  if (timer) {
    clearTimeout(timer)
  }
  option.series[0].force.friction = 0.15
  option.series[0].force.layoutAnimation = true
}

handleMouseUp.value = (_: ECElementEvent) => {
  timer = setTimeout(() => {
    option.series[0].force.layoutAnimation = false
  }, 6000)
}

onBeforeRouteLeave(() => {
  if (timer) {
    clearTimeout(timer)
  }
})
</script>

<template></template>
<style scoped></style>
