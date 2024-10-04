<script lang="ts" setup>
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { GraphChart } from 'echarts/charts'
import { TooltipComponent, TitleComponent } from 'echarts/components'
import VChart from 'vue-echarts'
import { ECElementEvent, ECharts } from 'echarts'
import { reactive, inject, ref } from 'vue'

import { Netmask } from 'netmask'

import { getBGP } from '../api/bgp'
import { prettierNet } from '../utils/colornet'
import { ASData } from '../api/meta'
import { ASDataKey } from '../inject/key'
import { selectItem } from './searchbar.vue'

const echarts = ref<ECharts | null>()

const loading = ref(true)

const asdata = inject(ASDataKey)?.value

const selectList = ref([] as Array<selectItem>)

interface Edge {
  source: string
  target: string
  value: number
  lineStyle?: any
  symbol?: string[]
  emphasis?: any
}

interface Node {
  name: string
  value: string
  meta?: any
  peer_num: number
  symbolSize?: number
  symbol?: string
  network: string[]
  itemStyle?: any
}

interface Params<T> {
  dataType: string
  data: T
}

use([CanvasRenderer, GraphChart, TooltipComponent, TitleComponent])

function mergeObjects(obj1: any, obj2: any): any {
  for (const key in obj2) {
    if (
      obj2.hasOwnProperty(key) &&
      (obj1.hasOwnProperty(key) || !(key in obj1))
    ) {
      if (
        typeof obj2[key] === 'object' &&
        obj2[key] !== null &&
        typeof obj1[key] === 'object' &&
        obj1[key] !== null
      ) {
        mergeObjects(obj1[key], obj2[key])
      } else {
        obj1[key] = obj2[key]
      }
    }
  }
}
const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  ? true
  : false

const option: any = reactive({
  title: {
    text: 'DN11 & Vidar Network',
    subtext: '',
  },
  tooltip: {
    trigger: 'item',
    triggerOn: 'mousemove',
    backgroundColor: isDark ? '#333' : 'white',
    textStyle: {
      color: isDark ? 'white' : 'black',
    },
    confine: true,
    enterable: true,
    formatter: (params: Params<any>) => {
      if (params.dataType === 'edge') {
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
          const {
            monitor: { appendix },
          } = metadata
          for (let key in appendix) {
            const value = appendix[key] as string | string[]
            if (typeof value === 'string') {
              output += `<br/>${key}: ${value}`
            } else if (Array.isArray(value)) {
              output += `<br/>${key}:`
              for (let i in value) {
                output += `<br/> - ${value[i]}`
              }
            }
          }
        }
      }
      output += `<br/> network:<br/>`
      if (asdata) {
        output += prettierNet(
          params.data.network,
          params.data.name,
          asdata.announcements,
        )
      } else {
        output += params.data.network.join('<br/>')
        output += `<br/>`
      }
      output += `Peer Count: <div class="peer_count"> ${params.data.peer_num} </div>`
      return output
    },
    position: function () {
      return [20, 50]
    },
  },
  series: [
    {
      type: 'graph',
      symbolSize: 50,
      layout: 'force',
      lineStyle: {
        color: 'source',
        opacity: 0.4,
        width: 0.5,
        curveness: 0.1,
      },
      force: {
        repulsion: 500,
        gravity: 0.3,
        friction: 1,
        edgeLength: [10, 140],
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
          if (params.data.meta && params.data.meta.display) {
            return params.data.meta.display
          }
          return params.data.name
        },
      },
      labelLayout: {
        hideOverlap: true,
      },
      draggable: true,
      data: [],
      links: [],
      emphasis: {
        focus: 'adjacency',
        lineStyle: {
          width: 10,
        },
      },
    },
  ],
})

const refreshData = async () => {
  const resp = await getBGP()
  if (!resp.as) {
    alert('no data')
    return
  }

  const nodes = resp.as.reduce((nodes, cur) => {
    nodes.push({
      name: cur.asn.toString(),
      value: cur.asn.toString(),
      peer_num: 0,
      network: cur.network
        .sort((a, b) => parseInt(a.split('/')[1]) - parseInt(b.split('/')[1]))
        .reduce(
          (network, cur) =>
            network.findIndex((net) => {
              let nmask = new Netmask(net)
              return nmask.contains(cur) || nmask.toString() === cur
            }) === -1
              ? [...network, cur]
              : network,
          [] as string[],
        )
        .sort((a, b) => {
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
        }),
    })
    return nodes
  }, [] as Node[])

  nodes.map((node) => {
    node = reactive(node)
    node.peer_num = resp.link.filter((lk) => {
      return lk.src === parseInt(node.name) || lk.dst === parseInt(node.name)
    }).length
    node.value = '' + node.peer_num
    node.symbolSize = Math.pow(node.peer_num + 3, 1 / 2) * 7
    if (asdata?.metadata && node.name in asdata?.metadata) {
      const customNode = asdata.metadata[node.name].monitor?.customNode
      if (customNode) {
        mergeObjects(node, asdata.metadata[node.name].monitor?.customNode)
      }
      node.meta = asdata.metadata[node.name]
    }
    if (node.peer_num === 0) {
      node.symbol =
        'path://M255.633,0C145.341,0.198,55.994,89.667,55.994,200.006v278.66c0,14.849,17.953,22.285,28.453,11.786l38.216-39.328 l54.883,55.994c6.51,6.509,17.063,6.509,23.572,0L256,451.124l54.883,55.994c6.509,6.509,17.062,6.509,23.571,0l54.884-55.994 l38.216,39.327c10.499,10.499,28.453,3.063,28.453-11.786V201.719C456.006,91.512,365.84-0.197,255.633,0z M172.664,266.674 c-27.572,0-50.001-22.429-50.001-50.001s22.43-50.001,50.001-50.001s50.001,22.43,50.001,50.001S200.236,266.674,172.664,266.674z M339.336,266.674c-27.572,0-50.001-22.429-50.001-50.001s22.43-50.001,50.001-50.001s50.001,22.43,50.001,50.001 S366.908,266.674,339.336,266.674z'
    }
  })

  const edges = resp.link.reduce((edges, cur) => {
    const src = nodes.find((node) => node.name === cur.src.toString())
    const dst = nodes.find((node) => node.name === cur.dst.toString())
    if (src == null || dst == null) {
      return edges
    }
    edges.push({
      source: cur.src.toString(),
      target: cur.dst.toString(),
      value: 100 / Math.min(src.peer_num, dst.peer_num) + 10,
    })
    return edges
  }, [] as Edge[])

  const setLoadingOnce = (() => {
    let once = false
    return () => {
      if (once) return
      once = true
      loading.value = true
      option.series[0].force.friction = 1
      return
    }
  })()

  // remove not existed edges
  for (let i = 0; i < option.series[0].links.length; i++) {
    if (
      edges.findIndex(
        (edge) =>
          edge.source === option.series[0].links[i].source &&
          edge.target === option.series[0].links[i].target,
      ) === -1
    ) {
      setLoadingOnce()
      option.series[0].links.splice(i, 1)
      i--
    }
  }
  // refresh nodes
  for (let i = 0; i < option.series[0].data.length; i++) {
    const idx = nodes.findIndex(
      (node) => node.name === option.series[0].data[i].name,
    )
    if (idx === -1) {
      setLoadingOnce()
      option.series[0].data.splice(i, 1)
      i--
      continue
    }
    if (
      option.series[0].data[i].peer_num !== nodes[idx].peer_num ||
      option.series[0].data[i].network.join('|') !==
        nodes[idx].network.join('|')
    ) {
      setLoadingOnce()
      option.series[0].data[i] = nodes[idx]
    }
  }
  // add new nodes
  for (let i = 0; i < nodes.length; i++) {
    if (
      option.series[0].data.findIndex(
        (node: Node) => node.name === nodes[i].name,
      ) === -1
    ) {
      setLoadingOnce()
      option.series[0].data.push(nodes[i])
    }
  }
  // add new edges
  for (let i = 0; i < edges.length; i++) {
    if (
      option.series[0].links.findIndex(
        (edge: Edge) =>
          edge.source === edges[i].source && edge.target === edges[i].target,
      ) === -1
    ) {
      setLoadingOnce()
      option.series[0].links.push(edges[i])
    }
  }

  option.series[0].force.edgeLength[1] = nodes.length * 3.5
  option.title.subtext = `Nodes: ${nodes.reduce(
    (p, c) => p + (c.peer_num === 0 ? 0 : 1),
    0,
  )} Peers: ${edges.length}`
  option.series[0].force.friction = 0.15
  loading.value = false
  selectList.value = nodes.map((n) => {
    return {
      label: n.meta?.display || n.name,
      asn: n.value,
      name: n.name,
      display: n.meta?.display || n.name,
      network: [
        ...n.network,
        ...(asdata?.announcements.assigned
          .filter((a) => a.asn === n.name)
          .map((a) => a.prefix) || []),
      ],
      value: n.name,
      onselected: () => {
        echarts.value?.dispatchAction({
          type: 'highlight',
          seriesIndex: 0,
          name: n.name,
        })
        echarts.value?.dispatchAction({
          type: 'showTip',
          seriesIndex: 0,
          name: n.name,
        })
      },
    }
  })
}

refreshData()
setInterval(() => {
  refreshData()
}, 60 * 1000)

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
  }, 6000)
}
</script>

<template>
  <div v-if="loading" class="graph dark-mode loading">Loading...</div>
  <v-chart
    ref="echarts"
    :option="option"
    class="graph"
    autoresize
    @mousedown="handle_mouse_down"
    @mouseup="handle_mouse_up"
  />
  <searchbar class="search-bar" :data="selectList"></searchbar>
</template>

<style scoped>
.search-bar {
  position: absolute;
  top: 2vh;
  right: 2vw;
  width: 12rem;
}

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

@media (prefers-color-scheme: dark) {
  .loading {
    color: gray;
  }
}
</style>
