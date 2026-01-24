<script setup lang="ts">
import { ref, onUnmounted } from 'vue'
import { getOSPFUptimeRecent } from '@/api/uptime'

const { asn, routerId } = defineProps<{
  asn: number
  routerId: string
}>()

const uptime10 = ref([
  true,
  true,
  true,
  true,
  true,
  true,
  true,
  true,
  true,
  true,
])

const update = async () => {
  const uptimes = await getOSPFUptimeRecent(asn, routerId)
  uptime10.value = uptimes.slice(0, 10)
}
update()
const ticker = setInterval(update, 1000 * 60)
onUnmounted(() => clearInterval(ticker))
</script>

<template>
  <LatestStatus :data="uptime10" />
</template>
