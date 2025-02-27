<script setup lang="ts">
import { useASMeta } from '@/state/meta'
import { ref } from 'vue'

const { grName, asn } = defineProps<{
  grName: string
  asn: number
}>()

const ASMeta = useASMeta()

const graph_mode = ref('24h')
</script>

<template>
  <div class="uptime-panel">
    <div class="uptime-head">
      <div class="title">
        {{
          ASMeta?.metadata?.[asn + '']?.display
            ? `${ASMeta.metadata[asn + ''].display} Network`
            : `AS ${asn}`
        }}
      </div>
      <BGPLatestStatus :grName="grName" class="status-bar" :asn="asn" />
    </div>
    <div class="uptime-body">
      <div class="wrap-graph">
        <el-select class="graph-selecter" v-model="graph_mode" placeholder="">
          <el-option value="24h" label="1 day" />
          <el-option value="168h" label="7 days" />
        </el-select>
        <LinksHistory
          class="links-graph"
          :grName="grName"
          :asn="asn"
          :time="graph_mode"
        />
      </div>
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
  display: flex;
  align-items: center;
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
  background-color: rgba(176, 190, 197, 0.5);
  padding: 20px;
  height: 3rem;
  font-weight: bold;
  color: black;
  font-size: 20px;
  line-height: 2.5;
  font-family: 'Microsoft YaHei';
}

.uptime-body {
  background-color: rgba(176, 190, 197, 0.5);
  padding: 2rem;
  height: calc(100% - 7rem - 40px);
  border-bottom-left-radius: 20px;
  border-bottom-right-radius: 20px;
  font-family: 'Microsoft YaHei';
  font-size: 16px;
  line-height: 1.5;
  text-align: justify;
  margin-top: 2px;
}

.graph-selecter {
  position: absolute;
  width: 6rem;
  margin-right: auto;
  z-index: 10;
  right: 2rem;
}

.status-bar {
  margin-left: auto;
  margin-right: 1rem;
  float: right;
}

.wrap-graph {
  display: relative;
  height: 100%;
  width: 100%;
  float: right;
}

html.dark .uptime-panel {
  background-color: rgba(48, 48, 48, 0.7);
}

html.dark .uptime-head {
  background-color: rgba(48, 48, 48, 0.5);
  color: #e5eaf3;
}

html.dark .uptime-body {
  background-color: rgba(48, 48, 48, 0.5);
}
</style>
