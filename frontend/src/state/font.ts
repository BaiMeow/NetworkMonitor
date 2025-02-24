import { useDark } from '@vueuse/core'
import { computed } from 'vue'
const isDark = useDark()

export const fontColor = computed(() => (isDark.value ? '#E5EAF3' : 'black'))
