import { ASData, loadASData } from '@/api/meta'
import { ref } from 'vue'

const ASMeta = ref<ASData>()
const doneLoading = ref<Boolean>(true)

async function loadMetadata() {
  const data = await loadASData()
  if (data) {
    ASMeta.value = data
  }
  doneLoading.value = false
}
loadMetadata()

export function useASMeta() {
  return ASMeta
}

export function useASMetaLoading() {
    return doneLoading
}
