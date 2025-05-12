<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script setup lang="ts">
import { getOSPFUptimeLinks, UptimeLinks } from '@/api/uptime'
import { ref, reactive, onUnmounted, computed, watch } from 'vue'
import VChart from 'vue-echarts'
import { graphic } from 'echarts/core'
import { LineChart } from 'echarts/charts'
import {
  GridComponent,
  DataZoomComponent,
  ToolboxComponent,
} from 'echarts/components'
import { use } from 'echarts/core'
import { parseDuration } from '@/utils/time'
import { useDark } from '@vueuse/core'
import { fontColor } from '@/state/font'
const { asn, routerId, time } = defineProps<{
  asn: number
  routerId: string
  time: string
}>()

const isDark = useDark()
const window = computed(() =>
  parseDuration(time) > parseDuration('24h') ? '1h' : '1m',
)

const data = ref(Array<UptimeLinks>())

use([LineChart, GridComponent, DataZoomComponent, ToolboxComponent])

const itemStyle = {
  color: computed(() =>
    isDark.value ? 'rgb(214, 36, 95)' : 'rgb(255, 70, 131)',
  ),
}
const areaStyle = {
  color: computed(() =>
    isDark.value
      ? new graphic.LinearGradient(0, 0, 0, 1, [
          {
            offset: 0,
            color: 'rgb(191, 119, 51)',
          },
          {
            offset: 1,
            color: 'rgb(214, 36, 95)',
          },
        ])
      : new graphic.LinearGradient(0, 0, 0, 1, [
          {
            offset: 0,
            color: 'rgb(255, 158, 68)',
          },
          {
            offset: 1,
            color: 'rgb(255, 70, 131)',
          },
        ]),
  ),
}

const option: any = reactive({
  title: {
    text: 'Peer Router',
    textStyle: {
      color: fontColor,
    },
  },
  tooltip: {
    trigger: 'axis',
    position: function (pt: Array<any>) {
      return [pt[0], '20%']
    },
  },
  dataZoom: [
    {
      type: 'inside',
      start: 0,
      minSpan: 1,
      end: 100,
    },
    {
      start: 0,
      minSpan: 1,
      end: 100,
    },
  ],
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: [],
  },
  yAxis: {
    type: 'value',
  },
  series: [
    {
      name: 'In',
      type: 'line',
      symbol: 'none',
      sampling: 'lttb',
      data: [],
      itemStyle: itemStyle,
      areaStyle: areaStyle,
    },
    {
      name: 'Out',
      type: 'line',
      symbol: 'none',
      sampling: 'lttb',
      data: [],
      itemStyle: itemStyle,
      areaStyle: areaStyle,
    },
  ],
})

const refreshData = async () => {
  data.value = await getOSPFUptimeLinks(asn, routerId, time, window.value)
  option.series[0].data = data.value.in.map((d) => d.links)
  option.series[1].data = data.value.out.map((d) => d.links)
  option.xAxis.data = data.value.in.map(
    (d) =>
      `${d.time.getMonth() + 1}/${d.time.getDate()} ${d.time.getHours()}:${d.time.getMinutes()}`,
  )
}

refreshData()
let ticker = setInterval(refreshData, parseDuration(window.value) * 1000)
watch(
  () => time,
  () => {
    clearInterval(ticker)
    refreshData()
    ticker = setInterval(refreshData, parseDuration(window.value) * 1000)
  },
)
onUnmounted(() => clearInterval(ticker))
</script>

<template>
  <v-chart :option="option" autoresize />
</template>

<style scoped></style>
