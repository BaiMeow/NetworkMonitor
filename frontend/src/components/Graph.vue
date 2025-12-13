<!-- eslint-disable vue/multi-word-component-names -->
<script lang="ts" setup>
import { listenEchartAction, useGraph, useGraphEvent } from '@/state/graph'
import { ECharts } from 'echarts'
import VChart from 'vue-echarts'
import { TooltipComponent, TitleComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import { GraphChart } from 'echarts/charts'
import { use, ECElementEvent, ElementEvent } from 'echarts/core'
import { ref } from 'vue'
import { useASMetaLoading } from '@/state/meta'

use([CanvasRenderer, GraphChart, TooltipComponent, TitleComponent])

const { option, selectList, loading } = useGraph()
const { handleMouseDown, handleMouseUp, handleClick, handleZrClick } =
  useGraphEvent()
const echarts = ref<ECharts | null>(null)
listenEchartAction((payload, opt) => {
  echarts.value?.dispatchAction(payload, opt)
})
const ASMetaLoading = useASMetaLoading()

// let showLoadingTimeout = setTimeout(() => {}, 0)
// const showLoading = ref(true)
// watch(
//   () => loading.value,
//   (cur, old) => {
//     if (old && !cur) {
//       clearTimeout(showLoadingTimeout)
//       showLoading.value = false
//     } else if (cur && !old) {
//       showLoadingTimeout = setTimeout(() => {
//         showLoading.value = true
//       }, 200)
//     } else {
//       showLoading.value = cur
//     }
//   },
// )
</script>
<template>
  <Transition>
    <div v-if="loading || ASMetaLoading" class="graph loading">Loading...</div>
  </Transition>
  <Transition>
    <v-chart
      v-if="!(loading || ASMetaLoading)"
      ref="echarts"
      :option="option"
      class="graph"
      autoresize
      @mousedown="(e: ECElementEvent) => handleMouseDown?.(e)"
      @mouseup="(e: ECElementEvent) => handleMouseUp?.(e)"
      @click="(e: ECElementEvent) => handleClick?.(e)"
      @zr:click="(e: ElementEvent) => handleZrClick?.(e)"
    />
  </Transition>
  <div class="top-bar">
    <Dark />
    <SearchBar class="search-bar" :data="selectList" />
  </div>
</template>

<style lang="css" scoped>
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

.top-bar {
  position: absolute;
  display: flex;
  top: 2vh;
  right: 2vw;
  width: 14rem;
  align-items: center;
  gap: 1rem;
}

.search-bar {
  flex-grow: 1;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
