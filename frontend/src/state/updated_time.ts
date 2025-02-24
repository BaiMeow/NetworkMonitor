import { ref } from 'vue'

const updated_time = ref<Date | null>()
export function setUpdatedTime (t: Date) {
  updated_time.value = t
}
export function useUpdatedTime(){
    return updated_time
}