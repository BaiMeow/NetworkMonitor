import { ApiHost } from '@/api/consts'
import { useEventSource } from '@vueuse/core'
import { computed } from 'vue'

export const { status, data } = useEventSource(`${ApiHost}/api/update`, ['update'], {
  autoReconnect: {
    retries: 3,
  },
})

export const updatedData = computed(() =>
  data.value
    ? (JSON.parse(data.value) as {
        type: string
        key: string
      })
    : null,
)
