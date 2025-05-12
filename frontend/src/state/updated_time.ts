import { ref } from 'vue'

const updated_time = ref<Date | null>(null)
export function setUpdatedTime(t: Date | null) {
  updated_time.value = t
}
export function useUpdatedTime() {
  return updated_time
}
