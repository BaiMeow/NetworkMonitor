<script setup lang="ts">
import { computed } from 'vue'

const upColor = '#00c853'
const downColor = '#455a64'

const props = defineProps<{
  data: Array<number>
}>()
const current = computed(() => props.data[props.data.length - 1])
</script>

<template>
  <div class="panel">
    <div
      class="status"
      :style="{
        backgroundColor: current ? upColor : downColor,
        color: current ? '#37474f' : '#cfd8dc',
      }"
    >
      {{ current ? 'UP' : 'DOWN' }}
    </div>
    <div class="bar">
      <div
        v-for="d in props.data"
        class="square"
        :style="{
          backgroundColor: d ? upColor : downColor,
        }"
      ></div>
    </div>
  </div>
</template>

<style scoped>
.panel {
  display: flex;
  align-items: center;
}

.square {
  display: inline-block;
  width: 20px;
  height: 20px;
  margin: 5px;
  border-radius: 2px;
  transition: background-color 0.5s;
  transition-timing-function: ease;
}

.status {
  display: inline-block;
  width: 80px;
  height: 40px;
  border-radius: 20px;
  align-content: center;
  text-align: center;
  font-weight: bolder;
  transition: background-color 0.8s;
  transition-timing-function: ease;
}

.bar {
  margin-left: auto;
  display: inline-block;
  right: 0;
}
</style>
