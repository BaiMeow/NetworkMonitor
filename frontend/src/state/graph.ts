import { reactive, ref } from 'vue'
import { selectItem } from '@/components/SearchBar.vue'
import { ECElementEvent, ElementEvent, ECharts } from 'echarts'
import { Payload } from 'echarts'

const option = reactive({} as any)
const selectList = ref([] as Array<selectItem>)
const loading = ref(true)
const handleMouseDown = ref<(e: ECElementEvent) => void>()
const handleMouseUp = ref<(e: ECElementEvent) => void>()
const handleClick = ref<(e: ECElementEvent) => void>()
const handleZrClick = ref<(e: ElementEvent) => void>()
let cbDispatchEchartsAction: ECharts['dispatchAction'] = () => {}

export function useGraph() {
  return {
    option,
    selectList,
    loading,
  }
}

export function useGraphEvent() {
  return {
    handleMouseDown,
    handleMouseUp,
    handleClick,
    handleZrClick,
  }
}

export function dispatchEchartAction(
  payload: Payload,
  opt?:
    | boolean
    | {
        silent?: boolean
        flush?: boolean | undefined
      },
): void {
  cbDispatchEchartsAction(payload, opt)
}

export function listenEchartAction(cb: ECharts['dispatchAction']): void {
  cbDispatchEchartsAction = cb
}
