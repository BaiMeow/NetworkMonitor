<script setup lang="ts">
import { ref, inject, onUnmounted } from 'vue'
import { ASDataKey } from '@/inject/key'
import { getUptimeRecent } from '@/api/uptime'

const { asn } = defineProps<{
  asn: number
}>()

const asdata = inject(ASDataKey)?.value

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
  const uptimes = await getUptimeRecent(asn)
  uptime10.value = uptimes.slice(0, 10)
}
update()
const ticker = setInterval(update, 1000 * 60)
onUnmounted(() => clearInterval(ticker))
</script>

<template>
  <div class="uptime-panel">
    <div class="uptime-head">
      {{
        asdata?.metadata?.[asn + '']?.display
          ? `${asdata.metadata[asn + ''].display} Network`
          : `AS ${asn}`
      }}
    </div>
    <div class="uptime-body">
      <LatestStatus :data="uptime10" />
    </div>
  </div>
</template>

<style scoped>
.uptime-panel {
  display: flex;
  flex-direction: column;
  flex-wrap: 2px;
  height: 80vh;
  width: 80vw;
  max-width: 800px;
  max-height: 600px;
  background-color: rgba(207, 216, 220, 0.7);
  backdrop-filter: blur(6px);
  border-radius: 20px;
}

.uptime-head {
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
  background-color: rgba(176, 190, 197, 0.5);
  padding: 10px;
  padding-left: 20px;
  height: 3rem;
  font-weight: bold;
  color: #37474f;
  font-size: 20px;
  line-height: 2.5;
  font-family: 'Microsoft YaHei';
}

.uptime-body {
  background-color: rgba(176, 190, 197, 0.5);
  padding: 2rem;
  height: calc(100% - 7rem - 22px);
  border-bottom-left-radius: 20px;
  border-bottom-right-radius: 20px;
  font-family: 'Microsoft YaHei';
  color: #37474f;
  font-size: 16px;
  line-height: 1.5;
  text-align: justify;
  margin-top: 2px;
}

html.dark .uptime-panel {
  background-color: rgba(48, 48, 48, 0.7);
}

html.dark .uptime-head {
  background-color: rgba(48, 48, 48, 0.5);
  color: #cfd8dc;
}

html.dark .uptime-body {
  background-color: rgba(48, 48, 48, 0.5);
  color: #cfd8dc;
}
</style>
